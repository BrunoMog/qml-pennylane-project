package user

import (
	"uuid"
)

type User struct {
	email Email
	name  Name
	role  Role
	id    uuid.UUID
}

func NewUser(name Name, email Email) (*User, error) {
	if !name.IsValid() {
		return nil, ErrEmptyName
	}
	if !email.IsValid() {
		return nil, ErrEmptyEmail
	}

	u := User{
		id:    uuid.New(),
		name:  name,
		email: email,
		role:  RoleUser,
	}
	return &u, nil
}

func RestoreUser(id uuid.UUID, name Name, email Email, role Role) (*User, error) {
	if id == uuid.Nil() {
		return nil, ErrNilID
	}
	if !name.IsValid() {
		return nil, ErrEmptyName
	}
	if !email.IsValid() {
		return nil, ErrEmptyEmail
	}
	if !role.IsValidRole() {
		return nil, ErrInvalidRole
	}

	u := User{
		id:    id,
		name:  name,
		email: email,
		role:  role,
	}
	return &u, nil
}

func (u *User) SetRole(newRole Role) error {
	if !newRole.IsValidRole() {
		return ErrInvalidRole
	}
	u.role = newRole
	return nil
}

func (u *User) SetName(newName Name) error {
	if !newName.IsValid() {
		return ErrEmptyName
	}
	u.name = newName
	return nil
}

func (u *User) SetEmail(newEmail Email) error {
	if !newEmail.IsValid() {
		return ErrEmptyEmail
	}
	u.email = newEmail
	return nil
}

func (u *User) IsAdmin() bool {
	return u.role == RoleAdmin
}

func (u *User) IsOwner() bool {
	return u.role == RoleOwner
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Name() Name {
	return u.name
}

func (u *User) Role() Role {
	return u.role
}

func (u *User) Email() Email {
	return u.email
}
