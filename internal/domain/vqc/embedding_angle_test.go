package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAngleEmbedding(t *testing.T) {
	testCases := []struct {
		testName  string
		rotation  EmbeddingRotation
		qubits    []Qubit
		expectErr error
	}{
		{testName: "Valid angle embedding", qubits: []Qubit{0, 1}, rotation: XRotation, expectErr: nil},
		{testName: "Invalid rotation", qubits: []Qubit{0, 1}, rotation: EmbeddingRotation("invalid_rotation"), expectErr: &InvalidRotationError{}},
		{testName: "Duplicate qubit", qubits: []Qubit{0, 1, 1}, rotation: XRotation, expectErr: &DuplicateQubitError{}},
		{testName: "Zero qubits", qubits: []Qubit{}, rotation: XRotation, expectErr: &ZeroQubitEmbeddingError{}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			embedding, err := NewAngleEmbedding(tc.qubits, tc.rotation)
			if tc.expectErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.qubits, embedding.Qubits())
				assert.Equal(t, tc.rotation, embedding.Rotation())
			}
		})
	}
}
