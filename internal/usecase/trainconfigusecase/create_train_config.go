package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/training"
	"time"

	"github.com/google/uuid"
)

type CreateTrainConfigInput struct {
	Training    *training.Training
	Name        string
	Description string
	CallerID    uuid.UUID
}

type CreateTrainConfigOutput struct {
	CreatedAt   time.Time
	Name        string
	Description string
	TrainingId  uuid.UUID
}

func (s *TrainConfigService) CreateTrainConfig(input CreateTrainConfigInput) (*CreateTrainConfigOutput, error) {
	exists, err := s.userRepository.ExistsByID(input.CallerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, &UserNotFoundError{}
	}

	exists, err = s.trainConfigRepository.ExistsByName(input.CallerID, input.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &TrainConfigNameAlreadyExistsError{Name: input.Name}
	}

	newConfig, err := trainconfig.NewTrainConfig(input.CallerID, input.Name, input.Description, input.Training)
	if err != nil {
		return nil, err
	}

	err = s.trainConfigRepository.Save(newConfig)
	if err != nil {
		return nil, err
	}

	output := &CreateTrainConfigOutput{
		Name:        newConfig.Name(),
		Description: newConfig.Description(),
		TrainingId:  newConfig.TrainConfigID(),
	}

	return output, nil
}
