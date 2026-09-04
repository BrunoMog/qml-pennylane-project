package trainingpipeline

import (
	"testing"
)

func TestCrossValidationConfig_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		config   CrossValidationConfig
		expected bool
	}{
		{name: "Cross-validation disabled", config: CrossValidationConfig{Enabled: false, Folds: 0}, expected: true},
		{name: "Valid cross-validation", config: CrossValidationConfig{Enabled: true, Folds: 5}, expected: true},
		{name: "Invalid cross-validation (folds <= 1)", config: CrossValidationConfig{Enabled: true, Folds: 1}, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.config.IsValid()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestCrossValidationConfig_EnabledCrossValidation(t *testing.T) {
	testCases := []struct {
		name     string
		config   CrossValidationConfig
		expected bool
	}{
		{name: "Cross-validation disabled", config: CrossValidationConfig{Enabled: false, Folds: 0}, expected: false},
		{name: "Cross-validation enabled", config: CrossValidationConfig{Enabled: true, Folds: 5}, expected: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.config.EnabledCrossValidation()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
