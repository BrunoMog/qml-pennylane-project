package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMeasurementRotation(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectedRotation MeasurementRotation
		expectedError    error
	}{
		{"Valid X rotation", "x", XMeasurementRotation, nil},
		{"Valid Y rotation", "y", YMeasurementRotation, nil},
		{"Valid Z rotation", "z", ZMeasurementRotation, nil},
		{"Invalid rotation", "invalid_rotation", "", &InvalidParseMeasurementRotationError{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotation, err := ParseMeasurementRotation(tt.input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRotation, rotation)
			}
		})
	}
}
