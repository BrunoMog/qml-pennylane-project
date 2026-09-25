package userusecase

import (
	"errors"
	"fmt"
	"pennylane_project_backend/internal/domain/user"
	"uuid"
)

type ChangeUserRoleInput struct {
	Role     string
	CallerID uuid.UUID
	TargetID uuid.UUID
}

func (s *UserService) ChangeUserRole(input ChangeUserRoleInput) error {
	if input.CallerID == uuid.Nil() || input.TargetID == uuid.Nil() {
		return ErrNilID
	}

	role, err := user.ParseRole(input.Role)
	if err != nil {
		return err
	}

	caller, err := s.repository.FindByID(input.CallerID)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return err
		}
		return fmt.Errorf("userusecase: find caller by ID: %w", err)
	}

	var target *user.User
	if input.CallerID != input.TargetID {
		target, err = s.repository.FindByID(input.TargetID)
		if err != nil {
			if errors.Is(err, user.ErrUserNotFound) {
				return err
			}
			return fmt.Errorf("userusecase: find target by ID: %w", err)
		}
	} else {
		target = caller
	}

	if reason, allowed := canAssignRole(caller.Role(), target.Role(), role); !allowed {
		return &PermissionDeniedError{reason: reason, callerID: caller.ID(), action: fmt.Sprintf("assign role %s to user %s", role, target.ID())}
	}

	if target.Role() == role {
		return nil
	}

	err = target.SetRole(role)
	if err != nil {
		return err
	}

	err = s.repository.Save(target)
	if err != nil {
		return fmt.Errorf("userusecase: save target user: %w", err)
	}

	return nil
}

func canAssignRole(callerRole, targetRole, newRole user.Role) (reason string, allowed bool) {
	if targetRole == user.RoleGuest {
		return "cannot change role of a guest user", false
	}

	switch callerRole {
	case user.RoleOwner:
		if newRole == user.RoleOwner || targetRole == user.RoleOwner {
			reason = "owner cannot assign owner role or change role of another owner"
			return reason, false
		}
		return "", true
	case user.RoleAdmin:
		if targetRole == user.RoleOwner || newRole == user.RoleOwner {
			reason = "admin cannot assign owner role or change role of an owner"
			return reason, false
		}
		return "", true
	default:
		return "only owners and admins have permission to assign roles", false
	}
}
