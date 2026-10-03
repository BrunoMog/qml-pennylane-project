package vqcconfig

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNameValidation(t *testing.T) {

	t.Run("valid name cases", func(t *testing.T) {
		validNames := []struct {
			testName string
			input    string
		}{
			{"simple name", "Valid Name"},
			{"name with spaces", "Valid Name With Spaces"},
			{"name with hyphen", "Valid-Name"},
			{"name with accents", "José María"},
			{"name with non-Latin characters", "李小龙李小龙"},
			{"name with mixed scripts", "Иван Иванович"},
			{"name with diacritics", "Anaïs Nin"},
			{"name with umlaut", "Chloë Sevigny"},
			{"name with multiple parts", "Robert Downey Jr."},
			{"name with ()", "Name (with parentheses)"},
			{"name with dots", "Name.with.dots"},
			{"name with underscores", "Name_with_underscores"},
			{"name with dashes", "Name-with-dashes"},
			{"max length name", strings.Repeat("a", maxNameLength)},
			{"max length name with accents", strings.Repeat("é", maxNameLength)},
			{"min length name", strings.Repeat("a", minNameLength)},
		}

		for _, testCase := range validNames {
			t.Run(testCase.testName, func(t *testing.T) {
				n, err := NewName(testCase.input)
				assert.NoError(t, err)
				assert.Equal(t, testCase.input, n.String())
				assert.True(t, n.isValid())
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
			{"too short", "ab"},
			{"too long", strings.Repeat("a", maxNameLength+1)},
			{"name with special characters", "John@Doe!"},
			{"name with newline", "John\nDoe"},
			{"name with tab", "John\tDoe"},
			{"name with null byte", "John\x00Doe"},
			{"name with only dashes", "---"},
		}

		for _, testCase := range invalidNames {
			t.Run(testCase.testName, func(t *testing.T) {
				n, err := NewName(testCase.input)
				assert.ErrorIs(t, err, ErrInvalidName)
				assert.Equal(t, "", n.String())
				assert.False(t, n.isValid())
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
				assert.ErrorIs(t, err, ErrInvalidName)
				assert.Equal(t, "", n.String())
				assert.False(t, n.isValid())
			})
		}
	})
}

func TestEqualsName(t *testing.T) {

	t.Run("equal names", func(t *testing.T) {
		inputs := []struct {
			testName string
			name1    string
			name2    string
		}{
			{"same name", "Valid Name", "Valid Name"},
			{"different case", "Valid Name", "valid name"},
			{"leading/trailing spaces", "Valid Name", " Valid Name "},
		}

		for _, testCase := range inputs {
			t.Run(testCase.testName, func(t *testing.T) {
				name1, err := NewName(testCase.name1)
				require.NoError(t, err)
				name2, err := NewName(testCase.name2)
				require.NoError(t, err)
				assert.True(t, name1.Equals(name2))
			})
		}
	})

	t.Run("different names", func(t *testing.T) {
		inputs := []struct {
			testName string
			name1    string
			name2    string
		}{
			{"different names", "Valid Name", "Another Name"},
			{"different lengths", "Valid Name", "Valid Names"},
			{"completely different", "Name One", "Name Two"},
		}

		for _, testCase := range inputs {
			t.Run(testCase.testName, func(t *testing.T) {
				name1, err := NewName(testCase.name1)
				require.NoError(t, err)
				name2, err := NewName(testCase.name2)
				require.NoError(t, err)
				assert.False(t, name1.Equals(name2))
			})
		}
	})
}
