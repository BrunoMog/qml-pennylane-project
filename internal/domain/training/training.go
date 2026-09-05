package training

import (
	"math"
	"slices"
)

const EPSILON = 1e-9

type Training struct {
	optimizer         Optimizer
	learningTask      LearningTask
	costFunction      CostFunction
	learningType      LearningType
	evaluationMetrics []EvalMetric
	earlyStopping     EarlyStopping
	crossValidation   CrossValidation
	trainRatio        float64
	validationRatio   float64
	testRatio         float64
	randomSeed        int
	maxEpochs         uint
	batchSize         uint
}

type TrainingInput struct {
	Optimizer         Optimizer
	LearningTask      LearningTask
	CostFunction      CostFunction
	LearningType      LearningType
	EvaluationMetrics []EvalMetric
	EarlyStopping     EarlyStopping
	CrossValidation   CrossValidation
	TrainRatio        float64
	ValidationRatio   float64
	TestRatio         float64
	RandomSeed        int
	MaxEpochs         uint
	BatchSize         uint
}

func NewTraining(input TrainingInput) (*Training, error) {
	err := validateInput(input)
	if err != nil {
		return nil, err
	}

	return &Training{
		learningType:      input.LearningType,
		learningTask:      input.LearningTask,
		costFunction:      input.CostFunction,
		evaluationMetrics: slices.Clone(input.EvaluationMetrics),
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

func validateInput(input TrainingInput) error {
	err := validateLearningSettings(input.LearningType, input.LearningTask)
	if err != nil {
		return err
	}

	err = validateFunctions(input.CostFunction, input.EvaluationMetrics)
	if err != nil {
		return err
	}

	err = validateFunctionsCompatibility(input.LearningTask, input.CostFunction, input.EvaluationMetrics)
	if err != nil {
		return err
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

	err = validateDataSplit(input.TrainRatio, input.ValidationRatio, input.TestRatio)
	if err != nil {
		return err
	}

	if input.Optimizer == nil {
		return &InvalidOptimizerError{optimizer: nil}
	}

	return nil
}

func validateLearningSettings(learningType LearningType, learningTask LearningTask) error {
	if !learningType.IsValid() {
		return &InvalidLearningTypeError{learningType: learningType}
	}
	if !learningTask.IsValid() {
		return &InvalidLearningTaskError{learningTask: learningTask}
	}

	return nil
}

func validateFunctions(costFunction CostFunction, evaluationMetrics []EvalMetric) error {
	if !costFunction.IsValid() {
		return &InvalidCostFunctionError{costFunction: costFunction}
	}

	for _, metric := range evaluationMetrics {
		if !metric.IsValid() {
			return &InvalidEvalMetricError{evaluationMetric: metric}
		}
	}

	if len(evaluationMetrics) == 0 {
		return &InvalidEvalMetricError{evaluationMetric: ""}
	}

	return nil
}

func validateFunctionsCompatibility(learningTask LearningTask, costFunction CostFunction, evaluationMetrics []EvalMetric) error {
	if !learningTask.IsCostFunctionCompatible(costFunction) {
		return &IncompatibleCostFunctionError{Task: learningTask, CostFunction: costFunction}
	}
	for _, metric := range evaluationMetrics {
		if !learningTask.IsEvalMetricCompatible(metric) {
			return &IncompatibleMetricError{Task: learningTask, Metric: metric}
		}
	}

	return nil
}

func validateDataSplit(trainRatio, validationRatio, testRatio float64) error {
	if !isFiniteFloat64(trainRatio) || trainRatio < 0 || trainRatio > 1 {
		return &InvalidTrainRatioError{trainRatio: trainRatio}
	}

	if !isFiniteFloat64(validationRatio) || validationRatio < 0 || validationRatio > 1 {
		return &InvalidValidationRatioError{validationRatio: validationRatio}
	}

	if !isFiniteFloat64(testRatio) || testRatio < 0 || testRatio > 1 {
		return &InvalidTestRatioError{testRatio: testRatio}
	}

	total := trainRatio + validationRatio + testRatio
	if math.Abs(total-1.0) > EPSILON {
		return &InvalidDataSplitError{
			trainRatio:      trainRatio,
			validationRatio: validationRatio,
			testRatio:       testRatio,
		}
	}

	return nil
}

func (t Training) Optimizer() Optimizer {
	return t.optimizer
}

func (t Training) LearningTask() LearningTask {
	return t.learningTask
}

func (t Training) CostFunction() CostFunction {
	return t.costFunction
}

func (t Training) LearningType() LearningType {
	return t.learningType
}

func (t Training) EvaluationMetrics() []EvalMetric {
	return slices.Clone(t.evaluationMetrics)
}

func (t Training) EarlyStopping() EarlyStopping {
	return t.earlyStopping
}

func (t Training) CrossValidation() CrossValidation {
	return t.crossValidation
}

func (t Training) TrainRatio() float64 {
	return t.trainRatio
}

func (t Training) ValidationRatio() float64 {
	return t.validationRatio
}

func (t Training) TestRatio() float64 {
	return t.testRatio
}

func (t Training) RandomSeed() int {
	return t.randomSeed
}

func (t Training) MaxEpochs() uint {
	return t.maxEpochs
}

func (t Training) BatchSize() uint {
	return t.batchSize
}
