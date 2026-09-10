package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMeasurementRotation(t *testing.T) {
	tests := []struct {
		expectedError    error
		testName         string
		input            string
		expectedRotation MeasurementRotation
	}{
		{testName: "Valid X rotation", input: "x", expectedRotation: XMeasurementRotation, expectedError: nil},
		{testName: "Valid Y rotation", input: "y", expectedRotation: YMeasurementRotation, expectedError: nil},
		{testName: "Valid Z rotation", input: "z", expectedRotation: ZMeasurementRotation, expectedError: nil},
		{testName: "Invalid rotation", input: "invalid_rotation", expectedRotation: "", expectedError: &InvalidParseMeasurementRotationError{}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
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
