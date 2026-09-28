package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateQubits(t *testing.T) {

	t.Run("has duplicate qubits", func(t *testing.T) {
		duplicatedQubit, duplicated := hasDuplicateQubits([]Qubit{0, 1, 2, 1})
		assert.True(t, duplicated)
		assert.Equal(t, Qubit(1), duplicatedQubit)
	})

	t.Run("no duplicate qubits", func(t *testing.T) {
		duplicatedQubit, duplicated := hasDuplicateQubits([]Qubit{0, 1, 2})
		assert.False(t, duplicated)
		assert.Equal(t, Qubit(0), duplicatedQubit) // Default value when no duplicate is found
	})
}

func TestNewQubit(t *testing.T) {

	t.Run("Valid qubit", func(t *testing.T) {
		qubit, err := NewQubit(1, 3)
		assert.NoError(t, err)
		assert.Equal(t, Qubit(1), qubit)
	})

	t.Run("Invalid qubit (out of range)", func(t *testing.T) {
		_, err := NewQubit(3, 3)
		assert.Error(t, err)
		assert.IsType(t, ErrInvalidQubit, err)
	})
}
