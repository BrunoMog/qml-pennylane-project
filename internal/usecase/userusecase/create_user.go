package userusecase

import (
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
		return nil, err
	}
	if exists {
		return nil, &EmailAlreadyExistsError{email}
	}

	newUser := user.NewUser(name, email)

	err = s.repository.Save(newUser)
	if err != nil {
		return nil, err
	}

	output := &UserOutput{
		ID:    newUser.ID(),
		Name:  newUser.Name().Value(),
		Email: newUser.Email().Value(),
		Role:  newUser.Role().Value(),
	}

	return output, nil
}
