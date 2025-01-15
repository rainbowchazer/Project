package data

import (
	"fmt"
)

type RelationError struct {
	Message  string
	Entity   string
	Original error
}

func (e *RelationError) Error() string {
	if e.Entity != "" {
		return fmt.Sprintf("%s: Invalid relation: %s", e.Entity, e.Message)
	}
	if e.Original != nil {
		return fmt.Sprintf("Invalid relation: %s. Original error: %v", e.Message, e.Original)
	}
	return fmt.Sprintf("Invalid relation: %s", e.Message)
}

func NewRelationError(message string) *RelationError {
	return &RelationError{
		Message: message,
	}
}

func NewRelationErrorWithEntity(entity, message string) *RelationError {
	return &RelationError{
		Message: message,
		Entity:  entity,
	}
}

func NewRelationErrorWithCause(message string, cause error) *RelationError {
	return &RelationError{
		Message:  message,
		Original: cause,
	}
}
