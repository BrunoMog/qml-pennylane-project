package vqcconfig

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescriptionValidation(t *testing.T) {
	tests := []struct {
		testName         string
		inputDescription string
		expectedError    error
	}{
		{
			testName:         "valid description",
			inputDescription: "This is a valid description.",
			expectedError:    nil,
		},
		{
			testName:         "empty description",
			inputDescription: "",
			expectedError:    nil,
		},
		{
			testName:         "description too long",
			inputDescription: strings.Repeat("a", MAX_DESCRIPTION_LENGTH+1),
			expectedError:    &InvalidDescriptionError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := NewDescription(tt.inputDescription)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
