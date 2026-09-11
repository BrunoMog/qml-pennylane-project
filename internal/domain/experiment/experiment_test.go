package experiment

import (
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
)

func TestNewExperiment(t *testing.T) {
	testCases := []struct {
		expectedError error
		setup         func() ExperimentInput
		testName      string
	}{
		{
			testName: "valid input",
			setup: func() ExperimentInput {
				name, _ := NewName("Valid Name")
				description, _ := NewDescription("Valid Description")
				return ExperimentInput{
					Name:          name,
					Description:   description,
					OwnerID:       uuid.New(),
					TrainConfigID: uuid.New(),
					VQCConfigID:   uuid.New(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "invalid owner ID",
			setup: func() ExperimentInput {
				name, _ := NewName("Valid Name")
				description, _ := NewDescription("Valid Description")
				return ExperimentInput{
					Name:          name,
					Description:   description,
					OwnerID:       uuid.Nil(),
					TrainConfigID: uuid.New(),
					VQCConfigID:   uuid.New(),
				}
			},
			expectedError: &InvalidOwnerIDError{},
		},
		{
			testName: "invalid train config ID",
			setup: func() ExperimentInput {
				name, _ := NewName("Valid Name")
				description, _ := NewDescription("Valid Description")
				return ExperimentInput{
					Name:          name,
					Description:   description,
					OwnerID:       uuid.New(),
					TrainConfigID: uuid.Nil(),
					VQCConfigID:   uuid.New(),
				}
			},
			expectedError: &InvalidTrainConfigIDError{},
		},
		{
			testName: "invalid VQC config ID",
			setup: func() ExperimentInput {
				name, _ := NewName("Valid Name")
				description, _ := NewDescription("Valid Description")
				return ExperimentInput{
					Name:          name,
					Description:   description,
					OwnerID:       uuid.New(),
					TrainConfigID: uuid.New(),
					VQCConfigID:   uuid.Nil(),
				}
			},
			expectedError: &InvalidVQCConfigIDError{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			input := tc.setup()
			experiment, err := NewExperiment(input)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, experiment)
				assert.Equal(t, input.Name, experiment.Name())
				assert.Equal(t, input.Description, experiment.Description())
				assert.Equal(t, input.OwnerID, experiment.ownerID)
				assert.Equal(t, input.TrainConfigID, experiment.trainConfigID)
				assert.Equal(t, input.VQCConfigID, experiment.vQCConfigID)
			}
		})
	}
}
