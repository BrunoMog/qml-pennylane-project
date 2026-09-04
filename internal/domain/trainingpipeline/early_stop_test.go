package trainingpipeline

import (
	"testing"
)

func TestEarlyStoppingConfig_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		config   EarlyStoppingConfig
		expected bool
	}{
		{name: "Early stopping disabled", config: EarlyStoppingConfig{enabled: false, patience: 0, minDelta: 0.0, validationMetric: EvaluationMetric("accuracy")}, expected: true},
		{name: "Valid early stopping", config: EarlyStoppingConfig{enabled: true, patience: 5, minDelta: 0.01, validationMetric: EvaluationMetric("accuracy")}, expected: true},
		{name: "Invalid early stopping (patience <= 0)", config: EarlyStoppingConfig{enabled: true, patience: 0, minDelta: 0.01, validationMetric: EvaluationMetric("accuracy")}, expected: false},
		{name: "Invalid early stopping (minDelta < 0)", config: EarlyStoppingConfig{enabled: true, patience: 5, minDelta: -0.01, validationMetric: EvaluationMetric("accuracy")}, expected: false},
		{name: "Invalid early stopping (invalid validation metric)", config: EarlyStoppingConfig{enabled: true, patience: 5, minDelta: 0.01, validationMetric: EvaluationMetric("invalid")}, expected: false},
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

func TestEarlyStoppingConfig_EnabledEarlyStopping(t *testing.T) {
	testCases := []struct {
		name     string
		config   EarlyStoppingConfig
		expected bool
	}{
		{name: "Early stopping disabled", config: EarlyStoppingConfig{enabled: false}, expected: false},
		{name: "Early stopping enabled", config: EarlyStoppingConfig{enabled: true}, expected: true},
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
