package user_usecase

import (
	"github.com/google/uuid"
)

type UpdateUserInput struct {
	UserID uuid.UUID
	Name   *string
	Email  *string
}

func (s *UserService) UpdateUser(input UpdateUserInput) error {
	user, err := s.repository.FindByID(input.UserID)
	if err != nil {
		return err
	}

	if input.Name != nil {
		if err := user.SetName(*input.Name); err != nil {
			return err
		}
	}

	if input.Email != nil {
		if err := user.SetEmail(*input.Email); err != nil {
			return err
		}
	}

	return s.repository.Save(user)
}
