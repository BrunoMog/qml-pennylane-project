package vqc

import (
	"testing"
)

func TestNewAngleEmbedding(t *testing.T) {
	testCases := []struct {
		name      string
		rotation  EmbeddingRotation
		qubits    []Qubit
		expectErr bool
	}{
		{name: "Valid angle embedding", qubits: []Qubit{0, 1}, rotation: XRotation, expectErr: false},
		{name: "Invalid rotation", qubits: []Qubit{0, 1}, rotation: EmbeddingRotation("invalid_rotation"), expectErr: true},
		{name: "Duplicate qubit index", qubits: []Qubit{0, 1, 1}, rotation: XRotation, expectErr: true},
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
		padwith   float64
		normalize bool
		expectErr bool
	}{
		{name: "Valid amplitude embedding", qubits: []Qubit{0, 1}, normalize: true, padwith: 0.0, expectErr: false},
		{name: "Duplicate qubit index", qubits: []Qubit{0, 1, 1}, padwith: 0.0, normalize: true, expectErr: true},
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
