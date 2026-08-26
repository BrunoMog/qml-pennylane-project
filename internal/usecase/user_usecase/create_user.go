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
	exists := s.repository.ExistsByEmail(email)
	if exists {
		return nil, &EmailAlreadyExistsError{email}
	}

	user, err := user.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	err = s.repository.Save(user)
	if err != nil {
		return nil, err
	}

	output := &UserOutput{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
		Role:  user.GetRole(),
	}

	return output, nil
}
