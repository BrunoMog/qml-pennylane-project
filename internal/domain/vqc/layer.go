package vqc

type Layer struct {
	gates []QuantumGate
}

func NewLayer(gates []QuantumGate) Layer {
	return Layer{
		gates: cloneGates(gates),
	}
}

func cloneGates(gates []QuantumGate) []QuantumGate {
	clonedGates := make([]QuantumGate, len(gates))
	for i, gate := range gates {
		clonedGates[i] = gate.Clone()
	}
	return clonedGates
}

func (l Layer) Gates() []QuantumGate {
	return cloneGates(l.gates)
}

func (l Layer) NumGates() uint {
	return uint(len(l.gates))
}

func (l Layer) HasParameterizedGates() bool {
	for _, gate := range l.gates {
		if gate.HasParameters() {
			return true
		}
	}
	return false
}

func (l Layer) Clone() Layer {
	return NewLayer(l.gates)
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
