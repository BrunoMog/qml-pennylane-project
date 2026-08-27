package user_usecase

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type ChangeUserRoleInput struct {
	CallerID uuid.UUID
	TargetID uuid.UUID
	Role     user.Role
}

func (s *UserService) ChangeUserRole(input ChangeUserRoleInput) error {
	caller, err := s.repository.FindByID(input.CallerID)
	if err != nil {
		return err
	}
	err = validateCallerRole(caller)
	if err != nil {
		return err
	}

	target, err := s.repository.FindByID(input.TargetID)
	if err != nil {
		return err
	}
	err = validateTargetRole(target)
	if err != nil {
		return err
	}
	err = target.SetRole(input.Role)
	if err != nil {
		return err
	}

	err = s.repository.Save(target)
	if err != nil {
		return err
	}

	return nil
}

func validateCallerRole(caller *user.User) error {
	if caller.GetRole() != user.RoleAdmin && caller.GetRole() != user.RoleOwner {
		return &UnauthorizedError{caller.GetName()}
	}
	return nil
}

func validateTargetRole(target *user.User) error {
	if target.GetRole() == user.RoleOwner {
		return &UnauthorizedError{target.GetName()}
	}
	return nil
}
