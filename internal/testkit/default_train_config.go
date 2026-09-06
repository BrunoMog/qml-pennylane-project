package testkit

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/training"
	"strconv"

	"github.com/google/uuid"
)

func DefaultTrainConfig() func(ownerID uuid.UUID) *trainconfig.TrainConfig {
	count := 0
	return func(ownerID uuid.UUID) *trainconfig.TrainConfig {
		count++
		trainConfig, err := trainconfig.NewTrainConfig(
			ownerID,
			"Test Train Config "+strconv.Itoa(count),
			"Test Train Config Description "+strconv.Itoa(count),
			&training.Training{},
		)
		if err != nil {
			panic(err)
		}
		return trainConfig
	}
}
