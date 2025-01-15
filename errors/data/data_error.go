package data

import "fmt"

type DataError struct {
	Message  string
	Entity   string
	Original error
}

func (e *DataError) Error() string {
	if e.Entity != "" {
		return fmt.Sprintf("%s: Data error %s", e.Entity, e.Message)
	}
	if e.Original != nil {
		return fmt.Sprintf("Data error: %s. Original error: %v", e.Message, e.Original)
	}
	return fmt.Sprintf("Data error: %s", e.Message)
}

func NewDataError(message string) *DataError {
	return &DataError{
		Message: message,
	}
}

func NewDataErrorWithEntity(entity, message string) *DataError {
	return &DataError{
		Message: message,
		Entity:  entity,
	}
}

func NewDataErrorWithCause(message string, cause error) *DataError {
	return &DataError{
		Message:  message,
		Original: cause,
	}
}
