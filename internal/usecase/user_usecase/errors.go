package user_usecase

import "fmt"

type EmailAlreadyExistsError struct {
	email string
}

func (e *EmailAlreadyExistsError) Error() string {
	return fmt.Sprintf("email already exists: %s", e.email)
}

type UnauthorizedError struct {
	name string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized: user %s is not authorized to perform this action", e.name)
}

type UserNotFoundError struct {
}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}
