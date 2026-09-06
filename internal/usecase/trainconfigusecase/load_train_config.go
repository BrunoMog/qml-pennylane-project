package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/training"
	"time"

	"github.com/google/uuid"
)

type LoadTrainConfigInput struct {
	TrainConfigID   *uuid.UUID
	TrainConfigName *string
	CallerID        uuid.UUID
}

type LoadTrainConfigOutput struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Description   string
	Training      training.Training
	OwnerID       uuid.UUID
	TrainConfigID uuid.UUID
}

func (s *TrainConfigService) LoadTrainConfig(input LoadTrainConfigInput) (*LoadTrainConfigOutput, error) {
	if input.TrainConfigID == nil && input.TrainConfigName == nil {
		return nil, &InvalidInputError{}
	}
	var config *trainconfig.TrainConfig
	var err error
	if input.TrainConfigID != nil {
		config, err = s.trainConfigRepository.FindByID(*input.TrainConfigID)
		if err != nil {
			return nil, err
		}
	} else if input.TrainConfigName != nil {
		config, err = s.trainConfigRepository.FindByName(input.CallerID, *input.TrainConfigName)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, &InvalidInputError{}
	}

	if !canLoadTrainConfig(input.CallerID, config) {
		return nil, &UnauthorizedError{}
	}

	output := &LoadTrainConfigOutput{
		Name:          config.Name(),
		Description:   config.Description(),
		Training:      config.Training(),
		CreatedAt:     config.CreatedAt(),
		UpdatedAt:     config.UpdatedAt(),
		OwnerID:       config.OwnerID(),
		TrainConfigID: config.TrainConfigID(),
	}
	return output, nil
}

func canLoadTrainConfig(callerID uuid.UUID, trainConfig *trainconfig.TrainConfig) bool {
	return callerID == trainConfig.OwnerID()
}

type LoadAllTrainConfigsInput struct {
	CallerID uuid.UUID
}

type LoadAllTrainConfigsOutput struct {
	TrainConfigs []LoadTrainConfigOutput
}

func (s *TrainConfigService) LoadAllTrainConfigs(input LoadAllTrainConfigsInput) (*LoadAllTrainConfigsOutput, error) {
	configs, err := s.trainConfigRepository.FindAllByOwnerID(input.CallerID)
	if err != nil {
		return nil, err
	}

	output := &LoadAllTrainConfigsOutput{
		TrainConfigs: make([]LoadTrainConfigOutput, len(configs)),
	}
	for i, config := range configs {
		output.TrainConfigs[i] = LoadTrainConfigOutput{
			Name:          config.Name(),
			Description:   config.Description(),
			Training:      config.Training(),
			CreatedAt:     config.CreatedAt(),
			UpdatedAt:     config.UpdatedAt(),
			OwnerID:       config.OwnerID(),
			TrainConfigID: config.TrainConfigID(),
		}
	}
	return output, nil
}
