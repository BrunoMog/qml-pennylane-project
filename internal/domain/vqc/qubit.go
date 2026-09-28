package vqc

type Qubit uint

func NewQubit(qubit uint, numQubits uint) (Qubit, error) {
	if qubit >= numQubits {
		return 0, ErrInvalidQubit
	}
	return Qubit(qubit), nil
}

func hasDuplicateQubits(qubits []Qubit) (Qubit, bool) {
	seen := make(map[Qubit]bool, len(qubits))

	for _, qubit := range qubits {
		if exists := seen[qubit]; exists {
			return qubit, true
		}
		seen[qubit] = true
	}

	return 0, false
}

func (q Qubit) Index() uint {
	return uint(q)
}
