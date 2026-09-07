package user

import (
	"fmt"
)

type InvalidRoleError struct {
	role Role
}

func (e *InvalidRoleError) Error() string {
	return fmt.Sprintf("invalid role: %s", e.role)
}

type PermissionDeniedError struct {
	name string
}

func (e *PermissionDeniedError) Error() string {
	return fmt.Sprintf("permission denied: user %s dont have permission to execut this action", e.name)
}

type InvalidNameError struct {
	Name   string
	Reason string
}

func (e *InvalidNameError) Error() string {
	return fmt.Sprintf("invalid name: %s - %s", e.Name, e.Reason)
}

type InvalidEmailError struct {
	Email  string
	Reason string
}

func (e *InvalidEmailError) Error() string {
	return fmt.Sprintf("invalid email: %s - %s", e.Email, e.Reason)
}
