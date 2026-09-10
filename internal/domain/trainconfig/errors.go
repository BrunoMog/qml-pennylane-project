package trainconfig

import (
	"fmt"
)

type InvalidNameError struct {
	name string
}

func (e *InvalidNameError) Error() string {
	return fmt.Sprintf("invalid name: %s", e.name)
}

type InvalidDescriptionError struct {
	description string
}

func (e *InvalidDescriptionError) Error() string {
	return fmt.Sprintf("invalid description: %s", e.description)
}

type InvalidTrainingError struct{}

func (e *InvalidTrainingError) Error() string {
	return "invalid training"
}

type InvalidOwnerIDError struct{}

func (e *InvalidOwnerIDError) Error() string {
	return "invalid owner ID"
}
