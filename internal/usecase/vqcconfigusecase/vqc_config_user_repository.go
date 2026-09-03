package vqcconfigusecase

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	ExistByID(id uuid.UUID) (bool, error)
}
