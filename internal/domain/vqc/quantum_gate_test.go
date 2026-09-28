package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewQuantumGate(t *testing.T) {

	t.Run("valid gate creation", func(t *testing.T) {
		testCases := []struct {
			testName      string
			gateType      GateType
			qubit         Qubit
			controlQubits []Qubit
		}{
			{
				testName:      "valid single-qubit HGate",
				gateType:      HGate,
				qubit:         Qubit(0),
				controlQubits: nil,
			},
			{
				testName:      "valid single-qubit RXGate",
				gateType:      RXGate,
				qubit:         Qubit(1),
				controlQubits: nil,
			},
			{
				testName:      "valid single-qubit RYGate",
				gateType:      RYGate,
				qubit:         Qubit(2),
				controlQubits: nil,
			},
			{
				testName:      "valid two-qubit CNOTGate",
				gateType:      CNOTGate,
				qubit:         Qubit(1),
				controlQubits: []Qubit{Qubit(0)},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.testName, func(t *testing.T) {
				gate, err := NewQuantumGate(tc.gateType, tc.qubit, tc.controlQubits)
				assert.NoError(t, err)
				assert.Equal(t, tc.gateType, gate.GateType())
				assert.Equal(t, tc.qubit, gate.qubit)
				assert.Equal(t, tc.controlQubits, gate.controlQubit)
			})
		}
	})

	t.Run("invalid gate creation", func(t *testing.T) {
		testCases := []struct {
			testName      string
			gateType      GateType
			qubit         Qubit
			controlQubits []Qubit
			expectErr     error
		}{
			{
				testName:      "invalid gate type",
				gateType:      GateType("invalid"),
				qubit:         Qubit(0),
				controlQubits: nil,
				expectErr:     ErrInvalidGateType,
			},
			{
				testName:      "single-qubit gate with control qubits",
				gateType:      HGate,
				qubit:         Qubit(0),
				controlQubits: []Qubit{Qubit(1)},
				expectErr:     ErrInvalidControlQubit,
			},
			{
				testName:      "two-qubit gate with no control qubits",
				gateType:      CNOTGate,
				qubit:         Qubit(1),
				controlQubits: nil,
				expectErr:     ErrInvalidControlQubit,
			},
			{
				testName:      "two-qubit gate with multiple control qubits",
				gateType:      CNOTGate,
				qubit:         Qubit(1),
				controlQubits: []Qubit{Qubit(0), Qubit(2)},
				expectErr:     ErrInvalidControlQubit,
			},
			{
				testName:      "duplicate qubits",
				gateType:      CNOTGate,
				qubit:         Qubit(0),
				controlQubits: []Qubit{Qubit(0)},
				expectErr:     ErrDuplicatedQubit,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.testName, func(t *testing.T) {
				quantumGate, err := NewQuantumGate(tc.gateType, tc.qubit, tc.controlQubits)
				assert.Error(t, err)
				assert.False(t, quantumGate.isValid())
				assert.ErrorIs(t, err, tc.expectErr)
			})
		}
	})
}

func TestQuantumGateIsValid(t *testing.T) {
	t.Run("valid gate cases", func(t *testing.T) {
		validGates := []QuantumGate{
			{gateType: HGate, qubit: Qubit(0), controlQubit: nil},
			{gateType: XGate, qubit: Qubit(1), controlQubit: nil},
			{gateType: CNOTGate, qubit: Qubit(1), controlQubit: []Qubit{Qubit(0)}},
		}

		for _, gate := range validGates {
			assert.True(t, gate.isValid())
		}
	})

	t.Run("invalid gate case", func(t *testing.T) {
		invalidGate := QuantumGate{}
		assert.False(t, invalidGate.isValid())
	})
}

