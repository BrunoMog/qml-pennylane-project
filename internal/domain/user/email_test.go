package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		expectedError error
		testName      string
		inputEmail    string
		expectedEmail string
	}{
		{
			testName:      "valid email",
			inputEmail:    "test@example.com",
			expectedEmail: "test@example.com",
			expectedError: nil,
		},
		{
			testName:      "invalid email",
			inputEmail:    "invalid-email",
			expectedEmail: "",
			expectedError: &InvalidEmailError{},
		},
		{
			testName:      "empty email",
			inputEmail:    "",
			expectedEmail: "",
			expectedError: &InvalidEmailError{},
		},
		{
			testName:      "email with spaces",
			inputEmail:    " test @example.com ",
			expectedEmail: "",
			expectedError: &InvalidEmailError{},
		},
		{
			testName:      "email with special characters",
			inputEmail:    "test!@example.com",
			expectedEmail: "test!@example.com",
			expectedError: nil,
		},
		{
			testName:      "email with spaces around",
			inputEmail:    " test@example.com ",
			expectedEmail: "test@example.com",
			expectedError: nil,
		},
		{
			testName:      "email with display name",
			inputEmail:    "John Doe <john.doe@example.com>",
			expectedEmail: "",
			expectedError: &InvalidEmailError{},
		},
		{
			testName:      "email with uppercase letters",
			inputEmail:    "JOHN.DOE@EXAMPLE.COM",
			expectedEmail: "john.doe@example.com",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			email, err := NewEmail(tt.inputEmail)
			assert.Equal(t, tt.expectedEmail, email.Value())
			assert.IsType(t, tt.expectedError, err)
		})
	}
}
