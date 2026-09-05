package trainingpipeline

import (
	"testing"
)

func TestRMSPropOptimizer_IsValid(t *testing.T) {
	testCases := []struct {
		name         string
		learningRate float64
		decay        float64
		epsilon      float64
		expected     bool
	}{
		{name: "Valid parameters", learningRate: 0.01, decay: 0.9, epsilon: 1e-8, expected: true},
		{name: "Invalid learning rate (negative)", learningRate: -0.01, decay: 0.9, epsilon: 1e-8, expected: false},
		{name: "Invalid decay (negative)", learningRate: 0.01, decay: -0.9, epsilon: 1e-8, expected: false},
		{name: "Invalid epsilon (negative)", learningRate: 0.01, decay: 0.9, epsilon: -1e-8, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			optimizer, err := NewRMSPropOptimizer(tc.learningRate, tc.decay, tc.epsilon)
			if err != nil && tc.expected {
				t.Errorf("Expected valid optimizer, but got error: %v", err)
			}
			if err == nil && !tc.expected {
				t.Errorf("Expected error for invalid parameters, but got valid optimizer")
			}
			if err == nil && (optimizer.LearningRate() != tc.learningRate || optimizer.Decay() != tc.decay || optimizer.Epsilon() != tc.epsilon) {
				t.Errorf("Expected learning rate %f, decay %f, and epsilon %f, but got learning rate %f, decay %f, and epsilon %f", tc.learningRate, tc.decay, tc.epsilon, optimizer.LearningRate(), optimizer.Decay(), optimizer.Epsilon())
			}
		})
	}
}
