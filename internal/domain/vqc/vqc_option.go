package vqc

type VQCOption interface {
	apply(*VQC)
}

type preLayerOption struct {
	layer Layer
}

func (o preLayerOption) apply(v *VQC) {
	layer := o.layer.Clone()
	v.pre_layer = layer
}

type layerOption struct {
	layer Layer
}

func (o layerOption) apply(v *VQC) {
	layer := o.layer.Clone()
	v.layer = layer
}

type postLayerOption struct {
	layer Layer
}

func (o postLayerOption) apply(v *VQC) {
	layer := o.layer.Clone()
	v.post_layer = layer
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
