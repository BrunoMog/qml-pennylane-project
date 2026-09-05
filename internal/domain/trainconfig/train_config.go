package trainconfig

import (
	"pennylane_project_backend/internal/domain/training"
	"time"

	"github.com/google/uuid"
)

const (
	MaxNameLength        = 100
	MaxDescriptionLength = 500
)

type TrainConfig struct {
	createdAt     time.Time
	updatedAt     time.Time
	training      *training.Training
	name          string
	description   string
	ownerID       uuid.UUID
	trainConfigID uuid.UUID
}

func NewTrainConfig(ownerID uuid.UUID, name string, description string, tr *training.Training) (*TrainConfig, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	err = validateDescription(description)
	if err != nil {
		return nil, err
	}
	err = validateTraining(tr)
	if err != nil {
		return nil, err
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

func validateName(name string) error {
	if len(name) == 0 {
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

func validateTraining(tr *training.Training) error {
	if tr == nil {
		return &TrainingMissingError{}
	}
	return nil
}

func (tc *TrainConfig) SetName(name string) error {
	err := validateName(name)
	if err != nil {
		return err
	}
	tc.name = name
	tc.updatedAt = time.Now()
	return nil
}

func (tc *TrainConfig) SetDescription(description string) error {
	err := validateDescription(description)
	if err != nil {
		return err
	}
	tc.description = description
	tc.updatedAt = time.Now()
	return nil
}

func (tc *TrainConfig) SetTraining(tr *training.Training) error {
	err := validateTraining(tr)
	if err != nil {
		return err
	}
	tc.training = tr
	tc.updatedAt = time.Now()
	return nil
}

func (tc TrainConfig) Name() string {
	return tc.name
}

func (tc TrainConfig) Description() string {
	return tc.description
}

func (tc TrainConfig) Training() training.Training {
	return *tc.training
}

func (tc TrainConfig) OwnerID() uuid.UUID {
	return tc.ownerID
}

func (tc TrainConfig) TrainConfigID() uuid.UUID {
	return tc.trainConfigID
}

func (tc TrainConfig) CreatedAt() time.Time {
	return tc.createdAt
}

func (tc TrainConfig) UpdatedAt() time.Time {
	return tc.updatedAt
}
