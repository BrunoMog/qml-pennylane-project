package trainingpipeline

import (
	"fmt"
)

type InvalidSeedError struct {
	seed int
}

func (e *InvalidSeedError) Error() string {
	return fmt.Sprintf("Invalid seed value: %d. Seed must be a non-negative integer.", e.seed)
}

type InvalidLearningTypeError struct {
	learningType LearningType
}

func (e *InvalidLearningTypeError) Error() string {
	return fmt.Sprintf("Invalid learning type: %s. Valid options are: supervised.", e.learningType)
}

type InvalidLearningTaskError struct {
	learningTask LearningTask
}

func (e *InvalidLearningTaskError) Error() string {
	return fmt.Sprintf("Invalid learning task: %s. Valid options are: classification, regression.", e.learningTask)
}

type InvalidCostFunctionError struct {
	costFunction CostFunction
}

func (e *InvalidCostFunctionError) Error() string {
	return fmt.Sprintf("Invalid cost function: %s. Valid options are: mean_squared_error, binary_cross_entropy.", e.costFunction)
}

type InvalidEvalMetricError struct {
	evaluationMetric EvalMetric
}

func (e *InvalidEvalMetricError) Error() string {
	return fmt.Sprintf("Invalid evaluation metric: %s. Valid options are: accuracy, f1_score, precision, recall.", e.evaluationMetric)
}

type InvalidTrainRatioError struct {
	trainRatio float64
}

func (e *InvalidTrainRatioError) Error() string {
	return fmt.Sprintf("Invalid train ratio: %f. Train ratio must be between 0 and 1.", e.trainRatio)
}

type InvalidValidationRatioError struct {
	validationRatio float64
}

func (e *InvalidValidationRatioError) Error() string {
	return fmt.Sprintf("Invalid validation ratio: %f. Validation ratio must be between 0 and 1.", e.validationRatio)
}

type InvalidTestRatioError struct {
	testRatio float64
}

func (e *InvalidTestRatioError) Error() string {
	return fmt.Sprintf("Invalid test ratio: %f. Test ratio must be between 0 and 1.", e.testRatio)
}

type InvalidDataSplitError struct {
	trainRatio      float64
	validationRatio float64
	testRatio       float64
}

func (e *InvalidDataSplitError) Error() string {
	return fmt.Sprintf("Invalid data split: train ratio (%f) + validation ratio (%f) + test ratio (%f) must equal 1.", e.trainRatio, e.validationRatio, e.testRatio)
}

type InvalidMaxEpochsError struct {
	maxEpochs uint
}

func (e *InvalidMaxEpochsError) Error() string {
	return fmt.Sprintf("Invalid max epochs: %d. Max epochs must be a positive integer.", e.maxEpochs)
}

type InvalidBatchSizeError struct {
	batchSize uint
}

func (e *InvalidBatchSizeError) Error() string {
	return fmt.Sprintf("Invalid batch size: %d. Batch size must be a positive integer.", e.batchSize)
}

type InvalidEarlyStoppingError struct {
	earlyStopping EarlyStopping
}

func (e *InvalidEarlyStoppingError) Error() string {
	return fmt.Sprintf("Invalid early stopping configuration: %+v. Please ensure the configuration is valid.", e.earlyStopping)
}

type ErrInvalidEarlyStopping struct{}

func (e *ErrInvalidEarlyStopping) Error() string {
	return "Invalid early stopping configuration. Please ensure the configuration is valid."
}

type InvalidCrossValidationConfigError struct {
	crossValidationConfig CrossValidation
}

func (e *InvalidCrossValidationConfigError) Error() string {
	return fmt.Sprintf("Invalid cross-validation configuration: %+v. Please ensure the configuration is valid.", e.crossValidationConfig)
}

type ErrInvalidCrossValidationConfig struct{}

func (e *ErrInvalidCrossValidationConfig) Error() string {
	return "Invalid cross-validation configuration. Please ensure the configuration is valid."
}

type InvalidLearningRateError struct {
	learningRate float64
}

func (e *InvalidLearningRateError) Error() string {
	return fmt.Sprintf("Invalid learning rate: %f. Learning rate must be a positive number.", e.learningRate)
}

type InvalidBeta1Error struct {
	beta1 float64
}

func (e *InvalidBeta1Error) Error() string {
	return fmt.Sprintf("Invalid beta1 value: %f. Beta1 must be in the range [0, 1).", e.beta1)
}

type InvalidBeta2Error struct {
	beta2 float64
}

func (e *InvalidBeta2Error) Error() string {
	return fmt.Sprintf("Invalid beta2 value: %f. Beta2 must be in the range [0, 1).", e.beta2)
}

type InvalidEpsilonError struct {
	epsilon float64
}

func (e *InvalidEpsilonError) Error() string {
	return fmt.Sprintf("Invalid epsilon value: %f. Epsilon must be a positive number.", e.epsilon)
}

type InvalidDecayError struct {
	decay float64
}

func (e *InvalidDecayError) Error() string {
	return fmt.Sprintf("Invalid decay value: %f. Decay must be a non-negative number.", e.decay)
}

type InvalidMomentumError struct {
	momentum float64
}

func (e *InvalidMomentumError) Error() string {
	return fmt.Sprintf("Invalid momentum value: %f. Momentum must be in the range [0, 1).", e.momentum)
}

type InvalidOptimizerError struct {
	optimizer Optimizer
}

func (e *InvalidOptimizerError) Error() string {
	return fmt.Sprintf("Invalid optimizer: %+v. Please ensure the optimizer is valid.", e.optimizer)
}

type IncompatibleCostFunctionError struct {
	Task         LearningTask
	CostFunction CostFunction
}

func (e *IncompatibleCostFunctionError) Error() string {
	return fmt.Sprintf("cost function '%s' is not compatible with learning task '%s'", e.CostFunction, e.Task)
}

type IncompatibleMetricError struct {
	Task   LearningTask
	Metric EvalMetric
}

func (e *IncompatibleMetricError) Error() string {
	return fmt.Sprintf("evaluation metric '%s' is not compatible with learning task '%s'", e.Metric, e.Task)
}
