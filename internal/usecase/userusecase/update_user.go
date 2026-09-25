package userusecase

import (
	"errors"
	"fmt"
	"pennylane_project_backend/internal/domain/user"

	"uuid"
)

type UpdateUserInput struct {
	Name     *string
	Email    *string
	CallerID uuid.UUID
	TargetID uuid.UUID
}

func (s *UserService) UpdateUser(input UpdateUserInput) error {
	if input.CallerID == uuid.Nil() || input.TargetID == uuid.Nil() {
		return ErrNilID
	}

	if input.Name == nil && input.Email == nil {
		return ErrNoFieldsToUpdate
	}

	caller, err := s.repository.FindByID(input.CallerID)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return err
		}
		return fmt.Errorf("userusecase: find caller by ID: %w", err)
	}

	var userToUpdate *user.User
	if input.CallerID != input.TargetID {
		userToUpdate, err = s.repository.FindByID(input.TargetID)
		if err != nil {
			if errors.Is(err, user.ErrUserNotFound) {
				return err
			}
			return fmt.Errorf("userusecase: find target by ID: %w", err)
		}
	} else {
		userToUpdate = caller
	}

	reason, allowed := canUpdateUser(caller, userToUpdate)
	if !allowed {
		return &PermissionDeniedError{reason: reason, callerID: caller.ID(), action: fmt.Sprintf("update user %s", userToUpdate.ID())}
	}

	var needToSave bool

	if input.Name != nil {
		name, err := user.NewName(*input.Name)
		if err != nil {
			return err
		}
		if userToUpdate.Name() != name {
			userToUpdate.SetName(name)
			needToSave = true
		}
	}

	if input.Email != nil {
		email, err := user.NewEmail(*input.Email)
		if err != nil {
			return err
		}
		if userToUpdate.Email() != email {
			exists, err := s.repository.ExistsByEmail(email)
			if err != nil {
				return fmt.Errorf("userusecase: check if email exists: %w", err)
			}
			if exists {
				return ErrEmailAlreadyExists
			}
			userToUpdate.SetEmail(email)
			needToSave = true
		}
	}

	if needToSave {
		err := s.repository.Save(userToUpdate)
		if err != nil {
			return fmt.Errorf("userusecase: save user: %w", err)
		}
	}

	return nil
}

func canUpdateUser(caller, target *user.User) (string, bool) {
	if caller.ID() == target.ID() {
		return "", true
	}

	switch caller.Role() {
	case user.RoleOwner:
		return "", true
	default:
		return "only owners can update other users", false
	}
}
