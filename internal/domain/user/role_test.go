package user

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRole(t *testing.T) {

	t.Run("valid role cases", func(t *testing.T) {
		validRoles := []struct {
			testName     string
			inputRole    string
			expectedRole Role
		}{
			{"owner role", "owner", RoleOwner},
			{"admin role", "admin", RoleAdmin},
			{"user role", "user", RoleUser},
			{"guest role", "guest", RoleGuest},
			{"role with trailing spaces", "  user  ", RoleUser}, // with spaces
			{"role with mixed case", "AdMIn", RoleAdmin},        // mixed case
		}

		for _, testCase := range validRoles {
			t.Run(testCase.testName, func(t *testing.T) {
				role, err := ParseRole(testCase.inputRole)
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedRole, role)
			})
		}
	})

	t.Run("invalid role cases", func(t *testing.T) {
		invalidRoles := []struct {
			testName string
			input    string
		}{
			{"invalid role", "invalid_role"},
			{"empty role", ""},
			{"superuser role", "superuser"},
			{"manager role", "manager"},
			{"numeric role", "123"},
		}

		for _, testCase := range invalidRoles {
			t.Run(testCase.testName, func(t *testing.T) {
				role, err := ParseRole(testCase.input)
				assert.Error(t, err)
				assert.Equal(t, ErrInvalidParseRole, err)
				assert.Equal(t, "", role.String())
			})
		}
	})

	t.Run("extreme role cases", func(t *testing.T) {
		extremeRoles := []struct {
			testName string
			input    string
		}{
			{"100,000 characters", strings.Repeat("a", 100000)},
			{"4,000,000 characters", strings.Repeat("a", 4000000)},
			{"10,000,000 characters", strings.Repeat("a", 10000000)},
		}

		for _, testCase := range extremeRoles {
			t.Run(testCase.testName, func(t *testing.T) {
				parsedRole, err := ParseRole(testCase.input)
				assert.Error(t, err)
				assert.IsType(t, ErrInvalidParseRole, err)
				assert.Empty(t, parsedRole.String())
			})
		}
	})
}

func TestIsValidRole(t *testing.T) {

	t.Run("correctly validates valid role cases", func(t *testing.T) {
		validRoles := []struct {
			testName string
			role     Role
		}{
			{"owner role", RoleOwner},
			{"admin role", RoleAdmin},
			{"user role", RoleUser},
			{"guest role", RoleGuest},
		}

		for _, testCase := range validRoles {
			t.Run(testCase.testName, func(t *testing.T) {
				assert.True(t, testCase.role.IsValidRole())
			})
		}
	})

	t.Run("correctly identifies invalid role cases", func(t *testing.T) {
		invalidRoles := []struct {
			testName string
			role     Role
		}{
			{"invalid role", "invalid_role"},
			{"empty role", ""},
			{"superuser role", "superuser"},
			{"manager role", "manager"},
			{"numeric role", "123"},
		}

		for _, testCase := range invalidRoles {
			t.Run(testCase.testName, func(t *testing.T) {
				assert.False(t, testCase.role.IsValidRole())
			})
		}
	})
}
