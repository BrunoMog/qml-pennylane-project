package user

import (
	"net/mail"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	if err := validateEmail(value); err != nil {
		return Email{}, err
	}
	return Email{value: value}, nil
}

func validateEmail(email string) error {
	if email == "" {
		return &InvalidEmailError{email, "email cannot be empty"}
	}
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return &InvalidEmailError{email, "invalid email format"}
	}
	if addr.Address != email {
		return &InvalidEmailError{email, "must not contain display name"}
	}

	return nil
}

func (e Email) Value() string {
	return e.value
}
