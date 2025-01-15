package data

import (
	"fmt"
)

type InvalidDataError struct {
	Message  string
	Entity   string
	Original error
}

func (e *InvalidDataError) Error() string {
	if e.Entity != "" {
		return fmt.Sprintf("%s: Invalid data given: %s", e.Entity, e.Message)
	}
	if e.Original != nil {
		return fmt.Sprintf("Invalid data given: %s. Original error: %v", e.Message, e.Original)
	}
	return fmt.Sprintf("Ivalid data given: %s", e.Message)
}

func NewInvalidDataError(message string) *InvalidDataError {
	return &InvalidDataError{
		Message: message,
	}
}

func NewInvalidDataErrorWithEntity(entity, message string) *InvalidDataError {
	return &InvalidDataError{
		Message: message,
		Entity:  entity,
	}
}

func NewInvalidDataErrorWithCause(message string, cause error) *InvalidDataError {
	return &InvalidDataError{
		Message:  message,
		Original: cause,
	}
}
