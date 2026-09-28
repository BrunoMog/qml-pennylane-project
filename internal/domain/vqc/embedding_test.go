package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmbeddingQubits(t *testing.T) {
	t.Run("valid qubit cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			qubits   []Qubit
		}{
			{"valid qubits with no duplicates", []Qubit{0, 1, 2}},
			{"valid qubits with single qubit", []Qubit{3}},
			{"valid qubits with non-consecutive indices", []Qubit{5, 7, 9}},
		}
		for _, testCase := range inputs {
			t.Run(testCase.testName, func(t *testing.T) {
				err := validateEmbeddingQubits(testCase.qubits)
				assert.NoError(t, err)
			})
		}
	})

	t.Run("invalid qubit cases", func(t *testing.T) {
		inputs := []struct {
			testName  string
			qubits    []Qubit
			expectErr error
		}{
			{"duplicate qubits", []Qubit{0, 1, 1}, ErrDuplicatedQubit},
			{"zero qubits", []Qubit{}, ErrZeroQubitEmbedding},
			{"nil qubits", nil, ErrZeroQubitEmbedding},
		}

		for _, testCase := range inputs {
			t.Run(testCase.testName, func(t *testing.T) {
				err := validateEmbeddingQubits(testCase.qubits)
				assert.Error(t, err)
				assert.ErrorIs(t, err, testCase.expectErr)
			})
		}
	})
}
