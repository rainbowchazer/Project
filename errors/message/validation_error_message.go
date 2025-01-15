package message

import "net/http"

type ValidationErrorMessage struct {
	AppErrorMessage
	Violations map[string]string `json:"violations"`
}

func NewValidationErrorMessage(violations map[string]string) *ValidationErrorMessage {
	return &ValidationErrorMessage{
		AppErrorMessage: AppErrorMessage{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Validation error",
		},
		Violations: violations,
	}
}
