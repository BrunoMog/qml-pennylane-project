package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"time"

	"github.com/google/uuid"
)

type GetVQCConfigInput struct {
	CallerID      uuid.UUID
	VQCConfigID   *uuid.UUID
	VQCConfigName *string
}

type GetVQCConfigOutput struct {
	Name        string
	Description string
	VQC         vqc.VQC
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *VQCConfigService) GetVQCConfig(input GetVQCConfigInput) (*GetVQCConfigOutput, error) {
	exists, err := s.userRepository.ExistByID(input.CallerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, &UserNotFoundError{}
	}

	var config *vqcconfig.VQCConfig
	if input.VQCConfigID != nil {
		config, err = s.vqcConfigRepository.FindByID(*input.VQCConfigID)
		if err != nil {
			return nil, err
		}
	} else if input.VQCConfigName != nil {
		config, err = s.vqcConfigRepository.FindByName(*input.VQCConfigName)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, &InvalidInputError{}
	}

	if !canGetVQCConfig(input.CallerID, config) {
		return nil, &UnauthorizedError{}
	}

	output := &GetVQCConfigOutput{
		Name:        config.GetName(),
		Description: config.GetDescription(),
		VQC:         config.GetVQC(),
		CreatedAt:   config.GetCreatedAt(),
		UpdatedAt:   config.GetUpdatedAt(),
	}
	return output, nil
}

func canGetVQCConfig(callerID uuid.UUID, config *vqcconfig.VQCConfig) bool {
	return callerID == config.GetOwnerID()
}

type ListVQCConfigsInput struct {
	CallerID uuid.UUID
}

func (s *VQCConfigService) ListVQCConfigs(input ListVQCConfigsInput) ([]*GetVQCConfigOutput, error) {
	exists, err := s.userRepository.ExistByID(input.CallerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, &UserNotFoundError{}
	}

	configs, err := s.vqcConfigRepository.FindAllByOwnerID(input.CallerID)
	if err != nil {
		return nil, err
	}

	outputs := make([]*GetVQCConfigOutput, len(configs))
	for i, config := range configs {
		outputs[i] = &GetVQCConfigOutput{
			Name:        config.GetName(),
			Description: config.GetDescription(),
			VQC:         config.GetVQC(),
			CreatedAt:   config.GetCreatedAt(),
			UpdatedAt:   config.GetUpdatedAt(),
		}
	}
	return outputs, nil
}
