package experimentusecase

import (
	"testing"
	"uuid"

	"pennylane_project_backend/internal/testkit"

	"github.com/stretchr/testify/assert"
)

func TestCreateExperiment(t *testing.T) {
	tests := []struct {
		testName      string
		setup         func(f *testFixture) CreateExperimentInput
		expectedError error
	}{
		{
			testName: "valid create experiment",
			setup: func(f *testFixture) CreateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())

				return CreateExperimentInput{
					Name:          "Test Experiment",
					Description:   "This is a test experiment",
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					VQCConfigID:   vqcConfig.VQCConfigID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent user",
			setup: func(f *testFixture) CreateExperimentInput {
				trainConfig := f.createTrainConfig(uuid.New())
				vqcConfig := f.createVQCConfig(uuid.New())

				return CreateExperimentInput{
					Name:          "Test Experiment",
					Description:   "This is a test experiment",
					CallerID:      uuid.New(),
					TrainConfigID: trainConfig.TrainConfigID(),
					VQCConfigID:   vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &UserNotFoundError{},
		},
		{
			testName: "other user train config",
			setup: func(f *testFixture) CreateExperimentInput {
				user := f.createUser()
				otherUser := f.createUser()
				trainConfig := f.createTrainConfig(otherUser.ID())
				vqcConfig := f.createVQCConfig(user.ID())

				return CreateExperimentInput{
					Name:          "Test Experiment",
					Description:   "This is a test experiment",
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					VQCConfigID:   vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "other user vqc config",
			setup: func(f *testFixture) CreateExperimentInput {
				user := f.createUser()
				otherUser := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(otherUser.ID())

				return CreateExperimentInput{
					Name:          "Test Experiment",
					Description:   "This is a test experiment",
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					VQCConfigID:   vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent train config",
			setup: func(f *testFixture) CreateExperimentInput {
				user := f.createUser()
				vqcConfig := f.createVQCConfig(user.ID())

				return CreateExperimentInput{
					Name:          "Test Experiment",
					Description:   "This is a test experiment",
					CallerID:      user.ID(),
					TrainConfigID: uuid.New(),
					VQCConfigID:   vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &testkit.ErrTrainConfigNotFound{},
		},
		{
			testName: "inexistent vqc config",
			setup: func(f *testFixture) CreateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())

				return CreateExperimentInput{
					Name:          "Test Experiment",
					Description:   "This is a test experiment",
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					VQCConfigID:   uuid.New(),
				}
			},
			expectedError: &testkit.ErrVQCConfigNotFound{},
		},
		{
			testName: "duplicate experiment name",
			setup: func(f *testFixture) CreateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				existingExperiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				return CreateExperimentInput{
					Name:          existingExperiment.Name().Value(),
					Description:   "This is a test experiment",
					CallerID:      user.ID(),
					TrainConfigID: trainConfig.TrainConfigID(),
					VQCConfigID:   vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &ExperimentNameAlreadyExistsError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)

			output, err := fixture.service.CreateExperiment(input)

			if tt.expectedError != nil {
				assert.IsType(t, tt.expectedError, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, input.Name, output.Name)
				assert.Equal(t, input.Description, output.Description)
				assert.NotEqual(t, uuid.Nil(), output.ExperimentID)
			}
		})
	}

}
