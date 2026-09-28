package vqc

type Embedding interface {
	Type() EmbeddingType
	Qubits() []Qubit
	Equals(other Embedding) bool
	isValid() bool
	validateQubits(numQubits uint) error

	isEmbedding()
}

func validateEmbeddingQubits(qubits []Qubit) error {
	if len(qubits) == 0 {
		return ErrZeroQubitEmbedding
	}
	if qubit, ok := hasDuplicateQubits(qubits); ok {
		return &DuplicateQubitError{QubitIndex: qubit.Index()}
	}
	return nil
}
