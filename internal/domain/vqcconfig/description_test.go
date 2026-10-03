package vqcconfig

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescriptionValidation(t *testing.T) {

	t.Run("valid description cases", func(t *testing.T) {
		validDescriptions := []struct {
			testName string
			input    string
		}{
			{"simple description", "This is a valid description."},
			{"empty description", ""},
			{"description with allowed special characters", "Description with special characters: !@#$%^&*()_+-=[]{}|;:',.<>?/"},
			{"description with spaces", "This description has spaces."},
			{"description with newlines", "This description\nhas newlines."},
			{"description with tabs", "This description\thas tabs."},
			{"description with unicode characters", "Description with unicode: 你好, мир, مرحبا"},
			{"max length description", strings.Repeat("a", maxDescriptionLength)},
			{"no description", ""},
		}
		for _, testCase := range validDescriptions {
			t.Run(testCase.testName, func(t *testing.T) {
				d, err := NewDescription(testCase.input)
				assert.NoError(t, err)
				assert.Equal(t, testCase.input, d.String())
			})
		}
	})

	t.Run("invalid description cases", func(t *testing.T) {
		invalidDescriptions := []struct {
			testName string
			input    string
		}{
			{"description too long", strings.Repeat("a", maxDescriptionLength+1)},
			{"description with invalid character", "This description has an invalid character: \x00"},
			{"another invalid character", "This description has another invalid character: \x1F"},
		}

		for _, testCase := range invalidDescriptions {
			t.Run(testCase.testName, func(t *testing.T) {
				d, err := NewDescription(testCase.input)
				assert.ErrorIs(t, err, ErrInvalidDescription)
				assert.Equal(t, "", d.String())
			})
		}
	})

	t.Run("extreme description cases", func(t *testing.T) {
		invalidDescriptions := []struct {
			testName string
			input    string
		}{
			{"100,000 characters", strings.Repeat("a", 100000)},
			{"4,000,000 characters", strings.Repeat("a", 4000000)},
			{"10,000,000 characters", strings.Repeat("a", 10000000)},
		}

		for _, testCase := range invalidDescriptions {
			t.Run(testCase.testName, func(t *testing.T) {
				d, err := NewDescription(testCase.input)
				assert.ErrorIs(t, err, ErrInvalidDescription)
				assert.Equal(t, "", d.String())
			})
		}
	})
}
