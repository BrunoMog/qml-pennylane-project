package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/training"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"uuid"

	"github.com/stretchr/testify/require"
)

type testFixture struct {
	t               *testing.T
	service         *TrainConfigService
	userRepo        *testkit.MockUserRepository
	trainConfigRepo *testkit.MockTrainConfigRepository
	makeUser        func() *user.User
	makeTrainConfig func(ownerID uuid.UUID) *trainconfig.TrainConfig
}

func newTestFixture(t *testing.T) *testFixture {
	t.Helper()
	userRepo := testkit.NewMockUserRepository()
	trainConfigRepo := testkit.NewMockTrainConfigRepository()
	service := NewTrainConfigService(trainConfigRepo, userRepo)
	makeUser := testkit.DefaultUser()
	makeTrainConfig := testkit.DefaultTrainConfig()

	return &testFixture{
		t:               t,
		service:         service,
		userRepo:        userRepo,
		trainConfigRepo: trainConfigRepo,
		makeUser:        makeUser,
		makeTrainConfig: makeTrainConfig,
	}
}

func (f *testFixture) createUser(role user.Role) *user.User {
	user := f.makeUser()
	user.SetRole(role)
	err := f.userRepo.Save(user)
	require.NoError(f.t, err)
	return user
}

func (f *testFixture) createTrainConfig(ownerID uuid.UUID) *trainconfig.TrainConfig {
	trainConfig := f.makeTrainConfig(ownerID)
	err := f.trainConfigRepo.Save(trainConfig)
	require.NoError(f.t, err)
	return trainConfig
}

func validTrainDTO() TrainingDTO {
	optimizerDTO := OptimizerDTO{
		OptimizerName: "adam",
		LearningRate:  0.001,
		Beta1:         0.9,
		Beta2:         0.999,
		Epsilon:       1e-8,
	}

	crossValidationDTO := CrossValidationDTO{
		Enabled: true,
		Folds:   5,
	}

	earlyStoppingDTO := EarlyStoppingDTO{
		Enabled:          true,
		ValidationMetric: "accuracy",
		Patience:         10,
		MinDelta:         0.001,
	}

	return TrainingDTO{
		Optimizer:         optimizerDTO,
		CrossValidation:   crossValidationDTO,
		EarlyStopping:     earlyStoppingDTO,
		LearningTask:      "binary_classification",
		CostFunction:      "binary_cross_entropy",
		LearningType:      "supervised",
		EvaluationMetrics: []string{"accuracy", "precision"},
		TrainRatio:        0.7,
		ValidationRatio:   0.15,
		TestRatio:         0.15,
		RandomSeed:        42,
		MaxEpochs:         100,
		BatchSize:         32,
	}
}

func validTrain() training.Training {
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
		RandomSeed:        42,
		MaxEpochs:         100,
		BatchSize:         32,
	}

	trainingConfig, err := training.NewTraining(input)
	if err != nil {
		panic(err)
	}

	return trainingConfig
}
