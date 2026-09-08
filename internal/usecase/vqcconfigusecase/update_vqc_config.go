package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqcconfig"

	"github.com/google/uuid"
)

type UpdateVQCConfigInput struct {
	Name        *string
	Description *string
	VQC         *VQCInputDTO
	CallerID    uuid.UUID
	VQCConfigID uuid.UUID
}

func (s *VQCConfigService) UpdateVQCConfig(input UpdateVQCConfigInput) error {
	if input.Name == nil && input.Description == nil && input.VQC == nil {
		return &NoFieldsToUpdateError{}
	}

	config, err := s.vqcConfigRepository.FindByID(input.VQCConfigID)
	if err != nil {
		return err
	}

	if !canUpdateVQCConfig(input.CallerID, config) {
		return &UnauthorizedError{}
	}

	if input.Name != nil {
		name, err := vqcconfig.NewName(*input.Name)
		if err != nil {
			return err
		}
		exists, err := s.vqcConfigRepository.ExistsByName(input.CallerID, name)
		if err != nil {
			return err
		}
		if exists {
			return &VQCConfigNameAlreadyExistsError{}
		}
		config.SetName(name)
	}
	if input.Description != nil {
		description, err := vqcconfig.NewDescription(*input.Description)
		if err != nil {
			return err
		}
		config.SetDescription(description)
	}
	if input.VQC != nil {
		vqc, err := buildVQC(*input.VQC)
		if err != nil {
			return err
		}
		config.SetVQC(vqc)
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
