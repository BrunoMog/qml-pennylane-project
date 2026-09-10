package user

import (
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestUserCreation(t *testing.T) {
	tests := []struct {
		testName     string
		setup        func() (Name, Email)
		expectedRole Role
	}{
		{
			testName: "create user with valid name and email",
			setup: func() (Name, Email) {
				name, _ := NewName("John Doe")
				email, _ := NewEmail("john.doe@example.com")
				return name, email
			},
			expectedRole: RoleUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			inputName, inputEmail := tt.setup()
			user := NewUser(inputName, inputEmail)
			assert.NotNil(t, user)
			assert.Equal(t, inputName, user.Name())
			assert.Equal(t, inputEmail, user.Email())
			assert.Equal(t, tt.expectedRole, user.Role())
			assert.NotEqual(t, uuid.Nil, user.ID())
		})
	}
}

func TestSetRole(t *testing.T) {
	tests := []struct {
		expectedErr error
		setup       func() *User
		testName    string
		newRole     Role
	}{
		{
			testName: "set role to admin",
			setup: func() *User {
				name, _ := NewName("John Doe")
				email, _ := NewEmail("john.doe@example.com")
				return NewUser(name, email)
			},
			newRole:     RoleAdmin,
			expectedErr: nil,
		},
		{
			testName: "try to set invalid role",
			setup: func() *User {
				name, _ := NewName("John Doe")
				email, _ := NewEmail("john.doe@example.com")
				return NewUser(name, email)
			},
			newRole:     "invalid",
			expectedErr: &InvalidRoleError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			user := tt.setup()
			err := user.SetRole(tt.newRole)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newRole, user.Role())
			}
		})
	}
}

func TestIsAdmin(t *testing.T) {
	tests := []struct {
		setup           func() *User
		testName        string
		expectedIsAdmin bool
	}{
		{
			testName: "user is admin",
			setup: func() *User {
				name, _ := NewName("Admin User")
				email, _ := NewEmail("admin@example.com")
				user := NewUser(name, email)
				user.SetRole(RoleAdmin)
				return user
			},
			expectedIsAdmin: true,
		},
		{
			testName: "user is not admin",
			setup: func() *User {
				name, _ := NewName("Regular User")
				email, _ := NewEmail("user@example.com")
				return NewUser(name, email)
			},
			expectedIsAdmin: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			user := tt.setup()
			assert.Equal(t, tt.expectedIsAdmin, user.IsAdmin())
		})
	}
}

func TestIsOwner(t *testing.T) {
	tests := []struct {
		setup           func() *User
		testName        string
		expectedIsOwner bool
	}{
		{
			testName: "user is owner",
			setup: func() *User {
				name, _ := NewName("Owner User")
				email, _ := NewEmail("owner@example.com")
				u := NewUser(name, email)
				u.SetRole(RoleOwner)
				return u
			},
			expectedIsOwner: true,
		},
		{
			testName: "user is not owner",
			setup: func() *User {
				name, _ := NewName("Regular User")
				email, _ := NewEmail("user@example.com")
				return NewUser(name, email)
			},
			expectedIsOwner: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			user := tt.setup()
			assert.Equal(t, tt.expectedIsOwner, user.IsOwner())
		})
	}
}
