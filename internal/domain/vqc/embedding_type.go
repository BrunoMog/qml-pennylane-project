package vqc

import "strings"

type EmbeddingType string

const (
	EmbeddingTypeAngle     EmbeddingType = "angle"
	EmbeddingTypeAmplitude EmbeddingType = "amplitude"
)

func (e EmbeddingType) Value() string {
	return string(e)
}

func ParseEmbeddingType(embeddingTypeStr string) (EmbeddingType, error) {
	switch strings.ToLower(strings.TrimSpace(embeddingTypeStr)) {
	case "angle":
		return EmbeddingTypeAngle, nil
	case "amplitude":
		return EmbeddingTypeAmplitude, nil
	default:
		return "", &InvalidParseEmbeddingError{embeddingTypeStr}
	}
}
