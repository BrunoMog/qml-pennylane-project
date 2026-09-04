package userusecase

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type ChangeUserRoleInput struct {
	Role     user.Role
	CallerID uuid.UUID
	TargetID uuid.UUID
}

func (s *UserService) ChangeUserRole(input ChangeUserRoleInput) error {
	caller, err := s.repository.FindByID(input.CallerID)
	if err != nil {
		return err
	}
	target, err := s.repository.FindByID(input.TargetID)
	if err != nil {
		return err
	}

	if !canAssignRole(caller.Role(), target.Role(), input.Role) {
		return &UnauthorizedError{caller.Name()}
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

func canAssignRole(callerRole, targetRole, newRole user.Role) bool {
	switch callerRole {
	case user.RoleOwner:
		if newRole == user.RoleOwner {
			return false
		}
		return true
	case user.RoleAdmin:
		if targetRole == user.RoleOwner || newRole == user.RoleOwner {
			return false
		}
		return true
	default:
		return false
	}
}
