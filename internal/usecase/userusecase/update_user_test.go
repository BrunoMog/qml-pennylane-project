package userusecase

import (
	"github.com/google/uuid"

	"pennylane_project_backend/internal/testkit"
	"testing"
)

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		newName       *string
		newEmail      *string
		name          string
		usersToSeed   []testkit.UserSeed
		callerUserRef uint8
		targetUserRef uint8
		expectErr     bool
	}{
		{
			name: "update existing user with valid name and email",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com"},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
			},
			callerUserRef: 1,
			targetUserRef: 1,
			newName:       strPtr("Jane Doe"),
			newEmail:      strPtr("jane.doe@example.com"),
			expectErr:     false,
		},
		{
			name: "update non-existing user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com"},
			},
			callerUserRef: 1,
			targetUserRef: 2,
			newName:       strPtr("Jane Doe"),
			newEmail:      strPtr("jane.doe@example.com"),
			expectErr:     true,
		},
		{
			name: "update only user name",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com"},
			},
			callerUserRef: 1,
			targetUserRef: 1,
			newName:       strPtr("Jane Doe"),
			newEmail:      nil,
			expectErr:     false,
		},
		{
			name: "update only user email",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com"},
			},
			callerUserRef: 1,
			targetUserRef: 1,
			newName:       nil,
			newEmail:      strPtr("jane.doe@example.com"),
			expectErr:     false,
		},
		{
			name: "unauthorized user trying to update another user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "John Doe", Email: "john.doe@example.com"},
				{Ref: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
			},
			callerUserRef: 1,
			targetUserRef: 2,
			newName:       strPtr("Jane Doe"),
			newEmail:      strPtr("jane.doe@example.com"),
			expectErr:     true,
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

			callerUser, ok := seedResult.ByRef[tt.callerUserRef]

			var callerID uuid.UUID
			if !ok || callerUser == nil {
				callerID = uuid.New()
			} else {
				callerID = callerUser.ID()
			}

			targetUser, ok := seedResult.ByRef[tt.targetUserRef]

			var targetID uuid.UUID
			if !ok || targetUser == nil {
				targetID = uuid.New()
			} else {
				targetID = targetUser.ID()
			}

			inpput := UpdateUserInput{
				CallerID: callerID,
				TargetID: targetID,
				Name:     tt.newName,
				Email:    tt.newEmail,
			}

			err = service.UpdateUser(inpput)

			if (err != nil) != tt.expectErr {
				t.Errorf("UpdateUser() error = %v, expectErr %v", err, tt.expectErr)
				return
			}

			if !tt.expectErr {
				updatedUser, err := repo.FindByID(callerID)
				if err != nil {
					t.Fatalf("Failed to retrieve updated user: %v", err)
				}

				if tt.newName != nil && updatedUser.Name() != *tt.newName {
					t.Errorf("Expected user name %s, got %s", *tt.newName, updatedUser.Name())
				}
				if tt.newEmail != nil && updatedUser.Email() != *tt.newEmail {
					t.Errorf("Expected user email %s, got %s", *tt.newEmail, updatedUser.Email())
				}
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}
