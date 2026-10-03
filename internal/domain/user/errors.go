package user

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyName        = errors.New("name cannot be empty")
	ErrEmptyEmail       = errors.New("email cannot be empty")
	ErrInvalidRole      = errors.New("role is not valid")
	ErrNilUserID        = errors.New("user ID cannot be nil")
	ErrInvalidParseRole = errors.New("role cannot be parsed")
	ErrInvalidName      = errors.New("name is not valid")
	ErrInvalidEmail     = errors.New("email is not valid")

	ErrUserNotFound = errors.New("user not found")
)

type InvalidNameError struct {
	Reason string
}

func (e *InvalidNameError) Error() string {
	return fmt.Sprintf("invalid name: %s", e.Reason)
}

func (e *InvalidNameError) Is(target error) bool {
	return target == ErrInvalidName
}

type InvalidEmailError struct {
	Reason string
}

func (e *InvalidEmailError) Error() string {
	return fmt.Sprintf("invalid email: %s", e.Reason)
}

func (e *InvalidEmailError) Is(target error) bool {
	return target == ErrInvalidEmail
}
