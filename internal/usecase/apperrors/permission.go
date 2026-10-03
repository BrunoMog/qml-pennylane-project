package apperrors

import (
	"errors"
	"fmt"
	"uuid"
)

var (
	ErrPermissionDenied = errors.New("permission denied")
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
	return target == ErrPermissionDenied
}

func NewPermissionDeniedError(reason string, action string, callerID uuid.UUID) *PermissionDeniedError {
	return &PermissionDeniedError{
		reason:   reason,
		action:   action,
		callerID: callerID,
	}
}
