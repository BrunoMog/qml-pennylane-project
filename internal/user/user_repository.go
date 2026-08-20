package user

import (
	"errors"

	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	Save(u *User) error
	FindByID(id uuid.UUID) (*User, error)
}
