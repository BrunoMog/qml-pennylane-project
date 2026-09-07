package vqc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAmplitudeEmbedding(t *testing.T) {
	testCases := []struct {
		testName  string
		qubits    []Qubit
		padwith   float64
		normalize bool
		expectErr error
	}{
		{testName: "Valid amplitude embedding", qubits: []Qubit{0, 1}, normalize: true, padwith: 0.0, expectErr: nil},
		{testName: "Duplicate qubit index", qubits: []Qubit{0, 1, 1}, padwith: 0.0, normalize: true, expectErr: &DuplicateQubitError{}},
		{testName: "Zero qubits", qubits: []Qubit{}, padwith: 0.0, normalize: true, expectErr: &ZeroQubitEmbeddingError{}},
		{testName: "NaN padwith value", qubits: []Qubit{0, 1}, padwith: math.NaN(), normalize: true, expectErr: &InvalidPadWithError{}},
		{testName: "Infinite padwith value", qubits: []Qubit{0, 1}, padwith: math.Inf(1), normalize: true, expectErr: &InvalidPadWithError{}},
		{testName: "Negative padwith value", qubits: []Qubit{0, 1}, padwith: -1.0, normalize: true, expectErr: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			embedding, err := NewAmplitudeEmbedding(tc.qubits, tc.normalize, tc.padwith)
			if tc.expectErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.qubits, embedding.Qubits())
				assert.Equal(t, tc.normalize, embedding.Normalize())
				assert.Equal(t, tc.padwith, embedding.PadWith())
			}
		})
	}
}
