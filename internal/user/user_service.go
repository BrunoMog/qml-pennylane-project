package user

import (
	"github.com/google/uuid"
)

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(name string) (*User, error) {
	user, err := NewUser(name)
	if err != nil {
		return nil, err
	}

	return user, s.repository.Save(user)
}

func (s *UserService) GetUserByID(id uuid.UUID) (*User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
