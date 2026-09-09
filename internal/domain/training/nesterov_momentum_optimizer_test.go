package training

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNesterovMomentumOptimizer_IsValid(t *testing.T) {
	testCases := []struct {
		name          string
		learningRate  float64
		momentum      float64
		expectedError error
	}{
		{name: "Valid learning rate and momentum", learningRate: 0.01, momentum: 0.9, expectedError: nil},
		{name: "Invalid learning rate (negative)", learningRate: -0.01, momentum: 0.9, expectedError: &InvalidLearningRateError{}},
		{name: "Invalid learning rate (zero)", learningRate: 0.0, momentum: 0.9, expectedError: &InvalidLearningRateError{}},
		{name: "Invalid momentum (negative)", learningRate: 0.01, momentum: -0.1, expectedError: &InvalidMomentumError{}},
		{name: "Invalid momentum (greater than 1)", learningRate: 0.01, momentum: 1.1, expectedError: &InvalidMomentumError{}},
		{name: "Invalid learning rate (NaN)", learningRate: math.NaN(), momentum: 0.9, expectedError: &InvalidLearningRateError{}},
		{name: "Invalid momentum (NaN)", learningRate: 0.01, momentum: math.NaN(), expectedError: &InvalidMomentumError{}},
		{name: "Invalid learning rate (Inf)", learningRate: math.Inf(1), momentum: 0.9, expectedError: &InvalidLearningRateError{}},
		{name: "Invalid momentum (Inf)", learningRate: 0.01, momentum: math.Inf(1), expectedError: &InvalidMomentumError{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			optimizer, err := NewNesterovMomentumOptimizer(tc.learningRate, tc.momentum)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.learningRate, optimizer.LearningRate())
				assert.Equal(t, tc.momentum, optimizer.Momentum())
			}
		})
	}
}
