package vqc

type VQC struct {
	embedding   Embedding
	measurement Measurement
	pre_layer   Layer
	layer       Layer
	post_layer  Layer
	num_qubits  uint
	num_layers  uint
}

type VQCBaseInput struct {
	embedding   Embedding
	measurement Measurement
	num_qubits  uint
	num_layers  uint
}

func NewVQC(input VQCBaseInput, options ...VQCOption) (*VQC, error) {
	if input.num_qubits == 0 {
		return nil, &ZeroQubitVQCError{num_qubits: input.num_qubits}
	}
	if input.embedding == nil {
		return nil, &NilEmbeddingError{}
	}

	vqc := &VQC{
		num_qubits:  input.num_qubits,
		embedding:   input.embedding,
		measurement: input.measurement,
		num_layers:  input.num_layers,
	}

	for _, option := range options {
		option.apply(vqc)
	}

	return vqc, nil
}

func (v VQC) NumQubits() uint {
	return v.num_qubits
}

func (v VQC) NumLayers() uint {
	return v.num_layers
}

func (v VQC) Embedding() Embedding {
	return v.embedding
}

func (v VQC) PreLayer() Layer {
	return v.pre_layer
}

func (v VQC) Layer() Layer {
	return v.layer
}

func (v VQC) PostLayer() Layer {
	return v.post_layer
}

func (v VQC) Measurement() Measurement {
	return v.measurement
}

func (v VQC) NumParameters() uint {
	num_parameters := v.pre_layer.NumParameterizedGates() + v.layer.NumParameterizedGates()*v.num_layers + v.post_layer.NumParameterizedGates()
	return num_parameters
}
