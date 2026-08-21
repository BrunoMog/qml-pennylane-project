package vqc

import "slices"

type GateType string

const (
	HGate    GateType = "h"
	XGate    GateType = "x"
	YGate    GateType = "y"
	RXGate   GateType = "rx"
	RYGate   GateType = "ry"
	RZGate   GateType = "rz"
	CNOTGate GateType = "cnot"
)

type QuantumGates struct {
	gate_type     GateType
	qubit         Qubit
	control_qubit []Qubit
}

func NewQuantumGate(gate_type GateType, qubit Qubit, control_qubit []Qubit) (*QuantumGates, error) {
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

func validateGate(gate_type GateType, qubit Qubit, control_qubit []Qubit) error {
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

func isPermitedGate(gate_type GateType) bool {
	switch gate_type {
	case HGate, XGate, YGate, RXGate, RYGate, RZGate, CNOTGate:
		return true
	default:
		return false
	}
}

func (q QuantumGates) Equal(other QuantumGates) bool {
	if q.gate_type != other.gate_type || q.qubit != other.qubit {
		return false
	}

	if !slices.Equal(q.control_qubit, other.control_qubit) {
		return false
	}

	return true
}
