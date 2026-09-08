package experimentusecase

import (
	"pennylane_project_backend/internal/domain/experiment"
	"time"

	"github.com/google/uuid"
)

type CreateExperimentInput struct {
	Name          string
	Description   string
	CallerID      uuid.UUID
	TrainConfigID uuid.UUID
	VQCConfigID   uuid.UUID
}

type CreateExperimentOutput struct {
	CreatedAt    time.Time
	Name         string
	Description  string
	ExperimentID uuid.UUID
}

func (s *ExperimentService) CreateExperiment(input CreateExperimentInput) (*CreateExperimentOutput, error) {
	userExists, err := s.userRepository.ExistsByID(input.CallerID)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, &UserNotFoundError{}
	}

	checkTrainConfigOwner, err := s.trainConfigRepository.CheckOwnership(input.CallerID, input.TrainConfigID)
	if err != nil {
		return nil, err
	}
	if !checkTrainConfigOwner {
		return nil, &UnauthorizedError{}
	}

	checkVQCConfigOwner, err := s.vqcConfigRepository.CheckOwnership(input.CallerID, input.VQCConfigID)
	if err != nil {
		return nil, err
	}
	if !checkVQCConfigOwner {
		return nil, &UnauthorizedError{}
	}

	experimentInput := experiment.ExperimentInput{
		Name:          input.Name,
		Description:   input.Description,
		OwnerID:       input.CallerID,
		TrainConfigID: input.TrainConfigID,
		VQCConfigID:   input.VQCConfigID,
	}

	experiment, err := experiment.NewExperiment(experimentInput)
	if err != nil {
		return nil, err
	}

	err = s.experimentRepository.Save(experiment)
	if err != nil {
		return nil, err
	}

	output := &CreateExperimentOutput{
		Name:         experiment.Name(),
		Description:  experiment.Description(),
		ExperimentID: experiment.ExperimentID(),
		CreatedAt:    experiment.CreatedAt(),
	}

	return output, nil
}
