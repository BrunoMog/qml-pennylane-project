package userusecase

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type DeleteUserInput struct {
	CallerID uuid.UUID
	TargetID uuid.UUID
}

func (s *UserService) DeleteUser(input DeleteUserInput) error {
	caller, err := s.repository.FindByID(input.CallerID)
	if err != nil {
		return err
	}

	var target *user.User
	if input.CallerID != input.TargetID {
		target, err = s.repository.FindByID(input.TargetID)
		if err != nil {
			return err
		}

	} else {
		target = caller
	}

	if !canDeleteUser(caller, target) {
		return &UnauthorizedError{name: caller.Name()}
	}

	err = s.repository.DeleteByID(input.TargetID)
	if err != nil {
		return err
	}

	return nil
}

func canDeleteUser(caller *user.User, target *user.User) bool {
	if caller.ID() == target.ID() && !caller.IsOwner() {
		return true
	}
	switch caller.Role() {
	case user.RoleOwner:
		return true
	case user.RoleAdmin:
		return target.Role() != user.RoleOwner
	default:
		return false
	}
}
