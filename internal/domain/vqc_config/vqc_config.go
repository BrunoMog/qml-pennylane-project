package vqc_config

import (
	"pennylane_project_backend/internal/domain/vqc"
	"time"

	"github.com/google/uuid"
)

type VQCConfig struct {
	userID      uuid.UUID
	vqcID       uuid.UUID
	name        string
	description string
	vqc         *vqc.VQC
	createdAt   time.Time
	updatedAt   time.Time
}

func NewVQCConfig(userID uuid.UUID, name string, description string, vqc *vqc.VQC) (*VQCConfig, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	err = validateDescription(description)
	if err != nil {
		return nil, err
	}

	return &VQCConfig{
		userID:      userID,
		vqcID:       uuid.New(),
		name:        name,
		description: description,
		vqc:         vqc,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func validateName(name string) error {
	if name == "" {
		return &InvalidNameError{name}
	}
	if len(name) > 20 {
		return &InvalidNameError{name}
	}
	return nil
}

func validateDescription(description string) error {
	if len(description) > 100 {
		return &InvalidDescriptionError{description}
	}
	return nil
}

func (vqcConfig *VQCConfig) UpdateName(name string) error {
	err := validateName(name)
	if err != nil {
		return err
	}
	vqcConfig.name = name
	vqcConfig.updatedAt = time.Now()
	return nil
}

func (vqcConfig *VQCConfig) UpdateDescription(description string) error {
	err := validateDescription(description)
	if err != nil {
		return err
	}
	vqcConfig.description = description
	vqcConfig.updatedAt = time.Now()
	return nil
}

func (vqcConfig *VQCConfig) UpdateVQC(vqc *vqc.VQC) {
	vqcConfig.vqc = vqc
	vqcConfig.updatedAt = time.Now()
}

func (vqcConfig VQCConfig) GetName() string {
	return vqcConfig.name
}

func (vqcConfig VQCConfig) GetDescription() string {
	return vqcConfig.description
}

func (vqcConfig VQCConfig) GetUserID() uuid.UUID {
	return vqcConfig.userID
}

func (vqcConfig VQCConfig) GetVQCID() uuid.UUID {
	return vqcConfig.vqcID
}

func (vqcConfig VQCConfig) GetCreatedAt() time.Time {
	return vqcConfig.createdAt
}

func (vqcConfig VQCConfig) GetUpdatedAt() time.Time {
	return vqcConfig.updatedAt
}

func (vqcConfig VQCConfig) GetVQC() vqc.VQC {
	return *vqcConfig.vqc
}
