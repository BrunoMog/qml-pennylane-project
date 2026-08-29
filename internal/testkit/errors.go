package testkit

import "fmt"

type ErrUserNotFound struct {
	Message string
}

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("User not found: %s", e.Message)
}

type ErrVQCConfigNotFound struct {
}

func (e ErrVQCConfigNotFound) Error() string {
	return "VQC config not found"
}
