package user

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Name struct {
	value string
}

const (
	maxNameLength = 50
	minNameLength = 3
)

func NewName(value string) (Name, error) {
	value = strings.TrimSpace(value)
	if err := validateName(value); err != nil {
		return Name{}, err
	}
	return Name{value: value}, nil
}

func (n Name) String() string {
	return n.value
}

func validateName(name string) error {
	runeCount := utf8.RuneCountInString(name)
	if runeCount < minNameLength || runeCount > maxNameLength {
		return &InvalidNameError{name, fmt.Sprintf("name must be between %d and %d characters", minNameLength, maxNameLength)}
	}
	for _, r := range name {
		if !validateNameRune(r) {
			return &InvalidNameError{name, fmt.Sprintf("name contains invalid character: %c", r)}
		}
	}

	return nil
}

func validateNameRune(r rune) bool {
	if unicode.IsLetter(r) {
		return true
	}
	if r == '-' || r == '\'' || r == ' ' {
		return true
	}
	return false
}
