package trainingpipeline

type RMSPropOptimizer struct {
	learningRate float64
	decay        float64
	epsilon      float64
}

func NewRMSPropOptimizer(learningRate, decay, epsilon float64) (RMSPropOptimizer, error) {
	if learningRate <= 0 {
		return RMSPropOptimizer{}, &InvalidLearningRateError{learningRate}
	}
	if decay < 0 || decay >= 1 {
		return RMSPropOptimizer{}, &InvalidDecayError{decay}
	}
	if epsilon <= 0 {
		return RMSPropOptimizer{}, &InvalidEpsilonError{epsilon}
	}

	return RMSPropOptimizer{
		learningRate: learningRate,
		decay:        decay,
		epsilon:      epsilon,
	}, nil
}

func (o RMSPropOptimizer) Name() OptimizerName {
	return OptimizerNameRMSProp
}

func (o RMSPropOptimizer) Equals(other Optimizer) bool {
	if otherRMSProp, ok := other.(RMSPropOptimizer); ok {
		return o.learningRate == otherRMSProp.LearningRate() &&
			o.decay == otherRMSProp.Decay() &&
			o.epsilon == otherRMSProp.Epsilon()
	}
	return false
}

func (o RMSPropOptimizer) LearningRate() float64 {
	return o.learningRate
}

func (o RMSPropOptimizer) Decay() float64 {
	return o.decay
}

func (o RMSPropOptimizer) Epsilon() float64 {
	return o.epsilon
}

func (o RMSPropOptimizer) isOptimizer() {}
