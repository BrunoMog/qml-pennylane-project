package vqc_config_usecase

import (
	"pennylane_project_backend/internal/domain/vqc_config"

	"github.com/google/uuid"
)

type VQCConfigRepository interface {
	Save(config *vqc_config.VQCConfig) error
	FindByID(id uuid.UUID) (*vqc_config.VQCConfig, error)
	FindByName(name string) (*vqc_config.VQCConfig, error)
	FindAll() ([]*vqc_config.VQCConfig, error)
	DeleteByID(id uuid.UUID) error
	DeleteByName(name string) error
	DeleteAll() error
}
