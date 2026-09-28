package vqc

import (
	"errors"
	"fmt"
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
		if errors.Is(err, ErrDuplicatedQubit) {
			return AmplitudeEmbedding{}, fmt.Errorf("embedding has %w", err)
		}
		return AmplitudeEmbedding{}, err
	}
	if math.IsNaN(padWith) || math.IsInf(padWith, 0) {
		return AmplitudeEmbedding{}, ErrInvalidPadWith
	}

	return AmplitudeEmbedding{qubits: slices.Clone(qubits), normalize: normalize, padWith: padWith}, nil
}

func (a AmplitudeEmbedding) isValid() bool {
	return len(a.qubits) > 0
}

func (a AmplitudeEmbedding) validateQubits(numQubits uint) error {
	for _, qubit := range a.qubits {
		if qubit.Index() >= numQubits {
			return &InvalidQubitError{QubitIndex: qubit.Index()}
		}
	}
	return nil
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

	return slices.Equal(a.qubits, otherAmplitude.qubits)
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
