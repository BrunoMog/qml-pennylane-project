package vqcconfig

import (
	"pennylane_project_backend/internal/domain/vqc"
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validVQC(t *testing.T) vqc.VQC {
	t.Helper()

	qubitZero, err := vqc.NewQubit(0, 1)
	require.NoError(t, err)
	qubits := []vqc.Qubit{qubitZero}
	embedding, err := vqc.NewAngleEmbedding(qubits, vqc.XRotation)
	require.NoError(t, err)
	measurement, err := vqc.NewMeasurement(qubits, vqc.ExpectationMeasurement, vqc.XMeasurementRotation)
	require.NoError(t, err)
	input := vqc.VQCBaseInput{
		NumQubits:   1,
		NumLayers:   1,
		Embedding:   embedding,
		Measurement: measurement,
	}
	vqc, err := vqc.NewVQC(input)
	require.NoError(t, err)
	return vqc
}

func TestNewVQCConfig(t *testing.T) {
	tests := []struct {
		expectError error
		setup       func(t *testing.T) (uuid.UUID, Name, Description, vqc.VQC)
		testName    string
	}{
		{
			testName: "valid VQCConfig",
			setup: func(t *testing.T) (uuid.UUID, Name, Description, vqc.VQC) {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				vqcInstance := validVQC(t)
				return userID, name, description, vqcInstance
			},
			expectError: nil,
		},
		{
			testName: "invalid VQCConfig with nil userID",
			setup: func(t *testing.T) (uuid.UUID, Name, Description, vqc.VQC) {
				userID := uuid.Nil()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				vqcInstance := validVQC(t)
				return userID, name, description, vqcInstance
			},
			expectError: ErrInvalidOwnerID,
		},
		{
			testName: "invalid VQCConfig with invalid VQC",
			setup: func(t *testing.T) (uuid.UUID, Name, Description, vqc.VQC) {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				return userID, name, description, vqc.VQC{}
			},
			expectError: ErrZeroValueVQC,
		},
		{
			testName: "invalid VQCConfig with empty name",
			setup: func(t *testing.T) (uuid.UUID, Name, Description, vqc.VQC) {
				userID := uuid.New()
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				vqcInstance := validVQC(t)
				return userID, Name{}, description, vqcInstance
			},
			expectError: ErrEmptyName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			userID, name, description, vqcInstance := tt.setup(t)
			vqcConfig, err := NewVQCConfig(userID, name, description, vqcInstance)
			if tt.expectError != nil {
				assert.ErrorIs(t, err, tt.expectError)
				assert.Nil(t, vqcConfig)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, vqcConfig)
				assert.Equal(t, userID, vqcConfig.userID)
				assert.Equal(t, name.String(), vqcConfig.name.String())
				assert.Equal(t, description.String(), vqcConfig.description.String())
				assert.Equal(t, vqcInstance, vqcConfig.vqc)
			}
		})
	}
}

func TestSetVQC(t *testing.T) {
	tests := []struct {
		expectError error
		setup       func(t *testing.T) *VQCConfig
		testName    string
		newVQC      vqc.VQC
	}{
		{
			testName: "valid SetVQC",
			setup: func(t *testing.T) *VQCConfig {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				vqcInstance := validVQC(t)
				vqcConfig, err := NewVQCConfig(userID, name, description, vqcInstance)
				require.NoError(t, err)
				return vqcConfig
			},
			newVQC:      validVQC(t),
			expectError: nil,
		},
		{
			testName: "invalid SetVQC with invalid VQC",
			setup: func(t *testing.T) *VQCConfig {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				vqcInstance := validVQC(t)
				vqcConfig, err := NewVQCConfig(userID, name, description, vqcInstance)
				require.NoError(t, err)
				return vqcConfig
			},
			newVQC:      vqc.VQC{},
			expectError: ErrZeroValueVQC,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			vqcConfig := tt.setup(t)
			err := vqcConfig.SetVQC(tt.newVQC)
			if tt.expectError != nil {
				assert.ErrorIs(t, err, tt.expectError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newVQC, vqcConfig.vqc)
			}
		})
	}
}

func TestSetName(t *testing.T) {
	tests := []struct {
		expectError error
		setup       func(t *testing.T) (*VQCConfig, Name)
		testName    string
	}{
		{
			testName: "valid SetName",
			setup: func(t *testing.T) (*VQCConfig, Name) {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				vqcInstance := validVQC(t)
				vqcConfig, err := NewVQCConfig(userID, name, description, vqcInstance)
				require.NoError(t, err)
				newName, err := NewName("New Valid Name")
				require.NoError(t, err)
				return vqcConfig, newName
			},
			expectError: nil,
		},
		{
			testName: "invalid SetName with empty name",
			setup: func(t *testing.T) (*VQCConfig, Name) {
				userID := uuid.New()
				name, err := NewName("Valid Name")
				require.NoError(t, err)
				description, err := NewDescription("Valid Description")
				require.NoError(t, err)
				vqcInstance := validVQC(t)
				vqcConfig, err := NewVQCConfig(userID, name, description, vqcInstance)
				require.NoError(t, err)
				return vqcConfig, Name{}
			},
			expectError: ErrEmptyName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			vqcConfig, newName := tt.setup(t)
			err := vqcConfig.SetName(newName)
			if tt.expectError != nil {
				assert.ErrorIs(t, err, tt.expectError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, newName.String(), vqcConfig.name.String())
			}
		})
	}
}
