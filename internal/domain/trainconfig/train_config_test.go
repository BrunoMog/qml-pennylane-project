package trainconfig

import (
	"pennylane_project_backend/internal/domain/training"

	"testing"
	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validTraining(t *testing.T) training.Training {
	adamOpt, err := training.NewAdamOptimizer(0.001, 0.9, 0.999, 1e-8)
	require.NoError(t, err)
	inputEarlyStopping := training.EarlyStoppingInput{
		ValidationMetric: training.EvalMetricAccuracy,
		Patience:         5,
		MinDelta:         0.001,
		Enabled:          true,
	}
	earlyStopping, err := training.NewEarlyStopping(inputEarlyStopping)
	require.NoError(t, err)
	inputCrossValidation := training.CrossValidationInput{
		Enabled: true,
		Folds:   5,
	}
	crossValidation, err := training.NewCrossValidation(inputCrossValidation)
	require.NoError(t, err)

	input := training.TrainingInput{
		LearningTask:      training.LearningTaskBinaryClassification,
		CostFunction:      training.CostFunctionBinaryCrossEntropy,
		LearningType:      training.LearningTypeSupervised,
		EvaluationMetrics: []training.EvalMetric{training.EvalMetricAccuracy},
		TrainRatio:        0.7,
		ValidationRatio:   0.2,
		TestRatio:         0.1,
		RandomSeed:        42,
		MaxEpochs:         10,
		BatchSize:         32,
		Optimizer:         adamOpt,
		EarlyStopping:     earlyStopping,
		CrossValidation:   crossValidation,
	}
	trainingInstance, err := training.NewTraining(input)
	require.NoError(t, err)
	return trainingInstance
}

func TestNewTrainConfig(t *testing.T) {
	tests := []struct {
		expectError error
		setup       func() (uuid.UUID, Name, Description, training.Training)
		testName    string
	}{
		{
			testName: "valid TrainConfig",
			setup: func() (uuid.UUID, Name, Description, training.Training) {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				trainingInstance := validTraining(t)
				return userID, name, description, trainingInstance
			},
			expectError: nil,
		},
		{
			testName: "invalid TrainConfig with nil userID",
			setup: func() (uuid.UUID, Name, Description, training.Training) {
				userID := uuid.Nil()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				trainingInstance := validTraining(t)
				return userID, name, description, trainingInstance
			},
			expectError: &InvalidOwnerIDError{},
		},
		{
			testName: "invalid TrainConfig with invalid Training",
			setup: func() (uuid.UUID, Name, Description, training.Training) {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				invalidTraining := training.Training{}
				return userID, name, description, invalidTraining
			},
			expectError: &InvalidTrainingError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			userID, name, description, trainingInstance := tt.setup()
			trainConfig, err := NewTrainConfig(userID, name, description, trainingInstance)

			if tt.expectError != nil {
				assert.Error(t, err)
				assert.Nil(t, trainConfig)
				assert.IsType(t, tt.expectError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, trainConfig)
				assert.Equal(t, userID, trainConfig.OwnerID())
				assert.Equal(t, name, trainConfig.Name())
				assert.Equal(t, description, trainConfig.Description())
				assert.Equal(t, trainingInstance, trainConfig.Training())
			}
		})
	}
}

func TestSetTraining(t *testing.T) {
	tests := []struct {
		expectError error
		setup       func() *TrainConfig
		testName    string
		newTraining training.Training
	}{
		{
			testName: "valid SetTraining",
			setup: func() *TrainConfig {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				trainingInstance := validTraining(t)
				trainConfig, err := NewTrainConfig(userID, name, description, trainingInstance)
				require.NoError(t, err)
				return trainConfig
			},
			newTraining: validTraining(t),
			expectError: nil,
		},
		{
			testName: "invalid SetTraining with invalid Training",
			setup: func() *TrainConfig {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				trainingInstance := validTraining(t)
				trainConfig, err := NewTrainConfig(userID, name, description, trainingInstance)
				require.NoError(t, err)
				return trainConfig
			},
			newTraining: training.Training{},
			expectError: &InvalidTrainingError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			trainConfig := tt.setup()
			err := trainConfig.SetTraining(tt.newTraining)
			if tt.expectError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newTraining, trainConfig.Training())
			}
		})
	}
}
