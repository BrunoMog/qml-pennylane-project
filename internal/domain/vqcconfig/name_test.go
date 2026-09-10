package vqcconfig

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNameValidation(t *testing.T) {
	tests := []struct {
		expectedError error
		testName      string
		inputName     string
		expectedName  string
	}{
		{
			testName:      "valid name",
			inputName:     "Valid Name",
			expectedName:  "Valid Name",
			expectedError: nil,
		},
		{
			testName:      "empty name",
			inputName:     "",
			expectedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name with only spaces",
			inputName:     "   ",
			expectedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name too short",
			inputName:     "abc",
			expectedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name too long",
			inputName:     strings.Repeat("a", maxNameLength+1),
			expectedName:  "",
			expectedError: &InvalidNameError{},
		},
		{
			testName:      "name with leading and trailing spaces",
			inputName:     "   Valid Name   ",
			expectedName:  "Valid Name",
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
				assert.Equal(t, tt.expectedName, name.Value())
			}
		})
	}
}

func TestEqualsName(t *testing.T) {
	tests := []struct {
		setup    func() (Name, Name)
		testName string
		expected bool
	}{
		{
			testName: "equal names",
			setup: func() (Name, Name) {
				name1, err := NewName("Valid Name")
				require.NoError(t, err)
				name2, err := NewName("Valid Name")
				require.NoError(t, err)
				return name1, name2
			},
			expected: true,
		},
		{
			testName: "different names",
			setup: func() (Name, Name) {
				name1, err := NewName("Name One")
				require.NoError(t, err)
				name2, err := NewName("Name Two")
				require.NoError(t, err)
				return name1, name2
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			name1, name2 := tt.setup()
			assert.Equal(t, tt.expected, name1.Equals(name2))
		})
	}
}
