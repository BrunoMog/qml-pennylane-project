package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithPreLayer(t *testing.T) {

	t.Run("valid pre-layer", func(t *testing.T) {
		layer, err := NewLayer([]QuantumGate{{gateType: HGate}})
		require.NoError(t, err)

		inputVQC := VQCBaseInput{
			Embedding:   validEmbeddingOneQubit(t),
			Measurement: validMeasurementOneQubit(t),
			NumQubits:   1,
			NumLayers:   1,
		}

		vqc, err := NewVQC(inputVQC, WithPreLayer(layer))
		assert.NoError(t, err)
		assert.Equal(t, layer, vqc.PreLayer())
	})

	t.Run("invalid pre-layer", func(t *testing.T) {
		layer, err := NewLayer([]QuantumGate{{gateType: HGate, qubit: Qubit(1)}})
		require.NoError(t, err)

		inputVQC := VQCBaseInput{
			Embedding:   validEmbeddingOneQubit(t),
			Measurement: validMeasurementOneQubit(t),
			NumQubits:   1,
			NumLayers:   1,
		}
		vqc, err := NewVQC(inputVQC, WithPreLayer(layer))
		assert.Error(t, err)
		assert.Equal(t, VQC{}, vqc)
	})
}

func TestWithLayer(t *testing.T) {

	t.Run("valid layer", func(t *testing.T) {
		layer, err := NewLayer([]QuantumGate{{gateType: XGate}})
		require.NoError(t, err)

		inputVQC := VQCBaseInput{
			Embedding:   validEmbeddingOneQubit(t),
			Measurement: validMeasurementOneQubit(t),
			NumQubits:   2,
			NumLayers:   1,
		}

		vqc, err := NewVQC(inputVQC, WithLayer(layer))
		assert.NoError(t, err)
		assert.Equal(t, layer, vqc.Layer())
	})

	t.Run("invalid layer", func(t *testing.T) {
		layer, err := NewLayer([]QuantumGate{{gateType: XGate, qubit: Qubit(2)}})
		require.NoError(t, err)

		inputVQC := VQCBaseInput{
			Embedding:   validEmbeddingOneQubit(t),
			Measurement: validMeasurementOneQubit(t),
			NumQubits:   2,
			NumLayers:   1,
		}
		vqc, err := NewVQC(inputVQC, WithLayer(layer))
		assert.Error(t, err)
		assert.Equal(t, VQC{}, vqc)
	})
}

func TestWithPostLayer(t *testing.T) {

	t.Run("valid post-layer", func(t *testing.T) {
		layer, err := NewLayer([]QuantumGate{{gateType: CNOTGate}})
		require.NoError(t, err)

		inputVQC := VQCBaseInput{
			Embedding:   validEmbeddingOneQubit(t),
			Measurement: validMeasurementOneQubit(t),
			NumQubits:   2,
			NumLayers:   1,
		}

		vqc, err := NewVQC(inputVQC, WithPostLayer(layer))
		assert.NoError(t, err)
		assert.Equal(t, layer, vqc.PostLayer())
	})

	t.Run("invalid post-layer", func(t *testing.T) {
		layer, err := NewLayer([]QuantumGate{{gateType: CNOTGate, qubit: Qubit(2)}})
		require.NoError(t, err)

		inputVQC := VQCBaseInput{
			Embedding:   validEmbeddingOneQubit(t),
			Measurement: validMeasurementOneQubit(t),
			NumQubits:   2,
			NumLayers:   1,
		}
		vqc, err := NewVQC(inputVQC, WithPostLayer(layer))
		assert.Error(t, err)
		assert.Equal(t, VQC{}, vqc)
	})
}

func TestWithMultipleOptions(t *testing.T) {
	preLayer, err := NewLayer([]QuantumGate{{gateType: HGate}})
	require.NoError(t, err)
	layer, err := NewLayer([]QuantumGate{{gateType: XGate}})
	require.NoError(t, err)
	postLayer, err := NewLayer([]QuantumGate{{gateType: CNOTGate}})
	require.NoError(t, err)

	inputVQC := VQCBaseInput{
		Embedding:   validEmbeddingOneQubit(t),
		Measurement: validMeasurementOneQubit(t),
		NumQubits:   2,
		NumLayers:   1,
	}

	vqc, err := NewVQC(inputVQC, WithPreLayer(preLayer), WithLayer(layer), WithPostLayer(postLayer))
	assert.NoError(t, err)
	assert.Equal(t, preLayer, vqc.PreLayer())
	assert.Equal(t, layer, vqc.Layer())
	assert.Equal(t, postLayer, vqc.PostLayer())
}
