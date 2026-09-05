package training

import (
	"math"
	"testing"
)

func TestNesterovMomentumOptimizer_IsValid(t *testing.T) {
	testCases := []struct {
		name         string
		learningRate float64
		momentum     float64
		expected     bool
	}{
		{name: "Valid learning rate and momentum", learningRate: 0.01, momentum: 0.9, expected: true},
		{name: "Invalid learning rate (negative)", learningRate: -0.01, momentum: 0.9, expected: false},
		{name: "Invalid learning rate (zero)", learningRate: 0.0, momentum: 0.9, expected: false},
		{name: "Invalid momentum (negative)", learningRate: 0.01, momentum: -0.1, expected: false},
		{name: "Invalid momentum (greater than 1)", learningRate: 0.01, momentum: 1.1, expected: false},
		{name: "Invalid learning rate (NaN)", learningRate: math.NaN(), momentum: 0.9, expected: false},
		{name: "Invalid momentum (NaN)", learningRate: 0.01, momentum: math.NaN(), expected: false},
		{name: "Invalid learning rate (Inf)", learningRate: math.Inf(1), momentum: 0.9, expected: false},
		{name: "Invalid momentum (Inf)", learningRate: 0.01, momentum: math.Inf(1), expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			optimizer, err := NewNesterovMomentumOptimizer(tc.learningRate, tc.momentum)
			if err != nil && tc.expected {
				t.Errorf("Expected valid optimizer, but got error: %v", err)
			}
			if err == nil && !tc.expected {
				t.Errorf("Expected error for invalid parameters, but got valid optimizer")
			}
			if err == nil && (optimizer.LearningRate() != tc.learningRate || optimizer.Momentum() != tc.momentum) {
				t.Errorf("Expected learning rate %f and momentum %f, but got learning rate %f and momentum %f", tc.learningRate, tc.momentum, optimizer.LearningRate(), optimizer.Momentum())
			}
		})
	}
}
