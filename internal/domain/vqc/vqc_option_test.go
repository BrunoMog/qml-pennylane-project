package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithPreLayer(t *testing.T) {
	layer := NewLayer([]QuantumGate{{gateType: HGate}})

	inputVQC := VQCBaseInput{
		Embedding:   validEmbedding(),
		Measurement: validMeasurement(),
		NumQubits:   2,
		NumLayers:   1,
	}

	vqc, err := NewVQC(inputVQC, WithPreLayer(layer))
	assert.NoError(t, err)
	assert.Equal(t, layer, vqc.PreLayer())
}

func TestWithLayer(t *testing.T) {
	layer := NewLayer([]QuantumGate{{gateType: XGate}})

	inputVQC := VQCBaseInput{
		Embedding:   validEmbedding(),
		Measurement: validMeasurement(),
		NumQubits:   2,
		NumLayers:   1,
	}

	vqc, err := NewVQC(inputVQC, WithLayer(layer))
	assert.NoError(t, err)
	assert.Equal(t, layer, vqc.Layer())
}

func TestWithPostLayer(t *testing.T) {
	layer := NewLayer([]QuantumGate{{gateType: CNOTGate}})

	inputVQC := VQCBaseInput{
		Embedding:   validEmbedding(),
		Measurement: validMeasurement(),
		NumQubits:   2,
		NumLayers:   1,
	}

	vqc, err := NewVQC(inputVQC, WithPostLayer(layer))
	assert.NoError(t, err)
	assert.Equal(t, layer, vqc.PostLayer())
}

func TestWithMultipleOptions(t *testing.T) {
	preLayer := NewLayer([]QuantumGate{{gateType: HGate}})
	layer := NewLayer([]QuantumGate{{gateType: XGate}})
	postLayer := NewLayer([]QuantumGate{{gateType: CNOTGate}})

	inputVQC := VQCBaseInput{
		Embedding:   validEmbedding(),
		Measurement: validMeasurement(),
		NumQubits:   2,
		NumLayers:   1,
	}

	vqc, err := NewVQC(inputVQC, WithPreLayer(preLayer), WithLayer(layer), WithPostLayer(postLayer))
	assert.NoError(t, err)
	assert.Equal(t, preLayer, vqc.PreLayer())
	assert.Equal(t, layer, vqc.Layer())
	assert.Equal(t, postLayer, vqc.PostLayer())
}
