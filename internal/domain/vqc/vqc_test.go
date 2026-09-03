package vqc

import (
	"testing"
)

func TestNewVQC(t *testing.T) {
	tests := []struct {
		name        string
		num_qubits  uint
		embedding   Embedding
		pre_layer   Layer
		layer       Layer
		post_layer  Layer
		measurement Measurement
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
			vqc, err := NewVQC(tt.num_qubits, tt.embedding, tt.pre_layer, tt.layer, tt.post_layer, tt.measurement, tt.num_layers)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewVQC() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr && vqc.GetNumQubits() != tt.num_qubits {
				t.Errorf("NewVQC() num_qubits = %v, want %v", vqc.GetNumQubits(), tt.num_qubits)
			}
		})
	}
}
