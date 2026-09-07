package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGateType(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedGate  GateType
		expectedError error
	}{
		{"Valid H gate", "h", HGate, nil},
		{"Valid X gate", "x", XGate, nil},
		{"Valid Y gate", "y", YGate, nil},
		{"Valid Z gate", "z", ZGate, nil},
		{"Valid RX gate", "rx", RXGate, nil},
		{"Valid RY gate", "ry", RYGate, nil},
		{"Valid RZ gate", "rz", RZGate, nil},
		{"Valid CNOT gate", "cnot", CNOTGate, nil},
		{"Invalid gate type", "invalid_gate", "", &InvalidParseGateTypeError{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gateType, err := ParseGateType(tt.input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedGate, gateType)
			}
		})
	}
}
