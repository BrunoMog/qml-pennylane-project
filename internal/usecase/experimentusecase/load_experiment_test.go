package experimentusecase

import (
	"pennylane_project_backend/internal/testkit"
	"testing"
	"uuid"

	"github.com/stretchr/testify/assert"
)

func TestLoadExperiment(t *testing.T) {
	tests := []struct {
		testName      string
		setup         func(f *testFixture) LoadExperimentInput
		expectedError error
	}{
		{
			testName: "valid load experiment by ID",
			setup: func(f *testFixture) LoadExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())
				experimentID := experiment.ExperimentID()
				return LoadExperimentInput{
					ExperimentID: &experimentID,
					CallerID:     user.ID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "valid load experiment by name",
			setup: func(f *testFixture) LoadExperimentInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				experiment := f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())
				experimentName := experiment.Name().Value()
				return LoadExperimentInput{
					ExperimentName: &experimentName,
					CallerID:       user.ID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent experiment",
			setup: func(f *testFixture) LoadExperimentInput {
				user := f.createUser()
				experimentID := uuid.New()
				return LoadExperimentInput{
					ExperimentID: &experimentID,
					CallerID:     user.ID(),
				}
			},
			expectedError: &testkit.ErrExperimentNotFound{},
		},
		{
			testName: "unauthorized load experiment",
			setup: func(f *testFixture) LoadExperimentInput {
				user := f.createUser()
				otherUser := f.createUser()
				trainConfig := f.createTrainConfig(otherUser.ID())
				vqcConfig := f.createVQCConfig(otherUser.ID())
				experiment := f.createExperiment(otherUser.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())
				experimentID := experiment.ExperimentID()
				return LoadExperimentInput{
					ExperimentID: &experimentID,
					CallerID:     user.ID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent user",
			setup: func(f *testFixture) LoadExperimentInput {
				trainConfig := f.createTrainConfig(uuid.New())
				vqcConfig := f.createVQCConfig(uuid.New())
				experiment := f.createExperiment(uuid.New(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())
				experimentID := experiment.ExperimentID()
				return LoadExperimentInput{
					ExperimentID: &experimentID,
					CallerID:     uuid.New(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "invalid input: both ExperimentID and ExperimentName are nil",
			setup: func(f *testFixture) LoadExperimentInput {
				user := f.createUser()
				return LoadExperimentInput{
					CallerID:       user.ID(),
					ExperimentID:   nil,
					ExperimentName: nil,
				}
			},
			expectedError: &InvalidInputError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)
			input := tt.setup(f)
			experiment, err := f.service.LoadExperiment(input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Nil(t, experiment)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, experiment)
				assert.Equal(t, input.CallerID, experiment.OwnerID)
				if input.ExperimentID != nil {
					assert.Equal(t, *input.ExperimentID, experiment.ExperimentID)
				}
				if input.ExperimentName != nil {
					assert.Equal(t, *input.ExperimentName, experiment.ExperimentName)
				}
				assert.Equal(t, experiment.TrainConfigID, experiment.TrainConfigID)
				assert.Equal(t, experiment.VQCConfigID, experiment.VQCConfigID)
			}
		})
	}
}

func TestLoadAllExperiments(t *testing.T) {
	tests := []struct {
		testName           string
		setup              func(f *testFixture) LoadAllExperimentsInput
		expectedSizeOutput int
	}{
		{
			testName: "valid load all experiments",
			setup: func(f *testFixture) LoadAllExperimentsInput {
				user := f.createUser()
				trainConfig := f.createTrainConfig(user.ID())
				vqcConfig := f.createVQCConfig(user.ID())
				f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())
				f.createExperiment(user.ID(), vqcConfig.VQCConfigID(), trainConfig.TrainConfigID())
				return LoadAllExperimentsInput{
					CallerID: user.ID(),
				}
			},
			expectedSizeOutput: 2,
		},
		{
			testName: "valid load all experiments with no experiments",
			setup: func(f *testFixture) LoadAllExperimentsInput {
				user := f.createUser()
				return LoadAllExperimentsInput{
					CallerID: user.ID(),
				}
			},
			expectedSizeOutput: 0,
		},
		{
			testName: "inexistent user",
			setup: func(f *testFixture) LoadAllExperimentsInput {
				return LoadAllExperimentsInput{
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
			experiments, err := f.service.LoadAllExperiments(input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedSizeOutput, len(experiments.Experiments))
		})
	}
}
