package training

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEarlyStoppingConfig_IsValid(t *testing.T) {
	testCases := []struct {
		expected error
		name     string
		config   EarlyStoppingInput
	}{
		{name: "Early stopping disabled", config: EarlyStoppingInput{Enabled: false, Patience: 0, MinDelta: 0.0, ValidationMetric: EvalMetric("accuracy")}, expected: nil},
		{name: "Valid early stopping", config: EarlyStoppingInput{Enabled: true, Patience: 5, MinDelta: 0.01, ValidationMetric: EvalMetric("accuracy")}, expected: nil},
		{name: "Invalid early stopping (patience <= 0)", config: EarlyStoppingInput{Enabled: true, Patience: 0, MinDelta: 0.01, ValidationMetric: EvalMetric("accuracy")}, expected: &ErrInvalidEarlyStopping{}},
		{name: "Invalid early stopping (minDelta < 0)", config: EarlyStoppingInput{Enabled: true, Patience: 5, MinDelta: -0.01, ValidationMetric: EvalMetric("accuracy")}, expected: &ErrInvalidEarlyStopping{}},
		{name: "Invalid early stopping (invalid validation metric)", config: EarlyStoppingInput{Enabled: true, Patience: 5, MinDelta: 0.01, ValidationMetric: EvalMetric("invalid")}, expected: &ErrInvalidEarlyStopping{}},
		{name: "Invalid early stopping (minDelta is NaN)", config: EarlyStoppingInput{Enabled: true, Patience: 5, MinDelta: math.NaN(), ValidationMetric: EvalMetric("accuracy")}, expected: &ErrInvalidEarlyStopping{}},
		{name: "Invalid early stopping (minDelta is Inf)", config: EarlyStoppingInput{Enabled: true, Patience: 5, MinDelta: math.Inf(1), ValidationMetric: EvalMetric("accuracy")}, expected: &ErrInvalidEarlyStopping{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config, err := NewEarlyStopping(tc.config)
			if tc.expected == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.config.Enabled, config.Enabled())
				assert.Equal(t, tc.config.Patience, config.Patience())
				assert.Equal(t, tc.config.MinDelta, config.MinDelta())
				assert.Equal(t, tc.config.ValidationMetric, config.ValidationMetric())
			} else {
				assert.Error(t, err)
				assert.IsType(t, tc.expected, err)
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
