package vqcconfig

import "strings"

const MAX_DESCRIPTION_LENGTH = 500

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
	if len(description) > MAX_DESCRIPTION_LENGTH {
		return &InvalidDescriptionError{description}
	}
	return nil
}

func (d Description) Value() string {
	return d.value
}
