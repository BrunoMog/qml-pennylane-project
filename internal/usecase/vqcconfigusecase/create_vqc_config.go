package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"time"

	"github.com/google/uuid"
)

type CreateVQCConfigInput struct {
	VQC         *vqc.VQC
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

	exists, err := s.userRepository.ExistsByID(input.CallerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, &UserNotFoundError{}
	}

	exists, err = s.vqcConfigRepository.ExistsByName(input.CallerID, input.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &VQCConfigNameAlreadyExistsError{Name: input.Name}
	}

	newConfig, err := vqcconfig.NewVQCConfig(input.CallerID, input.Name, input.Description, input.VQC)
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
