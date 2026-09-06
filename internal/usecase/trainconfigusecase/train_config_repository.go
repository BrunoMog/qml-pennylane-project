package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"

	"github.com/google/uuid"
)

type TrainConfigRepository interface {
	Save(trainConfig *trainconfig.TrainConfig) error
	FindByID(trainConfigID uuid.UUID) (*trainconfig.TrainConfig, error)
	FindByName(ownerID uuid.UUID, name string) (*trainconfig.TrainConfig, error)
	ExistsByID(trainConfigID uuid.UUID) (bool, error)
	ExistsByName(ownerID uuid.UUID, name string) (bool, error)
	FindAllByOwnerID(ownerID uuid.UUID) ([]*trainconfig.TrainConfig, error)
	CheckOwnership(trainConfigID uuid.UUID, ownerID uuid.UUID) (bool, error)
	DeleteByID(trainConfigID uuid.UUID) error
	DeleteAllByOwnerID(ownerID uuid.UUID) error
}
