package user

import (
	"testing"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{"valid name", "John Doe", false},
		{"empty name", "", true},
		{"name too long", "ThisNameIsWayTooLongToBeValid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateName(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateName() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	//TODO: Implement email validation tests when the validation logic is added
}

func TestIsValidRole(t *testing.T) {
	tests := []struct {
		name     string
		input    Role
		expected bool
	}{
		{"valid role owner", RoleOwner, true},
		{"valid role admin", RoleAdmin, true},
		{"valid role user", RoleUser, true},
		{"valid role guest", RoleGuest, true},
		{"invalid role", "invalid_role", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsValidRole()
			if result != tt.expected {
				t.Errorf("IsValidRole() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	tests := []struct {
		name       string
		inputName  string
		inputEmail string
		expectErr  bool
	}{
		{"valid user", "John Doe", "john.doe@example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.inputName, tt.inputEmail)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewUser() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if user != nil && (user.name != tt.inputName || user.email != tt.inputEmail) {
				t.Errorf("NewUser() returned user with name = %v, email = %v; want name = %v, email = %v", user.name, user.email, tt.inputName, tt.inputEmail)
			}
		})
	}
}

func TestUserSetters(t *testing.T) {
	user, _ := NewUser("John Doe", "john.doe@example.com")
	if user == nil {
		t.Fatal("Failed to create user")
	}

	user.SetName("Jane Doe")
	if user.name != "Jane Doe" {
		t.Errorf("Expected name 'Jane Doe', got '%s'", user.name)
	}

	user.SetEmail("jane.doe@example.com")
	if user.email != "jane.doe@example.com" {
		t.Errorf("Expected email 'jane.doe@example.com', got '%s'", user.email)
	}
}
