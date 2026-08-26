package testkit

import (
	"pennylane_project_backend/internal/domain/user"

	"github.com/google/uuid"
)

type MockUserRepository struct {
	users map[uuid.UUID]*user.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[uuid.UUID]*user.User),
	}
}

func (r *MockUserRepository) Save(u *user.User) error {
	r.users[u.GetID()] = u
	return nil
}

func (r *MockUserRepository) FindByID(id uuid.UUID) (*user.User, error) {
	if u, ok := r.users[id]; ok {
		return u, nil
	}
	return nil, &ErrUserNotFound{Message: id.String()}
}

func (r *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	for _, u := range r.users {
		if u.GetEmail() == email {
			return u, nil
		}
	}
	return nil, &ErrUserNotFound{Message: email}
}

func (r *MockUserRepository) ExistsByEmail(email string) bool {
	for _, u := range r.users {
		if u.GetEmail() == email {
			return true
		}
	}
	return false
}
