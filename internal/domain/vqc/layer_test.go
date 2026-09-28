package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLayer(t *testing.T) {

	t.Run("create layer with valid gates", func(t *testing.T) {
		testCases := []struct {
			testName      string
			gates         []QuantumGate
			expectedCount uint
		}{
			{
				testName:      "layer with no gates",
				gates:         []QuantumGate{},
				expectedCount: 0,
			},
			{
				testName:      "layer with single gate",
				gates:         []QuantumGate{QuantumGate{gateType: HGate, qubit: Qubit(0), controlQubit: []Qubit{}}},
				expectedCount: 1,
			},
			{
				testName:      "layer with three gates",
				gates:         []QuantumGate{QuantumGate{gateType: HGate, qubit: Qubit(0), controlQubit: []Qubit{}}, QuantumGate{gateType: XGate, qubit: Qubit(1), controlQubit: []Qubit{}}, QuantumGate{gateType: CNOTGate, qubit: Qubit(2), controlQubit: []Qubit{}}},
				expectedCount: 3,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.testName, func(t *testing.T) {
				layer, err := NewLayer(tc.gates)
				require.NoError(t, err)
				gates := layer.Gates()
				assert.Equal(t, tc.expectedCount, uint(len(gates)), "Expected gate count does not match actual count")
			})
		}
	})

	t.Run("create layer with invalid gates", func(t *testing.T) {
		layer, err := NewLayer([]QuantumGate{{}})
		assert.Error(t, err)
		assert.Equal(t, Layer{}, layer)
		assert.ErrorIs(t, err, ErrInvalidQuantumGate)
	})
}

func TestLayerGetNumParameterizedGates(t *testing.T) {
	t.Run("count parameterized gates", func(t *testing.T) {
		testCases := []struct {
			testName               string
			gates                  []QuantumGate
			expectedParameterCount uint
		}{
			{
				testName:               "empty layer",
				gates:                  []QuantumGate{},
				expectedParameterCount: 0,
			},
			{
				testName:               "non-parameterized gates only",
				gates:                  []QuantumGate{{gateType: HGate}, {gateType: XGate}, {gateType: CNOTGate}},
				expectedParameterCount: 0,
			},
			{
				testName:               "parameterized gates only",
				gates:                  []QuantumGate{{gateType: RXGate}, {gateType: RYGate}, {gateType: RZGate}},
				expectedParameterCount: 3,
			},
			{
				testName:               "mixed parameterized and non-parameterized gates",
				gates:                  []QuantumGate{{gateType: HGate}, {gateType: RXGate}, {gateType: XGate}, {gateType: RYGate}},
				expectedParameterCount: 2,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.testName, func(t *testing.T) {
				layer, err := NewLayer(tc.gates)
				require.NoError(t, err)
				count := layer.NumParameterizedGates()
				assert.Equal(t, tc.expectedParameterCount, count)
			})
		}
	})
}

func TestCloneLayer(t *testing.T) {
	t.Run("clone layer and modify original", func(t *testing.T) {
		originalGates := []QuantumGate{
			{gateType: HGate, qubit: Qubit(0), controlQubit: []Qubit{}},
			{gateType: XGate, qubit: Qubit(1), controlQubit: []Qubit{}},
		}

		layer, err := NewLayer(originalGates)
		require.NoError(t, err)
		clonedLayer := layer.Clone()

		// Modify the original layer's gates
		layerGates := layer.Gates()
		layerGates[0].gateType = YGate
		layerGates = append(layerGates, QuantumGate{gateType: RXGate, qubit: Qubit(2), controlQubit: []Qubit{}})

		// Check that the cloned layer's gates remain unchanged
		clonedGates := clonedLayer.Gates()
		assert.Equal(t, HGate, clonedGates[0].gateType)
		assert.Equal(t, 2, len(clonedGates))
	})
}

func TestLayerImmutability(t *testing.T) {
	t.Run("layer immutability", func(t *testing.T) {
		originalGates := []QuantumGate{
			{gateType: HGate, qubit: Qubit(0), controlQubit: []Qubit{}},
			{gateType: XGate, qubit: Qubit(1), controlQubit: []Qubit{}},
		}

		layer, err := NewLayer(originalGates)
		require.NoError(t, err)

		// Modify original gates slice
		originalGates[0].gateType = YGate
		originalGates = append(originalGates, QuantumGate{gateType: RXGate, qubit: Qubit(2), controlQubit: []Qubit{}})
		assert.Equal(t, HGate, layer.Gates()[0].gateType)
		assert.Equal(t, 2, len(layer.Gates()))

		layerGates := layer.Gates()
		// Modify the returned gates slice
		layerGates[1].gateType = ZGate
		layerGates = append(layerGates, QuantumGate{gateType: RYGate, qubit: Qubit(3), controlQubit: []Qubit{}})
		assert.Equal(t, XGate, layer.Gates()[1].gateType)
		assert.Equal(t, 2, len(layer.Gates()))
	})
}

func TestLayerDeepImmutability(t *testing.T) {
	t.Run("layer deep immutability", func(t *testing.T) {
		originalGates := []QuantumGate{
			{gateType: HGate, qubit: Qubit(0), controlQubit: []Qubit{Qubit(1)}},
		}

		layer, err := NewLayer(originalGates)
		require.NoError(t, err)

		// Modify the controlQubit slice of the original gate
		originalGates[0].controlQubit[0] = Qubit(2)
		assert.Equal(t, Qubit(1), layer.Gates()[0].controlQubit[0])

		layerGates := layer.Gates()
		// Modify the controlQubit slice of the returned gate
		layerGates[0].controlQubit[0] = Qubit(3)
		assert.Equal(t, Qubit(1), layer.Gates()[0].controlQubit[0])
	})
}
