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

type VQCInput struct {
	embedding   Embedding
	measurement Measurement
	pre_layer   Layer
	layer       Layer
	post_layer  Layer
	num_qubits  uint
	num_layers  uint
}

func NewVQC(input VQCInput) (*VQC, error) {
	if input.num_qubits == 0 {
		return nil, &ZeroQubitVQCError{num_qubits: input.num_qubits}
	}

	return &VQC{
		num_qubits:  input.num_qubits,
		embedding:   input.embedding,
		pre_layer:   input.pre_layer,
		layer:       input.layer,
		post_layer:  input.post_layer,
		measurement: input.measurement,
		num_layers:  input.num_layers,
	}, nil
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
