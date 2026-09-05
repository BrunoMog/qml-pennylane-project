package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"

	"github.com/google/uuid"
)

type TrainConfigRepository interface {
	Save(trainConfig *trainconfig.TrainConfig) error
	FindByID(trainConfigID uuid.UUID) (*trainconfig.TrainConfig, error)
	FindByName(trainConfigName string) (*trainconfig.TrainConfig, error)
	ExistsByID(trainConfigID uuid.UUID) (bool, error)
	ExistsByName(ownerID uuid.UUID, name string) (bool, error)
	FindAllByOwnerID(ownerID uuid.UUID) ([]*trainconfig.TrainConfig, error)
	DeleteByID(trainConfigID uuid.UUID) error
	DeleteByName(trainConfigName string) error
	DeleteAllByOwnerID(ownerID uuid.UUID) error
}
