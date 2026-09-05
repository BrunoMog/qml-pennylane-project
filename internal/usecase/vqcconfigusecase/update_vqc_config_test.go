package vqcconfigusecase

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"

	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/testkit"
)

func strPtr(s string) *string {
	return &s
}

func TestUpdateVQCConfig(t *testing.T) {
	dummyVQC1 := &vqc.VQC{}
	dummyVQC2 := &vqc.VQC{}

	tests := []struct {
		name             string
		usersToSeed      []testkit.UserSeed
		vqcConfigsToSeed []testkit.VQCConfigSeed
		callerRef        uint8
		vqcConfigRef     uint8
		newName          *string
		newDescription   *string
		newVQC           *vqc.VQC
		expectError      bool
	}{
		{
			name: "successfully update all fields",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial VQCConfig", Description: "Initial description", VQC: dummyVQC1},
			},
			callerRef:      1,
			vqcConfigRef:   1,
			newName:        strPtr("Updated VQCConfig"),
			newDescription: strPtr("Updated description"),
			newVQC:         dummyVQC2,
			expectError:    false,
		},
		{
			name: "successfully update only name",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial VQCConfig", Description: "Initial description", VQC: dummyVQC1},
			},
			callerRef:      1,
			vqcConfigRef:   1,
			newName:        strPtr("Only Name Updated"),
			newDescription: nil,
			newVQC:         nil,
			expectError:    false,
		},
		{
			name: "successfully update only description",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial VQCConfig", Description: "Initial description", VQC: dummyVQC1},
			},
			callerRef:      1,
			vqcConfigRef:   1,
			newName:        nil,
			newDescription: strPtr("Only Description Updated"),
			newVQC:         nil,
			expectError:    false,
		},
		{
			name: "successfully update only VQC",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial VQCConfig", Description: "Initial description", VQC: dummyVQC1},
			},
			callerRef:      1,
			vqcConfigRef:   1,
			newName:        nil,
			newDescription: nil,
			newVQC:         dummyVQC2,
			expectError:    false,
		},
		{
			name: "successfully update with no fields provided",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial VQCConfig", Description: "Initial description", VQC: dummyVQC1},
			},
			callerRef:      1,
			vqcConfigRef:   1,
			newName:        nil,
			newDescription: nil,
			newVQC:         nil,
			expectError:    false,
		},
		{
			name: "fail when caller does not exist",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial VQCConfig", Description: "Initial description", VQC: dummyVQC1},
			},
			callerRef:      2,
			vqcConfigRef:   1,
			newName:        strPtr("Updated Name"),
			expectError:    true,
		},
		{
			name: "fail when target VQCConfig does not exist",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{},
			callerRef:        1,
			vqcConfigRef:     1,
			newName:          strPtr("Updated Name"),
			expectError:      true,
		},
		{
			name: "fail when unauthorized user tries to update",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
				{Ref: 2, Name: "Bob", Email: "bob@example.com", Role: user.RoleUser},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Alice's VQCConfig", Description: "Alice's desc", VQC: dummyVQC1},
			},
			callerRef:      2,
			vqcConfigRef:   1,
			newName:        strPtr("Hacked Name"),
			expectError:    true,
		},
		{
			name: "fail with empty name",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial VQCConfig", Description: "Initial description", VQC: dummyVQC1},
			},
			callerRef:    1,
			vqcConfigRef: 1,
			newName:      strPtr(""),
			expectError:  true,
		},
		{
			name: "fail with name exceeding max length",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial VQCConfig", Description: "Initial description", VQC: dummyVQC1},
			},
			callerRef:    1,
			vqcConfigRef: 1,
			newName:      strPtr(strings.Repeat("a", 101)),
			expectError:  true,
		},
		{
			name: "fail with description exceeding max length",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial VQCConfig", Description: "Initial description", VQC: dummyVQC1},
			},
			callerRef:      1,
			vqcConfigRef:   1,
			newDescription: strPtr(strings.Repeat("d", 501)),
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := testkit.NewMockUserRepository()
			vqcConfigRepo := testkit.NewMockVQCConfigRepository()

			userSeedResult, err := testkit.SeedUsers(userRepo, tt.usersToSeed)
			if err != nil {
				t.Fatalf("Failed to seed users: %v", err)
			}

			vqcConfigSeedResult, err := testkit.SeedVQCConfigs(vqcConfigRepo, userSeedResult, tt.vqcConfigsToSeed)
			if err != nil {
				t.Fatalf("Failed to seed VQCConfigs: %v", err)
			}

			var callerID uuid.UUID
			caller, exists := userSeedResult.ByRef[tt.callerRef]
			if !exists {
				callerID = uuid.New()
			} else {
				callerID = caller.ID()
			}

			var vqcConfigID uuid.UUID
			vqcConfig, exists := vqcConfigSeedResult.ByRef[tt.vqcConfigRef]
			if !exists {
				vqcConfigID = uuid.New()
			} else {
				vqcConfigID = vqcConfig.VQCConfigID()
			}

			vqcConfigService := NewVQCConfigService(vqcConfigRepo, userRepo)

			input := UpdateVQCConfigInput{
				CallerID:    callerID,
				VQCConfigID: vqcConfigID,
				Name:        tt.newName,
				Description: tt.newDescription,
				VQC:         tt.newVQC,
			}

			err = vqcConfigService.UpdateVQCConfig(input)
			if (err != nil) != tt.expectError {
				t.Fatalf("UpdateVQCConfig() error = %v, expectError = %v", err, tt.expectError)
			}

			if !tt.expectError {
				updated, err := vqcConfigRepo.FindByID(vqcConfigID)
				if err != nil {
					t.Fatalf("Failed to find updated VQCConfig: %v", err)
				}

				if tt.newName != nil && updated.Name() != *tt.newName {
					t.Errorf("Expected Name %q, got %q", *tt.newName, updated.Name())
				}
				if tt.newDescription != nil && updated.Description() != *tt.newDescription {
					t.Errorf("Expected Description %q, got %q", *tt.newDescription, updated.Description())
				}
				if tt.newVQC != nil && !reflect.DeepEqual(updated.VQC(), *tt.newVQC) {
					t.Errorf("Expected VQC %+v, got %+v", *tt.newVQC, updated.VQC())
				}
			}
		})
	}
}
