package vqc

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidQubit               = errors.New("Invalid qubit index, must be less than the number of circuit qubits")
	ErrDuplicatedQubit            = errors.New("duplicated qubit")
	ErrInvalidPadWith             = errors.New("Invalid pad width: must be a finite number")
	ErrZeroQubitEmbedding         = errors.New("Embedding must have at least one qubit")
	ErrParseEmbeddingRotation     = errors.New("embedding rotation cannot be parsed, must be one of 'x', 'y', or 'z'")
	ErrParseEmbeddingType         = errors.New("embedding type cannot be parsed, must be one of 'angle' or 'amplitude'")
	ErrInvalidEmbeddingRotation   = errors.New("Invalid rotation: must be one of 'x', 'y', or 'z'")
	ErrParseMeasurementRotation   = errors.New("measurement rotation cannot be parsed, must be one of 'x', 'y', or 'z'")
	ErrParseMeasurementType       = errors.New("measurement type cannot be parsed, must be one of 'expectation', or 'probability'")
	ErrZeroQubitMeasurement       = errors.New("Measurement must have at least one qubit")
	ErrInvalidMeasurementType     = errors.New("Invalid measurement type must be one of 'expectation' or 'probability'")
	ErrInvalidMeasurementRotation = errors.New("Invalid measurement rotation must be one of 'x', 'y', or 'z'")
	ErrParseGateType              = errors.New("gate type cannot be parsed, must be a valid gate type")
	ErrInvalidGateType            = errors.New("Invalid gate type")
	ErrInvalidControlQubit        = errors.New("Incompatible control qubit")
	ErrZeroQubitVQC               = errors.New("VQC must have at least one qubit")
	ErrNilEmbedding               = errors.New("VQC embedding cannot be nil")
	ErrInvalidEmbedding           = errors.New("VQC embedding is invalid")
	ErrInvalidMeasurement         = errors.New("VQC measurement is invalid")
	ErrInvalidQuantumGate         = errors.New("quantum gate cannot be zero-valued")
)

type DuplicateQubitError struct {
	QubitIndex uint
}

func (e *DuplicateQubitError) Error() string {
	return fmt.Sprintf("duplicate qubit: %d", e.QubitIndex)
}

func (e *DuplicateQubitError) Is(target error) bool {
	return target == ErrDuplicatedQubit
}

type InvalidControlQubitError struct {
	Reason string
}

func (e *InvalidControlQubitError) Error() string {
	return e.Reason
}

func (e *InvalidControlQubitError) Is(target error) bool {
	return target == ErrInvalidControlQubit
}

type InvalidQubitError struct {
	QubitIndex uint
}

func (e *InvalidQubitError) Error() string {
	return fmt.Sprintf("invalid qubit index: %d", e.QubitIndex)
}

func (e *InvalidQubitError) Is(target error) bool {
	return target == ErrInvalidQubit
}
