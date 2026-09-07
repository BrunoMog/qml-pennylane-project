package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRole(t *testing.T) {
	tests := []struct {
		testName      string
		inputRole     string
		expectedRole  Role
		expectedError error
	}{
		{
			testName:      "valid role: user",
			inputRole:     "user",
			expectedRole:  RoleUser,
			expectedError: nil,
		},
		{
			testName:      "valid role: admin",
			inputRole:     "admin",
			expectedRole:  RoleAdmin,
			expectedError: nil,
		},
		{
			testName:      "invalid role",
			inputRole:     "invalid_role",
			expectedRole:  "",
			expectedError: &InvalidRoleError{},
		},
		{
			testName:      "empty role",
			inputRole:     "",
			expectedRole:  "",
			expectedError: &InvalidRoleError{},
		},
		{
			testName:      "role with spaces",
			inputRole:     "  user ",
			expectedRole:  RoleUser,
			expectedError: nil,
		},
		{
			testName:      "role with mixed case",
			inputRole:     "AdMiN",
			expectedRole:  RoleAdmin,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			role, err := ParseRole(tt.inputRole)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRole, role)
			}
		})
	}
}

func TestIsValidRole(t *testing.T) {
	tests := []struct {
		testName     string
		inputRole    Role
		expectedBool bool
	}{
		{
			testName:     "valid role user",
			inputRole:    RoleUser,
			expectedBool: true,
		},
		{
			testName:     "valid role admin",
			inputRole:    RoleAdmin,
			expectedBool: true,
		},
		{
			testName:     "valid role owner",
			inputRole:    RoleOwner,
			expectedBool: true,
		},
		{
			testName:     "valid role guest",
			inputRole:    RoleGuest,
			expectedBool: true,
		},
		{
			testName:     "invalid role",
			inputRole:    Role("invalid_role"),
			expectedBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result := tt.inputRole.IsValidRole()
			assert.Equal(t, tt.expectedBool, result)
		})
	}
}

func TestRoleValue(t *testing.T) {
	tests := []struct {
		testName      string
		inputRole     Role
		expectedValue string
	}{
		{
			testName:      "role user",
			inputRole:     RoleUser,
			expectedValue: "user",
		},
		{
			testName:      "role admin",
			inputRole:     RoleAdmin,
			expectedValue: "admin",
		},
		{
			testName:      "role owner",
			inputRole:     RoleOwner,
			expectedValue: "owner",
		},
		{
			testName:      "role guest",
			inputRole:     RoleGuest,
			expectedValue: "guest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			value := tt.inputRole.Value()
			assert.Equal(t, tt.expectedValue, value)
		})
	}
}
