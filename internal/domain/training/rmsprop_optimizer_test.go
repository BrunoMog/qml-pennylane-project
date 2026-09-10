package training

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRMSPropOptimizer_IsValid(t *testing.T) {
	testCases := []struct {
		expectedError error
		name          string
		learningRate  float64
		decay         float64
		epsilon       float64
	}{
		{name: "Valid parameters", learningRate: 0.01, decay: 0.9, epsilon: 1e-8, expectedError: nil},
		{name: "Invalid learning rate (negative)", learningRate: -0.01, decay: 0.9, epsilon: 1e-8, expectedError: &InvalidLearningRateError{}},
		{name: "Invalid decay (negative)", learningRate: 0.01, decay: -0.9, epsilon: 1e-8, expectedError: &InvalidDecayError{}},
		{name: "Invalid epsilon (negative)", learningRate: 0.01, decay: 0.9, epsilon: -1e-8, expectedError: &InvalidEpsilonError{}},
		{name: "Invalid learning rate NaN", learningRate: math.NaN(), decay: 0.9, epsilon: 1e-8, expectedError: &InvalidLearningRateError{}},
		{name: "Invalid decay NaN", learningRate: 0.01, decay: math.NaN(), epsilon: 1e-8, expectedError: &InvalidDecayError{}},
		{name: "Invalid epsilon NaN", learningRate: 0.01, decay: 0.9, epsilon: math.NaN(), expectedError: &InvalidEpsilonError{}},
		{name: "Invalid learning rate Inf", learningRate: math.Inf(1), decay: 0.9, epsilon: 1e-8, expectedError: &InvalidLearningRateError{}},
		{name: "Invalid decay Inf", learningRate: 0.01, decay: math.Inf(1), epsilon: 1e-8, expectedError: &InvalidDecayError{}},
		{name: "Invalid epsilon Inf", learningRate: 0.01, decay: 0.9, epsilon: math.Inf(1), expectedError: &InvalidEpsilonError{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			optimizer, err := NewRMSPropOptimizer(tc.learningRate, tc.decay, tc.epsilon)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.learningRate, optimizer.LearningRate())
				assert.Equal(t, tc.decay, optimizer.Decay())
				assert.Equal(t, tc.epsilon, optimizer.Epsilon())
			}
		})
	}
}
