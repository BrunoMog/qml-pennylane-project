package training

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validTrainingInput(t *testing.T) TrainingInput {
	t.Helper()

	optimizer, err := NewAdamOptimizer(0.001, 0.9, 0.999, 1e-8)
	require.NoError(t, err, "failed to create valid optimizer")

	return TrainingInput{
		Optimizer:         optimizer,
		LearningTask:      LearningTaskBinaryClassification,
		CostFunction:      CostFunctionBinaryCrossEntropy,
		LearningType:      LearningTypeSupervised,
		EvaluationMetrics: []EvalMetric{EvalMetricAccuracy, EvalMetricF1Score},
		EarlyStopping: EarlyStopping{
			enabled:          true,
			patience:         5,
			minDelta:         0.01,
			validationMetric: EvalMetricAccuracy,
		},
		CrossValidation: CrossValidation{enabled: true, folds: 5},
		TrainRatio:      0.7,
		ValidationRatio: 0.2,
		TestRatio:       0.1,
		RandomSeed:      42,
		MaxEpochs:       100,
		BatchSize:       32,
	}
}

func TestNewTraining(t *testing.T) {
	tests := []struct {
		expectedError error
		setupInput    func(t *testing.T) TrainingInput
		testName      string
	}{
		{
			testName: "Valid input",
			setupInput: func(t *testing.T) TrainingInput {
				return validTrainingInput(t)
			},
			expectedError: nil,
		},
		{
			testName: "Invalid learning type",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.LearningType = LearningType("invalid_learning_type")
				return input
			},
			expectedError: &InvalidLearningTypeError{},
		},
		{
			testName: "Invalid learning task",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.LearningTask = LearningTask("invalid_learning_task")
				return input
			},
			expectedError: &InvalidLearningTaskError{},
		},
		{
			testName: "Invalid cost function",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.CostFunction = CostFunction("invalid_cost_function")
				return input
			},
			expectedError: &InvalidCostFunctionError{},
		},
		{
			testName: "Zero eval metrics",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.EvaluationMetrics = []EvalMetric{}
				return input
			},
			expectedError: &InvalidEvalMetricError{},
		},
		{
			testName: "Invalid eval metric",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.EvaluationMetrics = []EvalMetric{EvalMetric("invalid_eval_metric")}
				return input
			},
			expectedError: &InvalidEvalMetricError{},
		},
		{
			testName: "Incompatible cost function",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.CostFunction = CostFunctionBinaryCrossEntropy
				input.LearningTask = LearningTaskRegression
				return input
			},
			expectedError: &IncompatibleCostFunctionError{},
		},
		{
			testName: "Incompatible eval metric",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.CostFunction = CostFunctionMSE
				input.EvaluationMetrics = []EvalMetric{EvalMetricPrecision}
				input.LearningTask = LearningTaskRegression
				return input
			},
			expectedError: &IncompatibleMetricError{},
		},
		{
			testName: "Invalid train ratio",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.TrainRatio = 1.5
				return input
			},
			expectedError: &InvalidTrainRatioError{},
		},
		{
			testName: "Invalid validation ratio",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.ValidationRatio = -0.1
				return input
			},
			expectedError: &InvalidValidationRatioError{},
		},
		{
			testName: "Invalid test ratio",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.TestRatio = 2.0
				return input
			},
			expectedError: &InvalidTestRatioError{},
		},
		{
			testName: "Invalid data split",
			setupInput: func(t *testing.T) TrainingInput {
				input := validTrainingInput(t)
				input.TrainRatio = 0.5
				input.ValidationRatio = 0.3
				input.TestRatio = 0.3
				return input
			},
			expectedError: &InvalidDataSplitError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			input := tt.setupInput(t)
			training, err := NewTraining(input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
				assert.Equal(t, Training{}, training)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, training)
				assert.Equal(t, input.Optimizer, training.Optimizer())
				assert.Equal(t, input.LearningTask, training.LearningTask())
				assert.Equal(t, input.CostFunction, training.CostFunction())
				assert.Equal(t, input.LearningType, training.LearningType())
				assert.True(t, reflect.DeepEqual(input.EvaluationMetrics, training.EvaluationMetrics()))
				assert.Equal(t, input.EarlyStopping.enabled, training.EarlyStopping().enabled)
				assert.Equal(t, input.EarlyStopping.patience, training.EarlyStopping().patience)
				assert.Equal(t, input.EarlyStopping.minDelta, training.EarlyStopping().minDelta)
				assert.Equal(t, input.EarlyStopping.validationMetric, training.EarlyStopping().validationMetric)
				assert.Equal(t, input.CrossValidation.enabled, training.CrossValidation().enabled)
				assert.Equal(t, input.CrossValidation.folds, training.CrossValidation().folds)
				assert.Equal(t, input.TrainRatio, training.TrainRatio())
				assert.Equal(t, input.ValidationRatio, training.ValidationRatio())
				assert.Equal(t, input.TestRatio, training.TestRatio())
			}
		})
	}
}

func TestEquals(t *testing.T) {
	input := validTrainingInput(t)
	training1, err := NewTraining(input)
	require.NoError(t, err)

	training2, err := NewTraining(input)
	require.NoError(t, err)

	assert.True(t, training1.Equals(training2))

	// Modify one field in training2
	training2Modified := training2
	training2Modified.optimizer, _ = NewRMSPropOptimizer(0.001, 0.9, 1e-8)

	assert.False(t, training1.Equals(training2Modified))
}

func TestImmutabilityEvaluationMetrics(t *testing.T) {
	input := validTrainingInput(t)
	training, err := NewTraining(input)
	require.NoError(t, err)

	// Modify the original input's evaluation metrics
	input.EvaluationMetrics[0] = EvalMetricF1Score

	// Check that the training's evaluation metrics remain unchanged
	assert.Equal(t, []EvalMetric{EvalMetricAccuracy, EvalMetricF1Score}, training.EvaluationMetrics())

	// Modify the returned evaluation metrics from the training
	trainingMetrics := training.EvaluationMetrics()
	trainingMetrics[0] = EvalMetricF1Score

	// Check that the training's evaluation metrics remain unchanged
	assert.Equal(t, []EvalMetric{EvalMetricAccuracy, EvalMetricF1Score}, training.EvaluationMetrics())
}
