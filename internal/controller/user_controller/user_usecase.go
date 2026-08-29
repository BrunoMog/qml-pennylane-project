package user_controller

import (
	"pennylane_project_backend/internal/usecase/user_usecase"
)

type UserUseCase interface {
	CreateUser(name, email string) (*user_usecase.UserOutput, error)
	ChangeUserRole(input user_usecase.ChangeUserRoleInput) error
	UpdateUser(input user_usecase.UpdateUserInput) error
}
