package experiment

import (
	"time"

	"uuid"
)

type Experiment struct {
	createdAt     time.Time
	updatedAt     time.Time
	name          Name
	description   Description
	experimentID  uuid.UUID
	ownerID       uuid.UUID
	trainConfigID uuid.UUID
	vQCConfigID   uuid.UUID
}

type ExperimentInput struct {
	Name          Name
	Description   Description
	OwnerID       uuid.UUID
	TrainConfigID uuid.UUID
	VQCConfigID   uuid.UUID
}

func NewExperiment(input ExperimentInput) (*Experiment, error) {
	if input.OwnerID == uuid.Nil() {
		return nil, &InvalidOwnerIDError{}
	}
	if input.TrainConfigID == uuid.Nil() {
		return nil, &InvalidTrainConfigIDError{}
	}
	if input.VQCConfigID == uuid.Nil() {
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

func (e *Experiment) SetName(newName Name) {
	e.name = newName
	e.updatedAt = time.Now()
}

func (e *Experiment) SetDescription(newDescription Description) {
	e.description = newDescription
	e.updatedAt = time.Now()
}

func (e *Experiment) SetTrainConfigID(newTrainConfigID uuid.UUID) {
	e.trainConfigID = newTrainConfigID
	e.updatedAt = time.Now()
}

func (e *Experiment) SetVQCConfigID(newVQCConfigID uuid.UUID) {
	e.vQCConfigID = newVQCConfigID
	e.updatedAt = time.Now()
}

func (e *Experiment) Name() Name {
	return e.name
}

func (e *Experiment) Description() Description {
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
