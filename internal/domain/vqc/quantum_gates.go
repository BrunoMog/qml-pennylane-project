package vqc

import (
	"slices"
)

type QuantumGate struct {
	gateType     GateType
	controlQubit []Qubit
	qubit        Qubit
}

func NewQuantumGate(gateType GateType, qubit Qubit, controlQubit []Qubit) (QuantumGate, error) {
	controlQubit = slices.Clone(controlQubit)
	err := validateGate(gateType, qubit, controlQubit)
	if err != nil {
		return QuantumGate{}, err
	}

	return QuantumGate{
		gateType:     gateType,
		qubit:        qubit,
		controlQubit: controlQubit,
	}, nil
}

func validateGate(gateType GateType, qubit Qubit, controlQubit []Qubit) error {
	if !isValidGate(gateType) {
		return &InvalidGateError{gateType}
	}

	if slices.Contains(singleQubitGates, gateType) && len(controlQubit) > 0 {
		return &InvalidControlQubitError{controlQubit}
	} else if slices.Contains(twoQubitGates, gateType) && len(controlQubit) != 1 {
		return &InvalidControlQubitError{controlQubit}
	}

	allQubits := append([]Qubit{qubit}, controlQubit...)
	if duplicatedQubit, duplicated := hasDuplicateQubits(allQubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit}
	}

	return nil
}

func isValidGate(gateType GateType) bool {
	switch gateType {
	case HGate, XGate, YGate, ZGate, RXGate, RYGate, RZGate, CNOTGate:
		return true
	default:
		return false
	}
}

func (q QuantumGate) Equals(other QuantumGate) bool {
	if q.gateType != other.gateType || q.qubit != other.qubit {
		return false
	}

	if !slices.Equal(q.controlQubit, other.controlQubit) {
		return false
	}

	return true
}

func (q QuantumGate) HasParameters() bool {
	switch q.gateType {
	case RXGate, RYGate, RZGate:
		return true
	default:
		return false
	}
}

func (q QuantumGate) Clone() QuantumGate {
	return QuantumGate{
		gateType:     q.gateType,
		qubit:        q.qubit,
		controlQubit: slices.Clone(q.controlQubit),
	}
}

func (q QuantumGate) GateType() GateType {
	return q.gateType
}

func (q QuantumGate) Qubit() Qubit {
	return q.qubit
}

func (q QuantumGate) ControlQubits() []Qubit {
	return slices.Clone(q.controlQubit)
}
