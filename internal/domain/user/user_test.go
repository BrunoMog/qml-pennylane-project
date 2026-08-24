package user

import (
	"testing"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{"valid name", "John Doe", false},
		{"empty name", "", true},
		{"name too long", "ThisNameIsWayTooLongToBeValid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateName(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateName() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
