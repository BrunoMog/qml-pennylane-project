package userusecase

import (
	"errors"
	"fmt"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/usecase/apperrors"

	"uuid"
)

type DeleteUserInput struct {
	CallerID uuid.UUID
	TargetID uuid.UUID
}

func (s *UserService) DeleteUser(input DeleteUserInput) error {
	if input.CallerID == uuid.Nil() || input.TargetID == uuid.Nil() {
		return user.ErrNilUserID
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

	reason, allowed := canDeleteUser(caller, target)
	if !allowed {
		return apperrors.NewPermissionDeniedError(reason, fmt.Sprintf("delete user %s", target.ID()), caller.ID())
	}

	err = s.repository.DeleteByID(input.TargetID)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return err
		}
		return fmt.Errorf("userusecase: delete user by ID: %w", err)
	}

	return nil
}

func canDeleteUser(caller *user.User, target *user.User) (string, bool) {
	if caller.ID() == target.ID() && caller.IsOwner() {
		return "owner cannot delete themselves", false
	} else if caller.ID() == target.ID() {
		return "", true
	}
	switch caller.Role() {
	case user.RoleOwner:
		return "", true
	case user.RoleAdmin:
		if target.Role() == user.RoleOwner {
			return "admin cannot delete owner", false
		}
		return "", true
	default:
		return "only owners and admins can delete users", false
	}
}
