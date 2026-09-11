package experimentusecase

import (
	"pennylane_project_backend/internal/testkit"
	"testing"
	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteExperiment(t *testing.T) {
	tests := []struct {
		testName      string
		setup         func(f *testFixture) DeleteExperimentInput
		expectedError error
	}{
		{
			testName: "valid delete experiment",
			setup: func(f *testFixture) DeleteExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				return DeleteExperimentInput{
					ExperimentID: experiment.ExperimentID(),
					CallerID:     user.ID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent experiment",
			setup: func(f *testFixture) DeleteExperimentInput {
				user := f.createUser()
				return DeleteExperimentInput{
					ExperimentID: uuid.New(),
					CallerID:     user.ID(),
				}
			},
			expectedError: &testkit.ErrExperimentNotFound{},
		},
		{
			testName: "unauthorized delete experiment",
			setup: func(f *testFixture) DeleteExperimentInput {
				user := f.createUser()
				otherUser := f.createUser()
				trainConfig := f.createTrainConfig(otherUser.ID())
				vqcConfig := f.createVQCConfig(otherUser.ID())
				experiment := f.createExperiment(otherUser.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				return DeleteExperimentInput{
					ExperimentID: experiment.ExperimentID(),
					CallerID:     user.ID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent user",
			setup: func(f *testFixture) DeleteExperimentInput {
				trainConfig := f.createTrainConfig(uuid.New())
				vqcConfig := f.createVQCConfig(uuid.New())
				experiment := f.createExperiment(uuid.New(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				return DeleteExperimentInput{
					ExperimentID: experiment.ExperimentID(),
					CallerID:     uuid.New(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)

			err := fixture.service.DeleteExperiment(input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				exists, err := fixture.experimentRepo.ExistsByID(input.ExperimentID)
				require.NoError(t, err)
				assert.False(t, exists)
			}
		})
	}
}
