package vqc

import (
	"slices"
)

type Qubit int

func hasDuplicateQubits(qubits []Qubit) (Qubit, bool) {
	seen := make(map[Qubit]struct{}, len(qubits))

	for _, qubit := range qubits {
		if _, exists := seen[qubit]; exists {
			return qubit, true
		}
		seen[qubit] = struct{}{}
	}

	return 0, false
}

func NewQubit(index int, num_qubits uint) (Qubit, error) {
	err := validateQubit(Qubit(index), num_qubits)
	if err != nil {
		return 0, err
	}
	return Qubit(index), nil
}

func validateQubit(qubit Qubit, num_qubits uint) error {
	if qubit < 0 || int(qubit) >= int(num_qubits) {
		return &InvalidQubitError{qubit}
	}
	return nil
}

type Embedding struct {
	embedding_type string
	qubits         []Qubit
	rotation       string
	normalize      bool
	padwith        float64
}

var permited_embeddings = []string{"amplitude", "angle"}
var permited_rotations = []string{"x", "y", "z", ""}

func NewAngleEmbedding(qubits []Qubit, rotation string) (*Embedding, error) {
	err := validateEmbedding("angle", qubits, rotation)
	if err != nil {
		return nil, err
	}

	return &Embedding{
		embedding_type: "angle",
		qubits:         qubits,
		rotation:       rotation,
	}, nil
}

func NewAmplitudeEmbedding(qubits []Qubit, normalize bool, padwith float64) (*Embedding, error) {
	err := validateEmbedding("amplitude", qubits, "")
	if err != nil {
		return nil, err
	}

	return &Embedding{
		embedding_type: "amplitude",
		qubits:         qubits,
		normalize:      normalize,
		padwith:        padwith,
	}, nil
}

func validateEmbedding(embedding_type string, qubits []Qubit, rotation string) error {
	if !isPermitedEmbedding(embedding_type) {
		return &InvalidEmbeddingError{embedding_type}
	}

	if len(qubits) == 0 {
		return &ZeroQubitEmbeddingError{qubits}
	}

	if duplicatedQubit, duplicated := hasDuplicateQubits(qubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit}
	}

	if !isPermitedRotation(rotation) {
		return &InvalidRotationError{rotation}
	}

	return nil
}

func isPermitedEmbedding(embedding_type string) bool {
	result := slices.Contains(permited_embeddings, embedding_type)
	return result
}

func isPermitedRotation(rotation string) bool {
	result := slices.Contains(permited_rotations, rotation)
	return result
}

type Layer struct {
	gates []QuantumGates
}

func NewLayer() *Layer {
	return &Layer{
		gates: []QuantumGates{},
	}
}

func (l *Layer) AddGate(gate QuantumGates) {
	l.gates = append(l.gates, gate)
}

func (l *Layer) RemoveGate(index int) error {
	if index < 0 || index >= len(l.gates) {
		return &InvalidIndexError{index}
	}
	l.gates = slices.Delete(l.gates, index, index+1)
	return nil
}

func (l *Layer) AddGateAtIndex(gate QuantumGates, index int) error {
	if index < 0 || index > len(l.gates) {
		return &InvalidIndexError{index}
	}
	l.gates = slices.Insert(l.gates, index, gate)
	return nil
}

func (l *Layer) UpdateGateAtIndex(gate QuantumGates, index int) error {
	if index < 0 || index >= len(l.gates) {
		return &InvalidIndexError{index}
	}
	l.gates[index] = gate
	return nil
}

type Measurement struct {
	qubits           []Qubit
	measurement_type string
}

var permited_measurements = []string{"expectation", "probability"}

func NewMeasurement(qubits []Qubit, measurement_type string) (*Measurement, error) {
	err := validateMeasurement(qubits, measurement_type)
	if err != nil {
		return nil, err
	}

	return &Measurement{
		qubits:           qubits,
		measurement_type: measurement_type,
	}, nil
}

func validateMeasurement(qubits []Qubit, measurement_type string) error {
	if len(qubits) == 0 {
		return &ZeroQubitMeasurementError{qubits}
	}

	if duplicatedQubit, duplicated := hasDuplicateQubits(qubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit}
	}

	if !isPermitedMeasurement(measurement_type) {
		return &InvalidMeasurementError{measurement_type}
	}

	return nil
}

func isPermitedMeasurement(measurement_type string) bool {
	result := slices.Contains(permited_measurements, measurement_type)
	return result
}
