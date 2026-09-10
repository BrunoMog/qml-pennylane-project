package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuantumGate(t *testing.T) {
	tests := []struct {
		testName      string
		setup         func(t *testing.T) (GateType, Qubit, []Qubit)
		expectedError error
	}{
		{
			testName: "valid single-qubit gate",
			setup: func(t *testing.T) (GateType, Qubit, []Qubit) {
				return HGate, Qubit(0), nil
			},
			expectedError: nil,
		},
		{
			testName: "valid two-qubit gate",
			setup: func(t *testing.T) (GateType, Qubit, []Qubit) {
				return CNOTGate, Qubit(1), []Qubit{Qubit(0)}
			},
			expectedError: nil,
		},
		{
			testName: "invalid gate type",
			setup: func(t *testing.T) (GateType, Qubit, []Qubit) {
				return GateType("invalid_gate"), Qubit(0), nil
			},
			expectedError: &InvalidGateError{},
		},
		{
			testName: "invalid single-qubit gate with control qubits",
			setup: func(t *testing.T) (GateType, Qubit, []Qubit) {
				return HGate, Qubit(0), []Qubit{Qubit(1)}
			},
			expectedError: &InvalidControlQubitError{},
		},
		{
			testName: "invalid two-qubit gate with no control qubits",
			setup: func(t *testing.T) (GateType, Qubit, []Qubit) {
				return CNOTGate, Qubit(1), nil
			},
			expectedError: &InvalidControlQubitError{},
		},
		{
			testName: "invalid two-qubit gate with multiple control qubits",
			setup: func(t *testing.T) (GateType, Qubit, []Qubit) {
				return CNOTGate, Qubit(1), []Qubit{Qubit(0), Qubit(2)}
			},
			expectedError: &InvalidControlQubitError{},
		},
		{
			testName: "duplicate qubits",
			setup: func(t *testing.T) (GateType, Qubit, []Qubit) {
				return CNOTGate, Qubit(0), []Qubit{Qubit(0)}
			},
			expectedError: &DuplicateQubitError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			gateType, qubit, controlQubit := tt.setup(t)
			gate, err := NewQuantumGate(gateType, qubit, controlQubit)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, gate)
				assert.Equal(t, gateType, gate.GateType())
				assert.Equal(t, qubit, gate.qubit)
				assert.Equal(t, controlQubit, gate.controlQubit)
			}
		})
	}
}

func TestQuantumGateEqual(t *testing.T) {
	tests := []struct {
		testName string
		gate1    QuantumGate
		gate2    QuantumGate
		expected bool
	}{
		{
			testName: "equal gates",
			gate1:    QuantumGate{gateType: HGate, qubit: Qubit(0), controlQubit: nil},
			gate2:    QuantumGate{gateType: HGate, qubit: Qubit(0), controlQubit: nil},
			expected: true,
		},
		{
			testName: "different gate types",
			gate1:    QuantumGate{gateType: HGate, qubit: Qubit(0), controlQubit: nil},
			gate2:    QuantumGate{gateType: XGate, qubit: Qubit(0), controlQubit: nil},
			expected: false,
		},
		{
			testName: "different qubits",
			gate1:    QuantumGate{gateType: HGate, qubit: Qubit(0), controlQubit: nil},
			gate2:    QuantumGate{gateType: HGate, qubit: Qubit(1), controlQubit: nil},
			expected: false,
		},
		{
			testName: "different control qubits",
			gate1:    QuantumGate{gateType: CNOTGate, qubit: Qubit(1), controlQubit: []Qubit{Qubit(0)}},
			gate2:    QuantumGate{gateType: CNOTGate, qubit: Qubit(1), controlQubit: []Qubit{Qubit(2)}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result := tt.gate1.Equals(tt.gate2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHasParameters(t *testing.T) {
	tests := []struct {
		testName string
		gate     QuantumGate
		expected bool
	}{
		{
			testName: "gate with no parameters",
			gate:     QuantumGate{gateType: HGate, qubit: Qubit(0), controlQubit: nil},
			expected: false,
		},
		{
			testName: "gate with parameters",
			gate:     QuantumGate{gateType: RXGate, qubit: Qubit(0), controlQubit: nil},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result := tt.gate.HasParameters()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCloneImmutability(t *testing.T) {
	originalGate := QuantumGate{
		gateType:     CNOTGate,
		qubit:        Qubit(1),
		controlQubit: []Qubit{Qubit(0)},
	}

	clonedGate := originalGate.Clone()

	// Modify the cloned gate's control qubits
	clonedGate.controlQubit[0] = Qubit(2)

	// The original gate's control qubits should remain unchanged
	assert.Equal(t, []Qubit{Qubit(0)}, originalGate.controlQubit)
	assert.Equal(t, []Qubit{Qubit(2)}, clonedGate.controlQubit)
}
