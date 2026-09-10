package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"strings"
	"testing"

	"uuid"

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
				newTrainingDTO := validTrainDTO()
				return UpdateTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					Name:          &newName,
					Description:   &newDescription,
					TrainingDTO:   &newTrainingDTO,
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
				newName := trainConfig2.Name().Value()
				return UpdateTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: trainConfig1.TrainConfigID(),
					Name:          &newName,
					Description:   nil,
				}
			},
			expectedError: &TrainConfigNameAlreadyExistsError{},
		},
		{
			testName: "cosmetic update name",
			setup: func(f *testFixture) UpdateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				name := strings.ToUpper(trainConfig.Name().Value())
				return UpdateTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					Name:          &name,
					Description:   nil,
				}
			},
			expectedError: nil,
		},
		{
			testName: "idepotent update with same name and description and training",
			setup: func(f *testFixture) UpdateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				name := trainConfig.Name().Value()
				description := trainConfig.Description().Value()
				trainingDTO := buildTrainingDTOFromTraining(trainConfig.Training())
				return UpdateTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					Name:          &name,
					Description:   &description,
					TrainingDTO:   &trainingDTO,
				}
			},
			expectedError: nil,
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
					assert.Equal(t, *input.Name, updatedTrainConfig.Name().Value())
				}
				if input.Description != nil {
					assert.Equal(t, *input.Description, updatedTrainConfig.Description().Value())
				}
				if input.TrainingDTO != nil {
					expectedTraining, err := buildTrainingFromDTO(*input.TrainingDTO)
					require.NoError(t, err)
					assert.True(t, expectedTraining.Equals(updatedTrainConfig.Training()))
				}
				assert.Equal(t, input.CallerID, updatedTrainConfig.OwnerID())
				assert.Equal(t, input.TrainConfigID, updatedTrainConfig.TrainConfigID())
			}
		})
	}
}
