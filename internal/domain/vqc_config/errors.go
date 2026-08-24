package vqc_config

import "fmt"

type InvalidNameError struct {
	name string
}

func (e *InvalidNameError) Error() string {
	return fmt.Sprintf("Invalid name: %s", e.name)
}

type InvalidDescriptionError struct {
	description string
}

func (e *InvalidDescriptionError) Error() string {
	return fmt.Sprintf("Invalid description: %s", e.description)
}
