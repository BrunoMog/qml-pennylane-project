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
	Embedding   Embedding
	Measurement Measurement
	NumQubits   uint
	NumLayers   uint
}

func NewVQC(input VQCBaseInput, options ...VQCOption) (VQC, error) {
	if input.NumQubits == 0 {
		return VQC{}, &ZeroQubitVQCError{num_qubits: input.NumQubits}
	}
	if input.Embedding == nil {
		return VQC{}, &NilEmbeddingError{}
	}

	vqc := &VQC{
		num_qubits:  input.NumQubits,
		embedding:   input.Embedding,
		measurement: input.Measurement,
		num_layers:  input.NumLayers,
	}

	for _, option := range options {
		option.apply(vqc)
	}

	return *vqc, nil
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
