package vqc

type Qubit uint

func hasDuplicateQubits(qubits []Qubit) (Qubit, bool) {
	seen := make(map[Qubit]struct{}, len(qubits))

	for _, qubit := range qubits {
		if _, exists := seen[qubit]; exists {
			return qubit, true
		}
		seen[qubit] = struct{}{}
	}

	return 0, false
}

func NewQubit(index uint, num_qubits uint) (Qubit, error) {
	err := validateQubit(index, num_qubits)
	if err != nil {
		return 0, err
	}
	return Qubit(index), nil
}

func validateQubit(qubit uint, num_qubits uint) error {
	if qubit >= num_qubits {
		return &InvalidQubitError{qubit}
	}
	return nil
}
