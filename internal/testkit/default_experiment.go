package testkit

import (
	"pennylane_project_backend/internal/domain/experiment"
	"strconv"
	"uuid"
)

func DefaultExperiment() func(ownerID uuid.UUID, vqcConfigID uuid.UUID, trainConfigID uuid.UUID) *experiment.Experiment {
	count := 0
	return func(ownerID uuid.UUID, vqcConfigID uuid.UUID, trainConfigID uuid.UUID) *experiment.Experiment {
		count++
		name, err := experiment.NewName("Test Experiment " + strconv.Itoa(count))
		if err != nil {
			panic(err)
		}
		description, err := experiment.NewDescription("Test Experiment Description " + strconv.Itoa(count))
		if err != nil {
			panic(err)
		}
		input := experiment.ExperimentInput{
			OwnerID:       ownerID,
			VQCConfigID:   vqcConfigID,
			TrainConfigID: trainConfigID,
			Name:          name,
			Description:   description,
		}
		exp, err := experiment.NewExperiment(input)
		if err != nil {
			panic(err)
		}
		return exp
	}
}
