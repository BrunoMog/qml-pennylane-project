package vqcconfigusecase

import (
	"errors"
	"fmt"
)

var (
	ErrNoFieldsToUpdate           = errors.New("no fields to update")
	ErrVQCConfigNameAlreadyExists = errors.New("VQC config name already exists")
)

type UnreachableEmbeddingTypeError struct {
	EmbeddingType string
}

func (e *UnreachableEmbeddingTypeError) Error() string {
	return fmt.Sprintf("Unreachable embedding type: %s", e.EmbeddingType)
}
