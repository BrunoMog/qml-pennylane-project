package vqc

import (
	"slices"
)

type Embedding interface {
	GetType() EmbeddingType
	GetQubits() []Qubit

	isEmbedding()
}

type EmbeddingType string

const (
	EmbeddingTypeAngle     EmbeddingType = "angle"
	EmbeddingTypeAmplitude EmbeddingType = "amplitude"
)

type EmbeddingRotation string

const (
	XRotation EmbeddingRotation = "x"
	YRotation EmbeddingRotation = "y"
	ZRotation EmbeddingRotation = "z"
)

func (e EmbeddingRotation) IsValid() bool {
	switch e {
	case XRotation, YRotation, ZRotation:
		return true
	default:
		return false
	}
}

type AngleEmbedding struct {
	qubits   []Qubit
	rotation EmbeddingRotation
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

func validateEmbeddingQubits(qubits []Qubit) error {
	if len(qubits) == 0 {
		return &ZeroQubitEmbeddingError{qubit: qubits}
	}
	if qubit, ok := hasDuplicateQubits(qubits); ok {
		return &DuplicateQubitError{qubit: qubit}
	}
	return nil
}

func (a AngleEmbedding) GetType() EmbeddingType {
	return EmbeddingTypeAngle
}

func (a AngleEmbedding) GetQubits() []Qubit {
	return slices.Clone(a.qubits)
}

func (a AngleEmbedding) GetRotation() EmbeddingRotation {
	return a.rotation
}

func (a AngleEmbedding) isEmbedding() {}

type AmplitudeEmbedding struct {
	qubits    []Qubit
	normalize bool
	padWidth  float64
}

func NewAmplitudeEmbedding(qubits []Qubit, normalize bool, padWidth float64) (AmplitudeEmbedding, error) {
	if err := validateEmbeddingQubits(qubits); err != nil {
		return AmplitudeEmbedding{}, err
	}

	return AmplitudeEmbedding{qubits: qubits, normalize: normalize, padWidth: padWidth}, nil
}

func (a AmplitudeEmbedding) GetType() EmbeddingType {
	return EmbeddingTypeAmplitude
}

func (a AmplitudeEmbedding) GetQubits() []Qubit {
	return slices.Clone(a.qubits)
}

func (a AmplitudeEmbedding) GetNormalize() bool {
	return a.normalize
}

func (a AmplitudeEmbedding) GetPadWidth() float64 {
	return a.padWidth
}

func (a AmplitudeEmbedding) isEmbedding() {}
