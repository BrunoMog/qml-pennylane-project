package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqcconfig"

	"github.com/google/uuid"
)

type DeleteVQCConfigInput struct {
	CallerID    uuid.UUID
	VQCConfigID uuid.UUID
}

func (s *VQCConfigService) DeleteVQCConfig(input DeleteVQCConfigInput) error {
	exists, err := s.userRepository.ExistByID(input.CallerID)
	if err != nil {
		return err
	}
	if !exists {
		return &UserNotFoundError{}
	}

	vqcConfig, err := s.vqcConfigRepository.FindByID(input.VQCConfigID)
	if err != nil {
		return err
	}

	if !canDeleteVQCConfig(input.CallerID, vqcConfig) {
		return &UnauthorizedError{}
	}

	err = s.vqcConfigRepository.DeleteByID(input.VQCConfigID)
	if err != nil {
		return err
	}

	return nil
}

func canDeleteVQCConfig(callerID uuid.UUID, vqcConfig *vqcconfig.VQCConfig) bool {
	return callerID == vqcConfig.OwnerID()
}
