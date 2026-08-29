package vqc_config_usecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestCreateVQCConfig(t *testing.T) {
	tests := []struct {
		name              string
		userToSeed        []testkit.UserSeed
		callerRef         uint8
		vqcConfigToSeed   []testkit.VQCConfigSeed
		vqcConfigToCreate testkit.VQCConfigSeed
		expectError       bool
	}{
		{
			name: "valid VQCConfig creation",
			userToSeed: []testkit.UserSeed{
				{
					Ref:   1,
					Name:  "Test User",
					Email: "test@example.com",
					Role:  user.RoleAdmin,
				},
			},
			callerRef:       1,
			vqcConfigToSeed: []testkit.VQCConfigSeed{},
			vqcConfigToCreate: testkit.VQCConfigSeed{
				Ref:         1,
				CallerRef:   1,
				Name:        "Test VQCConfig",
				Description: "This is a test VQCConfig",
				VQC:         &vqc.VQC{}, // Assuming a valid VQC object
			},
			expectError: false,
		},
		{
			name: "invalid VQCConfig name",
			userToSeed: []testkit.UserSeed{
				{
					Ref:   1,
					Name:  "Test User",
					Email: "test@example.com",
					Role:  user.RoleAdmin,
				},
			},
			callerRef:       1,
			vqcConfigToSeed: []testkit.VQCConfigSeed{},
			vqcConfigToCreate: testkit.VQCConfigSeed{
				Ref:         1,
				CallerRef:   1,
				Name:        "",
				Description: "This is a test VQCConfig",
				VQC:         &vqc.VQC{}, // Assuming a valid VQC object
			},
			expectError: true,
		},
		{
			name: "invalid VQCConfig description",
			userToSeed: []testkit.UserSeed{
				{
					Ref:   1,
					Name:  "Test User",
					Email: "test@example.com",
					Role:  user.RoleAdmin,
				},
			},
			callerRef:       1,
			vqcConfigToSeed: []testkit.VQCConfigSeed{},
			vqcConfigToCreate: testkit.VQCConfigSeed{
				Ref:         1,
				CallerRef:   1,
				Name:        "Test VQCConfig",
				Description: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				VQC:         &vqc.VQC{}, // Assuming a valid VQC object
			},
			expectError: true,
		},
		{
			name: "inexistent caller",
			userToSeed: []testkit.UserSeed{
				{
					Ref:   1,
					Name:  "Test User",
					Email: "test@example.com",
					Role:  user.RoleAdmin,
				},
			},
			callerRef:       2,
			vqcConfigToSeed: []testkit.VQCConfigSeed{},
			vqcConfigToCreate: testkit.VQCConfigSeed{
				CallerRef:   2,
				Ref:         1,
				Name:        "Test VQCConfig",
				Description: "This is a test VQCConfig",
				VQC:         &vqc.VQC{}, // Assuming a valid VQC object
			},
			expectError: true,
		},
		{
			name: "duplicate VQCConfig name",
			userToSeed: []testkit.UserSeed{
				{
					Ref:   1,
					Name:  "Test User",
					Email: "test@example.com",
					Role:  user.RoleAdmin,
				},
			},
			callerRef: 1,
			vqcConfigToSeed: []testkit.VQCConfigSeed{
				{
					Ref:         1,
					CallerRef:   1,
					Name:        "Test VQCConfig",
					Description: "This is a test VQCConfig",
					VQC:         &vqc.VQC{}, // Assuming a valid VQC object
				},
			},
			vqcConfigToCreate: testkit.VQCConfigSeed{
				Ref:         2,
				Name:        "Test VQCConfig",
				Description: "This is a test VQCConfig",
				VQC:         &vqc.VQC{}, // Assuming a valid VQC object
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := testkit.NewMockUserRepository()
			vqcConfigRepo := testkit.NewMockVQCConfigRepository()

			userSeedResult, err := testkit.SeedUsers(userRepo, tt.userToSeed)
			if err != nil {
				t.Fatalf("Failed to seed users: %v", err)
			}

			_, err = testkit.SeedVQCConfigs(vqcConfigRepo, userSeedResult, tt.vqcConfigToSeed)
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

			input := CreateVQCConfigInput{
				CallerID:    callerID,
				Name:        &tt.vqcConfigToCreate.Name,
				Description: &tt.vqcConfigToCreate.Description,
				VQC:         tt.vqcConfigToCreate.VQC,
			}

			output, err := vqcConfigService.CreateVQCConfig(input)
			if (err != nil) != tt.expectError {
				t.Errorf("CreateVQCConfig() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError {
				if output.Name != tt.vqcConfigToCreate.Name {
					t.Errorf("Expected name %v, got %v", tt.vqcConfigToCreate.Name, output.Name)
				}
				if output.Description != tt.vqcConfigToCreate.Description {
					t.Errorf("Expected description %v, got %v", tt.vqcConfigToCreate.Description, output.Description)
				}
			}
		})
	}
}
