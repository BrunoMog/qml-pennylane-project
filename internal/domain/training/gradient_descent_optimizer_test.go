package training

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradientDescentOptimizer_IsValid(t *testing.T) {
	testCases := []struct {
		expectedError error
		testName      string
		learningRate  float64
	}{
		{testName: "Valid learning rate", learningRate: 0.01, expectedError: nil},
		{testName: "Invalid learning rate (negative)", learningRate: -0.01, expectedError: &InvalidLearningRateError{}},
		{testName: "Invalid learning rate (zero)", learningRate: 0.0, expectedError: &InvalidLearningRateError{}},
		{testName: "Invalid learning rate (NaN)", learningRate: math.NaN(), expectedError: &InvalidLearningRateError{}},
		{testName: "Invalid learning rate (Inf)", learningRate: math.Inf(1), expectedError: &InvalidLearningRateError{}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			optimizer, err := NewGradientDescentOptimizer(tc.learningRate)
			assert.IsType(t, tc.expectedError, err)
			if err == nil {
				assert.Equal(t, tc.learningRate, optimizer.LearningRate())
			}
		})
	}
}
