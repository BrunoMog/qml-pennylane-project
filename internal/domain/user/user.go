package user

import (
	"github.com/google/uuid"
)

type User struct {
	email string
	name  string
	role  Role
	id    uuid.UUID
}

func NewUser(name string, email string) (*User, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	err = validateEmail(email)
	if err != nil {
		return nil, err
	}

	u := User{
		id:    uuid.New(),
		name:  name,
		email: email,
		role:  RoleUser,
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

func validateEmail(email string) error {
	if email == "" {
		return &InvalidEmailError{email}
	}
	// TODO: ajustar validação de email posterioremnte
	return nil
}

func (u *User) SetRole(newRole Role) error {
	if !newRole.IsValidRole() {
		return &InvalidRoleError{newRole}
	}
	u.role = newRole
	return nil
}

func (u *User) SetName(newName string) error {
	err := validateName(newName)
	if err != nil {
		return err
	}
	u.name = newName
	return nil
}

func (u *User) SetEmail(newEmail string) error {
	err := validateEmail(newEmail)
	if err != nil {
		return err
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

func (u *User) Name() string {
	return u.name
}

func (u *User) Role() Role {
	return u.role
}

func (u *User) Email() string {
	return u.email
}