func TestQuantumGateEquals(t *testing.T) {

	t.Run("equal gates", func(t *testing.T) {
		t.Run("single-qubit gates", func(t *testing.T) {
			gate1, err := NewQuantumGate(HGate, Qubit(0), nil)
			require.NoError(t, err)
			gate2, err := NewQuantumGate(HGate, Qubit(0), nil)
			require.NoError(t, err)

			assert.True(t, gate1.Equals(gate2))
		})

		t.Run("two-qubit gates", func(t *testing.T) {
			gate3, err := NewQuantumGate(CNOTGate, Qubit(1), []Qubit{Qubit(0)})
			require.NoError(t, err)
			gate4, err := NewQuantumGate(CNOTGate, Qubit(1), []Qubit{Qubit(0)})
			require.NoError(t, err)

			assert.True(t, gate3.Equals(gate4))
		})
	})

	t.Run("different gates", func(t *testing.T) {
		t.Run("single-qubit gates", func(t *testing.T) {
			gate1, err := NewQuantumGate(HGate, Qubit(0), nil)
			require.NoError(t, err)
			gate2, err := NewQuantumGate(XGate, Qubit(0), nil)
			require.NoError(t, err)

			assert.False(t, gate1.Equals(gate2))
		})

		t.Run("two-qubit gates", func(t *testing.T) {
			gate3, err := NewQuantumGate(CNOTGate, Qubit(1), []Qubit{Qubit(0)})
			require.NoError(t, err)
			gate4, err := NewQuantumGate(CNOTGate, Qubit(1), []Qubit{Qubit(2)})
			require.NoError(t, err)

			assert.False(t, gate3.Equals(gate4))
		})
	})
}

func TestHasParameters(t *testing.T) {

	t.Run("has parameters", func(t *testing.T) {
		inputGates := []struct {
			testName string
			gate     QuantumGate
		}{
			{
				testName: "RXGate",
				gate:     QuantumGate{gateType: RXGate, qubit: Qubit(0), controlQubit: nil},
			},
			{
				testName: "RYGate",
				gate:     QuantumGate{gateType: RYGate, qubit: Qubit(1), controlQubit: nil},
			},
			{
				testName: "RZGate",
				gate:     QuantumGate{gateType: RZGate, qubit: Qubit(2), controlQubit: nil},
			},
		}

		for _, testCase := range inputGates {
			t.Run(testCase.testName, func(t *testing.T) {
				result := testCase.gate.HasParameters()
				assert.True(t, result)
			})
		}
	})

	t.Run("does not have parameters", func(t *testing.T) {
		inputGates := []struct {
			testName string
			gate     QuantumGate
		}{
			{
				testName: "HGate",
				gate:     QuantumGate{gateType: HGate, qubit: Qubit(0), controlQubit: nil},
			},
			{
				testName: "XGate",
				gate:     QuantumGate{gateType: XGate, qubit: Qubit(1), controlQubit: nil},
			},
			{
				testName: "CNOTGate",
				gate:     QuantumGate{gateType: CNOTGate, qubit: Qubit(1), controlQubit: []Qubit{Qubit(0)}},
			},
		}

		for _, testCase := range inputGates {
			t.Run(testCase.testName, func(t *testing.T) {
				result := testCase.gate.HasParameters()
				assert.False(t, result)
			})
		}
	})
}

func TestCloneImmutability(t *testing.T) {
	t.Run("clone gate and modify original", func(t *testing.T) {
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
	})
}

func TestQuantumGateControlQubitsImmutability(t *testing.T) {
	t.Run("modify control qubits slice after creating gate", func(t *testing.T) {
		controlQubits := []Qubit{Qubit(0)}
		gate, err := NewQuantumGate(CNOTGate, Qubit(1), controlQubits)
		require.NoError(t, err)

		recivedControlQubits := gate.ControlQubits()

		recivedControlQubits[0] = Qubit(2)

		// The original control qubits in the gate should remain unchanged
		assert.NotEqual(t, recivedControlQubits, gate.ControlQubits())
	})
}

func TestQuantumGateValidateQubits(t *testing.T) {
	t.Run("valid qubit cases", func(t *testing.T) {
		gate, err := NewQuantumGate(CNOTGate, Qubit(1), []Qubit{Qubit(0)})
		require.NoError(t, err)
		err = gate.validateQubits(3)
		assert.NoError(t, err)
	})

	t.Run("invalid qubit cases", func(t *testing.T) {
		gate, err := NewQuantumGate(CNOTGate, Qubit(1), []Qubit{Qubit(0)})
		require.NoError(t, err)
		err = gate.validateQubits(1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidQubit)
	})
}
