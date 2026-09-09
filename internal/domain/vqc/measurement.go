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

func (m Measurement) IsValid() bool {
	if len(m.qubits) == 0 {
		return false
	}

	if _, duplicated := hasDuplicateQubits(m.qubits); duplicated {
		return false
	}

	if !isValidMeasurementType(m.measurementType) {
		return false
	}

	if !isValidMeasurementRotation(m.measurementRotation) {
		return false
	}

	return true
}

func (m Measurement) Equals(other Measurement) bool {
	if m.measurementType != other.measurementType {
		return false
	}

	if m.measurementRotation != other.measurementRotation {
		return false
	}

	if len(m.qubits) != len(other.qubits) {
		return false
	}

	for i, qubit := range m.qubits {
		if qubit != other.qubits[i] {
			return false
		}
	}

	return true
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
