package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name          string
		usersToCreate []testkit.UserSeed
		callerRef     uint8
		targetRef     uint8
		wantErr       bool
	}{
		{
			name: "Owner can delete any user",
			usersToCreate: []testkit.UserSeed{
				{Ref: 1, Name: "Owner", Email: "owner@example.com", Role: user.RoleOwner},
				{Ref: 2, Name: "Admin", Email: "admin@example.com", Role: user.RoleAdmin},
			},
			callerRef: 1,
			targetRef: 2,
			wantErr:   false,
		},
		{
			name: "Admin cannot delete owner",
			usersToCreate: []testkit.UserSeed{
				{Ref: 1, Name: "Owner", Email: "owner@example.com", Role: user.RoleOwner},
				{Ref: 2, Name: "Admin", Email: "admin@example.com", Role: user.RoleAdmin},
			},
			callerRef: 2,
			targetRef: 1,
			wantErr:   true,
		},
		{
			name: "User can delete themselves",
			usersToCreate: []testkit.UserSeed{
				{Ref: 1, Name: "User", Email: "user@example.com", Role: user.RoleUser},
			},
			callerRef: 1,
			targetRef: 1,
			wantErr:   false,
		},
		{
			name: "User cannot delete another user",
			usersToCreate: []testkit.UserSeed{
				{Ref: 1, Name: "User1", Email: "user1@example.com", Role: user.RoleUser},
				{Ref: 2, Name: "User2", Email: "user2@example.com", Role: user.RoleUser},
			},
			callerRef: 1,
			targetRef: 2,
			wantErr:   true,
		},
		{
			name: "Invalid caller ID",
			usersToCreate: []testkit.UserSeed{
				{Ref: 1, Name: "User", Email: "user@example.com", Role: user.RoleUser},
			},
			callerRef: 100,
			targetRef: 1,
			wantErr:   true,
		},
		{
			name: "Invalid target ID",
			usersToCreate: []testkit.UserSeed{
				{Ref: 1, Name: "User", Email: "user@example.com", Role: user.RoleUser},
			},
			callerRef: 1,
			targetRef: 2,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testkit.NewMockUserRepository()
			seedResult, err := testkit.SeedUsers(repo, tt.usersToCreate)
			if err != nil {
				t.Fatalf("Failed to seed users: %v", err)
			}

			caller, callerExists := seedResult.ByRef[tt.callerRef]
			target, targetExists := seedResult.ByRef[tt.targetRef]

			if !callerExists && !targetExists {
				t.Fatalf("Both caller and target do not exist in the seeded users")
			}

			input := DeleteUserInput{
				CallerID: uuid.Nil,
				TargetID: uuid.Nil,
			}

			if callerExists {
				input.CallerID = caller.GetID()
			} else {
				input.CallerID = uuid.New()
			}

			if targetExists {
				input.TargetID = target.GetID()
			} else {
				input.TargetID = uuid.New()
			}

			service := NewUserService(repo)
			err = service.DeleteUser(input)

			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
