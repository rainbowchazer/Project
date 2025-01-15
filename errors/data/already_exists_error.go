package data

import (
	"fmt"
)

type AlreadyExistsError struct {
	Message  string
	Entity   string
	Original error
}

func (e *AlreadyExistsError) Error() string {
	if e.Original != nil {
		return fmt.Sprintf("Entity already exists: %s. Details: %s. Original error: %v", e.Entity, e.Message, e.Original)
	}
	if e.Entity != "" {
		return fmt.Sprintf("Entity already exists: %s. Details: %s", e.Entity, e.Message)
	}
	return fmt.Sprintf("Entity already exists: %s", e.Message)
}

func NewAlreadyExistsError(message string) *AlreadyExistsError {
	return &AlreadyExistsError{
		Message: message,
	}
}

func NewAlreadyExistsErrorWithEntity(entity, message string) *AlreadyExistsError {
	return &AlreadyExistsError{
		Entity:  entity,
		Message: message,
	}
}

func NewAlreadyExistsErrorWithCause(message string, cause error) *AlreadyExistsError {
	return &AlreadyExistsError{
		Original: cause,
		Message:  message,
	}
}

func NewAlreadyExistErrorWithEntityAndCause(entity, message string, cause error) *AlreadyExistsError {
	return &AlreadyExistsError{
		Entity:   entity,
		Message:  message,
		Original: cause,
	}
}
