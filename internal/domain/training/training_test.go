package trainingpipeline

import (
	"reflect"
	"testing"
)

func validTrainingPipelineInput(t *testing.T) TrainingPipelineInput {
	t.Helper()

	optimizer, err := NewAdamOptimizer(0.001, 0.9, 0.999, 1e-8)
	if err != nil {
		t.Fatalf("failed to create test optimizer: %v", err)
	}

	return TrainingPipelineInput{
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

func TestNewTrainingPipeline(t *testing.T) {
	t.Run("creates a pipeline with the supplied configuration", func(t *testing.T) {
		input := validTrainingPipelineInput(t)

		pipeline, err := NewTrainingPipeline(input)
		if err != nil {
			t.Fatalf("NewTrainingPipeline() returned an unexpected error: %v", err)
		}

		if pipeline.Optimizer() != input.Optimizer {
			t.Errorf("Optimizer() = %v, want %v", pipeline.Optimizer(), input.Optimizer)
		}
		if pipeline.LearningTask() != input.LearningTask {
			t.Errorf("LearningTask() = %q, want %q", pipeline.LearningTask(), input.LearningTask)
		}
		if pipeline.CostFunction() != input.CostFunction {
			t.Errorf("CostFunction() = %q, want %q", pipeline.CostFunction(), input.CostFunction)
		}
		if pipeline.LearningType() != input.LearningType {
			t.Errorf("LearningType() = %q, want %q", pipeline.LearningType(), input.LearningType)
		}
		if pipeline.TrainRatio() != input.TrainRatio || pipeline.ValidationRatio() != input.ValidationRatio || pipeline.TestRatio() != input.TestRatio {
			t.Errorf("data split getters do not match input")
		}
		if pipeline.RandomSeed() != input.RandomSeed || pipeline.MaxEpochs() != input.MaxEpochs || pipeline.BatchSize() != input.BatchSize {
			t.Errorf("training parameter getters do not match input")
		}
		if pipeline.EarlyStopping() != input.EarlyStopping {
			t.Errorf("EarlyStopping() = %+v, want %+v", pipeline.EarlyStopping(), input.EarlyStopping)
		}
		if pipeline.CrossValidation() != input.CrossValidation {
			t.Errorf("CrossValidation() = %+v, want %+v", pipeline.CrossValidation(), input.CrossValidation)
		}
		if len(pipeline.EvaluationMetrics()) != len(input.EvaluationMetrics) {
			t.Fatalf("EvaluationMetrics() length = %d, want %d", len(pipeline.EvaluationMetrics()), len(input.EvaluationMetrics))
		}
		for index, metric := range input.EvaluationMetrics {
			if pipeline.EvaluationMetrics()[index] != metric {
				t.Errorf("EvaluationMetrics()[%d] = %q, want %q", index, pipeline.EvaluationMetrics()[index], metric)
			}
		}
	})

	t.Run("returns a copy of evaluation metrics", func(t *testing.T) {
		input := validTrainingPipelineInput(t)
		pipeline, err := NewTrainingPipeline(input)
		if err != nil {
			t.Fatalf("NewTrainingPipeline() returned an unexpected error: %v", err)
		}

		metrics := pipeline.EvaluationMetrics()
		metrics[0] = EvalMetricRecall
		if pipeline.EvaluationMetrics()[0] != EvalMetricAccuracy {
			t.Errorf("EvaluationMetrics() exposes the pipeline's internal slice")
		}
	})
}

func TestNewTrainingPipeline_ValidDataSplits(t *testing.T) {
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
			input := validTrainingPipelineInput(t)
			input.TrainRatio, input.ValidationRatio, input.TestRatio = tc.trainRatio, tc.validationRatio, tc.testRatio
			if _, err := NewTrainingPipeline(input); err != nil {
				t.Fatalf("NewTrainingPipeline() returned an unexpected error: %v", err)
			}
		})
	}
}

