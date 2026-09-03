package vqcconfig

import (
	"pennylane_project_backend/internal/domain/vqc"
	"testing"

	"github.com/google/uuid"
)

func TestNewVQCConfig(t *testing.T) {
	tests := []struct {
		vqc         *vqc.VQC
		name        string
		nameInput   string
		description string
		userID      uuid.UUID
		expectErr   bool
	}{
		{
			name:        "valid VQCConfig",
			userID:      uuid.New(),
			nameInput:   "Test Config",
			description: "This is a test VQC configuration.",
			vqc:         &vqc.VQC{},
			expectErr:   false,
		},
		{
			name:        "empty name",
			userID:      uuid.New(),
			nameInput:   "",
			description: "This is a test VQC configuration.",
			vqc:         &vqc.VQC{},
			expectErr:   true,
		},
		{
			name:        "name too long",
			userID:      uuid.New(),
			nameInput:   "This name is way too long for the validation rules. It exceeds the maximum allowed length of 100 characters.",
			description: "This is a test VQC configuration.",
			vqc:         &vqc.VQC{},
			expectErr:   true,
		},
		{
			name:        "description too long",
			userID:      uuid.New(),
			nameInput:   "Test Config",
			description: "050-----------------------------------------------100-----------------------------------------------150-----------------------------------------------200-----------------------------------------------250-----------------------------------------------300-----------------------------------------------350-----------------------------------------------400-----------------------------------------------450-----------------------------------------------500-----------------------------------------------X",
			vqc:         &vqc.VQC{},
			expectErr:   true,
		},
		{
			name:        "nil VQC",
			userID:      uuid.New(),
			nameInput:   "Test Config",
			description: "This is a test VQC configuration.",
			vqc:         nil,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewVQCConfig(tt.userID, tt.nameInput, tt.description, tt.vqc)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewVQCConfig() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
