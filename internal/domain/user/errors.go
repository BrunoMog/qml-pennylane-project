package user

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyName        = errors.New("name cannot be empty")
	ErrEmptyEmail       = errors.New("email cannot be empty")
	ErrInvalidRole      = errors.New("role is not valid")
	ErrNilID            = errors.New("id cannot be nil")
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
	if target == ErrInvalidName {
		return true
	}

	t, ok := target.(*InvalidNameError)
	if !ok {
		return false
	}

	return (e.Reason == t.Reason || t.Reason == "")
}

type InvalidEmailError struct {
	Reason string
}

func (e *InvalidEmailError) Error() string {
	return fmt.Sprintf("invalid email: %s", e.Reason)
}

func (e *InvalidEmailError) Is(target error) bool {
	if target == ErrInvalidEmail {
		return true
	}

	t, ok := target.(*InvalidEmailError)
	if !ok {
		return false
	}

	return (e.Reason == t.Reason || t.Reason == "")
}
