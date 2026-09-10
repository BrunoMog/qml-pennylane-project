package trainconfig

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		expectError  error
		testName     string
		name         string
		expectedName string
	}{
		{
			testName:     "valid name",
			name:         "Valid Name",
			expectedName: "Valid Name",
			expectError:  nil,
		},
		{
			testName:     "name too short",
			name:         "abc",
			expectedName: "",
			expectError:  &InvalidNameError{},
		},
		{
			testName:     "name too long",
			name:         strings.Repeat("a", maxNameLength+1),
			expectedName: "",
			expectError:  &InvalidNameError{},
		},
		{
			testName:     "name with leading and trailing spaces",
			name:         "  Valid Name  ",
			expectedName: "Valid Name",
			expectError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			name, err := NewName(tt.name)
			if tt.expectError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectError, err)
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
				name1, err := NewName("Equal Name")
				require.NoError(t, err)
				name2, err := NewName("Equal Name")
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
