package user

import (
	"fmt"
	"strings"
)

type Name struct {
	value string
}

const (
	MAX_NAME_LENGTH = 50
	MIN_NAME_LENGTH = 3
)

func NewName(value string) (Name, error) {
	value = strings.TrimSpace(value)
	if err := validateName(value); err != nil {
		return Name{}, err
	}
	return Name{value: value}, nil
}

func (n Name) Value() string {
	return n.value
}

func validateName(name string) error {
	if name == "" {
		return &InvalidNameError{name, "name cannot be empty"}
	}
	if len(name) < MIN_NAME_LENGTH || len(name) > MAX_NAME_LENGTH {
		return &InvalidNameError{name, fmt.Sprintf("name must be between %d and %d characters", MIN_NAME_LENGTH, MAX_NAME_LENGTH)}
	}
	return nil
}
