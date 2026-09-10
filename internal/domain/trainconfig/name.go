package trainconfig

import "strings"

type Name struct {
	value string
}

const (
	maxNameLength = 100
	minNameLength = 4
)

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
	if len(name) < minNameLength || len(name) > maxNameLength {
		return &InvalidNameError{name}
	}
	return nil
}

func (n Name) Equals(other Name) bool {
	return strings.EqualFold(n.value, other.value)
}

func (n Name) Value() string {
	return n.value
}
