package trainconfigusecase

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"

	"pennylane_project_backend/internal/domain/training"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
)

func strPtr(s string) *string {
	return &s
}

func TestUpdateTrainConfig(t *testing.T) {
	dummyTrain1 := &training.Training{}
	dummyTrain2 := &training.Training{}

	tests := []struct {
		name               string
		usersToSeed        []testkit.UserSeed
		trainConfigsToSeed []testkit.TrainConfigSeed
		callerRef          uint8
		trainConfigRef     uint8
		newName            *string
		newDescription     *string
		newTraining        *training.Training
		expectError        bool
	}{
		{
			name: "successfully update all fields",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial TrainConfig", Description: "Initial description", Train: dummyTrain1},
			},
			callerRef:      1,
			trainConfigRef: 1,
			newName:        strPtr("Updated TrainConfig"),
			newDescription: strPtr("Updated description"),
			newTraining:    dummyTrain2,
			expectError:    false,
		},
		{
			name: "successfully update only name",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial TrainConfig", Description: "Initial description", Train: dummyTrain1},
			},
			callerRef:      1,
			trainConfigRef: 1,
			newName:        strPtr("Only Name Updated"),
			newDescription: nil,
			newTraining:    nil,
			expectError:    false,
		},
		{
			name: "successfully update only description",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial TrainConfig", Description: "Initial description", Train: dummyTrain1},
			},
			callerRef:      1,
			trainConfigRef: 1,
			newName:        nil,
			newDescription: strPtr("Only Description Updated"),
			newTraining:    nil,
			expectError:    false,
		},
		{
			name: "successfully update only training",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial TrainConfig", Description: "Initial description", Train: dummyTrain1},
			},
			callerRef:      1,
			trainConfigRef: 1,
			newName:        nil,
			newDescription: nil,
			newTraining:    dummyTrain2,
			expectError:    false,
		},
		{
			name: "successfully update with no fields provided",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial TrainConfig", Description: "Initial description", Train: dummyTrain1},
			},
			callerRef:      1,
			trainConfigRef: 1,
			newName:        nil,
			newDescription: nil,
			newTraining:    nil,
			expectError:    false,
		},
		{
			name: "fail when caller does not exist",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial TrainConfig", Description: "Initial description", Train: dummyTrain1},
			},
			callerRef:      2,
			trainConfigRef: 1,
			newName:        strPtr("Updated Name"),
			expectError:    true,
		},
		{
			name: "fail when target TrainConfig does not exist",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{},
			callerRef:          1,
			trainConfigRef:     1,
			newName:            strPtr("Updated Name"),
			expectError:        true,
		},
		{
			name: "fail when unauthorized user tries to update",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
				{Ref: 2, Name: "Bob", Email: "bob@example.com", Role: user.RoleUser},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Alice's TrainConfig", Description: "Alice's desc", Train: dummyTrain1},
			},
			callerRef:      2,
			trainConfigRef: 1,
			newName:        strPtr("Hacked Name"),
			expectError:    true,
		},
		{
			name: "fail with empty name",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial TrainConfig", Description: "Initial description", Train: dummyTrain1},
			},
			callerRef:      1,
			trainConfigRef: 1,
			newName:        strPtr(""),
			expectError:    true,
		},
		{
			name: "fail with name exceeding max length",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial TrainConfig", Description: "Initial description", Train: dummyTrain1},
			},
			callerRef:      1,
			trainConfigRef: 1,
			newName:        strPtr(strings.Repeat("a", 101)),
			expectError:    true,
		},
		{
			name: "fail with description exceeding max length",
			usersToSeed: []testkit.UserSeed{
				{Ref: 1, Name: "Alice", Email: "alice@example.com", Role: user.RoleAdmin},
			},
			trainConfigsToSeed: []testkit.TrainConfigSeed{
				{Ref: 1, CallerRef: 1, Name: "Initial TrainConfig", Description: "Initial description", Train: dummyTrain1},
			},
			callerRef:      1,
			trainConfigRef: 1,
			newDescription: strPtr(strings.Repeat("d", 501)),
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
				t.Fatalf("Failed to seed TrainConfigs: %v", err)
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

			input := UpdateTrainConfigInput{
				CallerID:      callerID,
				TrainConfigID: trainConfigID,
				Name:          tt.newName,
				Description:   tt.newDescription,
				Training:      tt.newTraining,
			}

			err = trainConfigService.UpdateTrainConfig(input)
			if (err != nil) != tt.expectError {
				t.Fatalf("UpdateTrainConfig() error = %v, expectError = %v", err, tt.expectError)
			}

			if !tt.expectError {
				updated, err := trainConfigRepo.FindByID(trainConfigID)
				if err != nil {
					t.Fatalf("Failed to find updated TrainConfig: %v", err)
				}

				if tt.newName != nil && updated.Name() != *tt.newName {
					t.Errorf("Expected Name %q, got %q", *tt.newName, updated.Name())
				}
				if tt.newDescription != nil && updated.Description() != *tt.newDescription {
					t.Errorf("Expected Description %q, got %q", *tt.newDescription, updated.Description())
				}
				if tt.newTraining != nil && !reflect.DeepEqual(updated.Training(), *tt.newTraining) {
					t.Errorf("Expected Training %+v, got %+v", *tt.newTraining, updated.Training())
				}
			}
		})
	}
}
