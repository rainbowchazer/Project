package data

import "fmt"

type PermissionError struct {
	Message string
	Cause   error
}

func NewPermissionError(message string) *PermissionError {
	return &PermissionError{
		Message: "Permission violation: " + message,
	}
}

func NewPermissionErrorWithCause(message string, cause error) *PermissionError {
	return &PermissionError{
		Message: "Permission violation: " + message,
		Cause:   cause,
	}
}

func (e *PermissionError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s. Cause: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *PermissionError) Unwrap() error {
	return e.Cause
}
