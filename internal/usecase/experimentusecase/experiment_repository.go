package experimentusecase

import (
	"pennylane_project_backend/internal/domain/experiment"

	"uuid"
)

type ExperimentRepository interface {
	Save(experiment *experiment.Experiment) error
	FindByID(experimentID uuid.UUID) (*experiment.Experiment, error)
	FindByName(ownerID uuid.UUID, name experiment.Name) (*experiment.Experiment, error)
	ExistsByID(experimentID uuid.UUID) (bool, error)
	ExistsByName(ownerID uuid.UUID, name experiment.Name) (bool, error)
	FindAllByOwnerID(ownerID uuid.UUID) ([]*experiment.Experiment, error)
	CheckOwnership(ownerID, experimentID uuid.UUID) (bool, error)
	DeleteByID(experimentID uuid.UUID) error
	DeleteAllByOwnerID(ownerID uuid.UUID) error
}
