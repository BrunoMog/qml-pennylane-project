package vqcconfig

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

type VQCConfigMissingVQCError struct{}

func (e *VQCConfigMissingVQCError) Error() string {
	return "VQCConfig is missing VQC"
}
