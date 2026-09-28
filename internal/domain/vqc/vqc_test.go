package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validEmbeddingOneQubit(t *testing.T) Embedding {
	t.Helper()
	qubitZero, err := NewQubit(0, 1)
	require.NoError(t, err)
	qubits := []Qubit{qubitZero}
	embedding, err := NewAngleEmbedding(qubits, XRotation)
	require.NoError(t, err)
	return embedding
}

func validEmbeddingTwoQubits(t *testing.T) Embedding {
	t.Helper()
	qubitZero, err := NewQubit(0, 2)
	require.NoError(t, err)
	qubitOne, err := NewQubit(1, 2)
	require.NoError(t, err)
	qubits := []Qubit{qubitZero, qubitOne}
	embedding, err := NewAngleEmbedding(qubits, YRotation)
	require.NoError(t, err)
	return embedding
}

func validMeasurementOneQubit(t *testing.T) Measurement {
	t.Helper()
	qubitZero, err := NewQubit(0, 1)
	require.NoError(t, err)
	qubits := []Qubit{qubitZero}
	measurement, err := NewMeasurement(qubits, ExpectationMeasurement, XMeasurementRotation)
	require.NoError(t, err)
	return measurement
}

func validMeasurementTwoQubits(t *testing.T) Measurement {
	t.Helper()
	qubitZero, err := NewQubit(0, 2)
	require.NoError(t, err)
	qubitOne, err := NewQubit(1, 2)
	require.NoError(t, err)
	qubits := []Qubit{qubitZero, qubitOne}
	measurement, err := NewMeasurement(qubits, ExpectationMeasurement, XMeasurementRotation)
	require.NoError(t, err)
	return measurement
}

func TestNewVQC(t *testing.T) {

	t.Run("valid VQC", func(t *testing.T) {
		embedding := validEmbeddingOneQubit(t)
		measurement := validMeasurementOneQubit(t)

		inputVQC := VQCBaseInput{
			Embedding:   embedding,
			Measurement: measurement,
			NumQubits:   1,
			NumLayers:   1,
		}

		vqc, err := NewVQC(inputVQC)
		assert.NoError(t, err)
		assert.Equal(t, inputVQC.NumQubits, vqc.NumQubits())
		assert.Equal(t, inputVQC.NumLayers, vqc.NumLayers())
		assert.Equal(t, inputVQC.Embedding, vqc.Embedding())
		assert.Equal(t, inputVQC.Measurement, vqc.Measurement())
	})

	t.Run("invalid VQC cases", func(t *testing.T) {
		tests := []struct {
			testName  string
			inputVQC  VQCBaseInput
			expectErr error
		}{
			{
				testName: "zero qubits",
				inputVQC: VQCBaseInput{
					Embedding:   validEmbeddingOneQubit(t),
					Measurement: validMeasurementOneQubit(t),
					NumQubits:   0,
					NumLayers:   1,
				},
				expectErr: ErrZeroQubitVQC,
			},
			{
				testName: "nil embedding",
				inputVQC: VQCBaseInput{
					Embedding:   nil,
					Measurement: validMeasurementOneQubit(t),
					NumQubits:   1,
					NumLayers:   1,
				},
				expectErr: ErrNilEmbedding,
			},
			{
				testName: "invalid embedding",
				inputVQC: VQCBaseInput{
					Embedding:   AngleEmbedding{},
					Measurement: validMeasurementOneQubit(t),
					NumQubits:   1,
					NumLayers:   1,
				},
				expectErr: ErrInvalidEmbedding,
			},
			{
				testName: "invalid measurement",
				inputVQC: VQCBaseInput{
					Embedding:   validEmbeddingOneQubit(t),
					Measurement: Measurement{},
					NumQubits:   1,
					NumLayers:   1,
				},
				expectErr: ErrInvalidMeasurement,
			},
			{
				testName: "invalid measurement qubits",
				inputVQC: VQCBaseInput{
					Embedding:   validEmbeddingOneQubit(t),
					Measurement: validMeasurementTwoQubits(t),
					NumQubits:   1,
					NumLayers:   1,
				},
				expectErr: ErrInvalidQubit,
			},
			{
				testName: "invalid embedding qubits",
				inputVQC: VQCBaseInput{
					Embedding:   validEmbeddingTwoQubits(t),
					Measurement: validMeasurementOneQubit(t),
					NumQubits:   1,
					NumLayers:   1,
				},
				expectErr: ErrInvalidQubit,
			},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				vqc, err := NewVQC(tt.inputVQC)
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.expectErr)
				assert.False(t, vqc.IsValid())
			})
		}
	})
}

func TestEquals(t *testing.T) {

	t.Run("equal VQCs", func(t *testing.T) {
		qubitZero, err := NewQubit(0, 1)
		require.NoError(t, err)
		qubits := []Qubit{qubitZero}
		embedding, err := NewAngleEmbedding(qubits, XRotation)
		require.NoError(t, err)
		measurement, err := NewMeasurement(qubits, ExpectationMeasurement, XMeasurementRotation)
		require.NoError(t, err)

		vqcInput := VQCBaseInput{
			NumQubits:   1,
			NumLayers:   1,
			Embedding:   embedding,
			Measurement: measurement,
		}

		vqc1, err := NewVQC(vqcInput)
		require.NoError(t, err)

		vqc2, err := NewVQC(vqcInput)
		require.NoError(t, err)

		assert.True(t, vqc1.Equals(vqc2))
	})

	t.Run("different VQCs", func(t *testing.T) {
		qubitZero, err := NewQubit(0, 1)
		require.NoError(t, err)
		qubits := []Qubit{qubitZero}
		embedding1, err := NewAngleEmbedding(qubits, XRotation)
		require.NoError(t, err)
		embedding2, err := NewAngleEmbedding(qubits, YRotation)
		require.NoError(t, err)
		measurement, err := NewMeasurement(qubits, ExpectationMeasurement, XMeasurementRotation)
		require.NoError(t, err)

		vqcInput1 := VQCBaseInput{
			NumQubits:   1,
			NumLayers:   1,
			Embedding:   embedding1,
			Measurement: measurement,
		}

		vqcInput2 := VQCBaseInput{
			NumQubits:   1,
			NumLayers:   1,
			Embedding:   embedding2,
			Measurement: measurement,
		}

		vqc1, err := NewVQC(vqcInput1)
		require.NoError(t, err)

		vqc2, err := NewVQC(vqcInput2)
		require.NoError(t, err)

		assert.False(t, vqc1.Equals(vqc2))
	})
}
