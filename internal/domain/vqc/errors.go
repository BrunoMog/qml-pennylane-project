package vqc

import "fmt"

type InvalidGateError struct {
	gate_type GateType
}

func (e *InvalidGateError) Error() string {
	return fmt.Sprintf("Invalid gate type: %s", e.gate_type)
}

type InvalidQubitError struct {
	qubit uint
}

func (e *InvalidQubitError) Error() string {
	return fmt.Sprintf("Invalid qubit index: %d", e.qubit)
}

type InvalidControlQubitError struct {
	control_qubit []Qubit
}

func (e *InvalidControlQubitError) Error() string {
	return fmt.Sprintf("Invalid control qubit indices: %v", e.control_qubit)
}

type DuplicateQubitError struct {
	qubit Qubit
}

func (e *DuplicateQubitError) Error() string {
	return fmt.Sprintf("Duplicate qubit index: %d", e.qubit)
}

type InvalidEmbeddingError struct {
	embedding_type EmbeddingType
}

func (e *InvalidEmbeddingError) Error() string {
	return fmt.Sprintf("Invalid embedding type: %s", e.embedding_type)
}

type ZeroQubitEmbeddingError struct {
	qubit []Qubit
}

func (e *ZeroQubitEmbeddingError) Error() string {
	return fmt.Sprintf("Embedding must have at least one qubit, got: %v", e.qubit)
}

type InvalidRotationError struct {
	rotation EmbeddingRotation
}

func (e *InvalidRotationError) Error() string {
	return fmt.Sprintf("Invalid rotation type: %s", e.rotation)
}

type InvalidIndexError struct {
	index LayerIndex
}

func (e *InvalidIndexError) Error() string {
	return fmt.Sprintf("Invalid index: %d", e.index)
}

type InvalidMeasurementError struct {
	measurement_type MeasurementType
}

func (e *InvalidMeasurementError) Error() string {
	return fmt.Sprintf("Invalid measurement type: %s", e.measurement_type)
}

type ZeroQubitMeasurementError struct {
	qubit []Qubit
}

func (e *ZeroQubitMeasurementError) Error() string {
	return fmt.Sprintf("Measurement must have at least one qubit, got: %v", e.qubit)
}

type ZeroQubitVQCError struct {
	num_qubits uint
}

func (e *ZeroQubitVQCError) Error() string {
	return fmt.Sprintf("VQC must have at least one qubit, got: %d", e.num_qubits)
}

type InvalidMeasurementRotationError struct {
	measurement_rotation MeasurementRotation
}

func (e *InvalidMeasurementRotationError) Error() string {
	return fmt.Sprintf("Invalid measurement rotation: %s", e.measurement_rotation)
}
