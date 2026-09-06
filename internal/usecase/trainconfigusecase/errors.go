package trainconfigusecase

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}

type TrainConfigNameAlreadyExistsError struct {
	Name string
}

func (e *TrainConfigNameAlreadyExistsError) Error() string {
	return "train config name already exists: " + e.Name
}

type UnauthorizedError struct{}

func (e *UnauthorizedError) Error() string {
	return "unauthorized"
}

type InvalidInputError struct{}

func (e *InvalidInputError) Error() string {
	return "invalid input"
}

type NoFieldsToUpdateError struct{}

func (e *NoFieldsToUpdateError) Error() string {
	return "no fields to update"
}
