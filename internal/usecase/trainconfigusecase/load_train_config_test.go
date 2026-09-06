package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadTrainConfig(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(f *testFixture) LoadTrainConfigInput
		testName      string
	}{
		{
			testName: "load TrainConfig by ID successfully",
			setup: func(f *testFixture) LoadTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				trainConfigID := trainConfig.TrainConfigID()
				return LoadTrainConfigInput{
					CallerID:        user.ID(),
					TrainConfigID:   &trainConfigID,
					TrainConfigName: nil,
				}
			},
			expectedError: nil,
		},
		{
			testName: "load TrainConfig by name successfully",
			setup: func(f *testFixture) LoadTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				trainConfigName := trainConfig.Name()
				return LoadTrainConfigInput{
					CallerID:        user.ID(),
					TrainConfigID:   nil,
					TrainConfigName: &trainConfigName,
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) LoadTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				trainConfigID := trainConfig.TrainConfigID()
				return LoadTrainConfigInput{
					CallerID:        uuid.New(),
					TrainConfigID:   &trainConfigID,
					TrainConfigName: nil,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent TrainConfig",
			setup: func(f *testFixture) LoadTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfigID := uuid.New()
				return LoadTrainConfigInput{
					CallerID:        user.ID(),
					TrainConfigID:   &trainConfigID,
					TrainConfigName: nil,
				}
			},
			expectedError: &testkit.ErrTrainConfigNotFound{},
		},
		{
			testName: "invalid input: both TrainConfigID and TrainConfigName are nil",
			setup: func(f *testFixture) LoadTrainConfigInput {
				user := f.createUser(user.RoleUser)
				return LoadTrainConfigInput{
					CallerID:        user.ID(),
					TrainConfigID:   nil,
					TrainConfigName: nil,
				}
			},
			expectedError: &InvalidInputError{},
		},
		{
			testName: "unauthorized user",
			setup: func(f *testFixture) LoadTrainConfigInput {
				owner := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				trainConfigID := trainConfig.TrainConfigID()
				return LoadTrainConfigInput{
					CallerID:        unauthorizedUser.ID(),
					TrainConfigID:   &trainConfigID,
					TrainConfigName: nil,
				}
			},
			expectedError: &UnauthorizedError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)

			input := tt.setup(f)
			output, err := f.service.LoadTrainConfig(input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, input.CallerID, output.OwnerID)
				if input.TrainConfigID != nil {
					assert.Equal(t, input.TrainConfigID, &output.TrainConfigID)
				}
				if input.TrainConfigName != nil {
					assert.Equal(t, *input.TrainConfigName, output.Name)
				}
			}
		})
	}
}

func TestLoadAllTrainConfigs(t *testing.T) {
	tests := []struct {
		setup              func(f *testFixture) LoadAllTrainConfigsInput
		testName           string
		expectedSizeOutput int
	}{
		{
			testName: "load all TrainConfigs successfully",
			setup: func(f *testFixture) LoadAllTrainConfigsInput {
				user := f.createUser(user.RoleUser)
				f.createTrainConfig(user.ID())
				f.createTrainConfig(user.ID())
				return LoadAllTrainConfigsInput{
					CallerID: user.ID(),
				}
			},
			expectedSizeOutput: 2,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) LoadAllTrainConfigsInput {
				return LoadAllTrainConfigsInput{
					CallerID: uuid.New(),
				}
			},
			expectedSizeOutput: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)

			input := tt.setup(f)
			output, err := f.service.LoadAllTrainConfigs(input)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedSizeOutput, len(output.TrainConfigs))
		})
	}
}
