package message

import "net/http"

type AppErrorMessage struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewAppErrorMessage(message string, status int) *AppErrorMessage {
	return &AppErrorMessage{
		Message: message,
		Code:    status,
		Status:  http.StatusText(status),
	}
}
