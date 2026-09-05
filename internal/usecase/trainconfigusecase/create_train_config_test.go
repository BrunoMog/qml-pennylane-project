package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/training"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestCreateTrainConfig(t *testing.T) {
	tests := []struct {
		testConfigToCreate testkit.TrainConfigSeed
		name               string
		userToSeed         []testkit.UserSeed
		trainConfigToSeed  []testkit.TrainConfigSeed
		callerRef          uint8
		expectError        bool
	}{
		{
			name: "valid TrainConfig creation",
			userToSeed: []testkit.UserSeed{
				{
					Ref:   1,
					Name:  "Test User",
					Email: "test@example.com",
					Role:  user.RoleAdmin,
				},
			},
			callerRef:         1,
			trainConfigToSeed: []testkit.TrainConfigSeed{},
			testConfigToCreate: testkit.TrainConfigSeed{
				Ref:         1,
				CallerRef:   1,
				Name:        "Test TrainConfig",
				Description: "This is a test TrainConfig",
				Train:       &training.Training{}, // Assuming a valid Training object
			},
			expectError: false,
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
			callerRef:         2,
			trainConfigToSeed: []testkit.TrainConfigSeed{},
			testConfigToCreate: testkit.TrainConfigSeed{
				Ref:         1,
				CallerRef:   2,
				Name:        "Test TrainConfig",
				Description: "This is a test TrainConfig",
				Train:       &training.Training{}, // Assuming a valid Training object
			},
			expectError: true,
		},
		{
			name: "duplicate TrainConfig",
			userToSeed: []testkit.UserSeed{
				{
					Ref:   1,
					Name:  "Test User",
					Email: "test@example.com",
					Role:  user.RoleAdmin,
				},
			},
			callerRef: 1,
			trainConfigToSeed: []testkit.TrainConfigSeed{
				{
					Ref:         1,
					CallerRef:   1,
					Name:        "Test TrainConfig",
					Description: "This is a test TrainConfig",
					Train:       &training.Training{}, // Assuming a valid Training object
				},
			},
			testConfigToCreate: testkit.TrainConfigSeed{
				Ref:         2,
				CallerRef:   1,
				Name:        "Test TrainConfig",
				Description: "This is a test TrainConfig",
				Train:       &training.Training{}, // Assuming a valid Training object
			},
			expectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := testkit.NewMockUserRepository()
			trainConfigRepo := testkit.NewMockTrainConfigRepository()

			userSeedResult, err := testkit.SeedUsers(userRepo, tt.userToSeed)
			if err != nil {
				t.Fatalf("Failed to seed users: %v", err)
			}

			_, err = testkit.SeedTrainConfigs(trainConfigRepo, userSeedResult, tt.trainConfigToSeed)
			if err != nil {
				t.Fatalf("Failed to seed train configs: %v", err)
			}

			var callerID uuid.UUID
			caller, exists := userSeedResult.ByRef[tt.callerRef]
			if !exists {
				callerID = uuid.New()
			} else {
				callerID = caller.ID()
			}

			trainConfigService := NewTrainConfigService(trainConfigRepo, userRepo)

			input := CreateTrainConfigInput{
				CallerID:    callerID,
				Name:        &tt.testConfigToCreate.Name,
				Description: &tt.testConfigToCreate.Description,
				Training:    tt.testConfigToCreate.Train,
			}

			output, err := trainConfigService.CreateTrainConfig(input)
			if (err != nil) != tt.expectError {
				t.Errorf("CreateTrainConfig() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError {
				// Verify that the created TrainConfig matches the input
				if output.Name != tt.testConfigToCreate.Name || output.Description != tt.testConfigToCreate.Description {
					t.Errorf("Created TrainConfig does not match input. Got: %+v, Want: %+v", output, tt.testConfigToCreate)
				}
			}
		})
	}
}
