package vqc

import "slices"

type AngleEmbedding struct {
	rotation EmbeddingRotation
	qubits   []Qubit
}

func NewAngleEmbedding(qubits []Qubit, rotation EmbeddingRotation) (AngleEmbedding, error) {
	if !rotation.IsValid() {
		return AngleEmbedding{}, &InvalidRotationError{rotation}
	}
	if err := validateEmbeddingQubits(qubits); err != nil {
		return AngleEmbedding{}, err
	}

	return AngleEmbedding{qubits: qubits, rotation: rotation}, nil
}

func (a AngleEmbedding) IsValid() bool {
	if !a.rotation.IsValid() {
		return false
	}
	if err := validateEmbeddingQubits(a.qubits); err != nil {
		return false
	}
	return true
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

	for i, qubit := range a.qubits {
		if qubit != otherAngle.qubits[i] {
			return false
		}
	}

	return true
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
