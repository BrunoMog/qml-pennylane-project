package user

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyName        = errors.New("user: name cannot be empty")
	ErrEmptyEmail       = errors.New("user: email cannot be empty")
	ErrInvalidRole      = errors.New("user: role is not valid")
	ErrNilID            = errors.New("user: id cannot be nil")
	ErrInvalidParseRole = errors.New("role: role cannot be parsed")
	ErrInvalidName      = errors.New("user: name is not valid")
	ErrInvalidEmail     = errors.New("user: email is not valid")
)

type InvalidNameError struct {
	Reason string
}

func (e *InvalidNameError) Error() string {
	return fmt.Sprintf("invalid name: %s", e.Reason)
}

func (e *InvalidNameError) Is(target error) bool {
	if target == ErrInvalidName {
		return true
	}

	_, ok := target.(*InvalidNameError)
	return ok
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
