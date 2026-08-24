package vqc

import (
	"testing"
)

func TestNewAngleEmbedding(t *testing.T) {
	testCases := []struct {
		name      string
		qubits    []Qubit
		rotation  EmbeddingRotation
		expectErr bool
	}{
		{"Valid angle embedding", []Qubit{0, 1}, XRotation, false},
		{"Invalid rotation", []Qubit{0, 1}, EmbeddingRotation("invalid_rotation"), true},
		{"Duplicate qubit index", []Qubit{0, 1, 1}, XRotation, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewAngleEmbedding(tc.qubits, tc.rotation)
			if (err != nil) != tc.expectErr {
				t.Errorf("NewAngleEmbedding(%v, %s) returned error: %v, expected error: %v", tc.qubits, tc.rotation, err != nil, tc.expectErr)
			}
		})
	}
}

func TestNewAmplitudeEmbedding(t *testing.T) {
	testCases := []struct {
		name      string
		qubits    []Qubit
		normalize bool
		padwith   float64
		expectErr bool
	}{
		{"Valid amplitude embedding", []Qubit{0, 1}, true, 0.0, false},
		{"Duplicate qubit index", []Qubit{0, 1, 1}, true, 0.0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewAmplitudeEmbedding(tc.qubits, tc.normalize, tc.padwith)
			if (err != nil) != tc.expectErr {
				t.Errorf("NewAmplitudeEmbedding(%v, %v, %f) returned error: %v, expected error: %v", tc.qubits, tc.normalize, tc.padwith, err != nil, tc.expectErr)
			}
		})
	}
}
