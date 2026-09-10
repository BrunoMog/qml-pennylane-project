package training

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAdamOptimizer(t *testing.T) {
	testCases := []struct {
		expectErr    error
		name         string
		learningRate float64
		beta1        float64
		beta2        float64
		epsilon      float64
	}{
		{name: "Valid parameters", learningRate: 0.001, beta1: 0.9, beta2: 0.999, epsilon: 1e-8, expectErr: nil},
		{name: "Invalid learning rate", learningRate: -0.001, beta1: 0.9, beta2: 0.999, epsilon: 1e-8, expectErr: &InvalidLearningRateError{}},
		{name: "Invalid beta1", learningRate: 0.001, beta1: -0.9, beta2: 0.999, epsilon: 1e-8, expectErr: &InvalidBeta1Error{}},
		{name: "Invalid beta2", learningRate: 0.001, beta1: 0.9, beta2: 1.5, epsilon: 1e-8, expectErr: &InvalidBeta2Error{}},
		{name: "Invalid epsilon", learningRate: 0.001, beta1: 0.9, beta2: 0.999, epsilon: -1e-8, expectErr: &InvalidEpsilonError{}},
		{name: "NaN learning rate", learningRate: math.NaN(), beta1: 0.9, beta2: 0.999, epsilon: 1e-8, expectErr: &InvalidLearningRateError{}},
		{name: "NaN beta1", learningRate: 0.001, beta1: math.NaN(), beta2: 0.999, epsilon: 1e-8, expectErr: &InvalidBeta1Error{}},
		{name: "NaN beta2", learningRate: 0.001, beta1: 0.9, beta2: math.NaN(), epsilon: 1e-8, expectErr: &InvalidBeta2Error{}},
		{name: "NaN epsilon", learningRate: 0.001, beta1: 0.9, beta2: 0.999, epsilon: math.NaN(), expectErr: &InvalidEpsilonError{}},
		{name: "Infinite learning rate", learningRate: math.Inf(1), beta1: 0.9, beta2: 0.999, epsilon: 1e-8, expectErr: &InvalidLearningRateError{}},
		{name: "Infinite beta1", learningRate: 0.001, beta1: math.Inf(1), beta2: 0.999, epsilon: 1e-8, expectErr: &InvalidBeta1Error{}},
		{name: "Infinite beta2", learningRate: 0.001, beta1: 0.9, beta2: math.Inf(1), epsilon: 1e-8, expectErr: &InvalidBeta2Error{}},
		{name: "Infinite epsilon", learningRate: 0.001, beta1: 0.9, beta2: 0.999, epsilon: math.Inf(1), expectErr: &InvalidEpsilonError{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			optimizer, err := NewAdamOptimizer(tc.learningRate, tc.beta1, tc.beta2, tc.epsilon)
			if tc.expectErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.learningRate, optimizer.LearningRate())
				assert.Equal(t, tc.beta1, optimizer.Beta1())
				assert.Equal(t, tc.beta2, optimizer.Beta2())
				assert.Equal(t, tc.epsilon, optimizer.Epsilon())
			}
		})
	}
}
