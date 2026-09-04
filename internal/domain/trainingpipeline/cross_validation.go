package trainingpipeline

type CrossValidation struct {
	enabled bool
	folds   int
}

type CrossValidationInput struct {
	Enabled bool
	Folds   int
}

func NewCrossValidationConfig(input CrossValidationInput) (CrossValidation, error) {
	if !input.IsValid() {
		return CrossValidation{}, &ErrInvalidCrossValidationConfig{}
	}

	config := CrossValidation{
		enabled: input.Enabled,
		folds:   input.Folds,
	}
	return config, nil
}

func (cvc CrossValidationInput) IsValid() bool {
	if !cvc.Enabled {
		return true
	}

	if cvc.Folds <= 1 {
		return false
	}

	return true
}

func (cvc CrossValidation) Enabled() bool {
	return cvc.enabled
}

func (cvc CrossValidation) Folds() int {
	return cvc.folds
}
