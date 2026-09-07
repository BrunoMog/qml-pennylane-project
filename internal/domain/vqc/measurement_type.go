package vqc

import "strings"

type MeasurementType string

const (
	ExpectationMeasurement MeasurementType = "expectation"
	ProbabilityMeasurement MeasurementType = "probability"
)

func (m MeasurementType) Value() string {
	return string(m)
}

func isValidMeasurementType(measurementType MeasurementType) bool {
	switch measurementType {
	case ExpectationMeasurement, ProbabilityMeasurement:
		return true
	default:
		return false
	}
}

func ParseMeasurementType(measurementTypeStr string) (MeasurementType, error) {
	switch strings.ToLower(measurementTypeStr) {
	case "expectation":
		return ExpectationMeasurement, nil
	case "probability":
		return ProbabilityMeasurement, nil
	default:
		return "", &InvalidParseMeasurementTypeError{measurementType: MeasurementType(measurementTypeStr)}
	}
}
