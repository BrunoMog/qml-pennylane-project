package trainconfig

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDescription(t *testing.T) {
	tests := []struct {
		expectError  error
		testName     string
		description  string
		expectedDesc string
	}{
		{
			testName:     "valid description",
			description:  "Valid Description",
			expectedDesc: "Valid Description",
			expectError:  nil,
		},
		{
			testName:     "void description",
			description:  "",
			expectedDesc: "",
			expectError:  nil,
		},
		{
			testName:     "description too long",
			description:  strings.Repeat("a", maxDescriptionLength+1),
			expectedDesc: "",
			expectError:  &InvalidDescriptionError{},
		},
		{
			testName:     "description with leading and trailing spaces",
			description:  "  Valid Description  ",
			expectedDesc: "Valid Description",
			expectError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			desc, err := NewDescription(tt.description)
			if tt.expectError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedDesc, desc.Value())
			}
		})
	}
}
