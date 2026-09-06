package testkit

import (
	"pennylane_project_backend/internal/domain/trainconfig"

	"github.com/google/uuid"
)

func DefaultTrainConfig() func(ownerID uuid.UUID) trainconfig.TrainConfig {
	count := 0
	return func(ownerID uuid.UUID) trainconfig.TrainConfig {
		count++
		trainConfig, err := trainconfig.NewTrainConfig(
			ownerID,
			"Test Train Config "+string(rune(count)),
			"Test Train Config Description "+string(rune(count)),
			nil,
		)
		if err != nil {
			panic(err)
		}
		return *trainConfig
	}
}
