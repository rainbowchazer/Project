package data

import "fmt"

type AddToDefaultError struct {
	Message  string // Описание ошибки.
	Entity   string // Название сущности, связанной с ошибкой.
	Original error  // Исходная ошибка, если она есть.
}

func (e *AddToDefaultError) Error() string {
	if e.Entity != "" {
		return fmt.Sprintf("Could not add to default entity %s: %s", e.Entity, e.Message)
	}
	if e.Original != nil {
		return fmt.Sprintf("Could not add to default entity: %s. Original error: %v", e.Message, e.Original)
	}
	return fmt.Sprintf("Could not add to default entity: %s", e.Message)
}

func NewAddToDefaultError(message string) *AddToDefaultError {
	return &AddToDefaultError{Message: message}
}

func NewAddToDefaultErrorWithEntity(entity, message string) *AddToDefaultError {
	return &AddToDefaultError{Entity: entity, Message: message}
}

func NewAddToDefaultErrorWithCause(message string, cause error) *AddToDefaultError {
	return &AddToDefaultError{Message: message, Original: cause}
}
