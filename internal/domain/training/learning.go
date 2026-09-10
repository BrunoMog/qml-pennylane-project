package training

import "strings"

type LearningType string

const (
	LearningTypeSupervised LearningType = "supervised"
)

func (lt LearningType) Value() string {
	return string(lt)
}

func (lt LearningType) isValid() bool {
	switch lt {
	case LearningTypeSupervised:
		return true
	default:
		return false
	}
}

func ParseLearningType(lt string) (LearningType, error) {
	switch strings.ToLower(strings.TrimSpace(lt)) {
	case "supervised":
		return LearningTypeSupervised, nil
	default:
		return "", &InvalidLearningTypeError{learningType: LearningType(lt)}
	}
}

type LearningTask string

const (
	LearningTaskBinaryClassification LearningTask = "binary_classification"
	LearningTaskRegression           LearningTask = "regression"
)

func (tt LearningTask) Value() string {
	return string(tt)
}

func (tt LearningTask) isValid() bool {
	switch tt {
	case LearningTaskBinaryClassification, LearningTaskRegression:
		return true
	default:
		return false
	}
}

func ParseLearningTask(tt string) (LearningTask, error) {
	switch strings.ToLower(strings.TrimSpace(tt)) {
	case "binary_classification":
		return LearningTaskBinaryClassification, nil
	case "regression":
		return LearningTaskRegression, nil
	default:
		return "", &InvalidLearningTaskError{learningTask: LearningTask(tt)}
	}
}

var taskCostFunctionCompatibility = map[LearningTask]map[CostFunction]bool{
	LearningTaskBinaryClassification: {
		CostFunctionBinaryCrossEntropy: true,
		CostFunctionMSE:                true,
	},
	LearningTaskRegression: {
		CostFunctionMSE:  true,
		CostFunctionRMSE: true,
		CostFunctionMAE:  true,
	},
}

var taskMetricCompatibility = map[LearningTask]map[EvalMetric]bool{
	LearningTaskBinaryClassification: {
		EvalMetricAccuracy:  true,
		EvalMetricF1Score:   true,
		EvalMetricPrecision: true,
		EvalMetricRecall:    true,
	},
	LearningTaskRegression: {
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