func TestNewTrainingPipeline_InvalidInput(t *testing.T) {
	testCases := []struct {
		name        string
		mutate      func(*TrainingPipelineInput)
		expectedErr error
	}{
		{name: "learning type", mutate: func(input *TrainingPipelineInput) { input.LearningType = LearningType("invalid") }, expectedErr: &InvalidLearningTypeError{}},
		{name: "learning task", mutate: func(input *TrainingPipelineInput) { input.LearningTask = LearningTask("invalid") }, expectedErr: &InvalidLearningTaskError{}},
		{name: "cost function", mutate: func(input *TrainingPipelineInput) { input.CostFunction = CostFunction("invalid") }, expectedErr: &InvalidCostFunctionError{}},
		{name: "evaluation metric", mutate: func(input *TrainingPipelineInput) {
			input.EvaluationMetrics = []EvalMetric{EvalMetric("invalid")}
		}, expectedErr: &InvalidEvalMetricError{}},
		{name: "incompatible cost function", mutate: func(input *TrainingPipelineInput) { input.CostFunction = CostFunction("invalid") }, expectedErr: &IncompatibleCostFunctionError{}},
		{name: "incompatible evaluation metric", mutate: func(input *TrainingPipelineInput) {
			input.EvaluationMetrics = []EvalMetric{EvalMetric("invalid")}
		}, expectedErr: &IncompatibleMetricError{}},
		{name: "missing evaluation metrics", mutate: func(input *TrainingPipelineInput) { input.EvaluationMetrics = nil }, expectedErr: &InvalidEvalMetricError{}},
		{name: "early stopping", mutate: func(input *TrainingPipelineInput) { input.EarlyStopping.patience = 0 }, expectedErr: &InvalidEarlyStoppingError{}},
		{name: "cross-validation", mutate: func(input *TrainingPipelineInput) { input.CrossValidation.folds = 1 }, expectedErr: &InvalidCrossValidationConfigError{}},
		{name: "negative seed", mutate: func(input *TrainingPipelineInput) { input.RandomSeed = -1 }, expectedErr: &InvalidSeedError{}},
		{name: "zero max epochs", mutate: func(input *TrainingPipelineInput) { input.MaxEpochs = 0 }, expectedErr: &InvalidMaxEpochsError{}},
		{name: "zero batch size", mutate: func(input *TrainingPipelineInput) { input.BatchSize = 0 }, expectedErr: &InvalidBatchSizeError{}},
		{name: "negative train ratio", mutate: func(input *TrainingPipelineInput) { input.TrainRatio = -0.1 }, expectedErr: &InvalidTrainRatioError{}},
		{name: "train ratio greater than one", mutate: func(input *TrainingPipelineInput) { input.TrainRatio = 1.1 }, expectedErr: &InvalidTrainRatioError{}},
		{name: "negative validation ratio", mutate: func(input *TrainingPipelineInput) { input.ValidationRatio = -0.1 }, expectedErr: &InvalidValidationRatioError{}},
		{name: "validation ratio greater than one", mutate: func(input *TrainingPipelineInput) { input.ValidationRatio = 1.1 }, expectedErr: &InvalidValidationRatioError{}},
		{name: "negative test ratio", mutate: func(input *TrainingPipelineInput) { input.TestRatio = -0.1 }, expectedErr: &InvalidTestRatioError{}},
		{name: "test ratio greater than one", mutate: func(input *TrainingPipelineInput) { input.TestRatio = 1.1 }, expectedErr: &InvalidTestRatioError{}},
		{name: "data split below one", mutate: func(input *TrainingPipelineInput) {
			input.TrainRatio, input.ValidationRatio, input.TestRatio = 0.6, 0.2, 0.1
		}, expectedErr: &InvalidDataSplitError{}},
		{name: "nil optimizer", mutate: func(input *TrainingPipelineInput) { input.Optimizer = nil }, expectedErr: &InvalidOptimizerError{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := validTrainingPipelineInput(t)
			tc.mutate(&input)

			_, err := NewTrainingPipeline(input)
			if err == nil {
				t.Fatalf("NewTrainingPipeline() returned nil error, want %T", tc.expectedErr)
			}
			if reflect.TypeOf(err) != reflect.TypeOf(tc.expectedErr) {
				t.Errorf("NewTrainingPipeline() error type = %T, want %T", err, tc.expectedErr)
			}
		})
	}
}
