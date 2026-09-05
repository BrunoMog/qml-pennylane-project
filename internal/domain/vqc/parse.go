package vqc

import (
	"strings"
)

func ParseEmbeddingType(embeddingTypeStr string) (EmbeddingType, error) {
	switch strings.ToLower(embeddingTypeStr) {
	case "angle":
		return EmbeddingTypeAngle, nil
	case "amplitude":
		return EmbeddingTypeAmplitude, nil
	default:
		return "", &InvalidParseEmbeddingError{embeddingTypeStr}
	}
}

func ParseEmbeddingRotation(rotationStr string) (EmbeddingRotation, error) {
	switch strings.ToLower(rotationStr) {
	case "x":
		return XRotation, nil
	case "y":
		return YRotation, nil
	case "z":
		return ZRotation, nil
	default:
		return "", &InvalidParseEmbeddingRotationError{rotationStr}
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

func ParseGateType(gateTypeStr string) (GateType, error) {
	switch strings.ToLower(gateTypeStr) {
	case "h":
		return HGate, nil
	case "x":
		return XGate, nil
	case "y":
		return YGate, nil
	case "z":
		return ZGate, nil
	case "rx":
		return RXGate, nil
	case "ry":
		return RYGate, nil
	case "rz":
		return RZGate, nil
	case "cnot":
		return CNOTGate, nil
	default:
		return "", &InvalidParseGateTypeError{gateTypeStr}
	}
}
