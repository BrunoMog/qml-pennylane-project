package user_usecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name      string
		userName  string
		userEmail string
		expectErr bool
	}{
		{
			name:      "create user with valid name and email",
			userName:  "John Doe",
			userEmail: "john.doe@example.com",
			expectErr: false,
		},
		{
			name:      "create invalid user with empty name",
			userName:  "",
			userEmail: "john.doe@example.com",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testkit.NewMockUserRepository()
			service := NewUserService(repo)

			service_output, err := service.CreateUser(tt.userName, tt.userEmail)

			if (err != nil) != tt.expectErr {
				t.Errorf("CreateUser() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr {
				if service_output.Name != tt.userName {
					t.Errorf("Expected user name %s, got %s", tt.userName, service_output.Name)
				}
				if service_output.Email != tt.userEmail {
					t.Errorf("Expected user email %s, got %s", tt.userEmail, service_output.Email)
				}
				if service_output.Role != user.RoleUser {
					t.Errorf("Expected user role %s, got %s", user.RoleUser, service_output.Role)
				}
			}
		})
	}
}
