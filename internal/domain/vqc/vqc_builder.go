package vqc

type EmbeddingBuilderInput struct {
	Type      EmbeddingType
	Rotation  EmbeddingRotation
	Qubits    []Qubit
	PadWidth  float64
	Normalize bool
}

type MeasurementBuilderInput struct {
	Type     MeasurementType
	Rotation MeasurementRotation
	Qubits   []Qubit
}

type VQCBuilderInput struct {
	Measurement MeasurementBuilderInput
	Embedding   EmbeddingBuilderInput
	NumQubits   uint
	NumLayers   uint
}

type VQCBuilder struct {
	embedding   Embedding
	measurement Measurement
	preLayer    Layer
	layer       Layer
	postLayer   Layer
	numQubits   uint
	numLayers   uint
}

func NewVQCBuilder(input VQCBuilderInput) (*VQCBuilder, error) {
	builder := &VQCBuilder{
		numQubits: input.NumQubits,
		numLayers: input.NumLayers,
	}

	measurement, err := buildMeasurement(input.Measurement, input.NumQubits)
	if err != nil {
		return nil, err
	}
	builder.measurement = measurement

	embedding, err := buildEmbedding(input.Embedding, input.NumQubits)
	if err != nil {
		return nil, err
	}
	builder.embedding = embedding

	return builder, nil
}

func buildEmbedding(input EmbeddingBuilderInput, numQubits uint) (Embedding, error) {
	switch input.Type {
	case EmbeddingTypeAngle:
		return NewAngleEmbedding(input.Qubits, input.Rotation)
	case EmbeddingTypeAmplitude:
		return NewAmplitudeEmbedding(input.Qubits, input.Normalize, input.PadWidth)
	default:
		return nil, &InvalidEmbeddingError{EmbeddingType(input.Type)}
	}
}

func buildMeasurement(input MeasurementBuilderInput, numQubits uint) (Measurement, error) {
	return NewMeasurement(input.Qubits, input.Type, input.Rotation)
}

func (b *VQCBuilder) WithPreLayer(gates []QuantumGate) (*VQCBuilder, error) {
	layer := NewLayer(gates)
	b.preLayer = *layer
	return b, nil
}

func (b *VQCBuilder) WithLayer(gates []QuantumGate) (*VQCBuilder, error) {
	layer := NewLayer(gates)
	b.layer = *layer
	return b, nil
}

func (b *VQCBuilder) WithPostLayer(gates []QuantumGate) (*VQCBuilder, error) {
	layer := NewLayer(gates)
	b.postLayer = *layer
	return b, nil
}

func (b *VQCBuilder) Build() (*VQC, error) {
	input := VQCInput{
		embedding:   b.embedding,
		measurement: b.measurement,
		pre_layer:   b.preLayer,
		layer:       b.layer,
		post_layer:  b.postLayer,
		num_qubits:  b.numQubits,
		num_layers:  b.numLayers,
	}
	return NewVQC(input)
}
