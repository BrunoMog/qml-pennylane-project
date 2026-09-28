package vqc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmbeddingRotation(t *testing.T) {

	t.Run("valid rotation cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
			expected EmbeddingRotation
		}{
			{"valid X rotation", "x", XRotation},
			{"valid Y rotation", "y", YRotation},
			{"valid Z rotation", "z", ZRotation},
			{"valid X rotation uppercase", "X", XRotation},
			{"valid rotation with trailing whitespace", "x ", XRotation},
			{"valid rotation with leading whitespace", " x", XRotation},
		}
		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseEmbeddingRotation(tc.input)
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
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
		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseEmbeddingRotation(tc.input)
				assert.Error(t, err)
				assert.Equal(t, EmbeddingRotation(""), result)
			})
		}
	})

	t.Run("extreme rotation cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
		}{
			{"100,000 characters", strings.Repeat("x", 100000)},
			{"4,000,000 characters", strings.Repeat("y", 4000000)},
			{"10,000,000 characters", strings.Repeat("z", 10000000)},
		}
		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseEmbeddingRotation(tc.input)
				assert.Error(t, err)
				assert.Equal(t, EmbeddingRotation(""), result)
			})
		}
	})
}

func TestIsValid(t *testing.T) {

	t.Run("valid rotation cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			rotation EmbeddingRotation
		}{
			{"valid X rotation", XRotation},
			{"valid Y rotation", YRotation},
			{"valid Z rotation", ZRotation},
		}
		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result := tc.rotation.isValid()
				assert.True(t, result)
			})
		}
	})

	t.Run("invalid rotation case", func(t *testing.T) {
		rotation := EmbeddingRotation("invalid")
		result := rotation.isValid()
		assert.False(t, result)
	})
}
