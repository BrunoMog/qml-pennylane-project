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

func NewUser(name Name, email Email) *User {
	u := User{
		id:    uuid.New(),
		name:  name,
		email: email,
		role:  RoleUser,
	}
	return &u
}

func (u *User) SetRole(newRole Role) error {
	if !newRole.IsValidRole() {
		return &InvalidRoleError{newRole}
	}
	u.role = newRole
	return nil
}

func (u *User) SetName(newName Name) {
	u.name = newName
}

func (u *User) SetEmail(newEmail Email) {
	u.email = newEmail
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
