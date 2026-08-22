package vqc

import (
	"pennylane_project_backend/internal/user"
	"time"
)

type VQCConfig struct {
	name        string
	description string
	vqc         *VQC
	user        *user.User
	createdAt   time.Time
	updatedAt   time.Time
}

func NewVQCConfig(name string, description string, vqc *VQC, user *user.User) (*VQCConfig, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	err = validateDescription(description)
	if err != nil {
		return nil, err
	}

	return &VQCConfig{
		name:        name,
		description: description,
		vqc:         vqc,
		user:        user,
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
