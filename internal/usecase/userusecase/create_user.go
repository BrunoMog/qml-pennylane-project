package userusecase

import (
	"fmt"
	"pennylane_project_backend/internal/domain/user"

	"uuid"
)

type UserOutput struct {
	Name  string
	Email string
	Role  string
	ID    uuid.UUID
}

type CreateUserInput struct {
	Name  string
	Email string
}

func (s *UserService) CreateUser(input CreateUserInput) (*UserOutput, error) {
	name, err := user.NewName(input.Name)
	if err != nil {
		return nil, err
	}

	email, err := user.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	exists, err := s.repository.ExistsByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("userusecase: check email existence: %w", err)
	}
	if exists {
		return nil, ErrEmailAlreadyExists
	}

	newUser, err := user.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	err = s.repository.Save(newUser)
	if err != nil {
		return nil, fmt.Errorf("userusecase: save user: %w", err)
	}

	output := &UserOutput{
		ID:    newUser.ID(),
		Name:  newUser.Name().String(),
		Email: newUser.Email().String(),
		Role:  newUser.Role().String(),
	}

	return output, nil
}
