package vqc

type EmbeddingBuilderInput struct {
	Type      string  // "angle" or "amplitude"
	Qubits    []int   // Qubit indices for embedding
	Rotation  string  // Rotation type (for angle embedding)
	Normalize bool    // Normalize (for amplitude embedding)
	PadWidth  float64 // Pad width (for amplitude embedding)
}

type MeasurementBuilderInput struct {
	Type     string // Measurement type
	Rotation string // Measurement rotation
	Qubits   []int  // Qubits to measure
}

type GateBuilderInput struct {
	GateType      string // Type of gate (e.g., "RX", "CNOT")
	Qubit         int    // Target qubit index
	ControlQubits []int  // Control qubits (for controlled gates)
}

type VQCBuilderInput struct {
	NumQubits   int
	NumLayers   int
	Embedding   EmbeddingBuilderInput
	Measurement MeasurementBuilderInput
}

type VQCBuilder struct {
	numQubits   uint
	numLayers   uint
	embedding   *Embedding
	preLayer    Layer
	layer       Layer
	postLayer   Layer
	measurement *Measurement
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
func buildEmbedding(input EmbeddingBuilderInput, numQubits int) (*Embedding, error) {
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

func buildMeasurement(input MeasurementBuilderInput, numQubits int) (*Measurement, error) {
	qubits, err := buildQubits(input.Qubits, numQubits)
	if err != nil {
		return nil, err
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
	layer := NewLayer()
	for _, gateInput := range gates {
		gate, err := buildGate(gateInput, numQubits)
		if err != nil {
			return Layer{}, err
		}
		layer.AddGate(*gate)
	}
	return *layer, nil
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
