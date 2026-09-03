package vqc

type EmbeddingBuilderInput struct {
	Type      string
	Rotation  string
	Qubits    []int
	PadWidth  float64
	Normalize bool
}

type MeasurementBuilderInput struct {
	Type     string // Measurement type
	Rotation string // Measurement rotation
	Qubits   []int  // Qubits to measure
}

type GateBuilderInput struct {
	GateType      string
	ControlQubits []int
	Qubit         int
}

type VQCBuilderInput struct {
	Measurement MeasurementBuilderInput
	Embedding   EmbeddingBuilderInput
	NumQubits   int
	NumLayers   int
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
		numQubits: uint(input.NumQubits),
		numLayers: uint(input.NumLayers),
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

// Delegates validation to Embedding constructors (Secure by Design)
func buildEmbedding(input EmbeddingBuilderInput, numQubits int) (Embedding, error) {
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
		return NewAmplitudeEmbedding(qubits, input.Normalize, input.PadWidth)
	default:
		return nil, &InvalidEmbeddingError{EmbeddingType(input.Type)}
	}
}

func buildQubits(qubitIndices []int, numQubits int) ([]Qubit, error) {
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

func buildMeasurement(input MeasurementBuilderInput, numQubits int) (Measurement, error) {
	qubits, err := buildQubits(input.Qubits, numQubits)
	if err != nil {
		return Measurement{}, err
	}
	measurementType := MeasurementType(input.Type)
	rotation := MeasurementRotation(input.Rotation)
	return NewMeasurement(qubits, measurementType, rotation)
}

// WithPreLayer sets the initial state preparation layer
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
	return *NewLayer(layer), nil
}

func buildGate(gateInput GateBuilderInput, numQubits uint) (*QuantumGate, error) {
	gateType := GateType(gateInput.GateType)
	qubit, err := NewQubit(uint(gateInput.Qubit), numQubits)
	if err != nil {
		return nil, err
	}
	controlQubits, err := buildQubits(gateInput.ControlQubits, int(numQubits))
	if err != nil {
		return nil, err
	}
	return NewQuantumGate(gateType, qubit, controlQubits)
}

// WithLayer sets the parameterized variational layer
func (b *VQCBuilder) WithLayer(gates []GateBuilderInput) (*VQCBuilder, error) {
	layer, err := buildLayer(gates, b.numQubits)
	if err != nil {
		return nil, err
	}
	b.layer = layer
	return b, nil
}

// WithPostLayer sets the measurement basis rotation layer
func (b *VQCBuilder) WithPostLayer(gates []GateBuilderInput) (*VQCBuilder, error) {
	layer, err := buildLayer(gates, b.numQubits)
	if err != nil {
		return nil, err
	}
	b.postLayer = layer
	return b, nil
}

// Build returns the constructed VQC (validation delegated to NewVQC)
func (b *VQCBuilder) Build() (*VQC, error) {
	return NewVQC(b.numQubits, b.embedding, b.preLayer, b.layer, b.postLayer, b.measurement, b.numLayers)
}
