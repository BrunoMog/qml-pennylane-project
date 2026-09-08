package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqc"
)

type EmbeddingInputDTO struct {
	EmbeddingType string
	Qubits        []uint
	Rotation      string
	Normalize     bool
	PadWith       float64
}

type MeasurementInputDTO struct {
	MeasurementType     string
	MeasurementRotation string
	Qubits              []uint
}

type QuantumGateInputDTO struct {
	GateType string
	Qubit    uint
	Controls []uint
}

type LayerInputDTO struct {
	Gates []QuantumGateInputDTO
}

type VQCInputDTO struct {
	NumQubits   uint
	NumLayers   uint
	Embedding   EmbeddingInputDTO
	Measurement MeasurementInputDTO
	PreLayer    LayerInputDTO
	Layer       LayerInputDTO
	PostLayer   LayerInputDTO
}

type VQCOutputDTO struct {
	NumQubits   uint
	NumLayers   uint
	Embedding   EmbeddingInputDTO
	Measurement MeasurementInputDTO
	PreLayer    LayerInputDTO
	Layer       LayerInputDTO
	PostLayer   LayerInputDTO
}

func BuildVQC(input VQCInputDTO) (vqc.VQC, error) {
	embedding, err := buildEmbedding(input.Embedding, input.NumQubits)
	if err != nil {
		return vqc.VQC{}, err
	}

	measurement, err := buildMeasurement(input.Measurement, input.NumQubits)
	if err != nil {
		return vqc.VQC{}, err
	}

	var options []vqc.VQCOption

	if len(input.PreLayer.Gates) > 0 {
		preLayer, err := buildLayer(input.PreLayer, input.NumQubits)
		if err != nil {
			return vqc.VQC{}, err
		}
		options = append(options, vqc.WithPreLayer(preLayer))
	}

	if len(input.Layer.Gates) > 0 {
		layer, err := buildLayer(input.Layer, input.NumQubits)
		if err != nil {
			return vqc.VQC{}, err
		}
		options = append(options, vqc.WithLayer(layer))
	}

	if len(input.PostLayer.Gates) > 0 {
		postLayer, err := buildLayer(input.PostLayer, input.NumQubits)
		if err != nil {
			return vqc.VQC{}, err
		}
		options = append(options, vqc.WithPostLayer(postLayer))
	}

	vqcInput := vqc.VQCBaseInput{
		NumQubits:   input.NumQubits,
		NumLayers:   input.NumLayers,
		Embedding:   embedding,
		Measurement: measurement,
	}

	vqcInstance, err := vqc.NewVQC(vqcInput, options...)
	if err != nil {
		return vqc.VQC{}, err
	}

	return vqcInstance, nil
}

func buildEmbedding(input EmbeddingInputDTO, numQubits uint) (vqc.Embedding, error) {
	embeddingType, err := vqc.ParseEmbeddingType(input.EmbeddingType)
	if err != nil {
		return nil, err
	}
	qubits, err := buildQubits(input.Qubits, numQubits)
	if err != nil {
		return nil, err
	}

	switch embeddingType {
	case vqc.EmbeddingTypeAngle:
		return buildAngleEmbedding(input, qubits)
	case vqc.EmbeddingTypeAmplitude:
		return buildAmplitudeEmbedding(input, qubits)
	default:
		return nil, &UnreachableEmbeddingTypeError{embeddingType.Value()}
	}
}

func buildQubits(qubitIndices []uint, numQubits uint) ([]vqc.Qubit, error) {
	qubits := make([]vqc.Qubit, len(qubitIndices))
	for i, index := range qubitIndices {
		qubit, err := vqc.NewQubit(index, numQubits)
		if err != nil {
			return nil, err
		}
		qubits[i] = qubit
	}
	return qubits, nil
}

func buildAngleEmbedding(input EmbeddingInputDTO, qubits []vqc.Qubit) (vqc.Embedding, error) {
	rotation, err := vqc.ParseEmbeddingRotation(input.Rotation)
	if err != nil {
		return nil, err
	}
	return vqc.NewAngleEmbedding(qubits, rotation)
}

func buildAmplitudeEmbedding(input EmbeddingInputDTO, qubits []vqc.Qubit) (vqc.Embedding, error) {
	return vqc.NewAmplitudeEmbedding(qubits, input.Normalize, input.PadWith)
}

func buildMeasurement(input MeasurementInputDTO, numQubits uint) (vqc.Measurement, error) {
	measurementType, err := vqc.ParseMeasurementType(input.MeasurementType)
	if err != nil {
		return vqc.Measurement{}, err
	}
	qubits, err := buildQubits(input.Qubits, numQubits)
	if err != nil {
		return vqc.Measurement{}, err
	}
	rotation, err := vqc.ParseMeasurementRotation(input.MeasurementRotation)
	if err != nil {
		return vqc.Measurement{}, err
	}
	return vqc.NewMeasurement(qubits, measurementType, rotation)
}

