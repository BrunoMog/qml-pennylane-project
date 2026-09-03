package vqcconfigusecase

import "fmt"

type InvalidNameError struct {
}

func (e *InvalidNameError) Error() string {
	return "Invalid name"
}

type VQCConfigNameAlreadyExistsError struct {
	Name string
}

func (e *VQCConfigNameAlreadyExistsError) Error() string {
	return fmt.Sprintf("VQC config with name '%s' already exists", e.Name)
}

type UserNotFoundError struct {
}

func (e *UserNotFoundError) Error() string {
	return "User not found"
}

type VQCConfigNotFoundError struct {
}

func (e *VQCConfigNotFoundError) Error() string {
	return "VQC config not found"
}

type UnauthorizedError struct {
}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized"
}

type InvalidInputError struct {
}

func (e *InvalidInputError) Error() string {
	return "Invalid input"
}

type InvalidDescriptionError struct {
}

func (e *InvalidDescriptionError) Error() string {
	return "Invalid description"
}
