package training

type NesterovMomentumOptimizer struct {
	learningRate float64
	momentum     float64
}

func NewNesterovMomentumOptimizer(learningRate, momentum float64) (NesterovMomentumOptimizer, error) {
	if !isFiniteFloat64(learningRate) || (learningRate <= 0) {
		return NesterovMomentumOptimizer{}, &InvalidLearningRateError{learningRate}
	}
	if !isFiniteFloat64(momentum) || momentum < 0 || momentum >= 1 {
		return NesterovMomentumOptimizer{}, &InvalidMomentumError{momentum}
	}

	return NesterovMomentumOptimizer{
		learningRate: learningRate,
		momentum:     momentum,
	}, nil
}

func (o NesterovMomentumOptimizer) Name() OptimizerName {
	return OptimizerNameNesterovMomentum
}

func (o NesterovMomentumOptimizer) Equal(other Optimizer) bool {
	if otherNesterov, ok := other.(NesterovMomentumOptimizer); ok {
		return o.learningRate == otherNesterov.LearningRate() &&
			o.momentum == otherNesterov.Momentum()
	}
	return false
}

func (o NesterovMomentumOptimizer) LearningRate() float64 {
	return o.learningRate
}

func (o NesterovMomentumOptimizer) Momentum() float64 {
	return o.momentum
}

func (o NesterovMomentumOptimizer) isOptimizer() {}
