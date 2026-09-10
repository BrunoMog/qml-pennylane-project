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
			gates:                  []QuantumGate{{gateType: HGate}, {gateType: XGate}, {gateType: CNOTGate}},
			expectedParameterCount: 0,
		},
		{
			testName:               "Parameterized gates only",
			gates:                  []QuantumGate{{gateType: RXGate}, {gateType: RYGate}, {gateType: RZGate}},
			expectedParameterCount: 3,
		},
		{
			testName:               "Mixed parameterized and non-parameterized gates",
			gates:                  []QuantumGate{{gateType: HGate}, {gateType: RXGate}, {gateType: XGate}, {gateType: RYGate}},
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
		{gateType: HGate, qubit: Qubit(0), controlQubit: []Qubit{}},
		{gateType: XGate, qubit: Qubit(1), controlQubit: []Qubit{}},
	}

	layer := NewLayer(originalGates)
	clonedLayer := layer.Clone()

	// Modify the original layer's gates
	layerGates := layer.Gates()
	layerGates[0].gateType = YGate
	layerGates = append(layerGates, QuantumGate{gateType: RXGate, qubit: Qubit(2), controlQubit: []Qubit{}})

	// Check that the cloned layer's gates remain unchanged
	clonedGates := clonedLayer.Gates()
	assert.Equal(t, HGate, clonedGates[0].gateType, "Cloned layer was affected by modification of original layer")
	assert.Equal(t, 2, len(clonedGates), "Cloned layer size was affected by modification of original layer")
}

func TestLayerImmutability(t *testing.T) {
	originalGates := []QuantumGate{
		{gateType: HGate, qubit: Qubit(0), controlQubit: []Qubit{}},
		{gateType: XGate, qubit: Qubit(1), controlQubit: []Qubit{}},
	}

	layer := NewLayer(originalGates)

	// Modify original gates slice
	originalGates[0].gateType = YGate
	originalGates = append(originalGates, QuantumGate{gateType: RXGate, qubit: Qubit(2), controlQubit: []Qubit{}})
	assert.Equal(t, HGate, layer.Gates()[0].gateType, "Layer was affected by modification of original input slice")
	assert.Equal(t, 2, len(layer.Gates()), "Layer size was affected by modification of original input slice")

	layerGates := layer.Gates()
	// Modify the returned gates slice
	layerGates[1].gateType = ZGate
	layerGates = append(layerGates, QuantumGate{gateType: RYGate, qubit: Qubit(3), controlQubit: []Qubit{}})
	assert.Equal(t, XGate, layer.Gates()[1].gateType, "Layer was affected by modification of returned slice")
	assert.Equal(t, 2, len(layer.Gates()), "Layer size was affected by modification of returned slice")

}

func TestLayerDeepImmutability(t *testing.T) {
	originalGates := []QuantumGate{
		{gateType: HGate, qubit: Qubit(0), controlQubit: []Qubit{Qubit(1)}},
	}

	layer := NewLayer(originalGates)

	// Modify the controlQubit slice of the original gate
	originalGates[0].controlQubit[0] = Qubit(2)
	assert.Equal(t, Qubit(1), layer.Gates()[0].controlQubit[0], "Layer was affected by modification of original gate's controlQubit slice")

	layerGates := layer.Gates()
	// Modify the controlQubit slice of the returned gate
	layerGates[0].controlQubit[0] = Qubit(3)
	assert.Equal(t, Qubit(1), layer.Gates()[0].controlQubit[0], "Layer was affected by modification of returned gate's controlQubit slice")
}
