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
