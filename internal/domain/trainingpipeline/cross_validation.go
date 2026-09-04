package trainingpipeline

type CrossValidationConfig struct {
	Enabled bool
	Folds   int
}

func (cvc CrossValidationConfig) IsValid() bool {
	if !cvc.Enabled {
		return true
	}

	if cvc.Folds <= 1 {
		return false
	}

	return true
}

func (cvc CrossValidationConfig) EnabledCrossValidation() bool {
	return cvc.Enabled
}
