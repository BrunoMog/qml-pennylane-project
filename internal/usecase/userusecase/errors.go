package userusecase

import (
	"fmt"
	"pennylane_project_backend/internal/domain/user"
)

type EmailAlreadyExistsError struct {
	Email user.Email
}

func (e *EmailAlreadyExistsError) Error() string {
	return fmt.Sprintf("email already exists: %s", e.Email.Value())
}

type UnauthorizedError struct {
	Name user.Name
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized: user %s is not authorized to perform this action", e.Name.Value())
}

type UserNotFoundError struct {
}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}

type NoFieldsToUpdateError struct {
}

func (e *NoFieldsToUpdateError) Error() string {
	return "no fields to update"
}

type EmailAlreadyUsedError struct {
	Email user.Email
}

func (e *EmailAlreadyUsedError) Error() string {
	return fmt.Sprintf("email already used: %s", e.Email.Value())
}
