package user

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {

	t.Run("valid email cases", func(t *testing.T) {
		validEmails := []struct {
			testName string
			input    string
		}{
			{"simple email", "test@example.com"},
			{"email with dot in local part", "john.doe@example.com"},
			{"email with multiple dots in local part", "john.doe.smith@example.com"},
			{"email with mixed case", "JoHn.Doe@Example.Com"},
			{"email with tag", "user.name+tag@sub.domain.org"},
			{"email with quoted local part", `"john..doe"@example.com`},
		}

		for _, testCase := range validEmails {
			t.Run(testCase.testName, func(t *testing.T) {
				e, err := NewEmail(testCase.input)
				assert.NoError(t, err)
				assert.Equal(t, strings.ToLower(testCase.input), e.String())
			})
		}
	})

	t.Run("invalid email cases", func(t *testing.T) {
		invalidEmails := []struct {
			testName string
			input    string
		}{
			{"invalid email", "invalid-email"},
			{"empty email", ""},
			{"email with name", "John Doe <john.doe@example.com>"},
			{"email starting with dot", ".test@example.com"},
			{"email ending with dot", "test.@example.com"},
			{"email with double dots", "te..st@example.com"},
			{"email with special characters", "!@#$%^&*()@example.com"},
			{"email too long", strings.Repeat("a", 65) + "@example.com"},
		}

		for _, testCase := range invalidEmails {
			t.Run(testCase.testName, func(t *testing.T) {
				e, err := NewEmail(testCase.input)
				assert.Error(t, err)
				assert.IsType(t, &InvalidEmailError{}, err)
				assert.Equal(t, "", e.String())
			})
		}
	})

	t.Run("extreme email cases", func(t *testing.T) {
		extremeEmails := []struct {
			testName string
			input    string
		}{
			{"100,000 characters", strings.Repeat("a", 100000) + "@example.com"},
			{"4,000,000 characters", strings.Repeat("a", 4000000) + "@example.com"},
			{"10,000,000 characters", strings.Repeat("a", 10000000) + "@example.com"},
		}

		for _, testCase := range extremeEmails {
			t.Run(testCase.testName, func(t *testing.T) {
				e, err := NewEmail(testCase.input)
				assert.Error(t, err)
				assert.IsType(t, &InvalidEmailError{}, err)
				assert.Empty(t, e.String())
			})
		}
	})
}

func TestEmailIsValid(t *testing.T) {
	t.Run("correctly validates valid email", func(t *testing.T) {
		email, err := NewEmail("test@example.com")
		assert.NoError(t, err)
		assert.True(t, email.IsValid())
	})

	t.Run("correctly identifies invalid email", func(t *testing.T) {
		email, err := NewEmail("invalid-email")
		assert.Error(t, err)
		assert.False(t, email.IsValid())
	})
}
