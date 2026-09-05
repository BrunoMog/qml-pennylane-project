package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestLoadVQCConfig(t *testing.T) {
	tests := []struct {
		name             string
		usersToSeed      []testkit.UserSeed
		vqcConfigsToSeed []testkit.VQCConfigSeed
		callerRef        uint8
		vqcConfigRef     uint8
		findByID         bool
		expectError      bool
	}{
		{
			name: "successfully retrieve a VQCConfig By ID",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			callerRef: 1,
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test VQCConfig", Description: "This is a test VQCConfig", VQC: &vqc.VQC{}},
			},
			vqcConfigRef: 1,
			findByID:     true,
			expectError:  false,
		},
		{
			name: "successfully retrieve a VQCConfig By Name",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			callerRef: 1,
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test VQCConfig", Description: "This is a test VQCConfig", VQC: &vqc.VQC{}},
			},
			vqcConfigRef: 1,
			findByID:     false,
			expectError:  false,
		},
		{
			name: "fail to retrieve a non-existent VQCConfig",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			callerRef:        1,
			vqcConfigsToSeed: []testkit.VQCConfigSeed{},
			vqcConfigRef:     1,
			findByID:         true,
			expectError:      true,
		},
		{
			name: "fail to retrieve a VQCConfig with unauthorized user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleUser},
				{Ref: 2, Name: "Bob", Email: "bob@example.com", Role: user.RoleAdmin},
			},
			callerRef: 2,
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test VQCConfig", Description: "This is a test VQCConfig", VQC: &vqc.VQC{}},
			},
			vqcConfigRef: 1,
			findByID:     true,
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
				callerID = caller.ID()
			}

			var vqcConfigID uuid.UUID
			var vqcName *string
			vqcConfig, exists := vqcConfigSeedResult.ByRef[tt.vqcConfigRef]
			if !exists {
				vqcConfigID = uuid.New()
				vqcName = nil
			} else {
				vqcConfigID = vqcConfig.VQCConfigID()
				name := vqcConfig.Name()
				vqcName = &name
			}

			vqcConfigService := NewVQCConfigService(vqcConfigRepo, userRepo)

			var input LoadVQCConfigInput
			if tt.findByID {
				input = LoadVQCConfigInput{
					CallerID:    callerID,
					VQCConfigID: &vqcConfigID,
				}
			} else {
				input = LoadVQCConfigInput{
					CallerID:      callerID,
					VQCConfigName: vqcName,
				}
			}

			vqcConfigRetrived, err := vqcConfigService.LoadVQCConfig(input)
			if (err != nil) != tt.expectError {
				t.Errorf("LoadVQCConfig() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError {
				if vqcConfigRetrived.Name != vqcConfig.Name() {
					t.Errorf("Expected Name %v, got %v", vqcConfig.Name(), vqcConfigRetrived.Name)
				}
				if vqcConfigRetrived.Description != vqcConfig.Description() {
					t.Errorf("Expected Description %v, got %v", vqcConfig.Description(), vqcConfigRetrived.Description)
				}
			}
		})
	}
}

func TestLoadAllVQCConfigs(t *testing.T) {
	tests := []struct {
		name             string
		usersToSeed      []testkit.UserSeed
		vqcConfigsToSeed []testkit.VQCConfigSeed
		callerRef        uint8
		expectError      bool
	}{
		{
			name: "successfully list VQCConfigs",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			callerRef: 1,
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test VQCConfig", Description: "This is a test VQCConfig", VQC: &vqc.VQC{}},
				{Ref: 2, CallerRef: 1, Name: "Test VQCConfig 2", Description: "This is a test VQCConfig 2", VQC: &vqc.VQC{}},
				{Ref: 3, CallerRef: 1, Name: "Test VQCConfig 3", Description: "This is a test VQCConfig 3", VQC: &vqc.VQC{}},
				{Ref: 4, CallerRef: 1, Name: "Test VQCConfig 4", Description: "This is a test VQCConfig 4", VQC: &vqc.VQC{}},
			},
			expectError: false,
		},
		{
			name: "fail to list VQCConfigs with non-existent user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			callerRef: 2,
			vqcConfigsToSeed: []testkit.VQCConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test VQCConfig", Description: "This is a test VQCConfig", VQC: &vqc.VQC{}},
				{Ref: 2, CallerRef: 1, Name: "Test VQCConfig 2", Description: "This is a test VQCConfig 2", VQC: &vqc.VQC{}},
				{Ref: 3, CallerRef: 1, Name: "Test VQCConfig 3", Description: "This is a test VQCConfig 3", VQC: &vqc.VQC{}},
				{Ref: 4, CallerRef: 1, Name: "Test VQCConfig 4", Description: "This is a test VQCConfig 4", VQC: &vqc.VQC{}},
			},
			expectError: true,
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

			_, err = testkit.SeedVQCConfigs(vqcConfigRepo, userSeedResult, tt.vqcConfigsToSeed)
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

			vqcConfigService := NewVQCConfigService(vqcConfigRepo, userRepo)

			input := LoadAllVQCConfigsInput{
				CallerID: callerID,
			}

			vqcConfigsListed, err := vqcConfigService.LoadAllVQCConfigs(input)
			if (err != nil) != tt.expectError {
				t.Errorf("LoadAllVQCConfigs() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError {
				if len(vqcConfigsListed.VQCConfigs) != len(tt.vqcConfigsToSeed) {
					t.Errorf("Expected %d VQCConfigs, got %d", len(tt.vqcConfigsToSeed), len(vqcConfigsListed.VQCConfigs))
				}

				expectedByName := make(map[string]string, len(tt.vqcConfigsToSeed))
				for _, expectedVQCConfig := range tt.vqcConfigsToSeed {
					expectedByName[expectedVQCConfig.Name] = expectedVQCConfig.Description
				}

				for _, vqcConfig := range vqcConfigsListed.VQCConfigs {
					expectedDescription, exists := expectedByName[vqcConfig.Name]
					if !exists {
						t.Errorf("Unexpected VQCConfig name %q", vqcConfig.Name)
						continue
					}
					if vqcConfig.Description != expectedDescription {
						t.Errorf("Expected Description %v for %v, got %v", expectedDescription, vqcConfig.Name, vqcConfig.Description)
					}
					delete(expectedByName, vqcConfig.Name)
				}

				for name := range expectedByName {
					t.Errorf("VQCConfig %q was not returned", name)
				}
			}
		})
	}
}
