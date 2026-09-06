package testkit

import (
	"pennylane_project_backend/internal/domain/experiment"

	"github.com/google/uuid"
)

func DefaultExperiment() func(ownerID uuid.UUID, vqcConfigID uuid.UUID, trainConfigID uuid.UUID) experiment.Experiment {
	count := 0
	return func(ownerID uuid.UUID, vqcConfigID uuid.UUID, trainConfigID uuid.UUID) experiment.Experiment {
		count++
		input := experiment.ExperimentInput{
			OwnerID:       ownerID,
			VQCConfigID:   vqcConfigID,
			TrainConfigID: trainConfigID,
			Name:          "Test Experiment " + string(rune(count)),
			Description:   "Test Experiment Description " + string(rune(count)),
		}
		exp, err := experiment.NewExperiment(input)
		if err != nil {
			panic(err)
		}
		return *exp
	}
}
