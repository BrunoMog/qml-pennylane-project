package vqc_config_usecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestDeleteVQCConfig(t *testing.T) {
	tests := []struct {
		name             string
		usersToSeed      []testkit.UserSeed
		callerRef        uint8
		vqcConfigsToSeed []testkit.VQCConfigSeed
		vqcConfigRef     uint8
		expectError      bool
	}{
		{
			name: "successfully delete a VQCConfig",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			callerRef: 1,
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test VQCConfig", Description: "This is a test VQCConfig", VQC: &vqc.VQC{}},
				{Ref: 2, CallerRef: 1, Name: "Test VQCConfig 2", Description: "This is a test VQCConfig 2", VQC: &vqc.VQC{}},
			},
			vqcConfigRef: 1,
			expectError:  false,
		},
		{
			name: "fail to delete a non-existent VQCConfig",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			callerRef:        1,
			vqcConfigsToSeed: []testkit.VQCConfigSeed{},
			vqcConfigRef:     1,
			expectError:      true,
		},
		{
			name: "fail to delete a VQCConfig with unauthorized user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
				{Ref: 2, Name: "Bob", Email: "bob@example.com", Role: user.RoleUser},
			},
			callerRef: 2,
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test VQCConfig", Description: "This is a test VQCConfig", VQC: &vqc.VQC{}},
			},
			vqcConfigRef: 1,
			expectError:  true,
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
				callerID = caller.GetID()
			}

			var vqcConfigID uuid.UUID
			vqcConfig, exists := vqcConfigSeedResult.ByRef[tt.vqcConfigRef]
			if !exists {
				vqcConfigID = uuid.New()
			} else {
				vqcConfigID = vqcConfig.GetVQCConfigID()
			}

			vqcConfigService := NewVQCConfigService(vqcConfigRepo, userRepo)

			input := DeleteVQCConfigInput{
				CallerID:    callerID,
				VQCConfigID: vqcConfigID,
			}

			err = vqcConfigService.DeleteVQCConfig(input)
			if (err != nil) != tt.expectError {
				t.Errorf("DeleteVQCConfig() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError {
				_, err = vqcConfigRepo.FindByID(vqcConfigID)
				if err == nil {
					t.Errorf("VQCConfig with ID %v was not deleted", vqcConfigID)
				}
			}
		})
	}
}
