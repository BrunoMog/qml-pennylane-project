package vqc

import "strings"

type EmbeddingType string

const (
	EmbeddingTypeAngle     EmbeddingType = "angle"
	EmbeddingTypeAmplitude EmbeddingType = "amplitude"
)

const maxEmbeddingTypeLength = 20

func (e EmbeddingType) String() string {
	return string(e)
}

func ParseEmbeddingType(embeddingTypeStr string) (EmbeddingType, error) {
	if len(embeddingTypeStr) > maxEmbeddingTypeLength {
		return "", ErrParseEmbeddingType
	}
	switch strings.ToLower(strings.TrimSpace(embeddingTypeStr)) {
	case "angle":
		return EmbeddingTypeAngle, nil
	case "amplitude":
		return EmbeddingTypeAmplitude, nil
	default:
		return "", ErrParseEmbeddingType
	}
}
