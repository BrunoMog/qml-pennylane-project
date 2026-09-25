package userusecase

import (
	"errors"
	"fmt"
	"uuid"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrNoFieldsToUpdate   = errors.New("no fields to update")
	ErrPermissionDenied   = errors.New("unauthorized")
	ErrNilID              = errors.New("id cannot be nil")
)

type PermissionDeniedError struct {
	reason   string
	action   string
	callerID uuid.UUID
}

func (e *PermissionDeniedError) Error() string {
	return "permission denied: user is not permitted to perform this action"
}

func (e *PermissionDeniedError) Reason() string {
	return fmt.Sprintf("user %s tried to perform %s but was denied for the following reason: %s", e.callerID, e.action, e.reason)
}

func (e *PermissionDeniedError) Is(target error) bool {
	if target == ErrPermissionDenied {
		return true
	}

	_, ok := target.(*PermissionDeniedError)
	return ok
}
