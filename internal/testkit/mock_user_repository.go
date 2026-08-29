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
		copiedUser := *u
		return &copiedUser, nil
	}
	return nil, &ErrUserNotFound{Message: id.String()}
}

func (r *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	for _, u := range r.users {
		if u.GetEmail() == email {
			copiedUser := *u
			return &copiedUser, nil
		}
	}
	return nil, &ErrUserNotFound{Message: email}
}

func (r *MockUserRepository) ExistsByEmail(email string) (bool, error) {
	for _, u := range r.users {
		if u.GetEmail() == email {
			return true, nil
		}
	}
	return false, nil
}

func (r *MockUserRepository) ExistByID(id uuid.UUID) (bool, error) {
	if _, ok := r.users[id]; ok {
		return true, nil
	}
	return false, nil
}

func (r *MockUserRepository) DeleteByID(id uuid.UUID) error {
	if _, ok := r.users[id]; ok {
		delete(r.users, id)
		return nil
	}
	return &ErrUserNotFound{Message: id.String()}
}

func (r *MockUserRepository) ChangeOwner(callerID uuid.UUID, targetID uuid.UUID) error {
	caller, ok := r.users[callerID]
	if !ok {
		return &ErrUserNotFound{Message: callerID.String()}
	}

	target, ok := r.users[targetID]
	if !ok {
		return &ErrUserNotFound{Message: targetID.String()}
	}

	caller.SetRole(user.RoleAdmin)
	target.SetRole(user.RoleOwner)

	return nil
}
