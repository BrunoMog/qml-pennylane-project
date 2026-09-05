package training

import (
	"reflect"
	"testing"
)

func validTrainingInput(t *testing.T) TrainingInput {
	t.Helper()

	optimizer, err := NewAdamOptimizer(0.001, 0.9, 0.999, 1e-8)
	if err != nil {
		t.Fatalf("failed to create test optimizer: %v", err)
	}

	return TrainingInput{
		Optimizer:         optimizer,
		LearningTask:      BinaryClassification,
		CostFunction:      CostFunctionBinaryCrossEntropy,
		LearningType:      SupervisedLearning,
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
	t.Run("creates a training with the supplied configuration", func(t *testing.T) {
		input := validTrainingInput(t)

		training, err := NewTraining(input)
		if err != nil {
			t.Fatalf("NewTraining() returned an unexpected error: %v", err)
		}

		if training.Optimizer() != input.Optimizer {
			t.Errorf("Optimizer() = %v, want %v", training.Optimizer(), input.Optimizer)
		}
		if training.LearningTask() != input.LearningTask {
			t.Errorf("LearningTask() = %q, want %q", training.LearningTask(), input.LearningTask)
		}
		if training.CostFunction() != input.CostFunction {
			t.Errorf("CostFunction() = %q, want %q", training.CostFunction(), input.CostFunction)
		}
		if training.LearningType() != input.LearningType {
			t.Errorf("LearningType() = %q, want %q", training.LearningType(), input.LearningType)
		}
		if training.TrainRatio() != input.TrainRatio || training.ValidationRatio() != input.ValidationRatio || training.TestRatio() != input.TestRatio {
			t.Errorf("data split getters do not match input")
		}
		if training.RandomSeed() != input.RandomSeed || training.MaxEpochs() != input.MaxEpochs || training.BatchSize() != input.BatchSize {
			t.Errorf("training parameter getters do not match input")
		}
		if training.EarlyStopping() != input.EarlyStopping {
			t.Errorf("EarlyStopping() = %+v, want %+v", training.EarlyStopping(), input.EarlyStopping)
		}
		if training.CrossValidation() != input.CrossValidation {
			t.Errorf("CrossValidation() = %+v, want %+v", training.CrossValidation(), input.CrossValidation)
		}
		if len(training.EvaluationMetrics()) != len(input.EvaluationMetrics) {
			t.Fatalf("EvaluationMetrics() length = %d, want %d", len(training.EvaluationMetrics()), len(input.EvaluationMetrics))
		}
		for index, metric := range input.EvaluationMetrics {
			if training.EvaluationMetrics()[index] != metric {
				t.Errorf("EvaluationMetrics()[%d] = %q, want %q", index, training.EvaluationMetrics()[index], metric)
			}
		}
	})

	t.Run("returns a copy of evaluation metrics", func(t *testing.T) {
		input := validTrainingInput(t)
		training, err := NewTraining(input)
		if err != nil {
			t.Fatalf("NewTraining() returned an unexpected error: %v", err)
		}

		metrics := training.EvaluationMetrics()
		metrics[0] = EvalMetricRecall
		if training.EvaluationMetrics()[0] != EvalMetricAccuracy {
			t.Errorf("EvaluationMetrics() exposes the training internal slice")
		}
	})
}

func TestNewTraining_ValidDataSplits(t *testing.T) {
	testCases := []struct {
		name                                   string
		trainRatio, validationRatio, testRatio float64
	}{
		{name: "all training", trainRatio: 1},
		{name: "all validation", validationRatio: 1},
		{name: "all test", testRatio: 1},
		{name: "floating point tolerance", trainRatio: 0.7, validationRatio: 0.2, testRatio: 0.1000000005},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := validTrainingInput(t)
			input.TrainRatio, input.ValidationRatio, input.TestRatio = tc.trainRatio, tc.validationRatio, tc.testRatio
			if _, err := NewTraining(input); err != nil {
				t.Fatalf("NewTraining() returned an unexpected error: %v", err)
			}
		})
	}
}

