package vqcconfigusecase

import (
	"errors"
	"fmt"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"pennylane_project_backend/internal/usecase/apperrors"
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
		return ErrNoFieldsToUpdate
	}

	if input.CallerID == uuid.Nil() {
		return user.ErrNilUserID
	}

	if input.VQCConfigID == uuid.Nil() {
		return vqcconfig.ErrNilVQCConfigID
	}

	config, err := s.vqcConfigRepository.FindByID(input.VQCConfigID)
	if err != nil {
		if errors.Is(err, vqcconfig.ErrVQCConfigNotFound) {
			return err
		}

		return fmt.Errorf("vqcconfigusecase: find vqc config by ID: %w", err)
	}

	if !canUpdateVQCConfig(input.CallerID, config) {
		return apperrors.NewPermissionDeniedError(
			"user does not own the VQC config",
			"update VQC config",
			input.CallerID,
		)
	}

	var needToSave bool

	if input.Name != nil {
		name, err := vqcconfig.NewName(*input.Name)
		if err != nil {
			return err
		}

		if config.Name().String() != name.String() {
			if !name.Equals(config.Name()) {
				exists, err := s.vqcConfigRepository.ExistsByName(input.CallerID, name)
				if err != nil {
					return fmt.Errorf("vqcconfigusecase: check vqc config name existence: %w", err)
				}
				if exists {
					return ErrVQCConfigNameAlreadyExists
				}
			}
			err := config.SetName(name)
			if err != nil {
				return err
			}
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
		vqc, err := buildVQCFromDTO(*input.VQCDTO)
		if err != nil {
			return err
		}
		if !config.VQC().Equals(vqc) {
			needToSave = true
			err = config.SetVQC(vqc)
			if err != nil {
				return err
			}
		}
	}

	if needToSave {
		err = s.vqcConfigRepository.Save(config)
		if err != nil {
			return fmt.Errorf("vqcconfigusecase: save vqc config: %w", err)
		}
	}

	return nil
}

func canUpdateVQCConfig(callerID uuid.UUID, vqcConfig *vqcconfig.VQCConfig) bool {
	return callerID == vqcConfig.OwnerID()
}
