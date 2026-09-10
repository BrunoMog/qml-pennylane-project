package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		expectedError error
		testName      string
		inputName     string
		expectedName  string
	}{
		{
			testName:      "valid name",
			inputName:     "John Doe",
			expectedName:  "John Doe",
			expectedError: nil,
		},
		{
			testName:      "empty name",
			inputName:     "",
			expectedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name with spaces",
			inputName:     "   ",
			expectedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name with special characters",
			inputName:     "John@Doe!",
			expectedName:  "John@Doe!",
			expectedError: nil,
		},
		{
			testName:      "name with leading and trailing spaces",
			inputName:     "  John Doe  ",
			expectedName:  "John Doe",
			expectedError: nil,
		},
		{
			testName:      "too short name",
			inputName:     "J",
			expectedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "too long name",
			inputName:     "This is a very long name that exceeds the maximum allowed length for a name in this system",
			expectedName:  "",
			expectedError: &InvalidNameError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			name, err := NewName(tt.inputName)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedName, name.Value())
			}
		})
	}
}
