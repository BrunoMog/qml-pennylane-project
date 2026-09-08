package experimentusecase

import "fmt"

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user not found")
}

type UnauthorizedError struct{}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized")
}

type VQCConfigNotFoundError struct{}

func (e *VQCConfigNotFoundError) Error() string {
	return fmt.Sprintf("vqc config not found")
}

type TrainingConfigNotFoundError struct{}

func (e *TrainingConfigNotFoundError) Error() string {
	return fmt.Sprintf("training config not found")
}
