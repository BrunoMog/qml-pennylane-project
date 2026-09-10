package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/user"
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTrainConfig(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(f *testFixture) CreateTrainConfigInput
		testName      string
	}{
		{
			testName: "create TrainConfig successfully",
			setup: func(f *testFixture) CreateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				return CreateTrainConfigInput{
					Name:        "Test Train Config",
					Description: "Test Train Config Description",
					TrainingDTO: validTrainDTO(),
					CallerID:    user.ID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) CreateTrainConfigInput {
				return CreateTrainConfigInput{
					Name:        "Test Train Config",
					Description: "Test Train Config Description",
					TrainingDTO: validTrainDTO(),
					CallerID:    uuid.New(),
				}
			},
			expectedError: &UserNotFoundError{},
		},
		{
			testName: "duplicate TrainConfig name",
			setup: func(f *testFixture) CreateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				trainConfig := f.createTrainConfig(user.ID())
				name, err := trainconfig.NewName("Duplicate Name")
				require.NoError(t, err)
				trainConfig.SetName(name)
				return CreateTrainConfigInput{
					Name:        "Duplicate Name",
					Description: "Test Train Config Description",
					TrainingDTO: validTrainDTO(),
					CallerID:    user.ID(),
				}
			},
			expectedError: &TrainConfigNameAlreadyExistsError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)
			trainConfig, err := fixture.service.CreateTrainConfig(input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, trainConfig)
				assert.Equal(t, input.Name, trainConfig.Name)
				assert.Equal(t, input.Description, trainConfig.Description)
				assert.NotZero(t, trainConfig.CreatedAt)
			}
		})
	}

}
