package testkit

import (
	"pennylane_project_backend/internal/domain/vqcconfig"

	"github.com/google/uuid"
)

func DefaultVQCConfig() func(ownerID uuid.UUID) vqcconfig.VQCConfig {
	count := 0
	return func(ownerID uuid.UUID) vqcconfig.VQCConfig {
		count++
		vqcConfig, err := vqcconfig.NewVQCConfig(
			ownerID,
			"Test VQC Config "+string(rune(count)),
			"Test VQC Config Description "+string(rune(count)),
			nil,
		)
		if err != nil {
			panic(err)
		}
		return *vqcConfig
	}
}
