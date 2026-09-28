package vqc

import "fmt"

type VQCOption interface {
	apply(*VQC) error
}

type preLayerOption struct {
	layer Layer
}

func (o preLayerOption) apply(v *VQC) error {
	layer := o.layer.Clone()
	err := layer.validateQubits(v.numQubits)
	if err != nil {
		return fmt.Errorf("pre-layer has a gate with %w", err)
	}

	v.preLayer = layer
	return nil
}

type layerOption struct {
	layer Layer
}

func (o layerOption) apply(v *VQC) error {
	layer := o.layer.Clone()
	err := layer.validateQubits(v.numQubits)
	if err != nil {
		return fmt.Errorf("layer has a gate with %w", err)
	}

	v.layer = layer
	return nil
}

type postLayerOption struct {
	layer Layer
}

func (o postLayerOption) apply(v *VQC) error {
	layer := o.layer.Clone()
	err := layer.validateQubits(v.numQubits)
	if err != nil {
		return fmt.Errorf("post-layer has a gate with %w", err)
	}

	v.postLayer = layer
	return nil
}

func WithPreLayer(layer Layer) VQCOption {
	return preLayerOption{layer: layer}
}

func WithLayer(layer Layer) VQCOption {
	return layerOption{layer: layer}
}

func WithPostLayer(layer Layer) VQCOption {
	return postLayerOption{layer: layer}
}
