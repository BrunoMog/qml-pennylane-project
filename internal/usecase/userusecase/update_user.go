package userusecase

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type UpdateUserInput struct {
	Name     *string
	Email    *string
	CallerID uuid.UUID
	TargetID uuid.UUID
}

func (s *UserService) UpdateUser(input UpdateUserInput) error {
	if input.Name == nil && input.Email == nil {
		return &NoFieldsToUpdateError{}
	}

	caller, err := s.repository.FindByID(input.CallerID)
	if err != nil {
		return err
	}
	userToUpdate, err := s.repository.FindByID(input.TargetID)
	if err != nil {
		return err
	}

	if !canUpdateUser(caller, userToUpdate) {
		return &UnauthorizedError{caller.Name()}
	}

	if input.Name != nil {
		name, err := user.NewName(*input.Name)
		if err != nil {
			return err
		}
		userToUpdate.SetName(name)
	}

	if input.Email != nil {
		email, err := user.NewEmail(*input.Email)
		if err != nil {
			return err
		}
		exists, err := s.repository.ExistsByEmail(email)
		if err != nil {
			return err
		}
		if exists {
			return &EmailAlreadyExistsError{email}
		}
		userToUpdate.SetEmail(email)
	}

	return s.repository.Save(userToUpdate)
}

func canUpdateUser(caller, target *user.User) bool {
	if caller.ID() == target.ID() {
		return true
	}

	switch caller.Role() {
	case user.RoleOwner:
		return true
	default:
		return false
	}
}
