package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"

	"github.com/google/uuid"
)

type UpdateVQCConfigInput struct {
	Name        *string
	Description *string
	VQC         *vqc.VQC
	CallerID    uuid.UUID
	VQCConfigID uuid.UUID
}

func (s *VQCConfigService) UpdateVQCConfig(input UpdateVQCConfigInput) error {
	exists, err := s.userRepository.ExistsByID(input.CallerID)
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
		exists, err := s.vqcConfigRepository.ExistsByName(input.CallerID, *input.Name)
		if err != nil {
			return err
		}
		if exists {
			return &VQCConfigNameAlreadyExistsError{}
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

func canUpdateVQCConfig(callerID uuid.UUID, vqcConfig *vqcconfig.VQCConfig) bool {
	return callerID == vqcConfig.OwnerID()
}
