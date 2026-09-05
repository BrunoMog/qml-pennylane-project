package training

import (
	"math"
	"testing"
)

func TestGradientDescentOptimizer_IsValid(t *testing.T) {
	testCases := []struct {
		name         string
		learningRate float64
		expected     bool
	}{
		{name: "Valid learning rate", learningRate: 0.01, expected: true},
		{name: "Invalid learning rate (negative)", learningRate: -0.01, expected: false},
		{name: "Invalid learning rate (zero)", learningRate: 0.0, expected: false},
		{name: "Invalid learning rate (NaN)", learningRate: math.NaN(), expected: false},
		{name: "Invalid learning rate (Inf)", learningRate: math.Inf(1), expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			optimizer, err := NewGradientDescentOptimizer(tc.learningRate)
			if err != nil && tc.expected {
				t.Errorf("Expected valid optimizer, but got error: %v", err)
			}
			if err == nil && !tc.expected {
				t.Errorf("Expected error for invalid learning rate, but got valid optimizer")
			}
			if err == nil && optimizer.LearningRate() != tc.learningRate {
				t.Errorf("Expected learning rate %f, but got %f", tc.learningRate, optimizer.LearningRate())
			}
		})
	}
}
