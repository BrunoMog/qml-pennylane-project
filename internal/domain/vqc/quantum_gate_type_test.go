package vqc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGateType(t *testing.T) {

	t.Run("valid gate type cases", func(t *testing.T) {
		inputs := []struct {
			testName     string
			input        string
			expectedGate GateType
		}{
			{"valid H gate", "h", HGate},
			{"valid X gate", "x", XGate},
			{"valid Y gate", "y", YGate},
			{"valid Z gate", "z", ZGate},
			{"valid RX gate", "rx", RXGate},
			{"valid RY gate", "ry", RYGate},
			{"valid RZ gate", "rz", RZGate},
			{"valid CNOT gate", "cnot", CNOTGate},
			{"valid H gate with whitespace", " h ", HGate},
			{"valid X gate with uppercase", "X", XGate},
			{"valid RY gate with mixed case", "rY", RYGate},
		}

		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseGateType(tc.input)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedGate, result)
			})
		}
	})

	t.Run("invalid gate type cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
		}{
			{"invalid gate type", "invalid_gate"},
			{"empty string", ""},
			{"whitespace only", "   "},
			{"numeric string", "123"},
			{"special characters", "@#$%"},
			{"mixed valid and invalid characters", "h!x"},
			{"too long string", "this_is_a_very_long_gate_type_string"},
		}

		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseGateType(tc.input)
				assert.Error(t, err)
				assert.Equal(t, GateType(""), result)
			})
		}
	})

	t.Run("extreme gate type cases", func(t *testing.T) {
		inputs := []struct {
			testName string
			input    string
		}{
			{"100,000 characters", strings.Repeat("a", 100000)},
			{"4,000,000 characters", strings.Repeat("a", 4000000)},
			{"10,000,000 characters", strings.Repeat("a", 10000000)},
		}

		for _, tc := range inputs {
			t.Run(tc.testName, func(t *testing.T) {
				result, err := ParseGateType(tc.input)
				assert.Error(t, err)
				assert.Equal(t, GateType(""), result)
			})
		}
	})
}

func TestIsValidGateType(t *testing.T) {
	t.Run("valid gate type cases", func(t *testing.T) {
		validGateTypes := []GateType{
			HGate,
			XGate,
			YGate,
			ZGate,
			RXGate,
			RYGate,
			RZGate,
			CNOTGate,
		}

		for _, gateType := range validGateTypes {
			assert.True(t, gateType.isValid())
		}
	})

	t.Run("invalid gate type cases", func(t *testing.T) {
		invalidGateTypes := []GateType{
			GateType("invalid"),
			GateType(""),
			GateType("   "),
			GateType("123"),
			GateType("@#$%"),
			GateType("this_is_a_very_long_gate_type_string"),
		}

		for _, gateType := range invalidGateTypes {
			assert.False(t, gateType.isValid())
		}
	})
}

func TestIsSingleQubitGate(t *testing.T) {
	t.Run("single-qubit gate cases", func(t *testing.T) {
		singleQubitGates := []GateType{
			HGate,
			XGate,
			YGate,
			ZGate,
			RXGate,
			RYGate,
			RZGate,
		}

		for _, gateType := range singleQubitGates {
			assert.True(t, gateType.isSingleQubitGate())
		}
	})

	t.Run("non-single-qubit gate cases", func(t *testing.T) {
		nonSingleQubitGates := []GateType{
			CNOTGate,
			GateType("invalid"),
			GateType(""),
		}

		for _, gateType := range nonSingleQubitGates {
			assert.False(t, gateType.isSingleQubitGate())
		}
	})
}

func TestIsTwoQubitGate(t *testing.T) {
	t.Run("two-qubit gate cases", func(t *testing.T) {
		twoQubitGates := []GateType{
			CNOTGate,
		}

		for _, gateType := range twoQubitGates {
			assert.True(t, gateType.isTwoQubitGate())
		}
	})

	t.Run("non-two-qubit gate cases", func(t *testing.T) {
		nonTwoQubitGates := []GateType{
			HGate,
			XGate,
			YGate,
			ZGate,
			RXGate,
			RYGate,
			RZGate,
			GateType("invalid"),
			GateType(""),
		}

		for _, gateType := range nonTwoQubitGates {
			assert.False(t, gateType.isTwoQubitGate())
		}
	})
}
