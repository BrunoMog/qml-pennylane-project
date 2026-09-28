package vqc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAmplitudeEmbedding(t *testing.T) {

	t.Run("valid amplitude embedding cases", func(t *testing.T) {
		inputs := []struct {
			testName  string
			qubits    []Qubit
			normalize bool
			padWith   float64
		}{
			{"valid embedding with normalization and padding", []Qubit{0, 1}, true, 0.0},
			{"valid embedding without normalization and with padding", []Qubit{2, 3}, false, 1.0},
			{"valid embedding with normalization and no padding", []Qubit{4, 5}, true, 0.0},
			{"valid embedding without normalization and no padding", []Qubit{6, 7}, false, 0.0},
			{"valid embedding with negative padding", []Qubit{8, 9}, true, -1.0},
		}
		for _, testCase := range inputs {
			t.Run(testCase.testName, func(t *testing.T) {
				embedding, err := NewAmplitudeEmbedding(testCase.qubits, testCase.normalize, testCase.padWith)
				assert.NoError(t, err)
				assert.True(t, embedding.isValid())
				assert.Equal(t, testCase.qubits, embedding.Qubits())
				assert.Equal(t, testCase.normalize, embedding.Normalize())
				assert.Equal(t, testCase.padWith, embedding.PadWith())
			})
		}
	})

	t.Run("invalid amplitude embedding cases", func(t *testing.T) {
		inputs := []struct {
			testName  string
			qubits    []Qubit
			normalize bool
			padWith   float64
			expectErr error
		}{
			{"duplicate qubit index", []Qubit{0, 1, 1}, true, 0.0, ErrDuplicatedQubit},
			{"NaN padWith value", []Qubit{0, 1}, true, math.NaN(), ErrInvalidPadWith},
			{"Infinite padWith value", []Qubit{0, 1}, true, math.Inf(1), ErrInvalidPadWith},
			{"zero qubits", []Qubit{}, true, 0.0, ErrZeroQubitEmbedding},
		}

		for _, testCase := range inputs {
			t.Run(testCase.testName, func(t *testing.T) {
				embedding, err := NewAmplitudeEmbedding(testCase.qubits, testCase.normalize, testCase.padWith)
				assert.Error(t, err)
				assert.ErrorIs(t, err, testCase.expectErr)
				assert.False(t, embedding.isValid())
			})
		}
	})
}

func TestAmplitudeEmbeddingEquals(t *testing.T) {
	t.Run("equal amplitude embeddings", func(t *testing.T) {
		embedding1, _ := NewAmplitudeEmbedding([]Qubit{0, 1}, true, 0.0)
		embedding2, _ := NewAmplitudeEmbedding([]Qubit{0, 1}, true, 0.0)
		assert.True(t, embedding1.Equals(embedding2))
	})

	t.Run("unequal amplitude embeddings", func(t *testing.T) {
		embedding1, _ := NewAmplitudeEmbedding([]Qubit{0, 1}, true, 0.0)
		embedding2, _ := NewAmplitudeEmbedding([]Qubit{1, 2}, false, 1.0)
		assert.False(t, embedding1.Equals(embedding2))
	})
}

func TestAmplitudeEmbeddingValidateQubits(t *testing.T) {
	t.Run("valid qubit cases", func(t *testing.T) {
		embedding, err := NewAmplitudeEmbedding([]Qubit{0, 1}, true, 0.0)
		require.NoError(t, err)
		err = embedding.validateQubits(3)
		assert.NoError(t, err)
	})

	t.Run("invalid qubit cases", func(t *testing.T) {
		embedding, err := NewAmplitudeEmbedding([]Qubit{0, 1}, true, 0.0)
		require.NoError(t, err)
		err = embedding.validateQubits(1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidQubit)
	})
}
