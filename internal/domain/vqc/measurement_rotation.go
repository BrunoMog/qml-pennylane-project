package vqc

import "strings"

type MeasurementRotation string

const (
	XMeasurementRotation MeasurementRotation = "x"
	YMeasurementRotation MeasurementRotation = "y"
	ZMeasurementRotation MeasurementRotation = "z"
)

func (m MeasurementRotation) Value() string {
	return string(m)
}

func isValidMeasurementRotation(measurementRotation MeasurementRotation) bool {
	switch measurementRotation {
	case XMeasurementRotation, YMeasurementRotation, ZMeasurementRotation:
		return true
	default:
		return false
	}
}

func ParseMeasurementRotation(rotationStr string) (MeasurementRotation, error) {
	switch strings.ToLower(rotationStr) {
	case "x":
		return XMeasurementRotation, nil
	case "y":
		return YMeasurementRotation, nil
	case "z":
		return ZMeasurementRotation, nil
	default:
		return "", &InvalidParseMeasurementRotationError{rotationStr}
	}
}
