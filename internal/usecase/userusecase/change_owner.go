package userusecase

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type ChangeOwnerInput struct {
	CallerID uuid.UUID
	TargetID uuid.UUID
}

func (s *UserService) ChangeOwner(input ChangeOwnerInput) error {
	caller, err := s.repository.FindByID(input.CallerID)
	if err != nil {
		return err
	}

	target, err := s.repository.FindByID(input.TargetID)
	if err != nil {
		return err
	}

	if !canChangeOwner(caller, target) {
		return &UnauthorizedError{name: caller.GetName()}
	}

	err = s.repository.ChangeOwner(input.CallerID, input.TargetID)
	if err != nil {
		return err
	}

	return nil
}

func canChangeOwner(caller *user.User, target *user.User) bool {
	if caller.GetID() == target.GetID() {
		return false
	}
	if caller.GetRole() != user.RoleOwner {
		return false
	}
	return true
}
