package vqcconfig

import (
	"pennylane_project_backend/internal/domain/vqc"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
				vqcInstance := vqc.VQC{}
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
				vqcInstance := vqc.VQC{}
				return userID, name, description, vqcInstance
			},
			expectError: &InvalidOwnerIDError{},
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
