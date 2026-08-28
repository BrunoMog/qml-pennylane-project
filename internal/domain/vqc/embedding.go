package vqc

import "slices"

type Embedding struct {
	embeddingType EmbeddingType
	qubits        []Qubit
	rotation      EmbeddingRotation
	normalize     bool
	padwith       float64
}

type EmbeddingType string

const (
	AngleEmbedding     EmbeddingType = "angle"
	AmplitudeEmbedding EmbeddingType = "amplitude"
)

type EmbeddingRotation string

const (
	XRotation EmbeddingRotation = "x"
	YRotation EmbeddingRotation = "y"
	ZRotation EmbeddingRotation = "z"
)

func NewAngleEmbedding(qubits []Qubit, rotation EmbeddingRotation) (*Embedding, error) {
	err := validateEmbedding(AngleEmbedding, qubits, rotation)
	if err != nil {
		return nil, err
	}

	return &Embedding{
		embeddingType: AngleEmbedding,
		qubits:        qubits,
		rotation:      rotation,
	}, nil
}

func NewAmplitudeEmbedding(qubits []Qubit, normalize bool, padwith float64) (*Embedding, error) {
	err := validateEmbedding(AmplitudeEmbedding, qubits, "")
	if err != nil {
		return nil, err
	}

	return &Embedding{
		embeddingType: AmplitudeEmbedding,
		qubits:        qubits,
		normalize:     normalize,
		padwith:       padwith,
	}, nil
}

func validateEmbedding(embeddingType EmbeddingType, qubits []Qubit, rotation EmbeddingRotation) error {
	if !isPermitedEmbedding(embeddingType) {
		return &InvalidEmbeddingError{embeddingType}
	}

	if len(qubits) == 0 {
		return &ZeroQubitEmbeddingError{qubits}
	}

	if duplicatedQubit, duplicated := hasDuplicateQubits(qubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit}
	}

	if !isPermitedRotation(rotation) && embeddingType == AngleEmbedding {
		return &InvalidRotationError{rotation}
	}

	return nil
}

func isPermitedEmbedding(embeddingType EmbeddingType) bool {
	switch embeddingType {
	case AngleEmbedding, AmplitudeEmbedding:
		return true
	default:
		return false
	}
}

func isPermitedRotation(rotation EmbeddingRotation) bool {
	switch rotation {
	case XRotation, YRotation, ZRotation:
		return true
	default:
		return false
	}
}

func (e Embedding) GetType() EmbeddingType {
	return e.embeddingType
}

func (e Embedding) GetQubits() []Qubit {
	return slices.Clone(e.qubits)
}

func (e Embedding) GetRotation() (EmbeddingRotation, error) {
	if e.GetType() != AngleEmbedding {
		return "", &InvalidGetRotationError{}
	}
	return e.rotation, nil
}

func (e Embedding) GetNormalize() (bool, error) {
	if e.GetType() != AmplitudeEmbedding {
		return false, &InvalidGetNormalizeError{}
	}
	return e.normalize, nil
}

func (e Embedding) GetPadWith() (float64, error) {
	if e.GetType() != AmplitudeEmbedding {
		return 0, &InvalidGetPadWithError{}
	}
	return e.padwith, nil
}
