package data

import "fmt"

type JwtAuthenticationError struct {
	Message    string
	HttpStatus int
	Err        error
}

func NewJwtAuthenticationError(message string, status int, err error) *JwtAuthenticationError {
	return &JwtAuthenticationError{
		Message:    message,
		HttpStatus: status,
		Err:        err,
	}
}

func (e *JwtAuthenticationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("JWT Authentication Error: %s - HTTP Status: %d. Cause: %v", e.Message, e.HttpStatus, e.Err)
	}
	return fmt.Sprintf("JWT Authentication Error: %s - HTTP Status: %d", e.Message, e.HttpStatus)
}

func (e *JwtAuthenticationError) GetHttpStatus() int {
	return e.HttpStatus
}
