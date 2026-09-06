package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/training"

	"github.com/google/uuid"
)

type UpdateTrainConfigInput struct {
	Name          *string
	Description   *string
	Training      *training.Training
	CallerID      uuid.UUID
	TrainConfigID uuid.UUID
}

func (s *TrainConfigService) UpdateTrainConfig(input UpdateTrainConfigInput) error {
	if input.Name == nil && input.Description == nil && input.Training == nil {
		return &NoFieldsToUpdateError{}
	}
	config, err := s.trainConfigRepository.FindByID(input.TrainConfigID)
	if err != nil {
		return err
	}

	if !canUpdateTrainConfig(input.CallerID, config) {
		return &UnauthorizedError{}
	}

	if input.Name != nil {
		exits, err := s.trainConfigRepository.ExistsByName(input.CallerID, *input.Name)
		if err != nil {
			return err
		}
		if exits {
			return &TrainConfigNameAlreadyExistsError{}
		}
		err = config.SetName(*input.Name)
		if err != nil {
			return err
		}
	}
	if input.Description != nil {
		err = config.SetDescription(*input.Description)
		if err != nil {
			return err
		}
	}
	if input.Training != nil {
		err = config.SetTraining(input.Training)
		if err != nil {
			return err
		}
	}

	err = s.trainConfigRepository.Save(config)
	if err != nil {
		return err
	}

	return nil
}

func canUpdateTrainConfig(callerID uuid.UUID, config *trainconfig.TrainConfig) bool {
	return config.OwnerID() == callerID
}
