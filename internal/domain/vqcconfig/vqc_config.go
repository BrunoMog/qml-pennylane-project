package vqcconfig

import (
	"pennylane_project_backend/internal/domain/vqc"
	"time"

	"github.com/google/uuid"
)

const (
	MaxNameLength        = 100
	MaxDescriptionLength = 500
)

type VQCConfig struct {
	createdAt   time.Time
	updatedAt   time.Time
	vqc         *vqc.VQC
	name        string
	description string
	userID      uuid.UUID
	vqcID       uuid.UUID
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
	err = validateVQC(vqc)
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
	if len(name) > MaxNameLength {
		return &InvalidNameError{name}
	}
	return nil
}

func validateDescription(description string) error {
	if len(description) > MaxDescriptionLength {
		return &InvalidDescriptionError{description}
	}
	return nil
}

func validateVQC(vqc *vqc.VQC) error {
	if vqc == nil {
		return &VQCConfigMissingVQCError{}
	}
	return nil
}

func (vqcConfig *VQCConfig) SetName(name string) error {
	err := validateName(name)
	if err != nil {
		return err
	}
	vqcConfig.name = name
	vqcConfig.updatedAt = time.Now()
	return nil
}

func (vqcConfig *VQCConfig) SetDescription(description string) error {
	err := validateDescription(description)
	if err != nil {
		return err
	}
	vqcConfig.description = description
	vqcConfig.updatedAt = time.Now()
	return nil
}

func (vqcConfig *VQCConfig) SetVQC(vqc *vqc.VQC) error {
	err := validateVQC(vqc)
	if err != nil {
		return err
	}
	vqcConfig.vqc = vqc
	vqcConfig.updatedAt = time.Now()
	return nil
}

func (vqcConfig VQCConfig) GetName() string {
	return vqcConfig.name
}

func (vqcConfig VQCConfig) GetDescription() string {
	return vqcConfig.description
}

func (vqcConfig VQCConfig) GetOwnerID() uuid.UUID {
	return vqcConfig.userID
}

func (vqcConfig VQCConfig) GetVQCConfigID() uuid.UUID {
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
