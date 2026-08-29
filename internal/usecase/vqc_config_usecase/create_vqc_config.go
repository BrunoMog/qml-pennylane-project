package vqc_config_usecase

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqc_config"
	"time"

	"github.com/google/uuid"
)

type CreateVQCConfigInput struct {
	CallerID    uuid.UUID
	Name        *string
	Description *string
	VQC         *vqc.VQC
}

type CreateVQCConfigOutput struct {
	Name        string
	Description string
	VQCId       uuid.UUID
	CreatedAt   time.Time
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

	newConfig, err := vqc_config.NewVQCConfig(input.CallerID, *input.Name, *input.Description, input.VQC)
	if err != nil {
		return nil, err
	}

	err = s.vqcConfigRepository.Save(newConfig)
	if err != nil {
		return nil, err
	}

	output := &CreateVQCConfigOutput{
		Name:        newConfig.GetName(),
		Description: newConfig.GetDescription(),
		VQCId:       newConfig.GetVQCConfigID(),
		CreatedAt:   newConfig.GetCreatedAt(),
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
