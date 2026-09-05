package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/training"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
)

func TestDeleteTrainConfig(t *testing.T) {
	tests := []struct {
		name               string
		usersToSeed        []testkit.UserSeed
		trainConfigsToSeed []testkit.TrainConfigSeed
		callerRef          uint8
		trainConfigRef     uint8
		expectError        bool
	}{
		{
			name: "successfully delete a TrainConfig",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test TrainConfig", Description: "This is a test TrainConfig", Train: &training.Training{}},
			},
			callerRef:      1,
			trainConfigRef: 1,
			expectError:    false,
		},
		{
			name: "fail to delete a non-existent TrainConfig",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{},
			callerRef:          1,
			trainConfigRef:     1,
			expectError:        true,
		},
		{
			name: "fail to delete a TrainConfig with unauthorized user",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleUser},
				{Ref: 2, Name: "Bob", Email: "bob@example.com", Role: user.RoleUser},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Test TrainConfig", Description: "This is a test TrainConfig", Train: &training.Training{}},
			},
			callerRef:      2,
			trainConfigRef: 1,
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
			trainConfig, exists := trainConfigSeedResult.ByRef[tt.trainConfigRef]
			if !exists {
				trainConfigID = uuid.New()
			} else {
				trainConfigID = trainConfig.TrainConfigID()
			}

			trainConfigService := NewTrainConfigService(trainConfigRepo, userRepo)

			input := DeleteTrainConfigInput{
				CallerID:      callerID,
				TrainConfigID: trainConfigID,
			}

			err = trainConfigService.DeleteTrainConfig(input)
			if (err != nil) != tt.expectError {
				t.Errorf("DeleteTrainConfig() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !tt.expectError {
				// Verify that the TrainConfig has been deleted
				exists, err := trainConfigRepo.ExistsByID(trainConfigID)
				if err != nil {
					t.Fatalf("Failed to check existence of TrainConfig: %v", err)
				}
				if exists {
					t.Fatalf("Expected TrainConfig to be deleted, but it still exists")
				}
			}
		})
	}
}
