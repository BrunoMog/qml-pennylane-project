package vqcconfigusecase

import (
	"fmt"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"time"

	"uuid"
)

type CreateVQCConfigInput struct {
	Name        string
	Description string
	VQCDTO      VQCDTO
	CallerID    uuid.UUID
}

type CreateVQCConfigOutput struct {
	CreatedAt   time.Time
	Name        string
	Description string
	VQCConfigID uuid.UUID
}

func (s *VQCConfigService) CreateVQCConfig(input CreateVQCConfigInput) (*CreateVQCConfigOutput, error) {
	if input.CallerID == uuid.Nil() {
		return nil, user.ErrNilUserID
	}

	name, err := vqcconfig.NewName(input.Name)
	if err != nil {
		return nil, err
	}
	description, err := vqcconfig.NewDescription(input.Description)
	if err != nil {
		return nil, err
	}
	vqc, err := buildVQCFromDTO(input.VQCDTO)
	if err != nil {
		return nil, err
	}

	exists, err := s.userRepository.ExistsByID(input.CallerID)
	if err != nil {
		return nil, fmt.Errorf("vqcconfigusecase: check user existence: %w", err)
	}
	if !exists {
		return nil, user.ErrUserNotFound
	}
	exists, err = s.vqcConfigRepository.ExistsByName(input.CallerID, name)
	if err != nil {
		return nil, fmt.Errorf("vqcconfigusecase: check vqc config name existence: %w", err)
	}
	if exists {
		return nil, ErrVQCConfigNameAlreadyExists
	}

	newConfig, err := vqcconfig.NewVQCConfig(input.CallerID, name, description, vqc)
	if err != nil {
		return nil, err
	}

	err = s.vqcConfigRepository.Save(newConfig)
	if err != nil {
		return nil, fmt.Errorf("vqcconfigusecase: save vqc config: %w", err)
	}

	output := &CreateVQCConfigOutput{
		Name:        newConfig.Name().String(),
		Description: newConfig.Description().String(),
		VQCConfigID: newConfig.VQCConfigID(),
		CreatedAt:   newConfig.CreatedAt(),
	}

	return output, nil

}
