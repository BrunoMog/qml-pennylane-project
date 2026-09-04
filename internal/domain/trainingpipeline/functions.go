package trainingpipeline

type CostFunction string

const (
	MeanSquaredError   CostFunction = "mean_squared_error"
	BinaryCrossEntropy CostFunction = "binary_cross_entropy"
)

func (cf CostFunction) IsValid() bool {
	switch cf {
	case MeanSquaredError, BinaryCrossEntropy:
		return true
	default:
		return false
	}
}

type EvaluationMetric string

const (
	Accuracy  EvaluationMetric = "accuracy"
	F1Score   EvaluationMetric = "f1_score"
	Precision EvaluationMetric = "precision"
	Recall    EvaluationMetric = "recall"
)

func (em EvaluationMetric) IsValid() bool {
	switch em {
	case Accuracy, F1Score, Precision, Recall:
		return true
	default:
		return false
	}
}
