package training

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFiniteFloat64(t *testing.T) {
	testCases := []struct {
		testName string
		value    float64
		expected bool
	}{
		{"Finite positive number", 1.0, true},
		{"Finite negative number", -1.0, true},
		{"Zero", 0.0, true},
		{"Positive infinity", math.Inf(1), false},
		{"Negative infinity", math.Inf(-1), false},
		{"NaN", math.NaN(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := isFiniteFloat64(tc.value)
			assert.Equal(t, tc.expected, result)
		})
	}
}
