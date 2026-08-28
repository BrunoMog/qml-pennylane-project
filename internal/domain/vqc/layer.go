package vqc

import "slices"

type Layer struct {
	gates []QuantumGate
}

type LayerIndex uint

func (l *Layer) validateIndex(index LayerIndex) error {
	if index >= LayerIndex(len(l.gates)) {
		return &InvalidIndexError{index}
	}
	return nil
}

func NewLayer() *Layer {
	return &Layer{
		gates: []QuantumGate{},
	}
}

func (l *Layer) AddGate(gate QuantumGate) {
	l.gates = append(l.gates, gate)
}

func (l *Layer) RemoveGateAtIndex(index LayerIndex) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	l.gates = slices.Delete(l.gates, int(index), int(index)+1)
	return nil
}

func (l *Layer) AddGateAtIndex(gate QuantumGate, index LayerIndex) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	l.gates = slices.Insert(l.gates, int(index), gate)
	return nil
}

func (l *Layer) UpdateGateAtIndex(gate QuantumGate, index LayerIndex) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	l.gates[index] = gate
	return nil
}

func (l Layer) GetGates() []QuantumGate {
	return slices.Clone(l.gates)
}

func (l Layer) GetGateAtIndex(index LayerIndex) (QuantumGate, error) {
	if err := l.validateIndex(index); err != nil {
		return QuantumGate{}, err
	}

	return l.gates[index], nil
}

func (l Layer) GetNumParameterizedGates() uint {
	count := uint(0)
	for _, gate := range l.gates {
		if gate.HasParameters() {
			count++
		}
	}
	return count
}
