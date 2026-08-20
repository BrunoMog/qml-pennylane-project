package user

import (
	"github.com/google/uuid"
)

type UserController struct {
	useCase UserUseCase
}

func NewUserController(useCase UserUseCase) *UserController {
	return &UserController{
		useCase: useCase,
	}
}

func (c *UserController) CreateUser(name string) error {
	_, err := c.useCase.CreateUser(name)
	return err
}

func (c *UserController) SetUserRole(user *User, newRole Role) error {
	return c.useCase.SetUserRole(user, newRole)
}

func (c *UserController) GetUserByID(id uuid.UUID) (*User, error) {
	return c.useCase.GetUserByID(id)
}
