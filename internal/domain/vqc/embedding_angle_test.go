package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAngleEmbedding(t *testing.T) {

	t.Run("valid angle embedding cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			qubits   []Qubit
			rotation EmbeddingRotation
		}{
			{"valid embedding with X rotation", []Qubit{0, 1}, XRotation},
			{"valid embedding with Y rotation", []Qubit{2, 3}, YRotation},
			{"valid embedding with Z rotation", []Qubit{4, 5}, ZRotation},
		}

		for _, testCase := range inputs {
			t.Run(testCase.testName, func(t *testing.T) {
				embedding, err := NewAngleEmbedding(testCase.qubits, testCase.rotation)
				assert.NoError(t, err)
				assert.True(t, embedding.isValid())
				assert.Equal(t, testCase.qubits, embedding.Qubits())
				assert.Equal(t, testCase.rotation, embedding.Rotation())
			})
		}
	})

	t.Run("invalid angle embedding cases", func(t *testing.T) {
		inputs := []struct {
			testName  string
			qubits    []Qubit
			rotation  EmbeddingRotation
			expectErr error
		}{
			{"invalid rotation", []Qubit{0, 1}, EmbeddingRotation("invalid_rotation"), ErrInvalidEmbeddingRotation},
			{"duplicate qubit", []Qubit{0, 1, 1}, XRotation, ErrDuplicatedQubit},
			{"zero qubits", []Qubit{}, XRotation, ErrZeroQubitEmbedding},
		}

		for _, testCase := range inputs {
			t.Run(testCase.testName, func(t *testing.T) {
				embedding, err := NewAngleEmbedding(testCase.qubits, testCase.rotation)
				assert.Error(t, err)
				assert.ErrorIs(t, err, testCase.expectErr)
				assert.False(t, embedding.isValid())
			})
		}
	})
}

func TestAngleEmbeddingEquals(t *testing.T) {
	t.Run("equal angle embeddings", func(t *testing.T) {
		embedding1, err := NewAngleEmbedding([]Qubit{0, 1}, XRotation)
		require.NoError(t, err)
		embedding2, err := NewAngleEmbedding([]Qubit{0, 1}, XRotation)
		require.NoError(t, err)

		assert.True(t, embedding1.Equals(embedding2))
	})

	t.Run("different angle embeddings", func(t *testing.T) {
		embedding1, err := NewAngleEmbedding([]Qubit{0, 1}, XRotation)
		require.NoError(t, err)
		embedding2, err := NewAngleEmbedding([]Qubit{0, 1}, YRotation)
		require.NoError(t, err)
		embedding3, err := NewAngleEmbedding([]Qubit{0, 2}, XRotation)
		require.NoError(t, err)

		assert.False(t, embedding1.Equals(embedding2))
		assert.False(t, embedding1.Equals(embedding3))
	})
}

func TestAngleEmbeddingValidateQubits(t *testing.T) {
	t.Run("valid qubit cases", func(t *testing.T) {
		embedding, err := NewAngleEmbedding([]Qubit{0, 1}, XRotation)
		require.NoError(t, err)
		err = embedding.validateQubits(3)
		assert.NoError(t, err)
	})

	t.Run("invalid qubit cases", func(t *testing.T) {
		embedding, err := NewAngleEmbedding([]Qubit{0, 1}, XRotation)
		require.NoError(t, err)
		err = embedding.validateQubits(1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidQubit)
	})
}
