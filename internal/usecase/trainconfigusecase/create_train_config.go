package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"time"

	"uuid"
)

type CreateTrainConfigInput struct {
	Name        string
	Description string
	TrainingDTO TrainingDTO
	CallerID    uuid.UUID
}

type CreateTrainConfigOutput struct {
	CreatedAt     time.Time
	Name          string
	Description   string
	TrainConfigID uuid.UUID
}

func (s *TrainConfigService) CreateTrainConfig(input CreateTrainConfigInput) (*CreateTrainConfigOutput, error) {
	name, err := trainconfig.NewName(input.Name)
	if err != nil {
		return nil, err
	}

	description, err := trainconfig.NewDescription(input.Description)
	if err != nil {
		return nil, err
	}

	training, err := buildTrainingFromDTO(input.TrainingDTO)
	if err != nil {
		return nil, err
	}

	exists, err := s.userRepository.ExistsByID(input.CallerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, &UserNotFoundError{}
	}

	exists, err = s.trainConfigRepository.ExistsByName(input.CallerID, name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &TrainConfigNameAlreadyExistsError{Name: input.Name}
	}

	newConfig, err := trainconfig.NewTrainConfig(input.CallerID, name, description, training)
	if err != nil {
		return nil, err
	}

	err = s.trainConfigRepository.Save(newConfig)
	if err != nil {
		return nil, err
	}

	output := &CreateTrainConfigOutput{
		Name:          newConfig.Name().Value(),
		Description:   newConfig.Description().Value(),
		TrainConfigID: newConfig.TrainConfigID(),
		CreatedAt:     newConfig.CreatedAt(),
	}

	return output, nil
}
