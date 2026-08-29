package vqc_config_usecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestGetVQCConfig(t *testing.T) {
	tests := []struct {
		name             string
		usersToSeed      []testkit.UserSeed
		callerRef        uint8
		vqcConfigsToSeed []testkit.VQCConfigSeed
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
				callerID = caller.GetID()
			}

			var vqcConfigID uuid.UUID
			var vqcName *string
			vqcConfig, exists := vqcConfigSeedResult.ByRef[tt.vqcConfigRef]
			if !exists {
				vqcConfigID = uuid.New()
				vqcName = nil
			} else {
				vqcConfigID = vqcConfig.GetVQCConfigID()
				name := vqcConfig.GetName()
				vqcName = &name
			}

			vqcConfigService := NewVQCConfigService(vqcConfigRepo, userRepo)

			var input GetVQCConfigInput
			if tt.findByID {
				input = GetVQCConfigInput{
					CallerID:    callerID,
					VQCConfigID: &vqcConfigID,
				}
			} else {
				input = GetVQCConfigInput{
					CallerID:      callerID,
					VQCConfigName: vqcName,
				}
			}

			vqcConfigRetrived, err := vqcConfigService.GetVQCConfig(input)
			if (err != nil) != tt.expectError {
				t.Errorf("GetVQCConfig() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError {
				if vqcConfigRetrived.Name != vqcConfig.GetName() {
					t.Errorf("Expected Name %v, got %v", vqcConfig.GetName(), vqcConfigRetrived.Name)
				}
				if vqcConfigRetrived.Description != vqcConfig.GetDescription() {
					t.Errorf("Expected Description %v, got %v", vqcConfig.GetDescription(), vqcConfigRetrived.Description)
				}
			}
		})
	}
}

func TestListVQCConfigs(t *testing.T) {
	tests := []struct {
		name             string
		usersToSeed      []testkit.UserSeed
		callerRef        uint8
		vqcConfigsToSeed []testkit.VQCConfigSeed
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
				callerID = caller.GetID()
			}

			vqcConfigService := NewVQCConfigService(vqcConfigRepo, userRepo)

			input := ListVQCConfigsInput{
				CallerID: callerID,
			}

			vqcConfigsListed, err := vqcConfigService.ListVQCConfigs(input)
			if (err != nil) != tt.expectError {
				t.Errorf("ListVQCConfigs() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError {
				if len(vqcConfigsListed) != len(tt.vqcConfigsToSeed) {
					t.Errorf("Expected %d VQCConfigs, got %d", len(tt.vqcConfigsToSeed), len(vqcConfigsListed))
				}
				for i, vqcConfig := range vqcConfigsListed {
					expectedVQCConfig := tt.vqcConfigsToSeed[i]
					if vqcConfig.Name != expectedVQCConfig.Name {
						t.Errorf("Expected Name %v, got %v", expectedVQCConfig.Name, vqcConfig.Name)
					}
					if vqcConfig.Description != expectedVQCConfig.Description {
						t.Errorf("Expected Description %v, got %v", expectedVQCConfig.Description, vqcConfig.Description)
					}
				}
			}
		})
	}
}
