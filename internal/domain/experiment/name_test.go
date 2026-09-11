package experiment

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	testCases := []struct {
		testName           string
		name               string
		expectedNameString string
		expectedError      error
	}{
		{
			testName:           "valid name",
			name:               "Valid Name",
			expectedNameString: "Valid Name",
			expectedError:      nil,
		},
		{
			testName:           "empty name",
			name:               "",
			expectedNameString: "",
			expectedError:      &InvalidNameError{},
		},
		{
			testName:           "name too short",
			name:               "abc",
			expectedNameString: "",
			expectedError:      &InvalidNameError{},
		},
		{
			testName:           "name too long",
			name:               strings.Repeat("a", 101),
			expectedNameString: "",
			expectedError:      &InvalidNameError{},
		},
		{
			testName:           "name with leading and trailing spaces",
			name:               "   Valid Name   ",
			expectedNameString: "Valid Name",
			expectedError:      nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			name, err := NewName(tc.name)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedNameString, name.Value())
			}
		})
	}
}
