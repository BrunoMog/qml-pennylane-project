package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/training"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestLoadTrainConfig(t *testing.T) {
	tests := []struct {
		name               string
		usersToSeed        []testkit.UserSeed
		trainConfigsToSeed []testkit.TrainConfigSeed
		callerRef          uint8
		trainConfigRef     uint8
		findByID           bool
		expectError        bool
	}{
		{
			name: "successfully load a TrainConfig",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test TrainConfig", Description: "This is a test TrainConfig", Train: &training.Training{}},
			},
			callerRef:      1,
			trainConfigRef: 1,
			findByID:       true,
			expectError:    false,
		},
		{
			name: "successfully load a TrainConfig by name",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test TrainConfig", Description: "This is a test TrainConfig", Train: &training.Training{}},
			},
			callerRef:      1,
			trainConfigRef: 1,
			findByID:       false,
			expectError:    false,
		},
		{
			name: "fail to load a non-existent TrainConfig",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{},
			callerRef:          1,
			trainConfigRef:     1,
			findByID:           true,
			expectError:        true,
		},
		{
			name: "fail to load a TrainConfig with unauthorized user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleUser},
				{Ref: 2, Name: "Bob", Email: "bob@example.com", Role: user.RoleUser},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test TrainConfig", Description: "This is a test TrainConfig", Train: &training.Training{}},
			},
			callerRef:      2,
			trainConfigRef: 1,
			findByID:       true,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := testkit.NewMockUserRepository()
			trainConfigRepo := testkit.NewMockTrainConfigRepository()

			userSeedResult, err := testkit.SeedUsers(userRepo, tt.usersToSeed)
			if err != nil {
				t.Fatalf("Failed to seed users: %v", err)
			}

			trainConfigSeedResult, err := testkit.SeedTrainConfigs(trainConfigRepo, userSeedResult, tt.trainConfigsToSeed)
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

			var trainConfigID uuid.UUID
			var trainConfigName *string
			trainConfig, exists := trainConfigSeedResult.ByRef[tt.trainConfigRef]
			if !exists {
				trainConfigID = uuid.New()
				trainConfigName = nil
			} else {
				trainConfigID = trainConfig.TrainConfigID()
				name := trainConfig.Name()
				trainConfigName = &name
			}

			trainConfigService := NewTrainConfigService(trainConfigRepo, userRepo)

			var input LoadTrainConfigInput
			if tt.findByID {
				input = LoadTrainConfigInput{
					CallerID:        callerID,
					TrainConfigID:   &trainConfigID,
					TrainConfigName: nil,
				}
			} else {
				input = LoadTrainConfigInput{
					CallerID:        callerID,
					TrainConfigID:   nil,
					TrainConfigName: trainConfigName,
				}
			}

			trainConfigRetrieved, err := trainConfigService.LoadTrainConfig(input)
			if (err != nil) != tt.expectError {
				t.Errorf("LoadTrainConfig() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError {
				if trainConfigRetrieved.Name != trainConfig.Name() {
					t.Errorf("Expected TrainConfig name %q, got %q", trainConfig.Name(), trainConfigRetrieved.Name)
				}
				if trainConfigRetrieved.Description != trainConfig.Description() {
					t.Errorf("Expected TrainConfig description %q, got %q", trainConfig.Description(), trainConfigRetrieved.Description)
				}
			}
		})
	}
}

func TestLoadAllTrainConfigs(t *testing.T) {
	tests := []struct {
		name               string
		usersToSeed        []testkit.UserSeed
		trainConfigsToSeed []testkit.TrainConfigSeed
		callerRef          uint8
		expectError        bool
	}{
		{
			name: "successfully load all TrainConfigs",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleUser},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test TrainConfig 1", Description: "This is a test TrainConfig 1", Train: &training.Training{}},
				{Ref: 2, CallerRef: 1, Name: "Test TrainConfig 2", Description: "This is a test TrainConfig 2", Train: &training.Training{}},
			},
			callerRef:   1,
			expectError: false,
		},
		{
			name: "fail to load all TrainConfigs with non-existent user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleUser},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test TrainConfig 1", Description: "This is a test TrainConfig 1", Train: &training.Training{}},
				{Ref: 2, CallerRef: 1, Name: "Test TrainConfig 2", Description: "This is a test TrainConfig 2", Train: &training.Training{}},
			},
			callerRef:   2,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := testkit.NewMockUserRepository()
			trainConfigRepo := testkit.NewMockTrainConfigRepository()

			userSeedResult, err := testkit.SeedUsers(userRepo, tt.usersToSeed)
			if err != nil {
				t.Fatalf("Failed to seed users: %v", err)
			}

			_, err = testkit.SeedTrainConfigs(trainConfigRepo, userSeedResult, tt.trainConfigsToSeed)
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

			input := LoadAllTrainConfigsInput{
				CallerID: callerID,
			}
			trainConfigsRetrieved, err := trainConfigService.LoadAllTrainConfigs(input)
			if (err != nil) != tt.expectError {
				t.Errorf("LoadAllTrainConfigs() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError {
				if len(trainConfigsRetrieved.TrainConfigs) != len(tt.trainConfigsToSeed) {
					t.Errorf("Expected number of TrainConfigs %d, got %d", len(tt.trainConfigsToSeed), len(trainConfigsRetrieved.TrainConfigs))
				}
				expectedByName := make(map[string]string, len(tt.trainConfigsToSeed))
				for _, expectedTrainConfig := range tt.trainConfigsToSeed {
					expectedByName[expectedTrainConfig.Name] = expectedTrainConfig.Description
				}

				for _, trainConfig := range trainConfigsRetrieved.TrainConfigs {
					expectedDescription, exists := expectedByName[trainConfig.Name]
					if !exists {
						t.Errorf("Unexpected TrainConfig name %q", trainConfig.Name)
					} else if trainConfig.Description != expectedDescription {
						t.Errorf("Expected TrainConfig description %q for name %q, got %q", expectedDescription, trainConfig.Name, trainConfig.Description)
					}
				}
			}
		})
	}
}
