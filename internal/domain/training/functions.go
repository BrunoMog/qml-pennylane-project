package training

import (
	"strings"
)

type CostFunction string

const (
	CostFunctionBinaryCrossEntropy CostFunction = "binary_cross_entropy"
	CostFunctionMSE                CostFunction = "mse"
	CostFunctionRMSE               CostFunction = "rmse"
	CostFunctionMAE                CostFunction = "mae"
)

func (cf CostFunction) isValid() bool {
	switch cf {
	case CostFunctionBinaryCrossEntropy,
		CostFunctionMSE, CostFunctionRMSE, CostFunctionMAE:
		return true
	default:
		return false
	}
}

func ParseCostFunction(cf string) (CostFunction, error) {
	switch strings.ToLower(strings.TrimSpace(cf)) {
	case string(CostFunctionBinaryCrossEntropy):
		return CostFunctionBinaryCrossEntropy, nil
	case string(CostFunctionMSE):
		return CostFunctionMSE, nil
	case string(CostFunctionRMSE):
		return CostFunctionRMSE, nil
	case string(CostFunctionMAE):
		return CostFunctionMAE, nil
	default:
		return "", &InvalidCostFunctionError{costFunction: CostFunction(cf)}
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

func (em EvalMetric) isValid() bool {
	switch em {
	case EvalMetricAccuracy, EvalMetricF1Score, EvalMetricPrecision,
		EvalMetricRecall, EvalMetricRMSE, EvalMetricMAE:
		return true
	default:
		return false
	}
}

func ParseEvalMetric(em string) (EvalMetric, error) {
	switch strings.ToLower(strings.TrimSpace(em)) {
	case string(EvalMetricAccuracy):
		return EvalMetricAccuracy, nil
	case string(EvalMetricF1Score):
		return EvalMetricF1Score, nil
	case string(EvalMetricPrecision):
		return EvalMetricPrecision, nil
	case string(EvalMetricRecall):
		return EvalMetricRecall, nil
	case string(EvalMetricRMSE):
		return EvalMetricRMSE, nil
	case string(EvalMetricMAE):
		return EvalMetricMAE, nil
	default:
		return "", &InvalidEvalMetricError{evaluationMetric: EvalMetric(em)}
	}
}
