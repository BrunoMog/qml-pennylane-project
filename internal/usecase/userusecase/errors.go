package userusecase

import (
	"errors"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrNoFieldsToUpdate   = errors.New("no fields to update")
)
