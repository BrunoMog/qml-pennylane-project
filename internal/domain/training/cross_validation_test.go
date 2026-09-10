package training

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrossValidationConfig_IsValid(t *testing.T) {
	testCases := []struct {
		expected error
		name     string
		config   CrossValidationInput
	}{
		{name: "Cross-validation disabled", config: CrossValidationInput{Enabled: false, Folds: 0}, expected: nil},
		{name: "Valid cross-validation", config: CrossValidationInput{Enabled: true, Folds: 5}, expected: nil},
		{name: "Invalid cross-validation (folds <= 1)", config: CrossValidationInput{Enabled: true, Folds: 1}, expected: &ErrInvalidCrossValidationConfig{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config, err := NewCrossValidation(tc.config)
			if tc.expected == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.config.Enabled, config.Enabled())
				assert.Equal(t, tc.config.Folds, config.Folds())
			} else {
				assert.Error(t, err)
				assert.IsType(t, tc.expected, err)
			}
		})
	}
}

func TestCrossValidationConfig_Enabled(t *testing.T) {
	testCases := []struct {
		name     string
		config   CrossValidation
		expected bool
	}{
		{name: "Cross-validation disabled", config: CrossValidation{enabled: false, folds: 0}, expected: false},
		{name: "Cross-validation enabled", config: CrossValidation{enabled: true, folds: 5}, expected: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.config.Enabled()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
