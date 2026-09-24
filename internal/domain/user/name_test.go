package user

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	t.Run("valid name cases", func(t *testing.T) {
		validNames := []struct {
			testName string
			input    string
		}{
			{"simple name", "John Doe"},
			{"name with hyphen", "Mary-Jane"},
			{"name with apostrophe", "O'Connor"},
			{"name with spaces", "Jean-Luc Picard"},
			{"name with accents", "José María"},
			{"name with non-Latin characters", "李小龙"},
			{"name with mixed scripts", "Иван Иванович"},
			{"name with special characters", "Renée O'Connor"},
			{"name with diacritics", "Anaïs Nin"},
			{"name with umlaut", "Chloë Sevigny"},
			{"name with multiple parts", "Robert Downey Jr."},
			{"name with dashes", "d---"},
			{"max length name", strings.Repeat("a", maxNameLength)},
			{"min length name", "abc"},
		}

		for _, testCase := range validNames {
			t.Run(testCase.testName, func(t *testing.T) {
				n, err := NewName(testCase.input)
				assert.NoError(t, err)
				assert.Equal(t, testCase.input, n.String())
			})
		}
	})

	t.Run("invalid name cases", func(t *testing.T) {
		invalidNames := []struct {
			testName string
			input    string
		}{
			{"empty name", ""},
			{"whitespace only", "   "},
			{"too short", "Jo"},
			{"too long", "This is a very long name that exceeds the maximum allowed length for a name in this system"},
			{"name with special characters", "John@Doe!"},
			{"name with newline", "John\nDoe"},
			{"name with tab", "John\tDoe"},
			{"name with null byte", "John\x00Doe"},
			{"name with only dashes", "---"},
			{"name close to max length", strings.Repeat("a", maxNameLength+1)},
		}

		for _, testCase := range invalidNames {
			t.Run(testCase.testName, func(t *testing.T) {
				n, err := NewName(testCase.input)
				assert.Error(t, err)
				assert.IsType(t, &InvalidNameError{}, err)
				assert.Equal(t, "", n.String())
			})
		}
	})

	t.Run("extreme name cases", func(t *testing.T) {
		extremeNames := []struct {
			testName string
			input    string
		}{
			{"100,000 characters", strings.Repeat("a", 100000)},
			{"4,000,000 characters", strings.Repeat("a", 4000000)},
			{"10,000,000 characters", strings.Repeat("a", 10000000)},
		}

		for _, testCase := range extremeNames {
			t.Run(testCase.testName, func(t *testing.T) {
				n, err := NewName(testCase.input)
				assert.Error(t, err)
				assert.IsType(t, &InvalidNameError{}, err)
				assert.Empty(t, n.String())
			})
		}
	})
}

func TestNameIsValid(t *testing.T) {
	t.Run("correctly validates valid name", func(t *testing.T) {
		name, err := NewName("John Doe")
		assert.NoError(t, err)
		assert.True(t, name.IsValid())
	})

	t.Run("correctly identifies invalid name", func(t *testing.T) {
		name, err := NewName("")
		assert.Error(t, err)
		assert.False(t, name.IsValid())
	})
}
