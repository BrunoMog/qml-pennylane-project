package vqc

import "strings"

type MeasurementType string

const (
	ExpectationMeasurement MeasurementType = "expectation"
	ProbabilityMeasurement MeasurementType = "probability"
)

const maxMeasurementTypeLength = 20

func (m MeasurementType) String() string {
	return string(m)
}

func (m MeasurementType) isValid() bool {
	return m == ExpectationMeasurement || m == ProbabilityMeasurement
}

func ParseMeasurementType(measurementTypeStr string) (MeasurementType, error) {
	if len(measurementTypeStr) > maxMeasurementTypeLength {
		return "", ErrParseMeasurementType
	}
	switch strings.ToLower(strings.TrimSpace(measurementTypeStr)) {
	case "expectation":
		return ExpectationMeasurement, nil
	case "probability":
		return ProbabilityMeasurement, nil
	default:
		return "", ErrParseMeasurementType
	}
}
