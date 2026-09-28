package vqc

import "strings"

type EmbeddingRotation string

const (
	XRotation EmbeddingRotation = "x"
	YRotation EmbeddingRotation = "y"
	ZRotation EmbeddingRotation = "z"
)

const maxEmbeddingRotationLength = 10

func (e EmbeddingRotation) isValid() bool {
	switch e {
	case XRotation, YRotation, ZRotation:
		return true
	default:
		return false
	}
}

func (e EmbeddingRotation) String() string {
	return string(e)
}

func ParseEmbeddingRotation(rotationStr string) (EmbeddingRotation, error) {
	if len(rotationStr) > maxEmbeddingRotationLength {
		return "", ErrParseEmbeddingRotation
	}
	switch strings.ToLower(strings.TrimSpace(rotationStr)) {
	case "x":
		return XRotation, nil
	case "y":
		return YRotation, nil
	case "z":
		return ZRotation, nil
	default:
		return "", ErrParseEmbeddingRotation
	}
}
