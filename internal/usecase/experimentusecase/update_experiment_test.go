package experimentusecase

import (
	"pennylane_project_backend/internal/testkit"
	"testing"
	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateExperiment(t *testing.T) {
	tests := []struct {
		testName      string
		setup         func(f *testFixture) UpdateExperimentInput
		expectedError error
	}{
		{
			testName: "valid update experiment",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				newName := "new experiment name"
				newDescription := "new experiment description"
				newVQCConfig := f.createVQCConfig(user.ID())
				newVQCConfigID := newVQCConfig.VQCConfigID()
				newTrainConfig := f.createTrainConfig(user.ID())
				newTrainConfigID := newTrainConfig.TrainConfigID()

				return UpdateExperimentInput{
					CallerID:      user.ID(),
					ExperimentID:  experiment.ExperimentID(),
					Name:          &newName,
					Description:   &newDescription,
					VQCConfigID:   &newVQCConfigID,
					TrainConfigID: &newTrainConfigID,
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent experiment",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				newName := "new experiment name"
				return UpdateExperimentInput{
					CallerID:     user.ID(),
					ExperimentID: uuid.New(),
					Name:         &newName,
				}
			},
			expectedError: &testkit.ErrExperimentNotFound{},
		},
		{
			testName: "unauthorized update experiment",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				otherUser := f.createUser()
				trainConfig := f.createTrainConfig(otherUser.ID())
				vqcConfig := f.createVQCConfig(otherUser.ID())
				experiment := f.createExperiment(otherUser.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				newName := "new experiment name"
				return UpdateExperimentInput{
					CallerID:     user.ID(),
					ExperimentID: experiment.ExperimentID(),
					Name:         &newName,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent user",
			setup: func(f *testFixture) UpdateExperimentInput {
				trainConfig := f.createTrainConfig(uuid.New())
				vqcConfig := f.createVQCConfig(uuid.New())
				experiment := f.createExperiment(uuid.New(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				newName := "new experiment name"
				return UpdateExperimentInput{
					CallerID:     uuid.New(),
					ExperimentID: experiment.ExperimentID(),
					Name:         &newName,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "no fields to update",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				return UpdateExperimentInput{
					CallerID:     user.ID(),
					ExperimentID: experiment.ExperimentID(),
				}
			},
			expectedError: &NoFieldsToUpdateError{},
		},
		{
			testName: "update with existing name",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment1 := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())
				experiment2 := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				newName := experiment2.Name().Value()
				return UpdateExperimentInput{
					CallerID:     user.ID(),
					ExperimentID: experiment1.ExperimentID(),
					Name:         &newName,
				}
			},
			expectedError: &ExperimentNameAlreadyExistsError{},
		},
		{
			testName: "cosmetic name change (case insensitive)",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				newName := "EXPERIMENT NAME"
				return UpdateExperimentInput{
					CallerID:     user.ID(),
					ExperimentID: experiment.ExperimentID(),
					Name:         &newName,
				}
			},
			expectedError: nil,
		},
		{
			testName: "update with other user's train config",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				otherUser := f.createUser()
				trainConfig := f.createTrainConfig(otherUser.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				newTrainConfig := f.createTrainConfig(otherUser.ID())
				newTrainConfigID := newTrainConfig.TrainConfigID()

				return UpdateExperimentInput{
					CallerID:      user.ID(),
					ExperimentID:  experiment.ExperimentID(),
					TrainConfigID: &newTrainConfigID,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "update with other user's vqc config",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				otherUser := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(otherUser.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				newVQCConfig := f.createVQCConfig(otherUser.ID())
				newVQCConfigID := newVQCConfig.VQCConfigID()

				return UpdateExperimentInput{
					CallerID:     user.ID(),
					ExperimentID: experiment.ExperimentID(),
					VQCConfigID:  &newVQCConfigID,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "update with inexistent train config",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				newTrainConfigID := uuid.New()

				return UpdateExperimentInput{
					CallerID:      user.ID(),
					ExperimentID:  experiment.ExperimentID(),
					TrainConfigID: &newTrainConfigID,
				}
			},
			expectedError: &testkit.ErrTrainConfigNotFound{},
		},
		{
			testName: "update with inexistent vqc config",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				newVQCConfigID := uuid.New()

				return UpdateExperimentInput{
					CallerID:     user.ID(),
					ExperimentID: experiment.ExperimentID(),
					VQCConfigID:  &newVQCConfigID,
				}
			},
			expectedError: &testkit.ErrVQCConfigNotFound{},
		},
		{
			testName: "idepontent update (no actual changes)",
			setup: func(f *testFixture) UpdateExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())

				name := experiment.Name().Value()
				description := experiment.Description().Value()
				trainConfigID := experiment.TrainConfigID()
				vqcConfigID := experiment.VQCConfigID()

				return UpdateExperimentInput{
					CallerID:      user.ID(),
					ExperimentID:  experiment.ExperimentID(),
					Name:          &name,
					Description:   &description,
					TrainConfigID: &trainConfigID,
					VQCConfigID:   &vqcConfigID,
				}
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)
			input := tt.setup(f)
			err := f.service.UpdateExperiment(input)
			if tt.expectedError != nil {
				require.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				require.NoError(t, err)
				experiment, err := f.experimentRepo.FindByID(input.ExperimentID)
				require.NoError(t, err)
				require.NotNil(t, experiment)
				if input.Name != nil {
					assert.Equal(t, *input.Name, experiment.Name().Value())
				}
				if input.Description != nil {
					assert.Equal(t, *input.Description, experiment.Description().Value())
				}
				if input.TrainConfigID != nil {
					assert.Equal(t, *input.TrainConfigID, experiment.TrainConfigID())
				}
				if input.VQCConfigID != nil {
					assert.Equal(t, *input.VQCConfigID, experiment.VQCConfigID())
				}
			}
		})
	}
}
