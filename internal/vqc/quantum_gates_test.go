package vqc

import (
	"testing"
)

func TestNewQuantumGate(t *testing.T) {
	gate_type := CNOTGate
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
	const IGate GateType = "invalid_gate"
	gate_type := IGate
	var qubit Qubit = 1
	control_qubit := []Qubit{}

	_, err := NewQuantumGate(gate_type, qubit, control_qubit)
	if err == nil {
		t.Errorf("NewQuantumGate did not return an error for invalid gate type")
	}
}

func TestDuplicateQubit(t *testing.T) {
	gate_type := CNOTGate
	var qubit Qubit = 2
	control_qubit := []Qubit{2}

	_, err := NewQuantumGate(gate_type, qubit, control_qubit)
	if err == nil {
		t.Errorf("NewQuantumGate did not return an error for duplicate qubit index")
	}
}

func TestInvalidControlQubit(t *testing.T) {
	gate_type := CNOTGate
	var qubit Qubit = 1
	control_qubit := []Qubit{}

	_, err := NewQuantumGate(gate_type, qubit, control_qubit)
	if err == nil {
		t.Errorf("NewQuantumGate did not return an error for invalid control qubit")
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		name     string
		gate1    QuantumGates
		gate2    QuantumGates
		expected bool
	}{
		{"Equal gates", QuantumGates{HGate, 0, []Qubit{}}, QuantumGates{HGate, 0, []Qubit{}}, true},
		{"Different gate types", QuantumGates{HGate, 0, []Qubit{}}, QuantumGates{XGate, 0, []Qubit{}}, false},
		{"Different qubits", QuantumGates{HGate, 0, []Qubit{}}, QuantumGates{HGate, 1, []Qubit{}}, false},
		{"Different control qubits", QuantumGates{CNOTGate, 0, []Qubit{1}}, QuantumGates{CNOTGate, 0, []Qubit{2}}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.gate1.Equal(tc.gate2)
			if result != tc.expected {
				t.Errorf("Equal() returned %v, expected %v", result, tc.expected)
			}
		})
	}
}
