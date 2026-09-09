package vqc

import (
	"math"
	"slices"
)

type AmplitudeEmbedding struct {
	qubits    []Qubit
	normalize bool
	padWith   float64
}

func NewAmplitudeEmbedding(qubits []Qubit, normalize bool, padWith float64) (AmplitudeEmbedding, error) {
	if err := validateEmbeddingQubits(qubits); err != nil {
		return AmplitudeEmbedding{}, err
	}
	if math.IsNaN(padWith) || math.IsInf(padWith, 0) {
		return AmplitudeEmbedding{}, &InvalidPadWithError{padWith: padWith}
	}

	return AmplitudeEmbedding{qubits: qubits, normalize: normalize, padWith: padWith}, nil
}

func (a AmplitudeEmbedding) IsValid() bool {
	if err := validateEmbeddingQubits(a.qubits); err != nil {
		return false
	}
	if math.IsNaN(a.padWith) || math.IsInf(a.padWith, 0) {
		return false
	}
	return true
}

func (a AmplitudeEmbedding) Equals(other Embedding) bool {
	otherAmplitude, ok := other.(AmplitudeEmbedding)
	if !ok {
		return false
	}

	if a.normalize != otherAmplitude.normalize {
		return false
	}

	if a.padWith != otherAmplitude.padWith {
		return false
	}

	if len(a.qubits) != len(otherAmplitude.qubits) {
		return false
	}

	for i, qubit := range a.qubits {
		if qubit != otherAmplitude.qubits[i] {
			return false
		}
	}

	return true
}

func (a AmplitudeEmbedding) Type() EmbeddingType {
	return EmbeddingTypeAmplitude
}

func (a AmplitudeEmbedding) Qubits() []Qubit {
	return slices.Clone(a.qubits)
}

func (a AmplitudeEmbedding) Normalize() bool {
	return a.normalize
}

func (a AmplitudeEmbedding) PadWith() float64 {
	return a.padWith
}

func (a AmplitudeEmbedding) isEmbedding() {}
