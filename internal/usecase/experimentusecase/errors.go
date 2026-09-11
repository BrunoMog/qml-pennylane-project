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

type InvalidInputError struct{}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input")
}

type NoFieldsToUpdateError struct{}

func (e *NoFieldsToUpdateError) Error() string {
	return fmt.Sprintf("no fields to update")
}

type ExperimentNameAlreadyExistsError struct{}

func (e *ExperimentNameAlreadyExistsError) Error() string {
	return fmt.Sprintf("experiment name already exists")
}
