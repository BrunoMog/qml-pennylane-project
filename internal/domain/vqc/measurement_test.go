package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMeasurement(t *testing.T) {

	t.Run("valid measurement cases", func(t *testing.T) {
		inputs := []struct {
			testName             string
			qubits               []Qubit
			measurement_type     MeasurementType
			measurement_rotation MeasurementRotation
		}{
			{"valid measurement with expectation and x rotation", []Qubit{0, 1}, ExpectationMeasurement, XMeasurementRotation},
			{"valid measurement with probability and y rotation", []Qubit{2, 3}, ProbabilityMeasurement, YMeasurementRotation},
			{"valid measurement with single qubit", []Qubit{4}, ExpectationMeasurement, ZMeasurementRotation},
		}

		for _, input := range inputs {
			t.Run(input.testName, func(t *testing.T) {
				measurement, err := NewMeasurement(input.qubits, input.measurement_type, input.measurement_rotation)
				assert.NoError(t, err)
				assert.Equal(t, input.qubits, measurement.Qubits())
				assert.Equal(t, input.measurement_type, measurement.MeasurementType())
				assert.Equal(t, input.measurement_rotation, measurement.MeasurementRotation())
			})
		}
	})

	t.Run("invalid measurement cases", func(t *testing.T) {
		testCases := []struct {
			expectErr            error
			testName             string
			measurement_type     MeasurementType
			measurement_rotation MeasurementRotation
			qubits               []Qubit
		}{
			{testName: "zero qubits", qubits: []Qubit{}, measurement_type: ExpectationMeasurement, measurement_rotation: XMeasurementRotation, expectErr: ErrZeroQubitMeasurement},
			{testName: "duplicate qubits", qubits: []Qubit{0, 1, 1}, measurement_type: ExpectationMeasurement, measurement_rotation: XMeasurementRotation, expectErr: ErrDuplicatedQubit},
			{testName: "invalid measurement type", qubits: []Qubit{0, 1}, measurement_type: MeasurementType("invalid"), measurement_rotation: XMeasurementRotation, expectErr: ErrInvalidMeasurementType},
			{testName: "invalid measurement rotation", qubits: []Qubit{0, 1}, measurement_type: ExpectationMeasurement, measurement_rotation: MeasurementRotation("invalid"), expectErr: ErrInvalidMeasurementRotation},
		}

		for _, tc := range testCases {
			t.Run(tc.testName, func(t *testing.T) {
				measurement, err := NewMeasurement(tc.qubits, tc.measurement_type, tc.measurement_rotation)
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.expectErr)
				assert.False(t, measurement.isValid())
			})
		}
	})
}

func TestMeasurementQubitsImmutability(t *testing.T) {
	t.Run("modifying the original qubits slice does not affect the measurement's qubits", func(t *testing.T) {
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
	})
}

func TestMeasurementValidateQubits(t *testing.T) {
	t.Run("valid qubit cases", func(t *testing.T) {
		measurement, err := NewMeasurement([]Qubit{0, 1}, ExpectationMeasurement, XMeasurementRotation)
		require.NoError(t, err)
		err = measurement.validateQubits(3)
		assert.NoError(t, err)
	})

	t.Run("invalid qubit cases", func(t *testing.T) {
		measurement, err := NewMeasurement([]Qubit{0, 1}, ExpectationMeasurement, XMeasurementRotation)
		require.NoError(t, err)
		err = measurement.validateQubits(1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidQubit)
	})
}
