package trainingpipeline

type LearningType string

const (
	SupervisedLearning LearningType = "supervised"
)

func (lt LearningType) IsValid() bool {
	switch lt {
	case SupervisedLearning:
		return true
	default:
		return false
	}
}

type LearningTask string

const (
	BinaryClassification LearningTask = "classification"
	Regression           LearningTask = "regression"
)

func (tt LearningTask) IsValid() bool {
	switch tt {
	case BinaryClassification, Regression:
		return true
	default:
		return false
	}
}

var taskCostFunctionCompatibility = map[LearningTask]map[CostFunction]bool{
	BinaryClassification: {
		CostFunctionBinaryCrossEntropy: true,
		CostFunctionMeanSquaredError:   true,
	},
	Regression: {
		CostFunctionMSE:  true,
		CostFunctionRMSE: true,
		CostFunctionMAE:  true,
	},
}

var taskMetricCompatibility = map[LearningTask]map[EvalMetric]bool{
	BinaryClassification: {
		EvalMetricAccuracy:  true,
		EvalMetricF1Score:   true,
		EvalMetricPrecision: true,
		EvalMetricRecall:    true,
	},
	Regression: {
		EvalMetricRMSE: true,
		EvalMetricMAE:  true,
	},
}

func (tt LearningTask) IsCostFunctionCompatible(cf CostFunction) bool {
	if compatibleCostFunctions, ok := taskCostFunctionCompatibility[tt]; ok {
		return compatibleCostFunctions[cf]
	}
	return false
}

func (tt LearningTask) IsEvalMetricCompatible(em EvalMetric) bool {
	if compatibleMetrics, ok := taskMetricCompatibility[tt]; ok {
		return compatibleMetrics[em]
	}
	return false
}
