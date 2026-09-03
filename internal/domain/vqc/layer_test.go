package vqc

import (
	"testing"
)

func TestNewLayer(t *testing.T) {
	testCases := []struct {
		name          string
		gates         []QuantumGate
		expectedCount uint
	}{
		{
			name:          "Layer with no gates",
			gates:         []QuantumGate{},
			expectedCount: 0,
		},
		{
			name: "Layer with single gate",
			gates: []QuantumGate{
				{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			},
			expectedCount: 1,
		},
		{
			name: "Layer with multiple gates",
			gates: []QuantumGate{
				{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
				{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
				{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
			},
			expectedCount: 3,
		},
		{
			name: "Layer with controlled gates",
			gates: []QuantumGate{
				{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
				{gate_type: CNOTGate, qubit: Qubit(1), control_qubit: []Qubit{Qubit(0)}},
			},
			expectedCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			layer := NewLayer(tc.gates)
			if layer == nil {
				t.Error("NewLayer returned nil")
			}

			gates := layer.GetGates()
			if uint(len(gates)) != tc.expectedCount {
				t.Errorf("Expected %d gates, got %d", tc.expectedCount, len(gates))
			}
		})
	}
}

func TestLayerGetGates(t *testing.T) {
	originalGates := []QuantumGate{
		{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
		{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
	}

	layer := NewLayer(originalGates)
	retrievedGates := layer.GetGates()

	testCases := []struct {
		name     string
		checkFn  func() bool
		errorMsg string
	}{
		{
			name:     "Returns all gates",
			checkFn:  func() bool { return uint(len(retrievedGates)) == uint(len(originalGates)) },
			errorMsg: "Retrieved gates count doesn't match original",
		},
		{
			name: "Returns gates in correct order",
			checkFn: func() bool {
				for i := range retrievedGates {
					if !retrievedGates[i].Equal(originalGates[i]) {
						return false
					}
				}
				return true
			},
			errorMsg: "Gates order doesn't match",
		},
		{
			name: "Modification of returned slice doesn't affect original",
			checkFn: func() bool {
				if len(retrievedGates) > 0 {
					// Modify the returned slice
					retrievedGates[0].gate_type = YGate
					// Get gates again and check they're unchanged
					newRetrieval := layer.GetGates()
					return newRetrieval[0].gate_type == HGate
				}
				return true
			},
			errorMsg: "Layer was affected by external slice modification",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.checkFn() {
				t.Error(tc.errorMsg)
			}
		})
	}
}

func TestLayerGetNumParameterizedGates(t *testing.T) {
	testCases := []struct {
		name                   string
		gates                  []QuantumGate
		expectedParameterCount uint
	}{
		{
			name:                   "Empty layer",
			gates:                  []QuantumGate{},
			expectedParameterCount: 0,
		},
		{
			name: "Non-parameterized gates only",
			gates: []QuantumGate{
				{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
				{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
				{gate_type: YGate, qubit: Qubit(2), control_qubit: []Qubit{}},
			},
			expectedParameterCount: 0,
		},
		{
			name: "Parameterized gates only",
			gates: []QuantumGate{
				{gate_type: RXGate, qubit: Qubit(0), control_qubit: []Qubit{}},
				{gate_type: RYGate, qubit: Qubit(1), control_qubit: []Qubit{}},
				{gate_type: RZGate, qubit: Qubit(2), control_qubit: []Qubit{}},
			},
			expectedParameterCount: 3,
		},
		{
			name: "Mixed parameterized and non-parameterized gates",
			gates: []QuantumGate{
				{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
				{gate_type: RXGate, qubit: Qubit(1), control_qubit: []Qubit{}},
				{gate_type: XGate, qubit: Qubit(2), control_qubit: []Qubit{}},
				{gate_type: RYGate, qubit: Qubit(3), control_qubit: []Qubit{}},
			},
			expectedParameterCount: 2,
		},
		{
			name: "CNOT gate (non-parameterized controlled gate)",
			gates: []QuantumGate{
				{gate_type: CNOTGate, qubit: Qubit(0), control_qubit: []Qubit{Qubit(1)}},
				{gate_type: RXGate, qubit: Qubit(2), control_qubit: []Qubit{}},
			},
			expectedParameterCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			layer := NewLayer(tc.gates)
			count := layer.GetNumParameterizedGates()
			if count != tc.expectedParameterCount {
				t.Errorf("Expected %d parameterized gates, got %d", tc.expectedParameterCount, count)
			}
		})
	}
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

	// Get gates from layer and verify they're unchanged
	layerGates := layer.GetGates()
	if layerGates[0].gate_type != HGate {
		t.Error("Layer was affected by modification of original input slice")
	}
	if uint(len(layerGates)) != 2 {
		t.Error("Layer size was affected by modification of original input slice")
	}
}
