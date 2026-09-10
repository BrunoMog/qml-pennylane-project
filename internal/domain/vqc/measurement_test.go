package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMeasurement(t *testing.T) {
	testCases := []struct {
		expectErr            error
		testName             string
		measurement_type     MeasurementType
		measurement_rotation MeasurementRotation
		qubits               []Qubit
	}{
		{testName: "Valid measurement", qubits: []Qubit{0, 1}, measurement_type: ExpectationMeasurement, measurement_rotation: XMeasurementRotation, expectErr: nil},
		{testName: "zero qubits", qubits: []Qubit{}, measurement_type: ExpectationMeasurement, measurement_rotation: XMeasurementRotation, expectErr: &ZeroQubitMeasurementError{}},
		{testName: "duplicate qubits", qubits: []Qubit{0, 1, 1}, measurement_type: ExpectationMeasurement, measurement_rotation: XMeasurementRotation, expectErr: &DuplicateQubitError{}},
		{testName: "invalid measurement type", qubits: []Qubit{0, 1}, measurement_type: MeasurementType("invalid"), measurement_rotation: XMeasurementRotation, expectErr: &InvalidMeasurementError{}},
		{testName: "invalid measurement rotation", qubits: []Qubit{0, 1}, measurement_type: ExpectationMeasurement, measurement_rotation: MeasurementRotation("invalid"), expectErr: &InvalidMeasurementRotationError{}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			measurement, err := NewMeasurement(tc.qubits, tc.measurement_type, tc.measurement_rotation)
			if tc.expectErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.qubits, measurement.Qubits())
				assert.Equal(t, tc.measurement_type, measurement.MeasurementType())
				assert.Equal(t, tc.measurement_rotation, measurement.MeasurementRotation())
			}
		})
	}
}

func TestMeasurementQubitsImmutability(t *testing.T) {
	qubits := []Qubit{0, 1}
	measurement, err := NewMeasurement(qubits, ExpectationMeasurement, XMeasurementRotation)
	require.NoError(t, err)

	// Modify the original qubits slice
	qubits[0] = 2
	assert.Equal(t, []Qubit{0, 1}, measurement.Qubits())

	// Modify the returned qubits slice from the measurement
	returnedQubits := measurement.Qubits()
	returnedQubits[0] = 3
	assert.Equal(t, []Qubit{0, 1}, measurement.Qubits())
}