func buildLayer(input LayerInputDTO, numQubits uint) (vqc.Layer, error) {
	gates := make([]vqc.QuantumGate, len(input.Gates))
	for i, gateInput := range input.Gates {
		gate, err := buildQuantumGate(gateInput, numQubits)
		if err != nil {
			return vqc.Layer{}, err
		}
		gates[i] = gate
	}
	return vqc.NewLayer(gates), nil
}

func buildQuantumGate(input QuantumGateInputDTO, numQubits uint) (vqc.QuantumGate, error) {
	gateType, err := vqc.ParseGateType(input.GateType)
	if err != nil {
		return vqc.QuantumGate{}, err
	}

	qubit, err := vqc.NewQubit(input.Qubit, numQubits)
	if err != nil {
		return vqc.QuantumGate{}, err
	}

	controlQubits, err := buildQubits(input.Controls, numQubits)
	if err != nil {
		return vqc.QuantumGate{}, err
	}

	return vqc.NewQuantumGate(gateType, qubit, controlQubits)
}

func BuildVQCOutput(vqcInstance vqc.VQC) VQCOutputDTO {
	embedding := vqcInstance.Embedding()
	measurement := vqcInstance.Measurement()

	preLayer := vqcInstance.PreLayer()
	layer := vqcInstance.Layer()
	postLayer := vqcInstance.PostLayer()

	embeddingOutput := buildEmbeddingOutput(embedding)
	measurementOutput := buildMeasurementOutput(measurement)
	preLayerOutput := buildLayerOutput(preLayer)
	layerOutput := buildLayerOutput(layer)
	postLayerOutput := buildLayerOutput(postLayer)

	return VQCOutputDTO{
		NumQubits:   vqcInstance.NumQubits(),
		NumLayers:   vqcInstance.NumLayers(),
		Embedding:   embeddingOutput,
		Measurement: measurementOutput,
		PreLayer:    preLayerOutput,
		Layer:       layerOutput,
		PostLayer:   postLayerOutput,
	}
}

func buildEmbeddingOutput(embedding vqc.Embedding) EmbeddingInputDTO {
	qubits := make([]uint, len(embedding.Qubits()))
	for i, qubit := range embedding.Qubits() {
		qubits[i] = qubit.Index()
	}
	switch embedding.Type() {
	case vqc.EmbeddingTypeAngle:
		return buildAngleEmbeddingOutput(embedding, qubits)
	case vqc.EmbeddingTypeAmplitude:
		return buildAmplitudeEmbeddingOutput(embedding, qubits)
	default:
		return EmbeddingInputDTO{}
	}
}

func buildAngleEmbeddingOutput(embedding vqc.Embedding, qubits []uint) EmbeddingInputDTO {
	angleEmbedding := embedding.(vqc.AngleEmbedding)
	return EmbeddingInputDTO{
		EmbeddingType: angleEmbedding.Type().Value(),
		Qubits:        qubits,
		Rotation:      angleEmbedding.Rotation().Value(),
	}
}

func buildAmplitudeEmbeddingOutput(embedding vqc.Embedding, qubits []uint) EmbeddingInputDTO {
	amplitudeEmbedding := embedding.(vqc.AmplitudeEmbedding)
	return EmbeddingInputDTO{
		EmbeddingType: amplitudeEmbedding.Type().Value(),
		Qubits:        qubits,
		Normalize:     amplitudeEmbedding.Normalize(),
		PadWith:       amplitudeEmbedding.PadWith(),
	}
}

func buildMeasurementOutput(measurement vqc.Measurement) MeasurementInputDTO {
	qubits := make([]uint, len(measurement.Qubits()))
	for i, qubit := range measurement.Qubits() {
		qubits[i] = qubit.Index()
	}
	return MeasurementInputDTO{
		MeasurementType:     measurement.MeasurementType().Value(),
		MeasurementRotation: measurement.MeasurementRotation().Value(),
		Qubits:              qubits,
	}
}

func buildLayerOutput(layer vqc.Layer) LayerInputDTO {
	gates := make([]QuantumGateInputDTO, len(layer.Gates()))
	for i, gate := range layer.Gates() {
		gates[i] = buildQuantumGateOutput(gate)
	}
	return LayerInputDTO{
		Gates: gates,
	}
}

func buildQuantumGateOutput(gate vqc.QuantumGate) QuantumGateInputDTO {
	controlQubits := make([]uint, len(gate.ControlQubits()))
	for i, qubit := range gate.ControlQubits() {
		controlQubits[i] = qubit.Index()
	}
	return QuantumGateInputDTO{
		GateType: gate.GateType().Value(),
		Qubit:    gate.Qubit().Index(),
		Controls: controlQubits,
	}
}
