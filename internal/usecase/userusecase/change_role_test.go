package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestChangeUserRole(t *testing.T) {
	tests := []struct {
		name        string
		newRole     user.Role
		usersToSeed []testkit.UserSeed
		userCaller  uint8
		userTarget  uint8
		expectErr   bool
	}{
		{
			name: "change role of existing user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleOwner},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
			},
			userCaller: 1,
			userTarget: 2,
			newRole:    user.RoleAdmin,
			expectErr:  false,
		},
		{
			name: "change role of non-existing user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleOwner},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
			},
			userCaller: 1,
			userTarget: 3,
			newRole:    user.RoleAdmin,
			expectErr:  true,
		},
		{
			name: "change role with non-existing caller",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleOwner},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
			},
			userCaller: 3,
			userTarget: 2,
			newRole:    user.RoleAdmin,
			expectErr:  true,
		},
		{
			name: "change role with insufficient permissions",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleUser},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
			},
			userCaller: 1,
			userTarget: 2,
			newRole:    user.RoleAdmin,
			expectErr:  true,
		},
		{
			name: "change role of owner user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleOwner},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleAdmin},
			},
			userCaller: 2,
			userTarget: 1,
			newRole:    user.RoleAdmin,
			expectErr:  true,
		},
		{
			name: "change user role to invalid role",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleOwner},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
			},
			userCaller: 1,
			userTarget: 2,
			newRole:    user.Role("invalid_role"),
			expectErr:  true,
		},
		{
			name: "admin trying to change role of user to owner",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com", Role: user.RoleAdmin},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Role: user.RoleUser},
			},
			userCaller: 1,
			userTarget: 2,
			newRole:    user.RoleOwner,
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testkit.NewMockUserRepository()
			service := NewUserService(repo)

			seedResult, err := testkit.SeedUsers(repo, tt.usersToSeed)
			if err != nil {
				t.Fatalf("Failed to seed users: %v", err)
			}

			caller, ok := seedResult.ByRef[tt.userCaller]

			var callerID uuid.UUID
			if !ok || caller == nil {
				callerID = uuid.New() // Generate a random UUID for non-existing caller
			} else {
				callerID = caller.ID()
			}

			target, ok := seedResult.ByRef[tt.userTarget]

			var targetID uuid.UUID
			if !ok || target == nil {
				targetID = uuid.New() // Generate a random UUID for non-existing target
			} else {
				targetID = target.ID()
			}

			input := ChangeUserRoleInput{
				CallerID: callerID,
				TargetID: targetID,
				Role:     tt.newRole,
			}

			err = service.ChangeUserRole(input)
			if (err != nil) != tt.expectErr {
				t.Errorf("ChangeUserRole() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
