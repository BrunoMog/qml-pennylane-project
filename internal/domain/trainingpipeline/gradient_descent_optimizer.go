package trainingpipeline

type GradientDescentOptimizer struct {
	learningRate float64
}

func NewGradientDescentOptimizer(learningRate float64) (GradientDescentOptimizer, error) {
	if isFiniteFloat64(learningRate) && (learningRate <= 0) {
		return GradientDescentOptimizer{}, &InvalidLearningRateError{learningRate}
	}

	return GradientDescentOptimizer{
		learningRate: learningRate,
	}, nil
}

func (o GradientDescentOptimizer) Name() OptimizerName {
	return OptimizerNameGradientDescent
}

func (o GradientDescentOptimizer) Equal(other Optimizer) bool {
	if otherGD, ok := other.(GradientDescentOptimizer); ok {
		return o.learningRate == otherGD.LearningRate()
	}
	return false
}

func (o GradientDescentOptimizer) LearningRate() float64 {
	return o.learningRate
}

func (o GradientDescentOptimizer) isOptimizer() {}
