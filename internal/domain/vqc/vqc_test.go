package vqc

import (
	"testing"
)

func TestNewVQC(t *testing.T) {
	tests := []struct {
		embedding   Embedding
		measurement Measurement
		name        string
		pre_layer   Layer
		layer       Layer
		post_layer  Layer
		num_qubits  uint
		num_layers  uint
		expectErr   bool
	}{
		{
			name:        "valid VQC",
			num_qubits:  2,
			embedding:   AngleEmbedding{},
			pre_layer:   Layer{},
			layer:       Layer{},
			post_layer:  Layer{},
			measurement: Measurement{},
			num_layers:  1,
			expectErr:   false,
		},
		{
			name:        "zero qubits",
			num_qubits:  0,
			embedding:   AngleEmbedding{},
			pre_layer:   Layer{},
			layer:       Layer{},
			post_layer:  Layer{},
			measurement: Measurement{},
			num_layers:  1,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := VQCInput{
				embedding:   tt.embedding,
				measurement: tt.measurement,
				pre_layer:   tt.pre_layer,
				layer:       tt.layer,
				post_layer:  tt.post_layer,
				num_qubits:  tt.num_qubits,
				num_layers:  tt.num_layers,
			}
			vqc, err := NewVQC(input)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewVQC() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr && vqc.NumQubits() != tt.num_qubits {
				t.Errorf("NewVQC() num_qubits = %v, want %v", vqc.NumQubits(), tt.num_qubits)
			}
		})
	}
}
