package testkit

import (
	"pennylane_project_backend/internal/domain/user"

	"uuid"
)

type MockUserRepository struct {
	users map[uuid.UUID]*user.User

	ExistsByEmailErr error
	ExistsByIDErr    error
	SaveErr          error
	FindByIDErr      error
	FindByEmailErr   error
	DeleteByIDErr    error
	ChangeOwnerErr   error
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[uuid.UUID]*user.User),
	}
}

func (r *MockUserRepository) Save(u *user.User) error {
	if r.SaveErr != nil {
		return r.SaveErr
	}

	r.users[u.ID()] = u
	return nil
}

func (r *MockUserRepository) FindByID(id uuid.UUID) (*user.User, error) {
	if r.FindByIDErr != nil {
		return nil, r.FindByIDErr
	}

	if u, ok := r.users[id]; ok {
		copiedUser := *u
		return &copiedUser, nil
	}
	return nil, user.ErrUserNotFound
}

func (r *MockUserRepository) FindByEmail(email user.Email) (*user.User, error) {
	if r.FindByEmailErr != nil {
		return nil, r.FindByEmailErr
	}

	for _, u := range r.users {
		if u.Email() == email {
			copiedUser := *u
			return &copiedUser, nil
		}
	}
	return nil, user.ErrUserNotFound
}

func (r *MockUserRepository) ExistsByEmail(email user.Email) (bool, error) {
	if r.ExistsByEmailErr != nil {
		return false, r.ExistsByEmailErr
	}

	for _, u := range r.users {
		if u.Email() == email {
			return true, nil
		}
	}
	return false, nil
}

func (r *MockUserRepository) ExistsByID(id uuid.UUID) (bool, error) {
	if r.ExistsByIDErr != nil {
		return false, r.ExistsByIDErr
	}

	if _, ok := r.users[id]; ok {
		return true, nil
	}
	return false, nil
}

func (r *MockUserRepository) DeleteByID(id uuid.UUID) error {
	if r.DeleteByIDErr != nil {
		return r.DeleteByIDErr
	}

	if _, ok := r.users[id]; ok {
		delete(r.users, id)
		return nil
	}
	return user.ErrUserNotFound
}

func (r *MockUserRepository) ChangeOwner(callerID uuid.UUID, targetID uuid.UUID) error {
	if r.ChangeOwnerErr != nil {
		return r.ChangeOwnerErr
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
