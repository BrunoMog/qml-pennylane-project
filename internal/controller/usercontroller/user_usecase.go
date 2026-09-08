package usercontroller

import (
	"pennylane_project_backend/internal/usecase/userusecase"
)

type UserUseCase interface {
	CreateUser(input userusecase.CreateUserInput) (*userusecase.UserOutput, error)
	ChangeUserRole(input userusecase.ChangeUserRoleInput) error
	ChangeOwner(input userusecase.ChangeOwnerInput) error
	UpdateUser(input userusecase.UpdateUserInput) error
	DeleteUser(input userusecase.DeleteUserInput) error
}
