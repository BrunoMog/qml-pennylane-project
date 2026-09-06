package userusecase

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type UserOutput struct {
	Name  string
	Email string
	Role  user.Role
	ID    uuid.UUID
}

func (s *UserService) CreateUser(name string, email string) (*UserOutput, error) {
	exists, err := s.repository.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &EmailAlreadyExistsError{email}
	}

	newUser, err := user.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	err = s.repository.Save(newUser)
	if err != nil {
		return nil, err
	}

	output := &UserOutput{
		ID:    newUser.ID(),
		Name:  newUser.Name(),
		Email: newUser.Email(),
		Role:  newUser.Role(),
	}

	return output, nil
}
