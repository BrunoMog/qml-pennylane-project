package vqc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMeasurementType(t *testing.T) {

	t.Run("valid measurement type cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
			expected MeasurementType
		}{
			{"valid expectation measurement", "expectation", ExpectationMeasurement},
			{"valid probability measurement", "probability", ProbabilityMeasurement},
			{"valid expectation measurement with whitespace", " expectation ", ExpectationMeasurement},
			{"valid probability measurement with whitespace", " probability ", ProbabilityMeasurement},
			{"valid expectation measurement with uppercase", "EXPECTATION", ExpectationMeasurement},
			{"valid probability measurement with mixed case", "PrObAbIlItY", ProbabilityMeasurement},
		}

		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseMeasurementType(tc.input)
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			})
		}
	})

	t.Run("invalid measurement type cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
		}{
			{"invalid measurement type", "invalid"},
			{"empty string", ""},
			{"whitespace only", "   "},
			{"numeric string", "123"},
			{"special characters", "@#$%"},
			{"mixed valid and invalid characters", "expectation!probability"},
			{"too long string", "this_is_a_very_long_measurement_type_string"},
		}

		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseMeasurementType(tc.input)
				assert.Error(t, err)
				assert.Equal(t, MeasurementType(""), result)
			})
		}
	})

	t.Run("extreme measurement type cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
		}{
			{"100,000 characters", strings.Repeat("a", 100000)},
			{"4,000,000 characters", strings.Repeat("a", 4000000)},
			{"10,000,000 characters", strings.Repeat("a", 10000000)},
		}

		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseMeasurementType(tc.input)
				assert.Error(t, err)
				assert.Equal(t, MeasurementType(""), result)
			})
		}
	})
}

func TestIsValidMeasurementType(t *testing.T) {
	t.Run("valid measurement type cases", func(t *testing.T) {
		validTypes := []MeasurementType{
			ExpectationMeasurement,
			ProbabilityMeasurement,
		}

		for _, measurementType := range validTypes {
			t.Run(string(measurementType), func(t *testing.T) {
				assert.True(t, measurementType.isValid())
			})
		}
	})

	t.Run("invalid measurement type cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    MeasurementType
		}{
			{"invalid measurement type", MeasurementType("invalid")},
			{"empty string", MeasurementType("")},
			{"whitespace only", MeasurementType("   ")},
			{"numeric string", MeasurementType("123")},
			{"special characters", MeasurementType("@#$%")},
			{"too long string", MeasurementType(strings.Repeat("a", 100))},
		}

		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				assert.False(t, tc.input.isValid())
			})
		}
	})
}
