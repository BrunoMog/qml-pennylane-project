package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMeasurementType(t *testing.T) {
	tests := []struct {
		expectedError error
		testName      string
		input         string
		expectedType  MeasurementType
	}{
		{testName: "Valid expectation measurement", input: "expectation", expectedType: ExpectationMeasurement, expectedError: nil},
		{testName: "Valid probability measurement", input: "probability", expectedType: ProbabilityMeasurement, expectedError: nil},
		{testName: "Invalid measurement type", input: "invalid_measurement", expectedType: "", expectedError: &InvalidParseMeasurementTypeError{}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			measurementType, err := ParseMeasurementType(tt.input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedType, measurementType)
			}
		})
	}
}
