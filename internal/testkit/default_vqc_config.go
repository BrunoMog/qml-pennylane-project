package testkit

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"strconv"

	"github.com/google/uuid"
)

func DefaultVQCConfig() func(ownerID uuid.UUID) *vqcconfig.VQCConfig {
	count := 0
	return func(ownerID uuid.UUID) *vqcconfig.VQCConfig {
		count++
		vqcConfig, err := vqcconfig.NewVQCConfig(
			ownerID,
			"Test VQC Config "+strconv.Itoa(count),
			"Test VQC Config Description "+strconv.Itoa(count),
			&vqc.VQC{},
		)
		if err != nil {
			panic(err)
		}
		return vqcConfig
	}
}
