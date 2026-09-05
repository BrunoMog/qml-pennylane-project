package vqc

import "slices"

type GateType string

const (
	HGate    GateType = "h"
	XGate    GateType = "x"
	YGate    GateType = "y"
	ZGate    GateType = "z"
	RXGate   GateType = "rx"
	RYGate   GateType = "ry"
	RZGate   GateType = "rz"
	CNOTGate GateType = "cnot"
)

type QuantumGate struct {
	gate_type     GateType
	control_qubit []Qubit
	qubit         Qubit
}

func NewQuantumGate(gate_type GateType, qubit Qubit, control_qubit []Qubit) (*QuantumGate, error) {
	err := validateGate(gate_type, qubit, control_qubit)
	if err != nil {
		return nil, err
	}

	return &QuantumGate{
		gate_type:     gate_type,
		qubit:         qubit,
		control_qubit: control_qubit,
	}, nil
}

func validateGate(gate_type GateType, qubit Qubit, control_qubit []Qubit) error {
	if !isPermittedGate(gate_type) {
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

func isPermittedGate(gate_type GateType) bool {
	switch gate_type {
	case HGate, XGate, YGate, ZGate, RXGate, RYGate, RZGate, CNOTGate:
		return true
	default:
		return false
	}
}

func (q QuantumGate) Equal(other QuantumGate) bool {
	if q.gate_type != other.gate_type || q.qubit != other.qubit {
		return false
	}

	if !slices.Equal(q.control_qubit, other.control_qubit) {
		return false
	}

	return true
}

func (q QuantumGate) HasParameters() bool {
	switch q.gate_type {
	case RXGate, RYGate, RZGate:
		return true
	default:
		return false
	}
}

func (q QuantumGate) GateType() GateType {
	return q.gate_type
}

func (q QuantumGate) Qubit() Qubit {
	return q.qubit
}

func (q QuantumGate) ControlQubits() []Qubit {
	return slices.Clone(q.control_qubit)
}
