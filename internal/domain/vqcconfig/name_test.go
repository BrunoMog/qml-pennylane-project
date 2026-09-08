package vqcconfig

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameValidation(t *testing.T) {
	tests := []struct {
		testName      string
		inputName     string
		expextedName  string
		expectedError error
	}{
		{
			testName:      "valid name",
			inputName:     "Valid Name",
			expextedName:  "Valid Name",
			expectedError: nil,
		},
		{
			testName:      "empty name",
			inputName:     "",
			expextedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name with only spaces",
			inputName:     "   ",
			expextedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name too short",
			inputName:     "abc",
			expextedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name too long",
			inputName:     strings.Repeat("a", MAX_NAME_LENGTH+1),
			expextedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name with leading and trailing spaces",
			inputName:     "   Valid Name   ",
			expextedName:  "Valid Name",
			expectedError: nil,
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
				assert.Equal(t, tt.expextedName, name.Value())
			}
		})
	}
}
