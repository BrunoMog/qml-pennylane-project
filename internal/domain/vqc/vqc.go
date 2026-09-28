package vqc

import "fmt"

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
		err := option.apply(vqc)
		if err != nil {
			return VQC{}, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	return *vqc, nil
}

func validateVQCInput(input VQCBaseInput) error {
	if input.NumQubits == 0 {
		return ErrZeroQubitVQC
	}
	if input.Embedding == nil {
		return ErrNilEmbedding
	}
	if !input.Embedding.isValid() {
		return ErrInvalidEmbedding
	}
	if err := input.Embedding.validateQubits(input.NumQubits); err != nil {
		return fmt.Errorf("embedding has %w", err)
	}
	if !input.Measurement.isValid() {
		return ErrInvalidMeasurement
	}
	if err := input.Measurement.validateQubits(input.NumQubits); err != nil {
		return fmt.Errorf("measurement has %w", err)
	}

	return nil
}

func (v VQC) IsValid() bool {
	return v.numQubits > 0
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
	numParameters := v.preLayer.NumParameterizedGates() + v.layer.NumParameterizedGates()*v.numLayers + v.postLayer.NumParameterizedGates()
	return numParameters
}
