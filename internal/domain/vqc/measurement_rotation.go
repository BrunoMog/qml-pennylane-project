package vqc

import "strings"

type MeasurementRotation string

const (
	XMeasurementRotation MeasurementRotation = "x"
	YMeasurementRotation MeasurementRotation = "y"
	ZMeasurementRotation MeasurementRotation = "z"
)

const maxMeasurementRotationLength = 10

func (m MeasurementRotation) String() string {
	return string(m)
}

func (m MeasurementRotation) isValid() bool {
	switch m {
	case XMeasurementRotation, YMeasurementRotation, ZMeasurementRotation:
		return true
	default:
		return false
	}
}

func ParseMeasurementRotation(rotationStr string) (MeasurementRotation, error) {
	if len(rotationStr) > maxMeasurementRotationLength {
		return "", ErrParseMeasurementRotation
	}
	switch strings.ToLower(strings.TrimSpace(rotationStr)) {
	case "x":
		return XMeasurementRotation, nil
	case "y":
		return YMeasurementRotation, nil
	case "z":
		return ZMeasurementRotation, nil
	default:
		return "", ErrParseMeasurementRotation
	}
}
