package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"time"

	"github.com/google/uuid"
)

type LoadVQCConfigInput struct {
	VQCConfigID   *uuid.UUID
	VQCConfigName *string
	CallerID      uuid.UUID
}

type LoadVQCConfigOutput struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	VQC         vqc.VQC
	OwnerID     uuid.UUID
	VQCConfigID uuid.UUID
}

func (s *VQCConfigService) LoadVQCConfig(input LoadVQCConfigInput) (*LoadVQCConfigOutput, error) {
	if input.VQCConfigID == nil && input.VQCConfigName == nil {
		return nil, &InvalidInputError{}
	}
	var config *vqcconfig.VQCConfig
	var err error
	if input.VQCConfigID != nil {
		config, err = s.vqcConfigRepository.FindByID(*input.VQCConfigID)
		if err != nil {
			return nil, err
		}
	} else if input.VQCConfigName != nil {
		config, err = s.vqcConfigRepository.FindByName(input.CallerID, *input.VQCConfigName)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, &InvalidInputError{}
	}

	if !canLoadVQCConfig(input.CallerID, config) {
		return nil, &UnauthorizedError{}
	}

	output := &LoadVQCConfigOutput{
		Name:        config.Name(),
		Description: config.Description(),
		VQC:         config.VQC(),
		CreatedAt:   config.CreatedAt(),
		UpdatedAt:   config.UpdatedAt(),
		OwnerID:     config.OwnerID(),
		VQCConfigID: config.VQCConfigID(),
	}
	return output, nil
}

func canLoadVQCConfig(callerID uuid.UUID, config *vqcconfig.VQCConfig) bool {
	return callerID == config.OwnerID()
}

type LoadAllVQCConfigsInput struct {
	CallerID uuid.UUID
}

type LoadAllVQCConfigsOutput struct {
	VQCConfigs []LoadVQCConfigOutput
}

func (s *VQCConfigService) LoadAllVQCConfigs(input LoadAllVQCConfigsInput) (*LoadAllVQCConfigsOutput, error) {
	configs, err := s.vqcConfigRepository.FindAllByOwnerID(input.CallerID)
	if err != nil {
		return nil, err
	}

	output := &LoadAllVQCConfigsOutput{
		VQCConfigs: make([]LoadVQCConfigOutput, len(configs)),
	}
	for i, config := range configs {
		output.VQCConfigs[i] = LoadVQCConfigOutput{
			Name:        config.Name(),
			Description: config.Description(),
			VQC:         config.VQC(),
			CreatedAt:   config.CreatedAt(),
			UpdatedAt:   config.UpdatedAt(),
			OwnerID:     config.OwnerID(),
			VQCConfigID: config.VQCConfigID(),
		}
	}
	return output, nil
}