func TestNewTraining_InvalidInput(t *testing.T) {
	testCases := []struct {
		name        string
		mutate      func(*TrainingInput)
		expectedErr error
	}{
		{name: "learning type", mutate: func(input *TrainingInput) { input.LearningType = LearningType("invalid") }, expectedErr: &InvalidLearningTypeError{}},
		{name: "learning task", mutate: func(input *TrainingInput) { input.LearningTask = LearningTask("invalid") }, expectedErr: &InvalidLearningTaskError{}},
		{name: "cost function", mutate: func(input *TrainingInput) { input.CostFunction = CostFunction("invalid") }, expectedErr: &InvalidCostFunctionError{}},
		{name: "evaluation metric", mutate: func(input *TrainingInput) {
			input.EvaluationMetrics = []EvalMetric{EvalMetric("invalid")}
		}, expectedErr: &InvalidEvalMetricError{}},
		{name: "incompatible cost function", mutate: func(input *TrainingInput) {
			input.LearningTask = Regression
			input.CostFunction = CostFunctionBinaryCrossEntropy
		}, expectedErr: &IncompatibleCostFunctionError{}},
		{name: "incompatible evaluation metric", mutate: func(input *TrainingInput) {
			input.EvaluationMetrics = []EvalMetric{EvalMetricRMSE}
		}, expectedErr: &IncompatibleMetricError{}},
		{name: "missing evaluation metrics", mutate: func(input *TrainingInput) { input.EvaluationMetrics = nil }, expectedErr: &InvalidEvalMetricError{}},
		{name: "negative seed", mutate: func(input *TrainingInput) { input.RandomSeed = -1 }, expectedErr: &InvalidSeedError{}},
		{name: "zero max epochs", mutate: func(input *TrainingInput) { input.MaxEpochs = 0 }, expectedErr: &InvalidMaxEpochsError{}},
		{name: "zero batch size", mutate: func(input *TrainingInput) { input.BatchSize = 0 }, expectedErr: &InvalidBatchSizeError{}},
		{name: "negative train ratio", mutate: func(input *TrainingInput) { input.TrainRatio = -0.1 }, expectedErr: &InvalidTrainRatioError{}},
		{name: "train ratio greater than one", mutate: func(input *TrainingInput) { input.TrainRatio = 1.1 }, expectedErr: &InvalidTrainRatioError{}},
		{name: "negative validation ratio", mutate: func(input *TrainingInput) { input.ValidationRatio = -0.1 }, expectedErr: &InvalidValidationRatioError{}},
		{name: "validation ratio greater than one", mutate: func(input *TrainingInput) { input.ValidationRatio = 1.1 }, expectedErr: &InvalidValidationRatioError{}},
		{name: "negative test ratio", mutate: func(input *TrainingInput) { input.TestRatio = -0.1 }, expectedErr: &InvalidTestRatioError{}},
		{name: "test ratio greater than one", mutate: func(input *TrainingInput) { input.TestRatio = 1.1 }, expectedErr: &InvalidTestRatioError{}},
		{name: "data split below one", mutate: func(input *TrainingInput) {
			input.TrainRatio, input.ValidationRatio, input.TestRatio = 0.6, 0.2, 0.1
		}, expectedErr: &InvalidDataSplitError{}},
		{name: "nil optimizer", mutate: func(input *TrainingInput) { input.Optimizer = nil }, expectedErr: &InvalidOptimizerError{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := validTrainingInput(t)
			tc.mutate(&input)

			_, err := NewTraining(input)
			if err == nil {
				t.Fatalf("NewTraining() returned nil error, want %T", tc.expectedErr)
			}
			if reflect.TypeOf(err) != reflect.TypeOf(tc.expectedErr) {
				t.Errorf("NewTraining() error type = %T, want %T", err, tc.expectedErr)
			}
		})
	}
}
