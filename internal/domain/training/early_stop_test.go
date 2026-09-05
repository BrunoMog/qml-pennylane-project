package trainingpipeline

import (
	"testing"
)

func TestEarlyStoppingConfig_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		config   EarlyStoppingInput
		expected bool
	}{
		{name: "Early stopping disabled", config: EarlyStoppingInput{Enabled: false, Patience: 0, MinDelta: 0.0, ValidationMetric: EvalMetric("accuracy")}, expected: true},
		{name: "Valid early stopping", config: EarlyStoppingInput{Enabled: true, Patience: 5, MinDelta: 0.01, ValidationMetric: EvalMetric("accuracy")}, expected: true},
		{name: "Invalid early stopping (patience <= 0)", config: EarlyStoppingInput{Enabled: true, Patience: 0, MinDelta: 0.01, ValidationMetric: EvalMetric("accuracy")}, expected: false},
		{name: "Invalid early stopping (minDelta < 0)", config: EarlyStoppingInput{Enabled: true, Patience: 5, MinDelta: -0.01, ValidationMetric: EvalMetric("accuracy")}, expected: false},
		{name: "Invalid early stopping (invalid validation metric)", config: EarlyStoppingInput{Enabled: true, Patience: 5, MinDelta: 0.01, ValidationMetric: EvalMetric("invalid")}, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config, err := NewEarlyStopping(tc.config)
			if err != nil && tc.expected {
				t.Errorf("Expected no error, but got %v", err)
			}
			if err == nil && !tc.expected {
				t.Errorf("Expected error, but got none")
			}
			if err == nil && tc.expected {
				if config.Enabled() != tc.config.Enabled {
					t.Errorf("Expected Enabled to be %v, but got %v", tc.config.Enabled, config.Enabled())
				}
				if config.Patience() != tc.config.Patience {
					t.Errorf("Expected Patience to be %d, but got %d", tc.config.Patience, config.Patience())
				}
				if config.MinDelta() != tc.config.MinDelta {
					t.Errorf("Expected MinDelta to be %f, but got %f", tc.config.MinDelta, config.MinDelta())
				}
				if config.validationMetric != tc.config.ValidationMetric {
					t.Errorf("Expected ValidationMetric to be %v, but got %v", tc.config.ValidationMetric, config.validationMetric)
				}
			}
		})
	}
}

func TestEarlyStoppingConfig_EnabledEarlyStopping(t *testing.T) {
	testCases := []struct {
		name     string
		config   EarlyStopping
		expected bool
	}{
		{name: "Early stopping disabled", config: EarlyStopping{enabled: false}, expected: false},
		{name: "Early stopping enabled", config: EarlyStopping{enabled: true}, expected: true},
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
