package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"

	"uuid"
)

type UpdateTrainConfigInput struct {
	Name          *string
	Description   *string
	TrainingDTO   *TrainingDTO
	CallerID      uuid.UUID
	TrainConfigID uuid.UUID
}

func (s *TrainConfigService) UpdateTrainConfig(input UpdateTrainConfigInput) error {
	if input.Name == nil && input.Description == nil && input.TrainingDTO == nil {
		return &NoFieldsToUpdateError{}
	}
	config, err := s.trainConfigRepository.FindByID(input.TrainConfigID)
	if err != nil {
		return err
	}

	if !canUpdateTrainConfig(input.CallerID, config) {
		return &UnauthorizedError{}
	}

	var needToSave bool

	if input.Name != nil {
		name, err := trainconfig.NewName(*input.Name)
		if err != nil {
			return err
		}

		if config.Name().Value() != name.Value() {
			if !name.Equals(config.Name()) {
				exists, err := s.trainConfigRepository.ExistsByName(input.CallerID, name)
				if err != nil {
					return err
				}
				if exists {
					return &TrainConfigNameAlreadyExistsError{}
				}
			}
			config.SetName(name)
			needToSave = true
		}
	}

	if input.Description != nil {
		description, err := trainconfig.NewDescription(*input.Description)
		if err != nil {
			return err
		}
		if config.Description() != description {
			config.SetDescription(description)
			needToSave = true
		}
	}

	if input.TrainingDTO != nil {
		training, err := buildTrainingFromDTO(*input.TrainingDTO)
		if err != nil {
			return err
		}
		if !config.Training().Equals(training) {
			err = config.SetTraining(training)
			if err != nil {
				return err
			}
			needToSave = true
		}
	}

	if needToSave {
		err = s.trainConfigRepository.Save(config)
		if err != nil {
			return err
		}
	}

	return nil
}

func canUpdateTrainConfig(callerID uuid.UUID, config *trainconfig.TrainConfig) bool {
	return config.OwnerID() == callerID
}
