package vqc

import (
	"errors"
	"fmt"
	"slices"
)

type Measurement struct {
	measurementRotation MeasurementRotation
	measurementType     MeasurementType
	qubits              []Qubit
}

func NewMeasurement(qubits []Qubit, measurementType MeasurementType, measurementRotation MeasurementRotation) (Measurement, error) {
	err := validateMeasurement(qubits, measurementType, measurementRotation)
	if err != nil {
		if errors.Is(err, ErrDuplicatedQubit) {
			return Measurement{}, fmt.Errorf("measurement has %w", err)
		}
		return Measurement{}, err
	}

	return Measurement{
		qubits:              slices.Clone(qubits),
		measurementRotation: measurementRotation,
		measurementType:     measurementType,
	}, nil
}

func validateMeasurement(qubits []Qubit, measurementType MeasurementType, measurementRotation MeasurementRotation) error {
	if len(qubits) == 0 {
		return ErrZeroQubitMeasurement
	}

	if duplicatedQubit, duplicated := hasDuplicateQubits(qubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit.Index()}
	}

	if !measurementType.isValid() {
		return ErrInvalidMeasurementType
	}

	if !measurementRotation.isValid() {
		return ErrInvalidMeasurementRotation
	}

	return nil
}

func (m Measurement) isValid() bool {
	return len(m.qubits) > 0 && m.measurementType != "" && m.measurementRotation != ""
}

func (m Measurement) validateQubits(numQubits uint) error {
	for _, qubit := range m.qubits {
		if qubit.Index() >= numQubits {
			return &InvalidQubitError{QubitIndex: qubit.Index()}
		}
	}
	return nil
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

	return slices.Equal(m.qubits, other.qubits)
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
