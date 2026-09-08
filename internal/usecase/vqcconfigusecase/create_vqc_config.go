package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqcconfig"
	"time"

	"github.com/google/uuid"
)

type CreateVQCConfigInput struct {
	VQC         VQCInputDTO
	Name        string
	Description string
	CallerID    uuid.UUID
}

type CreateVQCConfigOutput struct {
	CreatedAt   time.Time
	Name        string
	Description string
	VQCId       uuid.UUID
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
	vqc, err := BuildVQC(input.VQC)
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
		VQCId:       newConfig.VQCConfigID(),
		CreatedAt:   newConfig.CreatedAt(),
	}
	return output, nil

}
