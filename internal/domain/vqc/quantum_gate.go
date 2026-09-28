package vqc

import (
	"errors"
	"fmt"
	"slices"
)

type QuantumGate struct {
	gateType     GateType
	controlQubit []Qubit
	qubit        Qubit
}

func NewQuantumGate(gateType GateType, qubit Qubit, controlQubit []Qubit) (QuantumGate, error) {
	err := validateGate(gateType, qubit, controlQubit)
	if err != nil {
		if errors.Is(err, ErrDuplicatedQubit) {
			return QuantumGate{}, fmt.Errorf("quantum gate has %w", err)
		}
		return QuantumGate{}, err
	}

	return QuantumGate{
		gateType:     gateType,
		qubit:        qubit,
		controlQubit: slices.Clone(controlQubit),
	}, nil
}

func validateGate(gateType GateType, qubit Qubit, controlQubit []Qubit) error {
	if !gateType.isValid() {
		return ErrInvalidGateType
	}

	if gateType.isSingleQubitGate() && len(controlQubit) > 0 {
		return &InvalidControlQubitError{Reason: fmt.Sprintf("single-qubit gate %s cannot have %d control qubits", gateType, len(controlQubit))}
	} else if gateType.isTwoQubitGate() && len(controlQubit) != 1 {
		return &InvalidControlQubitError{Reason: fmt.Sprintf("two-qubit gate %s must have exactly 1 control qubit", gateType)}
	}

	allQubits := append([]Qubit{qubit}, controlQubit...)
	if duplicatedQubit, duplicated := hasDuplicateQubits(allQubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit.Index()}
	}

	return nil
}

func (q QuantumGate) isValid() bool {
	return q.gateType != ""
}

func (q QuantumGate) validateQubits(numQubits uint) error {
	if q.qubit.Index() >= numQubits {
		return &InvalidQubitError{QubitIndex: q.qubit.Index()}
	}

	for _, control := range q.controlQubit {
		if control.Index() >= numQubits {
			return &InvalidQubitError{QubitIndex: control.Index()}
		}
	}

	return nil
}

func (q QuantumGate) Equals(other QuantumGate) bool {
	if q.gateType != other.gateType || q.qubit != other.qubit {
		return false
	}

	return slices.Equal(q.controlQubit, other.controlQubit)
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
