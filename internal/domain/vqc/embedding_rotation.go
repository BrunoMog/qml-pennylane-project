package vqc

import "strings"

type EmbeddingRotation string

const (
	XRotation EmbeddingRotation = "x"
	YRotation EmbeddingRotation = "y"
	ZRotation EmbeddingRotation = "z"
)

func (e EmbeddingRotation) IsValid() bool {
	switch e {
	case XRotation, YRotation, ZRotation:
		return true
	default:
		return false
	}
}

func (e EmbeddingRotation) Value() string {
	return string(e)
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
