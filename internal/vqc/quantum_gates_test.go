package vqc

import (
	"testing"
)

func TestNewQuantumGate(t *testing.T) {
	gate_type := "cnot"
	var qubit Qubit = 1
	control_qubit := []Qubit{2}

	gate, err := NewQuantumGate(gate_type, qubit, control_qubit)
	if err != nil {
		t.Errorf("NewQuantumGate returned an error: %v", err)
	}

	if gate.gate_type != gate_type || gate.qubit != qubit || len(gate.control_qubit) != len(control_qubit) {
		t.Errorf("NewQuantumGate returned unexpected values: got %v, want %v", gate, &QuantumGates{gate_type, qubit, control_qubit})
	}
}

func TestInvalidGateType(t *testing.T) {
	gate_type := "invalid_gate"
	var qubit Qubit = 1
	control_qubit := []Qubit{}

	_, err := NewQuantumGate(gate_type, qubit, control_qubit)
	if err == nil {
		t.Errorf("NewQuantumGate did not return an error for invalid gate type")
	}
}

func TestDuplicateQubit(t *testing.T) {
	gate_type := "cnot"
	var qubit Qubit = 2
	control_qubit := []Qubit{2}

	_, err := NewQuantumGate(gate_type, qubit, control_qubit)
	if err == nil {
		t.Errorf("NewQuantumGate did not return an error for duplicate qubit index")
	}
}
