package testkit

import (
	"pennylane_project_backend/internal/domain/user"

	"uuid"
)

type MockUserRepository struct {
	users map[uuid.UUID]*user.User

	ErrSave          error
	ErrExistsByID    error
	ErrFindByEmail   error
	ErrFindByID      error
	ErrExistsByEmail error
	ErrDeleteByID    error
	ErrChangeOwner   error
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[uuid.UUID]*user.User),
	}
}

func (r *MockUserRepository) Save(u *user.User) error {
	if r.ErrSave != nil {
		return r.ErrSave
	}

	r.users[u.ID()] = u
	return nil
}

func (r *MockUserRepository) FindByID(id uuid.UUID) (*user.User, error) {
	if r.ErrFindByID != nil {
		return nil, r.ErrFindByID
	}

	if u, ok := r.users[id]; ok {
		copiedUser := *u
		return &copiedUser, nil
	}
	return nil, user.ErrUserNotFound
}

func (r *MockUserRepository) FindByEmail(email user.Email) (*user.User, error) {
	if r.ErrFindByEmail != nil {
		return nil, r.ErrFindByEmail
	}

	for _, u := range r.users {
		if u.Email() == email {
			copiedUser := *u
			return &copiedUser, nil
		}
	}
	return nil, user.ErrUserNotFound
}

func (r *MockUserRepository) ExistsByID(id uuid.UUID) (bool, error) {
	if r.ErrExistsByID != nil {
		return false, r.ErrExistsByID
	}

	if _, ok := r.users[id]; ok {
		return true, nil
	}
	return false, nil
}

func (r *MockUserRepository) ExistsByEmail(email user.Email) (bool, error) {
	if r.ErrExistsByEmail != nil {
		return false, r.ErrExistsByEmail
	}

	for _, u := range r.users {
		if u.Email() == email {
			return true, nil
		}
	}
	return false, nil
}

func (r *MockUserRepository) DeleteByID(id uuid.UUID) error {
	if r.ErrDeleteByID != nil {
		return r.ErrDeleteByID
	}

	if _, ok := r.users[id]; ok {
		delete(r.users, id)
		return nil
	}
	return user.ErrUserNotFound
}

func (r *MockUserRepository) ChangeOwner(callerID uuid.UUID, targetID uuid.UUID) error {
	if r.ErrChangeOwner != nil {
		return r.ErrChangeOwner
	}

	caller, ok := r.users[callerID]
	if !ok {
		return user.ErrUserNotFound
	}

	target, ok := r.users[targetID]
	if !ok {
		return user.ErrUserNotFound
	}

	caller.SetRole(user.RoleAdmin)
	target.SetRole(user.RoleOwner)

	return nil
}
