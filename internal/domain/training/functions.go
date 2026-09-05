package trainingpipeline

type CostFunction string

const (
	CostFunctionMeanSquaredError   CostFunction = "mean_squared_error"
	CostFunctionBinaryCrossEntropy CostFunction = "binary_cross_entropy"
	CostFunctionMSE                CostFunction = "mse"
	CostFunctionRMSE               CostFunction = "rmse"
	CostFunctionMAE                CostFunction = "mae"
)

func (cf CostFunction) IsValid() bool {
	switch cf {
	case CostFunctionMeanSquaredError, CostFunctionBinaryCrossEntropy,
		CostFunctionMSE, CostFunctionRMSE, CostFunctionMAE:
		return true
	default:
		return false
	}
}

type EvalMetric string

const (
	EvalMetricAccuracy  EvalMetric = "accuracy"
	EvalMetricF1Score   EvalMetric = "f1_score"
	EvalMetricPrecision EvalMetric = "precision"
	EvalMetricRecall    EvalMetric = "recall"
	EvalMetricRMSE      EvalMetric = "rmse"
	EvalMetricMAE       EvalMetric = "mae"
)

func (em EvalMetric) IsValid() bool {
	switch em {
	case EvalMetricAccuracy, EvalMetricF1Score, EvalMetricPrecision,
		EvalMetricRecall, EvalMetricRMSE, EvalMetricMAE:
		return true
	default:
		return false
	}
}
