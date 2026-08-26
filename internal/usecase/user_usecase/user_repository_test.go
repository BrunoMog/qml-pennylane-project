package user_usecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"

	"github.com/google/uuid"

	"testing"
)

func TestSaveUser(t *testing.T) {
	repo := testkit.NewMockUserRepository()
	user, err := user.NewUser("John Doe", "john.doe@example.com")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	err = repo.Save(user)
	if err != nil {
		t.Fatalf("Failed to save user: %v", err)
	}
}

func TestFindUserByID(t *testing.T) {
	tests := []struct {
		name          string
		usersToSeed   []testkit.UserSeed
		refUserToFind uint8
		expectErr     bool
	}{
		{
			name: "find existing user by ID",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleUser},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
				{Ref: 3, Name: "Bob Johnson", Email: "bob.johnson@example.com", Role: user.RoleUser},
			},
			refUserToFind: 1,
			expectErr:     false,
		},
		{
			name: "find non-existing user by ID",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleUser},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
			},
			refUserToFind: 3,
			expectErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testkit.NewMockUserRepository()

			seedResult, err := testkit.SeedUsers(repo, tt.usersToSeed)
			if err != nil {
				t.Fatalf("Failed to seed users: %v", err)
			}

			userToFind, ok := seedResult.ByRef[tt.refUserToFind]

			var idToFind uuid.UUID
			if !ok || userToFind == nil {
				idToFind = uuid.New()
			} else {
				idToFind = userToFind.GetID()
			}

			_, err = repo.FindByID(idToFind)
			if (err != nil) != tt.expectErr {
				t.Errorf("FindByID() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

func TestFindUserByEmail(t *testing.T) {
	tests := []struct {
		name        string
		usersToSeed []testkit.UserSeed
		emailToFind string
		expectErr   bool
	}{
		{
			name: "find existing user by email",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleUser},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
			},
			emailToFind: "john.doe@example.com",
			expectErr:   false,
		},
		{
			name: "find non-existing user by email",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleUser},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
			},
			emailToFind: "bob.johnson@example.com",
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testkit.NewMockUserRepository()

			_, err := testkit.SeedUsers(repo, tt.usersToSeed)
			if err != nil {
				t.Fatalf("Failed to seed users: %v", err)
			}

			_, err = repo.FindByEmail(tt.emailToFind)
			if (err != nil) != tt.expectErr {
				t.Errorf("FindByEmail() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
