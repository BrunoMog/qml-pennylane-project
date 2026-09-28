package vqc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmbeddingType(t *testing.T) {

	t.Run("valid embedding type cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
			expected EmbeddingType
		}{
			{"valid angle embedding type", "angle", EmbeddingTypeAngle},
			{"valid amplitude embedding type", "amplitude", EmbeddingTypeAmplitude},
			{"valid angle embedding type with whitespace", " angle ", EmbeddingTypeAngle},
			{"valid amplitude embedding type with whitespace", " amplitude ", EmbeddingTypeAmplitude},
			{"valid angle embedding type with uppercase", "ANGLE", EmbeddingTypeAngle},
			{"valid angle embedding type with mixed case", "AnGlE", EmbeddingTypeAngle},
		}

		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseEmbeddingType(tc.input)
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			})
		}
	})

	t.Run("invalid embedding type cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
		}{
			{"invalid embedding type", "invalid"},
			{"empty string", ""},
			{"whitespace only", "   "},
			{"numeric string", "123"},
			{"special characters", "@#$%"},
			{"mixed valid and invalid characters", "angle!amplitude"},
			{"too long string", "this_is_a_very_long_embedding_type_string"},
		}

		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseEmbeddingType(tc.input)
				assert.Error(t, err)
				assert.Equal(t, EmbeddingType(""), result)
			})
		}
	})

	t.Run("extreme embedding type cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
		}{
			{"100,000 characters", strings.Repeat("angle", 100000)},
			{"4,000,000 characters", strings.Repeat("amplitude", 4000000)},
			{"10,000,000 characters", strings.Repeat("angle", 10000000)},
		}
		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseEmbeddingType(tc.input)
				assert.Error(t, err)
				assert.Equal(t, EmbeddingType(""), result)
			})
		}
	})
}
