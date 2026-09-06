package vqcconfigusecase

import (
	"github.com/google/uuid"
)

type DeleteVQCConfigInput struct {
	CallerID    uuid.UUID
	VQCConfigID uuid.UUID
}

func (s *VQCConfigService) DeleteVQCConfig(input DeleteVQCConfigInput) error {
	checkOwnership, err := s.vqcConfigRepository.CheckOwnership(input.CallerID, input.VQCConfigID)
	if err != nil {
		return err
	}

	if !checkOwnership {
		return &UnauthorizedError{}
	}

	err = s.vqcConfigRepository.DeleteByID(input.VQCConfigID)
	if err != nil {
		return err
	}

	return nil
}
