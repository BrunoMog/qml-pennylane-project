package training

import (
	"math"
	"testing"
)

func TestNewAdamOptimizer(t *testing.T) {
	testCases := []struct {
		name         string
		learningRate float64
		beta1        float64
		beta2        float64
		epsilon      float64
		expectErr    bool
	}{
		{name: "Valid parameters", learningRate: 0.001, beta1: 0.9, beta2: 0.999, epsilon: 1e-8, expectErr: false},
		{name: "Invalid learning rate", learningRate: -0.001, beta1: 0.9, beta2: 0.999, epsilon: 1e-8, expectErr: true},
		{name: "Invalid beta1", learningRate: 0.001, beta1: -0.9, beta2: 0.999, epsilon: 1e-8, expectErr: true},
		{name: "Invalid beta2", learningRate: 0.001, beta1: 0.9, beta2: 1.5, epsilon: 1e-8, expectErr: true},
		{name: "Invalid epsilon", learningRate: 0.001, beta1: 0.9, beta2: 0.999, epsilon: -1e-8, expectErr: true},
		{name: "NaN learning rate", learningRate: math.NaN(), beta1: 0.9, beta2: 0.999, epsilon: 1e-8, expectErr: true},
		{name: "NaN beta1", learningRate: 0.001, beta1: math.NaN(), beta2: 0.999, epsilon: 1e-8, expectErr: true},
		{name: "NaN beta2", learningRate: 0.001, beta1: 0.9, beta2: math.NaN(), epsilon: 1e-8, expectErr: true},
		{name: "NaN epsilon", learningRate: 0.001, beta1: 0.9, beta2: 0.999, epsilon: math.NaN(), expectErr: true},
		{name: "Infinite learning rate", learningRate: math.Inf(1), beta1: 0.9, beta2: 0.999, epsilon: 1e-8, expectErr: true},
		{name: "Infinite beta1", learningRate: 0.001, beta1: math.Inf(1), beta2: 0.999, epsilon: 1e-8, expectErr: true},
		{name: "Infinite beta2", learningRate: 0.001, beta1: 0.9, beta2: math.Inf(1), epsilon: 1e-8, expectErr: true},
		{name: "Infinite epsilon", learningRate: 0.001, beta1: 0.9, beta2: 0.999, epsilon: math.Inf(1), expectErr: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			optimizer, err := NewAdamOptimizer(tc.learningRate, tc.beta1, tc.beta2, tc.epsilon)
			if (err != nil) != tc.expectErr {
				t.Errorf("NewAdamOptimizer(%f, %f, %f, %f) returned error: %v, expected error: %v", tc.learningRate, tc.beta1, tc.beta2, tc.epsilon, err != nil, tc.expectErr)
			}
			if !tc.expectErr {
				if optimizer.Name() != OptimizerNameAdam {
					t.Errorf("Expected optimizer name %s, got %s", OptimizerNameAdam, optimizer.Name())
				}
				if optimizer.LearningRate() != tc.learningRate {
					t.Errorf("Expected learning rate %f, got %f", tc.learningRate, optimizer.LearningRate())
				}
				if optimizer.Beta1() != tc.beta1 {
					t.Errorf("Expected beta1 %f, got %f", tc.beta1, optimizer.Beta1())
				}
				if optimizer.Beta2() != tc.beta2 {
					t.Errorf("Expected beta2 %f, got %f", tc.beta2, optimizer.Beta2())
				}
				if optimizer.Epsilon() != tc.epsilon {
					t.Errorf("Expected epsilon %f, got %f", tc.epsilon, optimizer.Epsilon())
				}
			}
		})
	}
}
