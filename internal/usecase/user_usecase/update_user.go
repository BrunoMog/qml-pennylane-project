package user_usecase

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type UpdateUserInput struct {
	CallerID uuid.UUID
	TargetID uuid.UUID
	Name     *string
	Email    *string
}

func (s *UserService) UpdateUser(input UpdateUserInput) error {
	caller, err := s.repository.FindByID(input.CallerID)
	if err != nil {
		return err
	}
	user, err := s.repository.FindByID(input.TargetID)
	if err != nil {
		return err
	}

	if !canUpdateUser(caller, user) {
		return &UnauthorizedError{caller.GetName()}
	}

	if input.Name != nil {
		if err := user.SetName(*input.Name); err != nil {
			return err
		}
	}

	if input.Email != nil {
		if err := user.SetEmail(*input.Email); err != nil {
			return err
		}
	}

	return s.repository.Save(user)
}

func canUpdateUser(caller, target *user.User) bool {
	if caller.GetID() == target.GetID() {
		return true
	}

	switch caller.GetRole() {
	case user.RoleOwner:
		return true
	default:
		return false
	}
}
