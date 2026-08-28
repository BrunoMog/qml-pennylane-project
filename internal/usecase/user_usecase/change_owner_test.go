package user_usecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestChangeOwner(t *testing.T) {
	tests := []struct {
		name          string
		usersToCreate []testkit.UserSeed
		callerRef     uint8
		targetRef     uint8
		wantErr       bool
	}{
		{
			name: "Owner can change owner to another user",
			usersToCreate: []testkit.UserSeed{
				{Ref: 1, Name: "Owner", Email: "owner@example.com", Role: user.RoleOwner},
				{Ref: 2, Name: "User", Email: "user@example.com", Role: user.RoleUser},
			},
			callerRef: 1,
			targetRef: 2,
			wantErr:   false,
		},
		{
			name: "Admin cannot change owner",
			usersToCreate: []testkit.UserSeed{
				{Ref: 1, Name: "Owner", Email: "owner@example.com", Role: user.RoleOwner},
				{Ref: 2, Name: "Admin", Email: "admin@example.com", Role: user.RoleAdmin},
			},
			callerRef: 2,
			targetRef: 1,
			wantErr:   true,
		},
		{
			name: "Owner cannot change owner to themselves",
			usersToCreate: []testkit.UserSeed{
				{Ref: 1, Name: "Owner", Email: "owner@example.com", Role: user.RoleOwner},
			},
			callerRef: 1,
			targetRef: 1,
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

			input := ChangeOwnerInput{
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
			err = service.ChangeOwner(input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeOwner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
