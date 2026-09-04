package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"time"

	"github.com/google/uuid"
)

type CreateVQCConfigInput struct {
	Name        *string
	Description *string
	VQC         *vqc.VQC
	CallerID    uuid.UUID
}

type CreateVQCConfigOutput struct {
	CreatedAt   time.Time
	Name        string
	Description string
	VQCId       uuid.UUID
}

func (s *VQCConfigService) CreateVQCConfig(input CreateVQCConfigInput) (*CreateVQCConfigOutput, error) {
	err := validateName(input.Name)
	if err != nil {
		return nil, err
	}

	exists, err := s.userRepository.ExistByID(input.CallerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, &UserNotFoundError{}
	}

	exists, err = s.vqcConfigRepository.ExistsByName(input.CallerID, *input.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &VQCConfigNameAlreadyExistsError{Name: *input.Name}
	}

	newConfig, err := vqcconfig.NewVQCConfig(input.CallerID, *input.Name, *input.Description, input.VQC)
	if err != nil {
		return nil, err
	}

	err = s.vqcConfigRepository.Save(newConfig)
	if err != nil {
		return nil, err
	}

	output := &CreateVQCConfigOutput{
		Name:        newConfig.Name(),
		Description: newConfig.Description(),
		VQCId:       newConfig.VQCConfigID(),
		CreatedAt:   newConfig.CreatedAt(),
	}
	return output, nil

}

func validateName(name *string) error {
	if name == nil {
		return &InvalidNameError{}
	}
	if len(*name) > 20 {
		return &InvalidNameError{}
	}
	return nil
}

func validateDescription(description *string) error {
	if description != nil && len(*description) > 100 {
		return &InvalidDescriptionError{}
	}
	return nil
}
