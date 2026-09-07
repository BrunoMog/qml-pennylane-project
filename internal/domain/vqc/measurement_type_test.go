package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMeasurementType(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedType  MeasurementType
		expectedError error
	}{
		{"Valid expectation measurement", "expectation", ExpectationMeasurement, nil},
		{"Valid probability measurement", "probability", ProbabilityMeasurement, nil},
		{"Invalid measurement type", "invalid_measurement", "", &InvalidParseMeasurementTypeError{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
