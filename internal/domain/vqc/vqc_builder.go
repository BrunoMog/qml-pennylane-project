package vqc

type EmbeddingBuilderInput struct {
	Type      EmbeddingType
	Rotation  EmbeddingRotation
	Qubits    []uint
	PadWith   float64
	Normalize bool
}

type MeasurementBuilderInput struct {
	Type     MeasurementType
	Rotation MeasurementRotation
	Qubits   []uint
}

type GateBuilderInput struct {
	GateType      GateType
	ControlQubits []uint
	Qubit         uint
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
	case "angle":
		qubits, err := buildQubits(input.Qubits, numQubits)
		if err != nil {
			return nil, err
		}
		rotation := EmbeddingRotation(input.Rotation)
		return NewAngleEmbedding(qubits, rotation)
	case "amplitude":
		qubits, err := buildQubits(input.Qubits, numQubits)
		if err != nil {
			return nil, err
		}
		return NewAmplitudeEmbedding(qubits, input.Normalize, input.PadWith)
	default:
		return nil, &InvalidEmbeddingError{EmbeddingType(input.Type)}
	}
}

func buildQubits(qubitIndices []uint, numQubits uint) ([]Qubit, error) {
	qubits := make([]Qubit, len(qubitIndices))
	for i, index := range qubitIndices {
		qubit, err := NewQubit(uint(index), uint(numQubits))
		if err != nil {
			return nil, err
		}
		qubits[i] = qubit
	}
	return qubits, nil
}

func buildMeasurement(input MeasurementBuilderInput, numQubits uint) (Measurement, error) {
	qubits, err := buildQubits(input.Qubits, numQubits)
	if err != nil {
		return Measurement{}, err
	}
	measurementType := MeasurementType(input.Type)
	rotation := MeasurementRotation(input.Rotation)
	return NewMeasurement(qubits, measurementType, rotation)
}

func (b *VQCBuilder) WithPreLayer(gates []GateBuilderInput) (*VQCBuilder, error) {
	layer, err := buildLayer(gates, b.numQubits)
	if err != nil {
		return nil, err
	}
	b.preLayer = layer
	return b, nil
}

func buildLayer(gates []GateBuilderInput, numQubits uint) (Layer, error) {
	layer := make([]QuantumGate, 0, len(gates))
	for _, gateInput := range gates {
		gate, err := buildGate(gateInput, numQubits)
		if err != nil {
			return Layer{}, err
		}
		layer = append(layer, *gate)
	}
	return NewLayer(layer), nil
}

func buildGate(gateInput GateBuilderInput, numQubits uint) (*QuantumGate, error) {
	gateType := GateType(gateInput.GateType)
	qubit, err := NewQubit(uint(gateInput.Qubit), numQubits)
	if err != nil {
		return nil, err
	}
	controlQubits, err := buildQubits(gateInput.ControlQubits, numQubits)
	if err != nil {
		return nil, err
	}
	return NewQuantumGate(gateType, qubit, controlQubits)
}

func (b *VQCBuilder) WithLayer(gates []GateBuilderInput) (*VQCBuilder, error) {
	layer, err := buildLayer(gates, b.numQubits)
	if err != nil {
		return nil, err
	}
	b.layer = layer
	return b, nil
}

func (b *VQCBuilder) WithPostLayer(gates []GateBuilderInput) (*VQCBuilder, error) {
	layer, err := buildLayer(gates, b.numQubits)
	if err != nil {
		return nil, err
	}
	b.postLayer = layer
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
