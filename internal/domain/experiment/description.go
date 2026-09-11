package experiment

import "strings"

const (
	maxDescriptionLength = 500
)

type Description struct {
	value string
}

func NewDescription(value string) (Description, error) {
	value = strings.TrimSpace(value)
	if err := validateDescription(value); err != nil {
		return Description{}, err
	}
	return Description{value: value}, nil
}

func validateDescription(description string) error {
	if len(description) > maxDescriptionLength {
		return &InvalidDescriptionError{description}
	}
	return nil
}

func (d Description) Value() string {
	return d.value
}
