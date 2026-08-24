package vqc

import "slices"

type Layer struct {
	gates []QuantumGates
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
		gates: []QuantumGates{},
	}
}

func (l *Layer) AddGate(gate QuantumGates) error {
	l.gates = append(l.gates, gate)
	return nil
}

func (l *Layer) RemoveGateAtIndex(index LayerIndex) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	l.gates = slices.Delete(l.gates, int(index), int(index)+1)
	return nil
}

func (l *Layer) AddGateAtIndex(gate QuantumGates, index LayerIndex) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	l.gates = slices.Insert(l.gates, int(index), gate)
	return nil
}

func (l *Layer) UpdateGateAtIndex(gate QuantumGates, index LayerIndex) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	l.gates[index] = gate
	return nil
}
