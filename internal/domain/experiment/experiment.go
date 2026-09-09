package experiment

import (
	"time"

	"github.com/google/uuid"
)

const (
	MaxNameLength        = 50
	MaxDescriptionLength = 200
)

type Experiment struct {
	createdAt     time.Time
	updatedAt     time.Time
	name          string
	description   string
	experimentID  uuid.UUID
	ownerID       uuid.UUID
	trainConfigID uuid.UUID
	vQCConfigID   uuid.UUID
}

type ExperimentInput struct {
	Name          string
	Description   string
	OwnerID       uuid.UUID
	TrainConfigID uuid.UUID
	VQCConfigID   uuid.UUID
}

func NewExperiment(input ExperimentInput) (*Experiment, error) {
	if err := validateName(input.Name); err != nil {
		return nil, err
	}
	if err := validateDescription(input.Description); err != nil {
		return nil, err
	}
	if input.OwnerID == uuid.Nil {
		return nil, &InvalidOwnerIDError{}
	}
	if input.TrainConfigID == uuid.Nil {
		return nil, &InvalidTrainConfigIDError{}
	}
	if input.VQCConfigID == uuid.Nil {
		return nil, &InvalidVQCConfigIDError{}
	}

	return &Experiment{
		experimentID:  uuid.New(),
		ownerID:       input.OwnerID,
		trainConfigID: input.TrainConfigID,
		vQCConfigID:   input.VQCConfigID,
		name:          input.Name,
		description:   input.Description,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}, nil
}

func validateName(name string) error {
	if name == "" {
		return &InvalidNameError{name}
	}
	if len(name) > MaxNameLength {
		return &InvalidNameError{name}
	}
	return nil
}

func validateDescription(description string) error {
	if len(description) > MaxDescriptionLength {
		return &InvalidDescriptionError{description}
	}
	return nil
}

func (e *Experiment) SetName(newName string) error {
	if err := validateName(newName); err != nil {
		return err
	}
	e.name = newName
	e.updatedAt = time.Now()
	return nil
}

func (e *Experiment) SetDescription(newDescription string) error {
	if err := validateDescription(newDescription); err != nil {
		return err
	}
	e.description = newDescription
	e.updatedAt = time.Now()
	return nil
}

func (e *Experiment) SetTrainConfigID(newTrainConfigID uuid.UUID) {
	e.trainConfigID = newTrainConfigID
	e.updatedAt = time.Now()
}

func (e *Experiment) SetVQCConfigID(newVQCConfigID uuid.UUID) {
	e.vQCConfigID = newVQCConfigID
	e.updatedAt = time.Now()
}

func (e *Experiment) Name() string {
	return e.name
}

func (e *Experiment) Description() string {
	return e.description
}

func (e *Experiment) ExperimentID() uuid.UUID {
	return e.experimentID
}

func (e *Experiment) OwnerID() uuid.UUID {
	return e.ownerID
}

func (e *Experiment) TrainConfigID() uuid.UUID {
	return e.trainConfigID
}

func (e *Experiment) VQCConfigID() uuid.UUID {
	return e.vQCConfigID
}

func (e *Experiment) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Experiment) UpdatedAt() time.Time {
	return e.updatedAt
}
