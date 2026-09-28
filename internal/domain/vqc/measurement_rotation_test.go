package vqc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMeasurementRotation(t *testing.T) {

	t.Run("valid rotation cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
			expected MeasurementRotation
		}{
			{"valid X rotation", "x", XMeasurementRotation},
			{"valid Y rotation", "y", YMeasurementRotation},
			{"valid Z rotation", "z", ZMeasurementRotation},
			{"valid X rotation uppercase", "X", XMeasurementRotation},
			{"valid rotation with trailing whitespace", "x ", XMeasurementRotation},
			{"valid rotation with leading whitespace", " x", XMeasurementRotation},
			{"valid rotation in boundary case with max length", "x        ", XMeasurementRotation},
		}
		for _, tt := range inputs {
			t.Run(tt.testName, func(t *testing.T) {
				rotation, err := ParseMeasurementRotation(tt.input)
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, rotation)
			})
		}
	})

	t.Run("invalid rotation cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
		}{
			{"invalid rotation", "invalid"},
			{"empty string", ""},
			{"whitespace only", "   "},
			{"numeric string", "123"},
			{"special characters", "@#$%"},
			{"mixed valid and invalid characters", "x!y"},
			{"too long string", "            x   "},
		}
		for _, tt := range inputs {
			t.Run(tt.testName, func(t *testing.T) {
				rotation, err := ParseMeasurementRotation(tt.input)
				assert.Error(t, err)
				assert.Equal(t, MeasurementRotation(""), rotation)
			})
		}
	})

	t.Run("extreme rotation cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
		}{
			{"100,000 characters", strings.Repeat("a", 100000)},
			{"4,000,000 characters", strings.Repeat("a", 4000000)},
			{"10,000,000 characters", strings.Repeat("a", 10000000)},
		}
		for _, tt := range inputs {
			t.Run(tt.testName, func(t *testing.T) {
				rotation, err := ParseMeasurementRotation(tt.input)
				assert.Error(t, err)
				assert.Equal(t, MeasurementRotation(""), rotation)
			})
		}
	})
}

func TestIsValidMeasurementRotation(t *testing.T) {
	t.Run("valid measurement rotation cases", func(t *testing.T) {
		validRotations := []MeasurementRotation{
			XMeasurementRotation,
			YMeasurementRotation,
			ZMeasurementRotation,
		}

		for _, rotation := range validRotations {
			t.Run(string(rotation), func(t *testing.T) {
				assert.True(t, rotation.isValid())
			})
		}
	})

	t.Run("invalid measurement rotation cases", func(t *testing.T) {
		invalidRotations := []MeasurementRotation{
			"invalid",
			"",
			"   ",
			"123",
			"@#$%",
			"x!y",
		}

		for _, rotation := range invalidRotations {
			t.Run(string(rotation), func(t *testing.T) {
				assert.False(t, rotation.isValid())
			})
		}
	})
}
