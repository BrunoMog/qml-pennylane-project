package vqcconfig

import (
	"pennylane_project_backend/internal/domain/vqc"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validVQC() vqc.VQC {
	qubitZero, err := vqc.NewQubit(0, 1)
	if err != nil {
		panic(err)
	}
	qubits := []vqc.Qubit{qubitZero}
	embedding, err := vqc.NewAngleEmbedding(qubits, vqc.XRotation)
	if err != nil {
		panic(err)
	}
	measurement, err := vqc.NewMeasurement(qubits, vqc.ExpectationMeasurement, vqc.XMeasurementRotation)
	if err != nil {
		panic(err)
	}
	input := vqc.VQCBaseInput{
		NumQubits:   1,
		NumLayers:   1,
		Embedding:   embedding,
		Measurement: measurement,
	}
	vqc, err := vqc.NewVQC(input)
	if err != nil {
		panic(err)
	}
	return vqc
}

func TestNewVQCConfig(t *testing.T) {
	tests := []struct {
		testName    string
		setup       func() (uuid.UUID, Name, Description, vqc.VQC)
		expectError error
	}{
		{
			testName: "valid VQCConfig",
			setup: func() (uuid.UUID, Name, Description, vqc.VQC) {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				vqcInstance := validVQC()
				return userID, name, description, vqcInstance
			},
			expectError: nil,
		},
		{
			testName: "invalid VQCConfig with nil userID",
			setup: func() (uuid.UUID, Name, Description, vqc.VQC) {
				userID := uuid.Nil
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				vqcInstance := validVQC()
				return userID, name, description, vqcInstance
			},
			expectError: &InvalidOwnerIDError{},
		},
		{
			testName: "invalid VQCConfig with invalid VQC",
			setup: func() (uuid.UUID, Name, Description, vqc.VQC) {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				return userID, name, description, vqc.VQC{}
			},
			expectError: &InvalidVQCError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			userID, name, description, vqcInstance := tt.setup()
			vqcConfig, err := NewVQCConfig(userID, name, description, vqcInstance)
			if tt.expectError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectError, err)
				assert.Nil(t, vqcConfig)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, vqcConfig)
				assert.Equal(t, userID, vqcConfig.userID)
				assert.Equal(t, name.Value(), vqcConfig.name.Value())
				assert.Equal(t, description.Value(), vqcConfig.description.Value())
				assert.Equal(t, vqcInstance, vqcConfig.vqc)
			}
		})
	}
}
