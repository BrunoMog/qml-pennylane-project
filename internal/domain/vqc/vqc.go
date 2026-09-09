package vqc

type VQC struct {
	embedding   Embedding
	measurement Measurement
	preLayer    Layer
	layer       Layer
	postLayer   Layer
	numQubits   uint
	numLayers   uint
}

type VQCBaseInput struct {
	Embedding   Embedding
	Measurement Measurement
	NumQubits   uint
	NumLayers   uint
}

func NewVQC(input VQCBaseInput, options ...VQCOption) (VQC, error) {
	err := validateVQCInput(input)
	if err != nil {
		return VQC{}, err
	}

	vqc := &VQC{
		numQubits:   input.NumQubits,
		embedding:   input.Embedding,
		measurement: input.Measurement,
		numLayers:   input.NumLayers,
	}

	for _, option := range options {
		option.apply(vqc)
	}

	return *vqc, nil
}

func validateVQCInput(input VQCBaseInput) error {
	if input.NumQubits == 0 {
		return &ZeroQubitVQCError{numQubits: input.NumQubits}
	}
	if input.Embedding == nil {
		return &NilEmbeddingError{}
	}
	if !input.Embedding.IsValid() {
		return &InvalidEmbeddingError{}
	}
	if !input.Measurement.IsValid() {
		return &InvalidMeasurementError{}
	}
	return nil
}

func (v VQC) IsValid() bool {
	if v.numQubits == 0 {
		return false
	}
	if v.embedding == nil || !v.embedding.IsValid() {
		return false
	}
	if !v.measurement.IsValid() {
		return false
	}

	return true
}

func (v VQC) Equals(other VQC) bool {
	if v.numQubits != other.numQubits {
		return false
	}
	if v.numLayers != other.numLayers {
		return false
	}
	if !v.embedding.Equals(other.embedding) {
		return false
	}
	if !v.measurement.Equals(other.measurement) {
		return false
	}
	if !v.preLayer.Equals(other.preLayer) {
		return false
	}
	if !v.layer.Equals(other.layer) {
		return false
	}
	if !v.postLayer.Equals(other.postLayer) {
		return false
	}

	return true
}

func (v VQC) NumQubits() uint {
	return v.numQubits
}

func (v VQC) NumLayers() uint {
	return v.numLayers
}

func (v VQC) Embedding() Embedding {
	return v.embedding
}

func (v VQC) PreLayer() Layer {
	return v.preLayer
}

func (v VQC) Layer() Layer {
	return v.layer
}

func (v VQC) PostLayer() Layer {
	return v.postLayer
}

func (v VQC) Measurement() Measurement {
	return v.measurement
}

func (v VQC) NumParameters() uint {
	num_parameters := v.preLayer.NumParameterizedGates() + v.layer.NumParameterizedGates()*v.numLayers + v.postLayer.NumParameterizedGates()
	return num_parameters
}
