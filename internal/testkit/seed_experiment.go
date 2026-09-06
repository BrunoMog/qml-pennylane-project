package testkit

import (
	"pennylane_project_backend/internal/domain/experiment"

	"github.com/google/uuid"
)

type ExperimentSeed struct {
	Name           string
	Description    string
	Ref            uint8
	CallerRef      uint8
	TrainConfigRef uint8
	VQCConfigRef   uint8
}

type ExperimentSeedResult struct {
	ByRef map[uint8]*experiment.Experiment
}

func SeedExperiments(experimentRepository *MockExperimentRepository,
	userSeedResult UserSeedResult, trainConfigSeedResult TrainConfigSeedResult,
	vqcConfigSeedResult VQCConfigSeedResult, seeds []ExperimentSeed) (ExperimentSeedResult, error) {
	res := ExperimentSeedResult{ByRef: map[uint8]*experiment.Experiment{}}
	for _, s := range seeds {
		var callerID uuid.UUID
		caller, ok := userSeedResult.ByRef[s.CallerRef]
		if !ok {
			callerID = uuid.New()
		} else {
			callerID = caller.ID()
		}
		var trainConfigID uuid.UUID
		trainConfig, ok := trainConfigSeedResult.ByRef[s.TrainConfigRef]
		if !ok {
			trainConfigID = uuid.New()
		} else {
			trainConfigID = trainConfig.TrainConfigID()
		}
		var vqcConfigID uuid.UUID
		vqcConfig, ok := vqcConfigSeedResult.ByRef[s.VQCConfigRef]
		if !ok {
			vqcConfigID = uuid.New()
		} else {
			vqcConfigID = vqcConfig.VQCConfigID()
		}
		input := experiment.ExperimentInput{
			OwnerID:       callerID,
			Name:          s.Name,
			Description:   s.Description,
			TrainConfigID: trainConfigID,
			VQCConfigID:   vqcConfigID,
		}
		newExperiment, err := experiment.NewExperiment(input)
		if err != nil {
			return ExperimentSeedResult{}, err
		}
		if err := experimentRepository.Save(newExperiment); err != nil {
			return ExperimentSeedResult{}, err
		}
		res.ByRef[s.Ref] = newExperiment
	}
	return res, nil
}
