package userusecase

import (
	"errors"
	"fmt"
	"pennylane_project_backend/internal/domain/user"

	"uuid"
)

type ChangeOwnerInput struct {
	CallerID uuid.UUID
	TargetID uuid.UUID
}

func (s *UserService) ChangeOwner(input ChangeOwnerInput) error {
	if input.CallerID == uuid.Nil() || input.TargetID == uuid.Nil() {
		return ErrNilID
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

	if reason, allowed := canChangeOwner(caller, target); !allowed {
		return &PermissionDeniedError{reason: reason, callerID: caller.ID(), action: "change owner"}
	}

	if caller.ID() == target.ID() {
		return nil
	}

	err = s.repository.ChangeOwner(input.CallerID, input.TargetID)
	if err != nil {
		return fmt.Errorf("userusecase: change owner: %w", err)
	}

	return nil
}

func canChangeOwner(caller *user.User, target *user.User) (string, bool) {
	if caller.Role() != user.RoleOwner {
		return "Caller is not an owner", false
	}
	return "", true
}
