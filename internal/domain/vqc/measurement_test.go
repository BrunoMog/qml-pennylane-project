package vqc

import (
	"testing"
)

func TestNewMeasurement(t *testing.T) {
	testCases := []struct {
		name                 string
		measurement_type     MeasurementType
		measurement_rotation MeasurementRotation
		qubits               []Qubit
		expectErr            bool
	}{
		{name: "Valid measurement", qubits: []Qubit{0, 1}, measurement_type: ExpectationMeasurement, measurement_rotation: XMeasurementRotation, expectErr: false},
		{name: "Invalid measurement type", qubits: []Qubit{0, 1}, measurement_type: MeasurementType("invalid_measurement"), measurement_rotation: XMeasurementRotation, expectErr: true},
		{name: "Duplicate qubit index", qubits: []Qubit{0, 1, 1}, measurement_type: ExpectationMeasurement, measurement_rotation: XMeasurementRotation, expectErr: true},
		{name: "Invalid measurement rotation", qubits: []Qubit{0, 1}, measurement_type: ExpectationMeasurement, measurement_rotation: MeasurementRotation("invalid_rotation"), expectErr: true},
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
