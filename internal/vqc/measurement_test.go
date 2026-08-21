package vqc

import (
	"testing"
)

func TestNewMeasurement(t *testing.T) {
	testCases := []struct {
		name                 string
		qubits               []Qubit
		measurement_type     MeasurementType
		measurement_rotation MeasurementRotation
		expectErr            bool
	}{
		{"Valid measurement", []Qubit{0, 1}, ExpectationMeasurement, XMeasurementRotation, false},
		{"Invalid measurement type", []Qubit{0, 1}, MeasurementType("invalid_measurement"), XMeasurementRotation, true},
		{"Duplicate qubit index", []Qubit{0, 1, 1}, ExpectationMeasurement, XMeasurementRotation, true},
		{"Invalid measurement rotation", []Qubit{0, 1}, ExpectationMeasurement, MeasurementRotation("invalid_rotation"), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewMeasurement(tc.qubits, tc.measurement_type, tc.measurement_rotation)
			if (err != nil) != tc.expectErr {
				t.Errorf("NewMeasurement(%v, %s, %s) returned error: %v, expected error: %v", tc.qubits, tc.measurement_type, tc.measurement_rotation, err != nil, tc.expectErr)
			}
		})
	}
}
