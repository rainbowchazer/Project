package data

import (
	"fmt"
)

type EntityNotFoundError struct {
	Message  string
	Entity   string
	Original error
}

func (e *EntityNotFoundError) Error() string {
	if e.Entity != "" {
		return fmt.Sprintf("Entity not found %s: %s", e.Entity, e.Message)
	}
	if e.Original != nil {
		return fmt.Sprintf("Entity not found: %s. Original error: %v", e.Message, e.Original)
	}
	return fmt.Sprintf("Entity not found: %s", e.Message)
}

func NewEntityNotFoundError(message string) *EntityNotFoundError {
	return &EntityNotFoundError{
		Message: message,
	}
}

func NewEntityNotFoundErrorWhithEntity(entity, message string) *EntityNotFoundError {
	return &EntityNotFoundError{
		Message: message,
		Entity:  entity,
	}
}

func NewEntityNotFoundErrorWhithCause(message string, cause error) *EntityNotFoundError {
	return &EntityNotFoundError{
		Message:  message,
		Original: cause,
	}
}
