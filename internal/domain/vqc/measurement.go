package vqc

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
	qubits               []Qubit
	measurement_rotation MeasurementRotation
	measurement_type     MeasurementType
}

func NewMeasurement(qubits []Qubit, measurement_type MeasurementType, measurement_rotation MeasurementRotation) (*Measurement, error) {
	err := validateMeasurement(qubits, measurement_type, measurement_rotation)
	if err != nil {
		return nil, err
	}

	return &Measurement{
		qubits:               qubits,
		measurement_rotation: measurement_rotation,
		measurement_type:     measurement_type,
	}, nil
}

func validateMeasurement(qubits []Qubit, measurement_type MeasurementType, measurement_rotation MeasurementRotation) error {
	if len(qubits) == 0 {
		return &ZeroQubitMeasurementError{qubits}
	}

	if duplicatedQubit, duplicated := hasDuplicateQubits(qubits); duplicated {
		return &DuplicateQubitError{duplicatedQubit}
	}

	if !isPermitedMeasurement(measurement_type) {
		return &InvalidMeasurementError{measurement_type}
	}

	if !isPermitedMeasurementRotation(measurement_rotation) {
		return &InvalidMeasurementRotationError{measurement_rotation}
	}

	return nil
}

func isPermitedMeasurement(measurement_type MeasurementType) bool {
	switch measurement_type {
	case ExpectationMeasurement, ProbabilityMeasurement:
		return true
	default:
		return false
	}
}

func isPermitedMeasurementRotation(measurement_rotation MeasurementRotation) bool {
	switch measurement_rotation {
	case XMeasurementRotation, YMeasurementRotation, ZMeasurementRotation:
		return true
	default:
		return false
	}
}

func (m Measurement) GetQubits() []Qubit {
	return m.qubits
}

func (m Measurement) GetMeasurementType() MeasurementType {
	return m.measurement_type
}

func (m Measurement) GetMeasurementRotation() MeasurementRotation {
	return m.measurement_rotation
}
