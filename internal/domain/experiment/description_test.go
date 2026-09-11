package experiment

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDescription(t *testing.T) {
	testCases := []struct {
		testName            string
		description         string
		expectedDescription string
		expectedError       error
	}{
		{
			testName:            "valid description",
			description:         "This is a valid description.",
			expectedDescription: "This is a valid description.",
			expectedError:       nil,
		},
		{
			testName:            "empty description",
			description:         "",
			expectedDescription: "",
			expectedError:       nil,
		},
		{
			testName:            "description too long",
			description:         strings.Repeat("a", 501),
			expectedDescription: "",
			expectedError:       &InvalidDescriptionError{},
		},
		{
			testName:            "description with leading and trailing spaces",
			description:         "   This is a valid description.   ",
			expectedDescription: "This is a valid description.",
			expectedError:       nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			description, err := NewDescription(tc.description)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedDescription, description.Value())
			}
		})
	}
}
