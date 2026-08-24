package vqc

type VQC struct {
	num_qubits  uint
	embedding   *Embedding
	pre_layer   Layer
	layer       Layer
	post_layer  Layer
	measurement *Measurement
}

func NewVQC(num_qubits uint, embedding *Embedding, pre_layer Layer, layer Layer, post_layer Layer, measurement *Measurement, weights_seed int) (*VQC, error) {
	if num_qubits == 0 {
		return nil, &ZeroQubitVQCError{num_qubits}
	}

	return &VQC{
		num_qubits:  num_qubits,
		embedding:   embedding,
		pre_layer:   pre_layer,
		layer:       layer,
		post_layer:  post_layer,
		measurement: measurement,
	}, nil
}
