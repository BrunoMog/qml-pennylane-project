package vqcconfig

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

const maxDescriptionLength = 500

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
	runeCount := utf8.RuneCountInString(description)
	if runeCount > maxDescriptionLength {
		return &InvalidDescriptionError{fmt.Sprintf("description must be at most %d characters", maxDescriptionLength)}
	}

	for _, r := range description {
		if !validateDescriptionRune(r) {
			return &InvalidDescriptionError{fmt.Sprintf("description contains invalid character: %c", r)}
		}
	}

	return nil
}

func validateDescriptionRune(r rune) bool {
	if unicode.IsPrint(r) {
		return true
	}
	if r == '\n' || r == '\r' || r == '\t' {
		return true
	}

	return false
}

func (d Description) String() string {
	return d.value
}
