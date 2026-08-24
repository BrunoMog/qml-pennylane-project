package vqc_config

import (
	"pennylane_project_backend/internal/domain/user"
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

func NewVQCConfig(name string, description string, vqc *vqc.VQC, user *user.User) error {
	err := validateName(name)
	if err != nil {
		return err
	}
	err = validateDescription(description)
	if err != nil {
		return err
	}

	return nil
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
