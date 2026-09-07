package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateQubits(t *testing.T) {
	testCase := []struct {
		testName             string
		qubits               []Qubit
		expectedDuplicate    Qubit
		expectedHasDuplicate bool
	}{
		{"Duplicate qubit", []Qubit{0, 1, 2, 1}, 1, true},
		{"No duplicate qubits", []Qubit{0, 1, 2}, 0, false},
	}

	for _, tc := range testCase {
		t.Run(tc.testName, func(t *testing.T) {
			duplicatedQubit, duplicated := hasDuplicateQubits(tc.qubits)
			if tc.expectedHasDuplicate {
				assert.True(t, duplicated, "Expected duplicate qubit but found none")
				assert.Equal(t, tc.expectedDuplicate, duplicatedQubit, "Expected duplicate qubit does not match actual duplicate qubit")
			} else {
				assert.False(t, duplicated, "Expected no duplicate qubits but found one")
			}
		})
	}
}

func TestNewQubit(t *testing.T) {
	testCases := []struct {
		testName  string
		index     uint
		numQubits uint
		expectErr error
	}{
		{"Valid qubit", 1, 3, nil},
		{"Invalid qubit (out of range)", 3, 3, &InvalidQubitError{}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			qubit, err := NewQubit(tc.index, tc.numQubits)
			if tc.expectErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, Qubit(tc.index), qubit)
			}
		})
	}
}
