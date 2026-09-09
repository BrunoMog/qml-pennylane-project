package experiment

import (
	"fmt"
)

type InvalidNameError struct {
	Name string
}

func (e *InvalidNameError) Error() string {
	return fmt.Sprintf("invalid name: %s", e.Name)
}

type InvalidDescriptionError struct {
	Description string
}

func (e *InvalidDescriptionError) Error() string {
	return fmt.Sprintf("invalid description: %s", e.Description)
}

type InvalidOwnerIDError struct {
}

func (e *InvalidOwnerIDError) Error() string {
	return "invalid owner ID"
}

type InvalidTrainConfigIDError struct {
}

func (e *InvalidTrainConfigIDError) Error() string {
	return "invalid train config ID"
}

type InvalidVQCConfigIDError struct {
}

func (e *InvalidVQCConfigIDError) Error() string {
	return "invalid VQC config ID"
}
