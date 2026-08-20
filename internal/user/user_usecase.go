package user

import (
	"github.com/google/uuid"
)

type UserUseCase interface {
	CreateUser(name string) (*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	SetUserRole(user *User, newRole Role) error
}
