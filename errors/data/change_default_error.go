package data

import (
	"fmt"
)

type ChangeDefaultError struct {
	Message  string
	Entity   string
	Original error
}

func (e *ChangeDefaultError) Error() string {
	if e.Original != nil {
		return fmt.Sprintf("Could not change default entity %s: %s. Original error: %v", e.Entity, e.Message, e.Original)
	}
	return fmt.Sprintf("Could not change default entity %s: %s", e.Entity, e.Message)
}

func NewChangeDefaultError(message string) *ChangeDefaultError {
	return &ChangeDefaultError{
		Message: message,
	}
}

func NewChangeDefaultErrorWithEntity(entity, message string) *ChangeDefaultError {
	return &ChangeDefaultError{
		Entity:  entity,
		Message: message,
	}
}

func NewChangeDefaultErrorWithCause(message string, cause error) *ChangeDefaultError {
	return &ChangeDefaultError{
		Original: cause,
		Message:  message,
	}
}
