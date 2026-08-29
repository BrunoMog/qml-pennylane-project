package vqc_config_usecase

import (
	"pennylane_project_backend/internal/domain/vqc_config"

	"github.com/google/uuid"
)

type VQCConfigRepository interface {
	Save(vqcConfig *vqc_config.VQCConfig) error
	FindByID(vqcConfigID uuid.UUID) (*vqc_config.VQCConfig, error)
	FindByName(vqcConfigName string) (*vqc_config.VQCConfig, error)
	ExistsByID(vqcConfigID uuid.UUID) (bool, error)
	ExistsByName(ownerID uuid.UUID, name string) (bool, error)
	FindAllByOwnerID(ownerID uuid.UUID) ([]*vqc_config.VQCConfig, error)
	DeleteByID(vqcConfigID uuid.UUID) error
	DeleteByName(vqcConfigName string) error
	DeleteAllByOwnerID(ownerID uuid.UUID) error
}
