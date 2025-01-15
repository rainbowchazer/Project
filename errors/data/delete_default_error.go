package data

import (
	"fmt"
)

type DeleteDefaultError struct {
	Message  string
	Entity   string
	Original error
}

func (e *DeleteDefaultError) Error() string {
	if e.Entity != "" {
		return fmt.Sprintf("Can not delete default entity %s: %s", e.Entity, e.Message)
	}
	if e.Original != nil {
		return fmt.Sprintf("Can not delete default entity: %s. Original error: %v", e.Message, e.Original)
	}
	return fmt.Sprintf("Can not delete default entity: %s", e.Message)
}

func NewDeleteDefaultError(message string) *DeleteDefaultError {
	return &DeleteDefaultError{
		Message: message,
	}
}

func NewDeleteDefaultErrorWithEntity(entity, message string) *DeleteDefaultError {
	return &DeleteDefaultError{
		Message: message,
		Entity:  entity,
	}
}

func NewDeleteDefaultErrorWithCause(message string, cause error) *DeleteDefaultError {
	return &DeleteDefaultError{
		Message:  message,
		Original: cause,
	}
}
