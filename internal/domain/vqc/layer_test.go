package vqc

import (
	"testing"
)

type ModifyGate struct {
	index LayerIndex
	gate  QuantumGate
}

func TestLayer(t *testing.T) {
	testCases := []struct {
		name              string
		gatesToCreate     []QuantumGate
		gatesToRemove     []LayerIndex
		gatesToUpdate     ModifyGate
		gatesToAddAtIndex ModifyGate
		expectedGates     []QuantumGate
		expectErr         bool
	}{
		{"Adding gates to layer", []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, nil, ModifyGate{}, ModifyGate{}, []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, false,
		},
		{"Removing gates from layer", []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, []LayerIndex{1}, ModifyGate{}, ModifyGate{}, []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, false,
		},
		{"Updating gates in layer", []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, nil, ModifyGate{index: 1, gate: QuantumGate{gate_type: CNOTGate, qubit: Qubit(2), control_qubit: []Qubit{Qubit(3)}}}, ModifyGate{},
			[]QuantumGate{
				{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
				{gate_type: CNOTGate, qubit: Qubit(2), control_qubit: []Qubit{Qubit(3)}},
				{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
			}, false,
		},
		{"Adding gate at specific index in layer", []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, nil, ModifyGate{}, ModifyGate{index: 1, gate: QuantumGate{gate_type: CNOTGate, qubit: Qubit(2), control_qubit: []Qubit{Qubit(3)}}},
			[]QuantumGate{
				{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
				{gate_type: CNOTGate, qubit: Qubit(2), control_qubit: []Qubit{Qubit(3)}},
				{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
				{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
			}, false,
		},
		{"Invalid gate index for removal", []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, []LayerIndex{5}, ModifyGate{}, ModifyGate{}, []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, true,
		},
		{"Invalid gate index for update", []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, nil, ModifyGate{index: 5, gate: QuantumGate{gate_type: CNOTGate, qubit: Qubit(2), control_qubit: []Qubit{Qubit(3)}}}, ModifyGate{}, []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, true,
		},
		{"Invalid gate index for addition", []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, nil, ModifyGate{}, ModifyGate{index: 5, gate: QuantumGate{gate_type: CNOTGate, qubit: Qubit(2), control_qubit: []Qubit{Qubit(3)}}}, []QuantumGate{
			{gate_type: HGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(0), control_qubit: []Qubit{}},
			{gate_type: XGate, qubit: Qubit(1), control_qubit: []Qubit{}},
		}, true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			layer := NewLayer()
			for _, gate := range tc.gatesToCreate {
				err := layer.AddGate(gate)
				if err != nil {
					t.Fatalf("Failed to add gate: %v", err)
				}
			}

			for _, index := range tc.gatesToRemove {
				err := layer.RemoveGateAtIndex(index)
				if err != nil && !tc.expectErr {
					t.Fatalf("Failed to remove gate at index %d: %v", index, err)
				}
				if err == nil && tc.expectErr {
					t.Fatalf("Expected removing gate at index %d to fail", index)
				}
			}

			if !tc.gatesToUpdate.gate.Equal(QuantumGate{}) {
				err := layer.UpdateGateAtIndex(tc.gatesToUpdate.gate, tc.gatesToUpdate.index)
				if err != nil && !tc.expectErr {
					t.Fatalf("Failed to update gate at index %d: %v", tc.gatesToUpdate.index, err)
				}
				if err == nil && tc.expectErr {
					t.Fatalf("Expected updating gate at index %d to fail", tc.gatesToUpdate.index)
				}
			}

			if !tc.gatesToAddAtIndex.gate.Equal(QuantumGate{}) {
				err := layer.AddGateAtIndex(tc.gatesToAddAtIndex.gate, tc.gatesToAddAtIndex.index)
				if err != nil && !tc.expectErr {
					t.Fatalf("Failed to add gate at index %d: %v", tc.gatesToAddAtIndex.index, err)
				}
				if err == nil && tc.expectErr {
					t.Fatalf("Expected adding gate at index %d to fail", tc.gatesToAddAtIndex.index)
				}
			}

			if !slicesQuantumGateEqual(layer.gates, tc.expectedGates) {
				t.Errorf("Expected gates: %v, but got: %v", tc.expectedGates, layer.gates)
			}
		})
	}
}

func slicesQuantumGateEqual(a, b []QuantumGate) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equal(b[i]) {
			return false
		}
	}
	return true
}
