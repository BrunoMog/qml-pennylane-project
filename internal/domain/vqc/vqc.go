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

func NewVQC(num_qubits uint, embedding Embedding, pre_layer Layer, layer Layer, post_layer Layer, measurement Measurement, num_layers uint) (*VQC, error) {
	if num_qubits == 0 {
		return nil, &ZeroQubitVQCError{num_qubits: num_qubits}
	}

	return &VQC{
		num_qubits:  num_qubits,
		embedding:   embedding,
		pre_layer:   pre_layer,
		layer:       layer,
		post_layer:  post_layer,
		measurement: measurement,
		num_layers:  num_layers,
	}, nil
}

func (v VQC) GetNumQubits() uint {
	return v.num_qubits
}

func (v VQC) GetNumLayers() uint {
	return v.num_layers
}

func (v VQC) GetEmbedding() Embedding {
	return v.embedding
}

func (v VQC) GetPreLayer() Layer {
	return v.pre_layer
}

func (v VQC) GetLayer() Layer {
	return v.layer
}

func (v VQC) GetPostLayer() Layer {
	return v.post_layer
}

func (v VQC) GetMeasurement() Measurement {
	return v.measurement
}

func (v VQC) GetNumParameters() uint {
	num_parameters := v.pre_layer.GetNumParameterizedGates() + v.layer.GetNumParameterizedGates()*v.num_layers + v.post_layer.GetNumParameterizedGates()
	return num_parameters
}
