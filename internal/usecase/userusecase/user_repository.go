package userusecase

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(u *user.User) error
	FindByID(id uuid.UUID) (*user.User, error)
	FindByEmail(email user.Email) (*user.User, error)
	DeleteByID(id uuid.UUID) error
	ExistsByEmail(email user.Email) (bool, error)
	ChangeOwner(callerID uuid.UUID, targetID uuid.UUID) error
}
