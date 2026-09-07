package vqc

type Embedding interface {
	Type() EmbeddingType
	Qubits() []Qubit

	isEmbedding()
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
