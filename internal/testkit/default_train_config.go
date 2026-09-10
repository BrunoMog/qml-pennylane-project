package testkit

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/training"
	"strconv"

	"uuid"
)

func DefaultTrainConfig() func(ownerID uuid.UUID) *trainconfig.TrainConfig {
	count := 0
	return func(ownerID uuid.UUID) *trainconfig.TrainConfig {
		count++
		name, err := trainconfig.NewName("Test Train Config " + strconv.Itoa(count))
		if err != nil {
			panic(err)
		}
		description, err := trainconfig.NewDescription("Test Train Config Description " + strconv.Itoa(count))
		if err != nil {
			panic(err)
		}
		trainConfig, err := trainconfig.NewTrainConfig(
			ownerID,
			name,
			description,
			validTraining(count+2),
		)
		if err != nil {
			panic(err)
		}
		return trainConfig
	}
}

func validTraining(count int) training.Training {
	optimizer, err := training.NewAdamOptimizer(0.001, 0.9, 0.999, 1e-8)
	if err != nil {
		panic(err)
	}

	crossValidation, err := training.NewCrossValidation(training.CrossValidationInput{Enabled: true, Folds: 5})
	if err != nil {
		panic(err)
	}

	earlyStopping, err := training.NewEarlyStopping(training.EarlyStoppingInput{Enabled: true, ValidationMetric: "accuracy", Patience: 10, MinDelta: 0.001})
	if err != nil {
		panic(err)
	}

	input := training.TrainingInput{
		Optimizer:         optimizer,
		CrossValidation:   crossValidation,
		EarlyStopping:     earlyStopping,
		LearningTask:      training.LearningTaskBinaryClassification,
		CostFunction:      training.CostFunctionBinaryCrossEntropy,
		LearningType:      training.LearningTypeSupervised,
		EvaluationMetrics: []training.EvalMetric{training.EvalMetricAccuracy, training.EvalMetricPrecision},
		TrainRatio:        0.7,
		ValidationRatio:   0.15,
		TestRatio:         0.15,
		RandomSeed:        count,
		MaxEpochs:         100,
		BatchSize:         32,
	}

	trainingConfig, err := training.NewTraining(input)
	if err != nil {
		panic(err)
	}

	return trainingConfig
}
