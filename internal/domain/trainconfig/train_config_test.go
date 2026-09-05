package trainconfig

import (
	"pennylane_project_backend/internal/domain/training"
	"testing"

	"github.com/google/uuid"
)

func TestNewTrainConfig(t *testing.T) {
	tests := []struct {
		training    *training.Training
		name        string
		nameInput   string
		description string
		ownerID     uuid.UUID
		expectErr   bool
	}{
		{
			name:        "valid TrainConfig",
			ownerID:     uuid.New(),
			nameInput:   "Test Config",
			description: "This is a test training configuration.",
			training:    &training.Training{},
			expectErr:   false,
		},
		{
			name:        "empty name",
			ownerID:     uuid.New(),
			nameInput:   "",
			description: "This is a test training configuration.",
			training:    &training.Training{},
			expectErr:   true,
		},
		{
			name:        "name too long",
			ownerID:     uuid.New(),
			nameInput:   "This name is way too long for the validation rules. It exceeds the maximum allowed length of 100 characters.",
			description: "This is a test training configuration.",
			training:    &training.Training{},
			expectErr:   true,
		},
		{
			name:        "description too long",
			ownerID:     uuid.New(),
			nameInput:   "Test Config",
			description: "050-----------------------------------------------100-----------------------------------------------150-----------------------------------------------200-----------------------------------------------250-----------------------------------------------300-----------------------------------------------350-----------------------------------------------400-----------------------------------------------450-----------------------------------------------500-----------------------------------------------X",
			training:    &training.Training{},
			expectErr:   true,
		},
		{
			name:        "nil Training",
			ownerID:     uuid.New(),
			nameInput:   "Test Config",
			description: "This is a test training configuration.",
			training:    nil,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc, err := NewTrainConfig(tt.ownerID, tt.nameInput, tt.description, tt.training)
			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				} else {
					if tc.Name() != tt.nameInput {
						t.Errorf("Expected name %v, got %v", tt.nameInput, tc.Name())
					}
					if tc.Description() != tt.description {
						t.Errorf("Expected description %v, got %v", tt.description, tc.Description())
					}
					if tc.OwnerID() != tt.ownerID {
						t.Errorf("Expected ownerID %v, got %v", tt.ownerID, tc.OwnerID())
					}
				}
			}
		})
	}
}
