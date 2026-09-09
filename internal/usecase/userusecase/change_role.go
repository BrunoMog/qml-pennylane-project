package userusecase

import (
	"pennylane_project_backend/internal/domain/user"

	"uuid"
)

type ChangeUserRoleInput struct {
	Role     string
	CallerID uuid.UUID
	TargetID uuid.UUID
}

func (s *UserService) ChangeUserRole(input ChangeUserRoleInput) error {
	role, err := user.ParseRole(input.Role)
	if err != nil {
		return err
	}
	caller, err := s.repository.FindByID(input.CallerID)
	if err != nil {
		return err
	}
	target, err := s.repository.FindByID(input.TargetID)
	if err != nil {
		return err
	}

	if !canAssignRole(caller.Role(), target.Role(), role) {
		return &UnauthorizedError{caller.Name()}
	}

	err = target.SetRole(role)
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
		if newRole == user.RoleOwner || targetRole == user.RoleOwner {
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
