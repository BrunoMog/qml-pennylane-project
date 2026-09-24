package user

import (
	"net/mail"
	"strings"
)

type Email struct {
	value string
}

const (
	maxEmailLength     = 254
	maxLocalPartLength = 64
)

func NewEmail(value string) (Email, error) {
	value = strings.TrimSpace(value)
	if err := validateEmail(value); err != nil {
		return Email{}, err
	}
	value = strings.ToLower(value)
	return Email{value: value}, nil
}

func validateEmail(email string) error {
	if email == "" {
		return &InvalidEmailError{"email cannot be empty"}
	}

	if len(email) > maxEmailLength {
		return &InvalidEmailError{"email exceeds maximum length"}
	}

	addr, err := mail.ParseAddress(email)
	if err != nil {
		return &InvalidEmailError{err.Error()}
	}

	localPart := addr.Address[:strings.Index(addr.Address, "@")]
	if len(localPart) > maxLocalPartLength {
		return &InvalidEmailError{"local part exceeds maximum length"}
	}

	if len(addr.Name) > 0 {
		return &InvalidEmailError{"must not contain display name"}
	}

	return nil
}

func (e Email) IsValid() bool {
	return len(e.value) > 0
}

func (e Email) String() string {
	return e.value
}
