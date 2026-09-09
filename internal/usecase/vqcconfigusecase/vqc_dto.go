package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/vqc"
)

type EmbeddingDTO struct {
	EmbeddingType string
	Qubits        []uint
	Rotation      string
	Normalize     bool
	PadWith       float64
}

type MeasurementDTO struct {
	MeasurementType     string
	MeasurementRotation string
	Qubits              []uint
}

type QuantumGateDTO struct {
	GateType string
	Qubit    uint
	Controls []uint
}

type LayerDTO struct {
	Gates []QuantumGateDTO
}

type VQCDTO struct {
	NumQubits   uint
	NumLayers   uint
	Embedding   EmbeddingDTO
	Measurement MeasurementDTO
	PreLayer    LayerDTO
	Layer       LayerDTO
	PostLayer   LayerDTO
}

func buildVQCDTOToVQC(input VQCDTO) (vqc.VQC, error) {
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

func buildEmbedding(input EmbeddingDTO, numQubits uint) (vqc.Embedding, error) {
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

func buildAngleEmbedding(input EmbeddingDTO, qubits []vqc.Qubit) (vqc.Embedding, error) {
	rotation, err := vqc.ParseEmbeddingRotation(input.Rotation)
	if err != nil {
		return nil, err
	}
	return vqc.NewAngleEmbedding(qubits, rotation)
}

func buildAmplitudeEmbedding(input EmbeddingDTO, qubits []vqc.Qubit) (vqc.Embedding, error) {
	return vqc.NewAmplitudeEmbedding(qubits, input.Normalize, input.PadWith)
}

func buildMeasurement(input MeasurementDTO, numQubits uint) (vqc.Measurement, error) {
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

func buildLayer(input LayerDTO, numQubits uint) (vqc.Layer, error) {
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

func buildQuantumGate(input QuantumGateDTO, numQubits uint) (vqc.QuantumGate, error) {
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

func buildVQCToVQCDTO(vqcInstance vqc.VQC) VQCDTO {
	embedding := vqcInstance.Embedding()
	measurement := vqcInstance.Measurement()

	preLayer := vqcInstance.PreLayer()
	layer := vqcInstance.Layer()
	postLayer := vqcInstance.PostLayer()

	embeddingOutput := buildEmbeddingDTO(embedding)
	measurementOutput := buildMeasurementDTO(measurement)
	preLayerOutput := buildLayerDTO(preLayer)
	layerOutput := buildLayerDTO(layer)
	postLayerOutput := buildLayerDTO(postLayer)

	return VQCDTO{
		NumQubits:   vqcInstance.NumQubits(),
		NumLayers:   vqcInstance.NumLayers(),
		Embedding:   embeddingOutput,
		Measurement: measurementOutput,
		PreLayer:    preLayerOutput,
		Layer:       layerOutput,
		PostLayer:   postLayerOutput,
	}
}

func buildEmbeddingDTO(embedding vqc.Embedding) EmbeddingDTO {
	qubits := make([]uint, len(embedding.Qubits()))
	for i, qubit := range embedding.Qubits() {
		qubits[i] = qubit.Index()
	}
	switch embedding.Type() {
	case vqc.EmbeddingTypeAngle:
		return buildAngleEmbeddingDTO(embedding, qubits)
	case vqc.EmbeddingTypeAmplitude:
		return buildAmplitudeEmbeddingDTO(embedding, qubits)
	default:
		return EmbeddingDTO{}
	}
}

func buildAngleEmbeddingDTO(embedding vqc.Embedding, qubits []uint) EmbeddingDTO {
	angleEmbedding, ok := embedding.(vqc.AngleEmbedding)
	if !ok {
		return EmbeddingDTO{}
	}
	return EmbeddingDTO{
		EmbeddingType: angleEmbedding.Type().Value(),
		Qubits:        qubits,
		Rotation:      angleEmbedding.Rotation().Value(),
	}
}

func buildAmplitudeEmbeddingDTO(embedding vqc.Embedding, qubits []uint) EmbeddingDTO {
	amplitudeEmbedding, ok := embedding.(vqc.AmplitudeEmbedding)
	if !ok {
		return EmbeddingDTO{}
	}
	return EmbeddingDTO{
		EmbeddingType: amplitudeEmbedding.Type().Value(),
		Qubits:        qubits,
		Normalize:     amplitudeEmbedding.Normalize(),
		PadWith:       amplitudeEmbedding.PadWith(),
	}
}

func buildMeasurementDTO(measurement vqc.Measurement) MeasurementDTO {
	qubits := make([]uint, len(measurement.Qubits()))
	for i, qubit := range measurement.Qubits() {
		qubits[i] = qubit.Index()
	}
	return MeasurementDTO{
		MeasurementType:     measurement.MeasurementType().Value(),
		MeasurementRotation: measurement.MeasurementRotation().Value(),
		Qubits:              qubits,
	}
}

func buildLayerDTO(layer vqc.Layer) LayerDTO {
	gates := make([]QuantumGateDTO, len(layer.Gates()))
	for i, gate := range layer.Gates() {
		gates[i] = buildQuantumGateDTO(gate)
	}
	return LayerDTO{
		Gates: gates,
	}
}

func buildQuantumGateDTO(gate vqc.QuantumGate) QuantumGateDTO {
	controlQubits := make([]uint, len(gate.ControlQubits()))
	for i, qubit := range gate.ControlQubits() {
		controlQubits[i] = qubit.Index()
	}
	return QuantumGateDTO{
		GateType: gate.GateType().Value(),
		Qubit:    gate.Qubit().Index(),
		Controls: controlQubits,
	}
}
