package testkit

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"

	"github.com/google/uuid"
)

type VQCConfigSeed struct {
	VQC         *vqc.VQC
	Name        string
	Description string
	Ref         uint8
	CallerRef   uint8
}

type VQCConfigSeedResult struct {
	ByRef map[uint8]*vqcconfig.VQCConfig
}

func SeedVQCConfigs(vqcConfigRepository *MockVQCConfigRepository, userSeedResult UserSeedResult, seeds []VQCConfigSeed) (VQCConfigSeedResult, error) {
	res := VQCConfigSeedResult{ByRef: map[uint8]*vqcconfig.VQCConfig{}}
	for _, s := range seeds {
		var callerID uuid.UUID
		caller, ok := userSeedResult.ByRef[s.CallerRef]
		if !ok {
			callerID = uuid.New()
		} else {
			callerID = caller.ID()
		}
		newConfig, err := vqcconfig.NewVQCConfig(callerID, s.Name, s.Description, s.VQC)
		if err != nil {
			return VQCConfigSeedResult{}, err
		}
		if err := vqcConfigRepository.Save(newConfig); err != nil {
			return VQCConfigSeedResult{}, err
		}
		res.ByRef[s.Ref] = newConfig
	}
	return res, nil
}
