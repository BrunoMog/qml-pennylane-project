package vqc_config_usecase

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqc_config"

	"github.com/google/uuid"
)

type UpdateVQCConfigInput struct {
	CallerID    uuid.UUID
	VQCConfigID uuid.UUID
	Name        *string
	Description *string
	VQC         *vqc.VQC
}

func (s *VQCConfigService) UpdateVQCConfig(input UpdateVQCConfigInput) error {
	exists, err := s.userRepository.ExistByID(input.CallerID)
	if err != nil {
		return err
	}
	if !exists {
		return &UserNotFoundError{}
	}

	config, err := s.vqcConfigRepository.FindByID(input.VQCConfigID)
	if err != nil {
		return err
	}

	if !canUpdateVQCConfig(input.CallerID, config) {
		return &UnauthorizedError{}
	}

	if input.Name != nil {
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
	if input.VQC != nil {
		err = config.SetVQC(input.VQC)
		if err != nil {
			return err
		}
	}

	err = s.vqcConfigRepository.Save(config)
	if err != nil {
		return err
	}

	return nil
}

func canUpdateVQCConfig(callerID uuid.UUID, vqcConfig *vqc_config.VQCConfig) bool {
	return callerID == vqcConfig.GetOwnerID()
}
