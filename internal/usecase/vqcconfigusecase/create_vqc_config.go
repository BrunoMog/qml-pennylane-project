package vqcconfigusecase

import (
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
		return nil, err
	}
	if !exists {
		return nil, &UserNotFoundError{}
	}
	exists, err = s.vqcConfigRepository.ExistsByName(input.CallerID, name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &VQCConfigNameAlreadyExistsError{Name: input.Name}
	}

	newConfig, err := vqcconfig.NewVQCConfig(input.CallerID, name, description, vqc)
	if err != nil {
		return nil, err
	}

	err = s.vqcConfigRepository.Save(newConfig)
	if err != nil {
		return nil, err
	}

	output := &CreateVQCConfigOutput{
		Name:        newConfig.Name().Value(),
		Description: newConfig.Description().Value(),
		VQCConfigID: newConfig.VQCConfigID(),
		CreatedAt:   newConfig.CreatedAt(),
	}
	return output, nil

}
