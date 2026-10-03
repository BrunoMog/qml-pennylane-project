package vqcconfig

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidName        = errors.New("invalid name for VQCConfig")
	ErrInvalidDescription = errors.New("invalid description for VQCConfig")
	ErrInvalidOwnerID     = errors.New("owner ID cannot be nil")
	ErrZeroValueVQC       = errors.New("VQC zero value is invalid")
	ErrEmptyName          = errors.New("name cannot be empty")
	ErrNilVQCConfigID     = errors.New("VQC config ID cannot be nil")

	ErrVQCConfigNotFound = errors.New("VQC config not found")
)

type InvalidNameError struct {
	Reason string
}

func (e *InvalidNameError) Error() string {
	return fmt.Sprintf("invalid name for VQCConfig: %s", e.Reason)
}

func (e *InvalidNameError) Is(target error) bool {
	return target == ErrInvalidName
}

type InvalidDescriptionError struct {
	Reason string
}

func (e *InvalidDescriptionError) Error() string {
	return fmt.Sprintf("invalid description for VQCConfig: %s", e.Reason)
}

func (e *InvalidDescriptionError) Is(target error) bool {
	return target == ErrInvalidDescription
}
