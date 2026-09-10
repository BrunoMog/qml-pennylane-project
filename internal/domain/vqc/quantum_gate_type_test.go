package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGateType(t *testing.T) {
	tests := []struct {
		expectedError error
		testName      string
		input         string
		expectedGate  GateType
	}{
		{testName: "Valid H gate", input: "h", expectedGate: HGate, expectedError: nil},
		{testName: "Valid X gate", input: "x", expectedGate: XGate, expectedError: nil},
		{testName: "Valid Y gate", input: "y", expectedGate: YGate, expectedError: nil},
		{testName: "Valid Z gate", input: "z", expectedGate: ZGate, expectedError: nil},
		{testName: "Valid RX gate", input: "rx", expectedGate: RXGate, expectedError: nil},
		{testName: "Valid RY gate", input: "ry", expectedGate: RYGate, expectedError: nil},
		{testName: "Valid RZ gate", input: "rz", expectedGate: RZGate, expectedError: nil},
		{testName: "Valid CNOT gate", input: "cnot", expectedGate: CNOTGate, expectedError: nil},
		{testName: "Invalid gate type", input: "invalid_gate", expectedGate: "", expectedError: &InvalidParseGateTypeError{}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
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
