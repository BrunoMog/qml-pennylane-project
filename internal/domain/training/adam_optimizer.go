package trainingpipeline

type AdamOptimizer struct {
	learningRate float64
	beta1        float64
	beta2        float64
	epsilon      float64
}

func NewAdamOptimizer(learningRate, beta1, beta2, epsilon float64) (AdamOptimizer, error) {
	if isFiniteFloat64(learningRate) && (learningRate <= 0) {
		return AdamOptimizer{}, &InvalidLearningRateError{learningRate}
	}
	if isFiniteFloat64(beta1) && (beta1 < 0 || beta1 >= 1) {
		return AdamOptimizer{}, &InvalidBeta1Error{beta1}
	}
	if isFiniteFloat64(beta2) && (beta2 < 0 || beta2 >= 1) {
		return AdamOptimizer{}, &InvalidBeta2Error{beta2}
	}
	if isFiniteFloat64(epsilon) && (epsilon <= 0) {
		return AdamOptimizer{}, &InvalidEpsilonError{epsilon}
	}

	return AdamOptimizer{
		learningRate: learningRate,
		beta1:        beta1,
		beta2:        beta2,
		epsilon:      epsilon,
	}, nil
}

func (o AdamOptimizer) Name() OptimizerName {
	return OptimizerNameAdam
}

func (o AdamOptimizer) Equal(other Optimizer) bool {
	if otherAdam, ok := other.(AdamOptimizer); ok {
		return o.learningRate == otherAdam.LearningRate() &&
			o.beta1 == otherAdam.Beta1() &&
			o.beta2 == otherAdam.Beta2() &&
			o.epsilon == otherAdam.Epsilon()
	}
	return false
}

func (o AdamOptimizer) LearningRate() float64 {
	return o.learningRate
}

func (o AdamOptimizer) Beta1() float64 {
	return o.beta1
}

func (o AdamOptimizer) Beta2() float64 {
	return o.beta2
}

func (o AdamOptimizer) Epsilon() float64 {
	return o.epsilon
}

func (o AdamOptimizer) isOptimizer() {}
