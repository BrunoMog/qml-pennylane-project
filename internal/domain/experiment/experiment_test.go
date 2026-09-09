package experiment

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewExperiment(t *testing.T) {
	testCases := []struct {
		expectedError error
		name          string
		input         ExperimentInput
	}{
		{
			name: "valid input",
			input: ExperimentInput{
				Name:          "Test Experiment",
				Description:   "This is a test experiment.",
				OwnerID:       uuid.New(),
				TrainConfigID: uuid.New(),
				VQCConfigID:   uuid.New(),
			},
			expectedError: nil,
		},
		{
			name: "invalid name",
			input: ExperimentInput{
				Name:          "",
				Description:   "This is a test experiment.",
				OwnerID:       uuid.New(),
				TrainConfigID: uuid.New(),
				VQCConfigID:   uuid.New(),
			},
			expectedError: &InvalidNameError{""},
		},
		{
			name: "invalid description",
			input: ExperimentInput{
				Name:          "Test Experiment",
				Description:   "Description that exceeds the maximum length of 200 characters. This description is intentionally made very long to test the validation logic in the NewExperiment function. It should trigger an InvalidDescriptionError because it is too long.",
				OwnerID:       uuid.New(),
				TrainConfigID: uuid.New(),
				VQCConfigID:   uuid.New(),
			},
			expectedError: &InvalidDescriptionError{""},
		},
		{
			name: "invalid owner ID",
			input: ExperimentInput{
				Name:          "Test Experiment",
				Description:   "This is a test experiment.",
				OwnerID:       uuid.Nil,
				TrainConfigID: uuid.New(),
				VQCConfigID:   uuid.New(),
			},
			expectedError: &InvalidOwnerIDError{},
		},
		{
			name: "invalid train config ID",
			input: ExperimentInput{
				Name:          "Test Experiment",
				Description:   "This is a test experiment.",
				OwnerID:       uuid.New(),
				TrainConfigID: uuid.Nil,
				VQCConfigID:   uuid.New(),
			},
			expectedError: &InvalidTrainConfigIDError{},
		},
		{
			name: "invalid VQC config ID",
			input: ExperimentInput{
				Name:          "Test Experiment",
				Description:   "This is a test experiment.",
				OwnerID:       uuid.New(),
				TrainConfigID: uuid.New(),
				VQCConfigID:   uuid.Nil,
			},
			expectedError: &InvalidVQCConfigIDError{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			experiment, err := NewExperiment(tc.input)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Nil(t, experiment)
				assert.IsType(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, experiment)
				assert.Equal(t, tc.input.Name, experiment.name)
				assert.Equal(t, tc.input.Description, experiment.description)
				assert.Equal(t, tc.input.OwnerID, experiment.ownerID)
				assert.Equal(t, tc.input.TrainConfigID, experiment.trainConfigID)
				assert.Equal(t, tc.input.VQCConfigID, experiment.vQCConfigID)
			}
		})
	}
}
