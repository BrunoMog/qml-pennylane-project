package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/training"
	"pennylane_project_backend/internal/domain/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
					Training:    &training.Training{},
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
					Training:    &training.Training{},
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
				trainConfig.SetName("Duplicate Name")
				return CreateTrainConfigInput{
					Name:        "Duplicate Name",
					Description: "Test Train Config Description",
					Training:    &training.Training{},
					CallerID:    user.ID(),
				}
			},
			expectedError: &TrainConfigNameAlreadyExistsError{},
		},
		{
			testName: "nil Training input",
			setup: func(f *testFixture) CreateTrainConfigInput {
				user := f.createUser(user.RoleUser)
				return CreateTrainConfigInput{
					Name:        "Test Train Config",
					Description: "Test Train Config Description",
					Training:    nil,
					CallerID:    user.ID(),
				}
			},
			expectedError: &trainconfig.TrainingMissingError{},
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
			}
		})
	}

}
