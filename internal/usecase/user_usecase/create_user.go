package user_usecase

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type UserOutput struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Role  user.Role `json:"role"`
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
		ID:    newUser.GetID(),
		Name:  newUser.GetName(),
		Email: newUser.GetEmail(),
		Role:  newUser.GetRole(),
	}

	return output, nil
}
