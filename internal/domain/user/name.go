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

func validateName(name string) error {
	runeCount := utf8.RuneCountInString(name)
	if runeCount < minNameLength || runeCount > maxNameLength {
		return &InvalidNameError{fmt.Sprintf("name must be between %d and %d characters", minNameLength, maxNameLength)}
	}
	var hasLetter bool
	for _, r := range name {
		if !validateNameRune(r) {
			return &InvalidNameError{fmt.Sprintf("name contains invalid character: %c", r)}
		}

		if !hasLetter && unicode.IsLetter(r) {
			hasLetter = true
		}
	}
	if !hasLetter {
		return &InvalidNameError{"name must contain at least one letter"}
	}

	return nil
}

func validateNameRune(r rune) bool {
	if unicode.IsLetter(r) {
		return true
	}
	if r == '-' || r == '\'' || r == ' ' || r == '.' {
		return true
	}
	return false
}

func (n Name) IsValid() bool {
	return len(n.value) > 0
}

func (n Name) String() string {
	return n.value
}
