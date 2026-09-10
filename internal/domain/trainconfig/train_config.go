package trainconfig

import (
	"pennylane_project_backend/internal/domain/training"
	"time"

	"uuid"
)

type TrainConfig struct {
	createdAt     time.Time
	updatedAt     time.Time
	name          Name
	description   Description
	training      training.Training
	ownerID       uuid.UUID
	trainConfigID uuid.UUID
}

func NewTrainConfig(ownerID uuid.UUID, name Name, description Description, tr training.Training) (*TrainConfig, error) {
	if !tr.IsValid() {
		return nil, &InvalidTrainingError{}
	}
	if ownerID == uuid.Nil() {
		return nil, &InvalidOwnerIDError{}
	}

	return &TrainConfig{
		ownerID:       ownerID,
		trainConfigID: uuid.New(),
		name:          name,
		description:   description,
		training:      tr,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}, nil
}

func (tc *TrainConfig) SetName(name Name) {
	tc.name = name
	tc.updatedAt = time.Now()
}

func (tc *TrainConfig) SetDescription(description Description) {
	tc.description = description
	tc.updatedAt = time.Now()
}

func (tc *TrainConfig) SetTraining(tr training.Training) error {
	if !tr.IsValid() {
		return &InvalidTrainingError{}
	}
	tc.training = tr
	tc.updatedAt = time.Now()
	return nil
}

func (tc *TrainConfig) Name() Name {
	return tc.name
}

func (tc *TrainConfig) Description() Description {
	return tc.description
}

func (tc *TrainConfig) Training() training.Training {
	return tc.training
}

func (tc *TrainConfig) OwnerID() uuid.UUID {
	return tc.ownerID
}

func (tc *TrainConfig) TrainConfigID() uuid.UUID {
	return tc.trainConfigID
}

func (tc *TrainConfig) CreatedAt() time.Time {
	return tc.createdAt
}

func (tc *TrainConfig) UpdatedAt() time.Time {
	return tc.updatedAt
}
