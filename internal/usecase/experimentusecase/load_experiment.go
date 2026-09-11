package experimentusecase

import (
	"pennylane_project_backend/internal/domain/experiment"
	"time"
	"uuid"
)

type LoadExperimentInput struct {
	ExperimentName *string
	ExperimentID   *uuid.UUID
	CallerID       uuid.UUID
}

type LoadExperimentOutput struct {
	ExperimentName        string
	ExperimentDescription string
	ExperimentID          uuid.UUID
	OwnerID               uuid.UUID
	TrainConfigID         uuid.UUID
	VQCConfigID           uuid.UUID
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (u *ExperimentService) LoadExperiment(input LoadExperimentInput) (*LoadExperimentOutput, error) {
	if input.ExperimentID == nil && input.ExperimentName == nil {
		return nil, &InvalidInputError{}
	}

	var exp *experiment.Experiment
	var err error
	if input.ExperimentID != nil {
		exp, err = u.experimentRepository.FindByID(*input.ExperimentID)
		if err != nil {
			return nil, err
		}
	} else if input.ExperimentName != nil {
		name, err := experiment.NewName(*input.ExperimentName)
		if err != nil {
			return nil, err
		}
		exp, err = u.experimentRepository.FindByName(input.CallerID, name)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, &InvalidInputError{}
	}

	if !canLoadExperiment(input.CallerID, exp) {
		return nil, &UnauthorizedError{}
	}

	output := &LoadExperimentOutput{
		ExperimentName:        exp.Name().Value(),
		ExperimentDescription: exp.Description().Value(),
		ExperimentID:          exp.ExperimentID(),
		OwnerID:               exp.OwnerID(),
		TrainConfigID:         exp.TrainConfigID(),
		VQCConfigID:           exp.VQCConfigID(),
		CreatedAt:             exp.CreatedAt(),
		UpdatedAt:             exp.UpdatedAt(),
	}

	return output, nil
}

func canLoadExperiment(callerID uuid.UUID, exp *experiment.Experiment) bool {
	return exp.OwnerID() == callerID
}

type LoadAllExperimentsInput struct {
	CallerID uuid.UUID
}

type LoadAllExperimentsOutput struct {
	Experiments []LoadExperimentOutput
}

func (u *ExperimentService) LoadAllExperiments(input LoadAllExperimentsInput) (*LoadAllExperimentsOutput, error) {
	experiments, err := u.experimentRepository.FindAllByOwnerID(input.CallerID)
	if err != nil {
		return nil, err
	}

	output := &LoadAllExperimentsOutput{
		Experiments: make([]LoadExperimentOutput, len(experiments)),
	}

	for i, exp := range experiments {
		output.Experiments[i] = LoadExperimentOutput{
			ExperimentName:        exp.Name().Value(),
			ExperimentDescription: exp.Description().Value(),
			ExperimentID:          exp.ExperimentID(),
			OwnerID:               exp.OwnerID(),
			TrainConfigID:         exp.TrainConfigID(),
			VQCConfigID:           exp.VQCConfigID(),
			CreatedAt:             exp.CreatedAt(),
			UpdatedAt:             exp.UpdatedAt(),
		}
	}

	return output, nil
}
