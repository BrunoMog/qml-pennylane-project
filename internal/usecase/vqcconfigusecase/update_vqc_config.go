package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqcconfig"

	"uuid"
)

type UpdateVQCConfigInput struct {
	Name        *string
	Description *string
	VQCDTO      *VQCDTO
	CallerID    uuid.UUID
	VQCConfigID uuid.UUID
}

func (s *VQCConfigService) UpdateVQCConfig(input UpdateVQCConfigInput) error {
	if input.Name == nil && input.Description == nil && input.VQCDTO == nil {
		return &NoFieldsToUpdateError{}
	}

	config, err := s.vqcConfigRepository.FindByID(input.VQCConfigID)
	if err != nil {
		return err
	}

	if !canUpdateVQCConfig(input.CallerID, config) {
		return &UnauthorizedError{}
	}

	var needToSave bool

	if input.Name != nil {
		name, err := vqcconfig.NewName(*input.Name)
		if err != nil {
			return err
		}

		if config.Name().Value() != name.Value() {
			if !name.Equals(config.Name()) {
				exists, err := s.vqcConfigRepository.ExistsByName(input.CallerID, name)
				if err != nil {
					return err
				}
				if exists {
					return &VQCConfigNameAlreadyExistsError{}
				}
			}
			config.SetName(name)
			needToSave = true
		}
	}

	if input.Description != nil {
		description, err := vqcconfig.NewDescription(*input.Description)
		if err != nil {
			return err
		}
		if config.Description() != description {
			needToSave = true
			config.SetDescription(description)
		}
	}

	if input.VQCDTO != nil {
		vqc, err := buildVQCDTOToVQC(*input.VQCDTO)
		if err != nil {
			return err
		}
		if !config.VQC().Equals(vqc) {
			needToSave = true
			config.SetVQC(vqc)
		}
	}

	if needToSave {
		err = s.vqcConfigRepository.Save(config)
		if err != nil {
			return err
		}
	}

	return nil
}

func canUpdateVQCConfig(callerID uuid.UUID, vqcConfig *vqcconfig.VQCConfig) bool {
	return callerID == vqcConfig.OwnerID()
}
