package trainingpipeline

import "slices"

const EPSILON = 1e-9

type TrainingPipeline struct {
	optimizer         Optimizer
	learningTask      LearningTask
	costFunction      CostFunction
	learningType      LearningType
	evaluationMetrics []EvaluationMetric
	earlyStopping     EarlyStoppingConfig
	crossValidation   CrossValidationConfig
	trainRatio        float64
	validationRatio   float64
	testRatio         float64
	randomSeed        int
	maxEpochs         uint
	batchSize         uint
}

type TrainingPipelineInput struct {
	Optimizer         Optimizer
	LearningTask      LearningTask
	CostFunction      CostFunction
	LearningType      LearningType
	EvaluationMetrics []EvaluationMetric
	EarlyStopping     EarlyStoppingConfig
	CrossValidation   CrossValidationConfig
	TrainRatio        float64
	ValidationRatio   float64
	TestRatio         float64
	RandomSeed        int
	MaxEpochs         uint
	BatchSize         uint
}

func NewTrainingPipeline(input TrainingPipelineInput) (*TrainingPipeline, error) {
	err := validateInput(input)
	if err != nil {
		return nil, err
	}

	return &TrainingPipeline{
		learningType:      input.LearningType,
		learningTask:      input.LearningTask,
		costFunction:      input.CostFunction,
		evaluationMetrics: input.EvaluationMetrics,
		trainRatio:        input.TrainRatio,
		validationRatio:   input.ValidationRatio,
		testRatio:         input.TestRatio,
		randomSeed:        input.RandomSeed,
		maxEpochs:         input.MaxEpochs,
		batchSize:         input.BatchSize,
		earlyStopping:     input.EarlyStopping,
		crossValidation:   input.CrossValidation,
		optimizer:         input.Optimizer,
	}, nil
}

func validateInput(input TrainingPipelineInput) error {
	if !input.LearningType.IsValid() {
		return &InvalidLearningTypeError{learningType: input.LearningType}
	}

	if !input.LearningTask.IsValid() {
		return &InvalidLearningTaskError{learningTask: input.LearningTask}
	}

	if !input.CostFunction.IsValid() {
		return &InvalidCostFunctionError{costFunction: input.CostFunction}
	}

	for _, metric := range input.EvaluationMetrics {
		if !metric.IsValid() {
			return &InvalidEvaluationMetricError{evaluationMetric: metric}
		}
	}

	if len(input.EvaluationMetrics) == 0 {
		return &InvalidEvaluationMetricError{evaluationMetric: ""}
	}

	if !input.EarlyStopping.IsValid() {
		return &InvalidEarlyStoppingConfigError{earlyStoppingConfig: input.EarlyStopping}
	}

	if !input.CrossValidation.IsValid() {
		return &InvalidCrossValidationConfigError{crossValidationConfig: input.CrossValidation}
	}

	if input.RandomSeed < 0 {
		return &InvalidSeedError{seed: input.RandomSeed}
	}

	if input.MaxEpochs == 0 {
		return &InvalidMaxEpochsError{maxEpochs: input.MaxEpochs}
	}

	if input.BatchSize == 0 {
		return &InvalidBatchSizeError{batchSize: input.BatchSize}
	}

	if input.TrainRatio < 0 || input.TrainRatio > 1 {
		return &InvalidTrainRatioError{trainRatio: input.TrainRatio}
	}

	if input.ValidationRatio < 0 || input.ValidationRatio > 1 {
		return &InvalidValidationRatioError{validationRatio: input.ValidationRatio}
	}

	if input.TestRatio < 0 || input.TestRatio > 1 {
		return &InvalidTestRatioError{testRatio: input.TestRatio}
	}

	if input.TrainRatio+input.ValidationRatio+input.TestRatio < 1-EPSILON || input.TrainRatio+input.ValidationRatio+input.TestRatio > 1+EPSILON {
		return &InvalidDataSplitError{
			trainRatio:      input.TrainRatio,
			validationRatio: input.ValidationRatio,
			testRatio:       input.TestRatio,
		}
	}

	if input.Optimizer == nil {
		return &InvalidOptimizerError{optimizer: nil}
	}

	return nil
}

func (tp TrainingPipeline) Optimizer() Optimizer {
	return tp.optimizer
}

func (tp TrainingPipeline) LearningTask() LearningTask {
	return tp.learningTask
}

func (tp TrainingPipeline) CostFunction() CostFunction {
	return tp.costFunction
}

func (tp TrainingPipeline) LearningType() LearningType {
	return tp.learningType
}

func (tp TrainingPipeline) EvaluationMetrics() []EvaluationMetric {
	return slices.Clone(tp.evaluationMetrics)
}

func (tp TrainingPipeline) EarlyStopping() EarlyStoppingConfig {
	return tp.earlyStopping
}

func (tp TrainingPipeline) CrossValidation() CrossValidationConfig {
	return tp.crossValidation
}

func (tp TrainingPipeline) TrainRatio() float64 {
	return tp.trainRatio
}

func (tp TrainingPipeline) ValidationRatio() float64 {
	return tp.validationRatio
}

func (tp TrainingPipeline) TestRatio() float64 {
	return tp.testRatio
}

func (tp TrainingPipeline) RandomSeed() int {
	return tp.randomSeed
}

func (tp TrainingPipeline) MaxEpochs() uint {
	return tp.maxEpochs
}

func (tp TrainingPipeline) BatchSize() uint {
	return tp.batchSize
}
