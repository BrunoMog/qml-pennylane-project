package trainingpipeline

type EarlyStopping struct {
	validationMetric EvalMetric
	patience         int
	minDelta         float64
	enabled          bool
}

type EarlyStoppingInput struct {
	ValidationMetric EvalMetric
	Patience         int
	MinDelta         float64
	Enabled          bool
}

func NewEarlyStopping(input EarlyStoppingInput) (EarlyStopping, error) {
	if !input.IsValid() {
		return EarlyStopping{}, &ErrInvalidEarlyStopping{}
	}

	stopping := EarlyStopping{
		validationMetric: input.ValidationMetric,
		patience:         input.Patience,
		minDelta:         input.MinDelta,
		enabled:          input.Enabled,
	}
	return stopping, nil
}

func (esc EarlyStoppingInput) IsValid() bool {
	if !esc.Enabled {
		return true
	}
	if esc.Patience <= 0 {
		return false
	}
	if isFiniteFloat64(esc.MinDelta) && (esc.MinDelta < 0) {
		return false
	}
	if !esc.ValidationMetric.IsValid() {
		return false
	}

	return true
}

func (esc EarlyStopping) Enabled() bool {
	return esc.enabled
}

func (esc EarlyStopping) Patience() int {
	return esc.patience
}

func (esc EarlyStopping) MinDelta() float64 {
	return esc.minDelta
}

func (esc EarlyStopping) ValidationMetric() EvalMetric {
	return esc.validationMetric
}
