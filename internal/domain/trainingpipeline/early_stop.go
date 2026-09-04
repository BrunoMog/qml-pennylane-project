package trainingpipeline

type EarlyStoppingConfig struct {
	validationMetric EvaluationMetric
	patience         int
	minDelta         float64
	enabled          bool
}

func (esc EarlyStoppingConfig) IsValid() bool {
	if !esc.enabled {
		return true
	}

	if esc.patience <= 0 {
		return false
	}

	if esc.minDelta < 0 {
		return false
	}

	if !esc.validationMetric.IsValid() {
		return false
	}

	return true
}

func (esc EarlyStoppingConfig) Enabled() bool {
	return esc.enabled
}

func (esc EarlyStoppingConfig) Patience() int {
	return esc.patience
}

func (esc EarlyStoppingConfig) MinDelta() float64 {
	return esc.minDelta
}

func (esc EarlyStoppingConfig) ValidationMetric() EvaluationMetric {
	return esc.validationMetric
}
