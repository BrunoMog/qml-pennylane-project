package vqc

type Embedding struct {
	embedding_type EmbeddingType
	qubits         []Qubit
	rotation       EmbeddingRotation
	normalize      bool
	padwith        float64
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
		embedding_type: AngleEmbedding,
		qubits:         qubits,
		rotation:       rotation,
	}, nil
}

func NewAmplitudeEmbedding(qubits []Qubit, normalize bool, padwith float64) (*Embedding, error) {
	err := validateEmbedding(AmplitudeEmbedding, qubits, "")
	if err != nil {
		return nil, err
	}

	return &Embedding{
		embedding_type: AmplitudeEmbedding,
		qubits:         qubits,
		normalize:      normalize,
		padwith:        padwith,
	}, nil
}

func validateEmbedding(embedding_type EmbeddingType, qubits []Qubit, rotation EmbeddingRotation) error {
	if !isPermitedEmbedding(embedding_type) {
		return &InvalidEmbeddingError{embedding_type}
	}

	if len(qubits) == 0 {
		return &ZeroQubitEmbeddingError{qubits}
	}

	if duplicatedQubit, duplicated := hasDuplicateQubits(qubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit}
	}

	if !isPermitedRotation(rotation) && embedding_type == AngleEmbedding {
		return &InvalidRotationError{rotation}
	}

	return nil
}

func isPermitedEmbedding(embedding_type EmbeddingType) bool {
	switch embedding_type {
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
