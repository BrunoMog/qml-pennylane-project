package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTrainConfig(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(f *testFixture) DeleteTrainConfigInput
		testName      string
	}{
		{
			testName: "delete TrainConfig successfully",
			setup: func(f *testFixture) DeleteTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				return DeleteTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) DeleteTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				return DeleteTrainConfigInput{
					CallerID:      uuid.New(),
					TrainConfigID: trainConfig.TrainConfigID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent TrainConfig",
			setup: func(f *testFixture) DeleteTrainConfigInput {
				user := f.createUser(user.RoleUser)
				return DeleteTrainConfigInput{
					CallerID:      user.ID(),
					TrainConfigID: uuid.New(),
				}
			},
			expectedError: &testkit.ErrTrainConfigNotFound{},
		},
		{
			testName: "unauthorized user",
			setup: func(f *testFixture) DeleteTrainConfigInput {
				owner := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				return DeleteTrainConfigInput{
					CallerID:      unauthorizedUser.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)
			input := tt.setup(f)
			err := f.service.DeleteTrainConfig(input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
