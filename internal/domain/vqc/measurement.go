package vqc

import (
	"slices"
)

type MeasurementType string

const (
	ExpectationMeasurement MeasurementType = "expectation"
	ProbabilityMeasurement MeasurementType = "probability"
)

type MeasurementRotation string

const (
	XMeasurementRotation MeasurementRotation = "x"
	YMeasurementRotation MeasurementRotation = "y"
	ZMeasurementRotation MeasurementRotation = "z"
)

type Measurement struct {
	measurementRotation MeasurementRotation
	measurementType     MeasurementType
	qubits              []Qubit
}

func NewMeasurement(qubits []Qubit, measurementType MeasurementType, measurementRotation MeasurementRotation) (Measurement, error) {
	err := validateMeasurement(qubits, measurementType, measurementRotation)
	if err != nil {
		return Measurement{}, err
	}

	return Measurement{
		qubits:              qubits,
		measurementRotation: measurementRotation,
		measurementType:     measurementType,
	}, nil
}

func validateMeasurement(qubits []Qubit, measurementType MeasurementType, measurementRotation MeasurementRotation) error {
	if len(qubits) == 0 {
		return &ZeroQubitMeasurementError{qubits}
	}

	if duplicatedQubit, duplicated := hasDuplicateQubits(qubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit}
	}

	if !isPermitedMeasurement(measurementType) {
		return &InvalidMeasurementError{measurementType}
	}

	if !isPermitedMeasurementRotation(measurementRotation) {
		return &InvalidMeasurementRotationError{measurementRotation}
	}

	return nil
}

func isPermitedMeasurement(measurementType MeasurementType) bool {
	switch measurementType {
	case ExpectationMeasurement, ProbabilityMeasurement:
		return true
	default:
		return false
	}
}

func isPermitedMeasurementRotation(measurementRotation MeasurementRotation) bool {
	switch measurementRotation {
	case XMeasurementRotation, YMeasurementRotation, ZMeasurementRotation:
		return true
	default:
		return false
	}
}

func (m Measurement) Qubits() []Qubit {
	return slices.Clone(m.qubits)
}

func (m Measurement) MeasurementType() MeasurementType {
	return m.measurementType
}

func (m Measurement) MeasurementRotation() MeasurementRotation {
	return m.measurementRotation
}
