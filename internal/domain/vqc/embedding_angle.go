package vqc

import (
	"errors"
	"fmt"
	"slices"
)

type AngleEmbedding struct {
	rotation EmbeddingRotation
	qubits   []Qubit
}

func NewAngleEmbedding(qubits []Qubit, rotation EmbeddingRotation) (AngleEmbedding, error) {
	if !rotation.isValid() {
		return AngleEmbedding{}, ErrInvalidEmbeddingRotation
	}
	if err := validateEmbeddingQubits(qubits); err != nil {
		if errors.Is(err, ErrDuplicatedQubit) {
			return AngleEmbedding{}, fmt.Errorf("embedding has %w", err)
		}
		return AngleEmbedding{}, err
	}

	return AngleEmbedding{qubits: slices.Clone(qubits), rotation: rotation}, nil
}

func (a AngleEmbedding) isValid() bool {
	return len(a.qubits) > 0 && a.rotation != ""
}

func (a AngleEmbedding) validateQubits(numQubits uint) error {
	for _, qubit := range a.qubits {
		if qubit.Index() >= numQubits {
			return &InvalidQubitError{QubitIndex: qubit.Index()}
		}
	}
	return nil
}

func (a AngleEmbedding) Equals(other Embedding) bool {
	otherAngle, ok := other.(AngleEmbedding)
	if !ok {
		return false
	}

	if a.rotation != otherAngle.rotation {
		return false
	}

	if len(a.qubits) != len(otherAngle.qubits) {
		return false
	}

	return slices.Equal(a.qubits, otherAngle.qubits)
}

func (a AngleEmbedding) Type() EmbeddingType {
	return EmbeddingTypeAngle
}

func (a AngleEmbedding) Qubits() []Qubit {
	return slices.Clone(a.qubits)
}

func (a AngleEmbedding) Rotation() EmbeddingRotation {
	return a.rotation
}

func (a AngleEmbedding) isEmbedding() {}
