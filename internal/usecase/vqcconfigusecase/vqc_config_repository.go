package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqcconfig"

	"github.com/google/uuid"
)

type VQCConfigRepository interface {
	Save(vqcConfig *vqcconfig.VQCConfig) error
	FindByID(vqcConfigID uuid.UUID) (*vqcconfig.VQCConfig, error)
	FindByName(vqcConfigName string) (*vqcconfig.VQCConfig, error)
	ExistsByID(vqcConfigID uuid.UUID) (bool, error)
	ExistsByName(ownerID uuid.UUID, name string) (bool, error)
	FindAllByOwnerID(ownerID uuid.UUID) ([]*vqcconfig.VQCConfig, error)
	DeleteByID(vqcConfigID uuid.UUID) error
	DeleteByName(vqcConfigName string) error
	DeleteAllByOwnerID(ownerID uuid.UUID) error
}
