package vqc

import "slices"

type Layer struct {
	gates []QuantumGate
}

func NewLayer(gates []QuantumGate) Layer {
	return Layer{
		gates: slices.Clone(gates),
	}
}

func (l Layer) Gates() []QuantumGate {
	return slices.Clone(l.gates)
}

func (l Layer) NumParameterizedGates() uint {
	count := uint(0)
	for _, gate := range l.gates {
		if gate.HasParameters() {
			count++
		}
	}
	return count
}
