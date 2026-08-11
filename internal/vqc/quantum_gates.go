package vqc

import (
	"slices"
)

type QuantumGates struct {
	gate_type     string
	qubit         Qubit
	control_qubit []Qubit
}

var permited_gates = []string{"h", "x", "y", "rx", "ry", "rz", "cnot"}

func NewQuantumGate(gate_type string, qubit Qubit, control_qubit []Qubit) (*QuantumGates, error) {
	err := validateGate(gate_type, qubit, control_qubit)
	if err != nil {
		return nil, err
	}

	return &QuantumGates{
		gate_type:     gate_type,
		qubit:         qubit,
		control_qubit: control_qubit,
	}, nil
}

func validateGate(gate_type string, qubit Qubit, control_qubit []Qubit) error {
	if !isPermitedGate(gate_type) {
		return &InvalidGateError{gate_type}
	}

	if gate_type == "cnot" {
		if len(control_qubit) != 1 {
			return &InvalidControlQubitError{control_qubit}
		}
	} else if len(control_qubit) != 0 {
		return &InvalidControlQubitError{control_qubit}
	}

	allQubits := append([]Qubit{qubit}, control_qubit...)
	if duplicatedQubit, duplicated := hasDuplicateQubits(allQubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit}
	}

	return nil
}

func isPermitedGate(gate_type string) bool {
	result := slices.Contains(permited_gates, gate_type)
	return result
}
