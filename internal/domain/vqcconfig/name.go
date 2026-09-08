package vqcconfig

import "strings"

const (
	MAX_NAME_LENGTH = 100
	MIN_NAME_LENGTH = 4
)

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	value = strings.TrimSpace(value)
	if err := validateName(value); err != nil {
		return Name{}, err
	}
	return Name{value: value}, nil
}

func validateName(name string) error {
	if name == "" {
		return &InvalidNameError{name}
	}
	if len(name) < MIN_NAME_LENGTH || len(name) > MAX_NAME_LENGTH {
		return &InvalidNameError{name}
	}
	return nil
}

func (n Name) Value() string {
	return n.value
}
