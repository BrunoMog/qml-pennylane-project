package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithPreLayer(t *testing.T) {
	layer := NewLayer([]QuantumGate{{gate_type: HGate}})

	inputVQC := VQCBaseInput{
		Embedding:   AmplitudeEmbedding{},
		Measurement: Measurement{},
		NumQubits:   2,
		NumLayers:   1,
	}

	vqc, err := NewVQC(inputVQC, WithPreLayer(layer))
	assert.NoError(t, err)
	assert.Equal(t, layer, vqc.PreLayer())
}

func TestWithLayer(t *testing.T) {
	layer := NewLayer([]QuantumGate{{gate_type: XGate}})

	inputVQC := VQCBaseInput{
		Embedding:   AmplitudeEmbedding{},
		Measurement: Measurement{},
		NumQubits:   2,
		NumLayers:   1,
	}

	vqc, err := NewVQC(inputVQC, WithLayer(layer))
	assert.NoError(t, err)
	assert.Equal(t, layer, vqc.Layer())
}

func TestWithPostLayer(t *testing.T) {
	layer := NewLayer([]QuantumGate{{gate_type: CNOTGate}})

	inputVQC := VQCBaseInput{
		Embedding:   AmplitudeEmbedding{},
		Measurement: Measurement{},
		NumQubits:   2,
		NumLayers:   1,
	}

	vqc, err := NewVQC(inputVQC, WithPostLayer(layer))
	assert.NoError(t, err)
	assert.Equal(t, layer, vqc.PostLayer())
}

func TestWithMultipleOptions(t *testing.T) {
	preLayer := NewLayer([]QuantumGate{{gate_type: HGate}})
	layer := NewLayer([]QuantumGate{{gate_type: XGate}})
	postLayer := NewLayer([]QuantumGate{{gate_type: CNOTGate}})

	inputVQC := VQCBaseInput{
		Embedding:   AmplitudeEmbedding{},
		Measurement: Measurement{},
		NumQubits:   2,
		NumLayers:   1,
	}

	vqc, err := NewVQC(inputVQC, WithPreLayer(preLayer), WithLayer(layer), WithPostLayer(postLayer))
	assert.NoError(t, err)
	assert.Equal(t, preLayer, vqc.PreLayer())
	assert.Equal(t, layer, vqc.Layer())
	assert.Equal(t, postLayer, vqc.PostLayer())
}
