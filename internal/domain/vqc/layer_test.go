package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLayer(t *testing.T) {
	testCases := []struct {
		testName      string
		gates         []QuantumGate
		expectedCount uint
	}{
		{
			testName:      "Layer with no gates",
			gates:         []QuantumGate{},
			expectedCount: 0,
		},
		{
			testName:      "Layer with single gate",
			gates:         []QuantumGate{QuantumGate{}},
			expectedCount: 1,
		},
		{
			testName:      "Layer with multiple gates",
			gates:         []QuantumGate{QuantumGate{}, QuantumGate{}, QuantumGate{}},
			expectedCount: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			layer := NewLayer(tc.gates)
			gates := layer.Gates()
			assert.Equal(t, tc.expectedCount, uint(len(gates)), "Expected gate count does not match actual count")
		})
	}
}

func TestLayerGetNumParameterizedGates(t *testing.T) {
	testCases := []struct {
		testName               string
		gates                  []QuantumGate
		expectedParameterCount uint
	}{
		{
			testName:               "Empty layer",
			gates:                  []QuantumGate{},
			expectedParameterCount: 0,
		},
		{
			testName:               "Non-parameterized gates only",
			gates:                  []QuantumGate{{gate_type: HGate}, {gate_type: XGate}, {gate_type: CNOTGate}},
			expectedParameterCount: 0,
		},
		{
			testName:               "Parameterized gates only",
			gates:                  []QuantumGate{{gate_type: RXGate}, {gate_type: RYGate}, {gate_type: RZGate}},
			expectedParameterCount: 3,
		},
		{
			testName:               "Mixed parameterized and non-parameterized gates",
			gates:                  []QuantumGate{{gate_type: HGate}, {gate_type: RXGate}, {gate_type: XGate}, {gate_type: RYGate}},
			expectedParameterCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			layer := NewLayer(tc.gates)
			count := layer.NumParameterizedGates()
			assert.Equal(t, tc.expectedParameterCount, count, "Expected parameterized gate count does not match actual count")
		})
	}
}

func TestCloneLayer(t *testing.T) {
	originalGates := []QuantumGate{
		{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
		{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
	}

	layer := NewLayer(originalGates)
	clonedLayer := layer.Clone()

	// Modify the original layer's gates
	layerGates := layer.Gates()
	layerGates[0].gate_type = YGate
	layerGates = append(layerGates, QuantumGate{gate_type: RXGate, qubit: Qubit(2), control_qubit: []Qubit{}})

	// Check that the cloned layer's gates remain unchanged
	clonedGates := clonedLayer.Gates()
	assert.Equal(t, HGate, clonedGates[0].gate_type, "Cloned layer was affected by modification of original layer")
	assert.Equal(t, 2, len(clonedGates), "Cloned layer size was affected by modification of original layer")
}

func TestLayerImmutability(t *testing.T) {
	originalGates := []QuantumGate{
		{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
		{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
	}

	layer := NewLayer(originalGates)

	// Modify original gates slice
	originalGates[0].gate_type = YGate
	originalGates = append(originalGates, QuantumGate{gate_type: RXGate, qubit: Qubit(2), control_qubit: []Qubit{}})
	assert.Equal(t, HGate, layer.Gates()[0].gate_type, "Layer was affected by modification of original input slice")
	assert.Equal(t, 2, len(layer.Gates()), "Layer size was affected by modification of original input slice")

	layerGates := layer.Gates()
	// Modify the returned gates slice
	layerGates[1].gate_type = ZGate
	layerGates = append(layerGates, QuantumGate{gate_type: RYGate, qubit: Qubit(3), control_qubit: []Qubit{}})
	assert.Equal(t, XGate, layer.Gates()[1].gate_type, "Layer was affected by modification of returned slice")
	assert.Equal(t, 2, len(layer.Gates()), "Layer size was affected by modification of returned slice")

}

func TestLayerDeepImmutability(t *testing.T) {
	originalGates := []QuantumGate{
		{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{Qubit(1)}},
	}

	layer := NewLayer(originalGates)

	// Modify the control_qubit slice of the original gate
	originalGates[0].control_qubit[0] = Qubit(2)
	assert.Equal(t, Qubit(1), layer.Gates()[0].control_qubit[0], "Layer was affected by modification of original gate's control_qubit slice")

	layerGates := layer.Gates()
	// Modify the control_qubit slice of the returned gate
	layerGates[0].control_qubit[0] = Qubit(3)
	assert.Equal(t, Qubit(1), layer.Gates()[0].control_qubit[0], "Layer was affected by modification of returned gate's control_qubit slice")
}
