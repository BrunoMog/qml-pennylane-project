package vqc

import (
	"testing"
)

func TestDuplicateQubits(t *testing.T) {
	testCase := []struct {
		name                 string
		qubits               []Qubit
		expectedDuplicate    Qubit
		expectedHasDuplicate bool
	}{
		{"Duplicate qubit", []Qubit{0, 1, 2, 1}, 1, true},
		{"No duplicate qubits", []Qubit{0, 1, 2}, 0, false},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			duplicatedQubit, duplicated := hasDuplicateQubits(tc.qubits)
			if duplicated != tc.expectedHasDuplicate || duplicatedQubit != tc.expectedDuplicate {
				t.Errorf("For qubits %v: expected duplicate %v (has duplicate: %v), got duplicate %v (has duplicate: %v)",
					tc.qubits, tc.expectedDuplicate, tc.expectedHasDuplicate, duplicatedQubit, duplicated)
			}
		})
	}
}

func TestNewQubit(t *testing.T) {
	testCases := []struct {
		name      string
		index     int
		numQubits uint
		expectErr bool
	}{
		{"Valid qubit", 1, 3, false},
		{"Invalid qubit (out of range)", 3, 3, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewQubit(tc.index, tc.numQubits)
			if (err != nil) != tc.expectErr {
				t.Errorf("NewQubit(%d, %d) returned error: %v, expected error: %v", tc.index, tc.numQubits, err != nil, tc.expectErr)
			}
		})
	}
}
