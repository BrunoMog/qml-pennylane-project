package testkit

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/training"

	"github.com/google/uuid"
)

type TrainConfigSeed struct {
	Train       *training.Training
	Name        string
	Description string
	Ref         uint8
	CallerRef   uint8
}

type TrainConfigSeedResult struct {
	ByRef map[uint8]*trainconfig.TrainConfig
}

func SeedTrainConfigs(trainConfigRepository *MockTrainConfigRepository, userSeedResult UserSeedResult, seeds []TrainConfigSeed) (TrainConfigSeedResult, error) {
	res := TrainConfigSeedResult{ByRef: map[uint8]*trainconfig.TrainConfig{}}
	for _, s := range seeds {
		var callerID uuid.UUID
		caller, ok := userSeedResult.ByRef[s.CallerRef]
		if !ok {
			callerID = uuid.New()
		} else {
			callerID = caller.ID()
		}
		newConfig, err := trainconfig.NewTrainConfig(callerID, s.Name, s.Description, s.Train)
		if err != nil {
			return TrainConfigSeedResult{}, err
		}
		if err := trainConfigRepository.Save(newConfig); err != nil {
			return TrainConfigSeedResult{}, err
		}
		res.ByRef[s.Ref] = newConfig
	}
	return res, nil
}
