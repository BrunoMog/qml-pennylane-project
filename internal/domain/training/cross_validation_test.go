package trainingpipeline

import (
	"testing"
)

func TestCrossValidationConfig_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		config   CrossValidationInput
		expected bool
	}{
		{name: "Cross-validation disabled", config: CrossValidationInput{Enabled: false, Folds: 0}, expected: true},
		{name: "Valid cross-validation", config: CrossValidationInput{Enabled: true, Folds: 5}, expected: true},
		{name: "Invalid cross-validation (folds <= 1)", config: CrossValidationInput{Enabled: true, Folds: 1}, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config, err := NewCrossValidationConfig(tc.config)
			if (err != nil) != !tc.expected {
				t.Errorf("NewCrossValidationConfig() error = %v, expected error = %v", err, !tc.expected)
			}
			if err == nil && tc.expected {
				if config.Enabled() != tc.config.Enabled {
					t.Errorf("Expected Enabled() = %v, got %v", tc.config.Enabled, config.Enabled())
				}
				if config.Folds() != tc.config.Folds {
					t.Errorf("Expected Folds() = %v, got %v", tc.config.Folds, config.Folds())
				}
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
