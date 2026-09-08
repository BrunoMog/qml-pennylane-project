package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqcconfig"

	"github.com/google/uuid"
)

type VQCConfigRepository interface {
	Save(vqcConfig *vqcconfig.VQCConfig) error
	FindByID(vqcConfigID uuid.UUID) (*vqcconfig.VQCConfig, error)
	FindByName(ownerID uuid.UUID, name vqcconfig.Name) (*vqcconfig.VQCConfig, error)
	ExistsByID(vqcConfigID uuid.UUID) (bool, error)
	ExistsByName(ownerID uuid.UUID, name vqcconfig.Name) (bool, error)
	FindAllByOwnerID(ownerID uuid.UUID) ([]*vqcconfig.VQCConfig, error)
	CheckOwnership(ownerID uuid.UUID, vqcConfigID uuid.UUID) (bool, error)
	DeleteByID(vqcConfigID uuid.UUID) error
	DeleteAllByOwnerID(ownerID uuid.UUID) error
}
