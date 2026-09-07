package vqc

import (
	"slices"
)

type Measurement struct {
	measurementRotation MeasurementRotation
	measurementType     MeasurementType
	qubits              []Qubit
}

func NewMeasurement(qubits []Qubit, measurementType MeasurementType, measurementRotation MeasurementRotation) (Measurement, error) {
	qubits = slices.Clone(qubits)
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

	if !isValidMeasurementType(measurementType) {
		return &InvalidMeasurementError{measurementType}
	}

	if !isValidMeasurementRotation(measurementRotation) {
		return &InvalidMeasurementRotationError{measurementRotation}
	}

	return nil
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
