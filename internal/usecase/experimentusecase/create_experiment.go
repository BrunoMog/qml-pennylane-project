package experimentusecase

import (
	"pennylane_project_backend/internal/domain/experiment"
	"time"

	"uuid"
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
	name, err := experiment.NewName(input.Name)
	if err != nil {
		return nil, err
	}

	description, err := experiment.NewDescription(input.Description)
	if err != nil {
		return nil, err
	}

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

	experimentNameExists, err := s.experimentRepository.ExistsByName(input.CallerID, name)
	if err != nil {
		return nil, err
	}
	if experimentNameExists {
		return nil, &ExperimentNameAlreadyExistsError{}
	}

	experimentInput := experiment.ExperimentInput{
		Name:          name,
		Description:   description,
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
		Name:         experiment.Name().Value(),
		Description:  experiment.Description().Value(),
		ExperimentID: experiment.ExperimentID(),
		CreatedAt:    experiment.CreatedAt(),
	}

	return output, nil
}
