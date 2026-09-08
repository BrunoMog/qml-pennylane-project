package experimentusecase

import (
	"pennylane_project_backend/internal/domain/experiment"

	"github.com/google/uuid"
)

type ExperimentRepository interface {
	Save(experiment *experiment.Experiment) error
	FindByID(experimentID uuid.UUID) (*experiment.Experiment, error)
	FindByName(ownerID uuid.UUID, name string) (*experiment.Experiment, error)
	ExistsByID(experimentID uuid.UUID) (bool, error)
	ExistsByName(ownerID uuid.UUID, name string) (bool, error)
	FindAllByOwnerID(ownerID uuid.UUID) ([]*experiment.Experiment, error)
	DeleteByID(experimentID uuid.UUID) error
	DeleteAllByOwnerID(ownerID uuid.UUID) error
}
