package user

import (
	"github.com/google/uuid"
)

type User struct {
	id   uuid.UUID
	name string
	role Role
}

func NewUser(name string) (*User, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	u := User{
		id:   uuid.New(),
		name: name,
		role: RoleUser,
	}

	return &u, nil
}

func validateName(name string) error {
	if name == "" {
		return &InvalidNameError{name}
	}
	if len(name) > 20 {
		return &InvalidNameError{name}
	}
	return nil
}

func (u *User) SetRole(targetUser *User, newRole Role) error {
	if !u.IsAdmin() && !u.IsOwner() {
		return &PermissionDeniedError{u.name}
	}
	if targetUser == nil {
		return &PermissionDeniedError{u.name}
	}
	if targetUser.IsOwner() {
		return &PermissionDeniedError{u.name}
	}
	if !IsValidRole(newRole) {
		return &InvalidRoleError{newRole}
	}
	if newRole == RoleOwner && !u.IsOwner() {
		return &PermissionDeniedError{u.name}
	}

	targetUser.role = newRole
	return nil
}

func (u *User) IsAdmin() bool {
	return u.role == RoleAdmin
}

func (u *User) IsOwner() bool {
	return u.role == RoleOwner
}

func (u *User) GetID() uuid.UUID {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetRole() Role {
	return u.role
}
