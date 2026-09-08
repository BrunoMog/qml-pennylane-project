package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func validEmbedding() Embedding {
	qubitZero, err := NewQubit(0, 1)
	if err != nil {
		panic(err)
	}
	qubits := []Qubit{qubitZero}
	embedding, err := NewAngleEmbedding(qubits, XRotation)
	if err != nil {
		panic(err)
	}
	return embedding
}

func validMeasurement() Measurement {
	qubitZero, err := NewQubit(0, 1)
	if err != nil {
		panic(err)
	}
	qubits := []Qubit{qubitZero}
	measurement, err := NewMeasurement(qubits, ExpectationMeasurement, XMeasurementRotation)
	if err != nil {
		panic(err)
	}
	return measurement
}

func TestNewVQC(t *testing.T) {
	tests := []struct {
		embedding   Embedding
		measurement Measurement
		testName    string
		pre_layer   Layer
		layer       Layer
		post_layer  Layer
		num_qubits  uint
		num_layers  uint
		expectErr   error
	}{
		{
			testName:    "valid VQC",
			num_qubits:  2,
			embedding:   validEmbedding(),
			measurement: validMeasurement(),
			num_layers:  1,
			expectErr:   nil,
		},
		{
			testName:    "zero qubits",
			num_qubits:  0,
			embedding:   validEmbedding(),
			measurement: validMeasurement(),
			num_layers:  1,
			expectErr:   &ZeroQubitVQCError{},
		},
		{
			testName:    "nil embedding",
			num_qubits:  2,
			embedding:   nil,
			measurement: validMeasurement(),
			num_layers:  1,
			expectErr:   &NilEmbeddingError{},
		},
		{
			testName:    "invalid embedding",
			num_qubits:  2,
			embedding:   AngleEmbedding{},
			measurement: validMeasurement(),
			num_layers:  1,
			expectErr:   &InvalidEmbeddingError{},
		},
		{
			testName:    "invalid measurement",
			num_qubits:  2,
			embedding:   validEmbedding(),
			measurement: Measurement{},
			num_layers:  1,
			expectErr:   &InvalidMeasurementError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			input := VQCBaseInput{
				Embedding:   tt.embedding,
				Measurement: tt.measurement,
				NumQubits:   tt.num_qubits,
				NumLayers:   tt.num_layers,
			}
			vqc, err := NewVQC(input)
			if tt.expectErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.num_qubits, vqc.NumQubits())
				assert.Equal(t, tt.num_layers, vqc.NumLayers())
				assert.Equal(t, tt.embedding, vqc.Embedding())
				assert.Equal(t, tt.pre_layer, vqc.PreLayer())
				assert.Equal(t, tt.layer, vqc.Layer())
				assert.Equal(t, tt.post_layer, vqc.PostLayer())
			}
		})
	}
}
