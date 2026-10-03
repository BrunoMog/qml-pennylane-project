package vqcconfigusecase

import (
	"errors"
	"fmt"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"pennylane_project_backend/internal/usecase/apperrors"
	"uuid"
)

type DeleteVQCConfigInput struct {
	CallerID    uuid.UUID
	VQCConfigID uuid.UUID
}

func (s *VQCConfigService) DeleteVQCConfig(input DeleteVQCConfigInput) error {
	if input.CallerID == uuid.Nil() {
		return user.ErrNilUserID
	}

	if input.VQCConfigID == uuid.Nil() {
		return vqcconfig.ErrNilVQCConfigID
	}

	checkOwnership, err := s.vqcConfigRepository.CheckOwnership(input.CallerID, input.VQCConfigID)
	if err != nil {
		if errors.Is(err, vqcconfig.ErrVQCConfigNotFound) {
			return err
		}
		return fmt.Errorf("vqcconfigusecase: check ownership: %w", err)
	}

	if !checkOwnership {
		return apperrors.NewPermissionDeniedError(
			"user does not own the VQC config",
			"delete VQC config",
			input.CallerID,
		)
	}

	err = s.vqcConfigRepository.DeleteByID(input.VQCConfigID)
	if err != nil {
		return fmt.Errorf("vqcconfigusecase: delete vqc config by ID: %w", err)
	}

	return nil
}
