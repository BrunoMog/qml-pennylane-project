package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateTrainConfig(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(f *testFixture) UpdateTrainConfigInput
		testName      string
	}{
		{
			testName: "update TrainConfig successfully",
			setup: func(f *testFixture) UpdateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				newName := "Updated Name"
				newDescription := "Updated Description"
				return UpdateTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					Name:          &newName,
					Description:   &newDescription,
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) UpdateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				newName := "Updated Name"
				return UpdateTrainConfigInput{
					CallerID:      uuid.New(),
					TrainConfigID: trainConfig.TrainConfigID(),
					Name:          &newName,
					Description:   nil,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent TrainConfig",
			setup: func(f *testFixture) UpdateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				newName := "Updated Name"
				return UpdateTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: uuid.New(),
					Name:          &newName,
					Description:   nil,
				}
			},
			expectedError: &testkit.ErrTrainConfigNotFound{},
		},
		{
			testName: "unauthorized user",
			setup: func(f *testFixture) UpdateTrainConfigInput {
				owner := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				newName := "Updated Name"
				return UpdateTrainConfigInput{
					CallerID:      unauthorizedUser.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					Name:          &newName,
					Description:   nil,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "no fields to update",
			setup: func(f *testFixture) UpdateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				return UpdateTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					Name:          nil,
					Description:   nil,
				}
			},
			expectedError: &NoFieldsToUpdateError{},
		},
		{
			testName: "try to update with existing name",
			setup: func(f *testFixture) UpdateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig1 := f.createTrainConfig(user.ID())
				trainConfig2 := f.createTrainConfig(user.ID())
				newName := trainConfig2.Name()
				return UpdateTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: trainConfig1.TrainConfigID(),
					Name:          &newName,
					Description:   nil,
				}
			},
			expectedError: &TrainConfigNameAlreadyExistsError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)
			input := tt.setup(f)
			err := f.service.UpdateTrainConfig(input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				updatedTrainConfig, err := f.trainConfigRepo.FindByID(input.TrainConfigID)
				require.NoError(t, err)
				if input.Name != nil {
					assert.Equal(t, *input.Name, updatedTrainConfig.Name())
				}
				if input.Description != nil {
					assert.Equal(t, *input.Description, updatedTrainConfig.Description())
				}
				assert.Equal(t, input.CallerID, updatedTrainConfig.OwnerID())
				assert.Equal(t, input.TrainConfigID, updatedTrainConfig.TrainConfigID())
			}
		})
	}
}
