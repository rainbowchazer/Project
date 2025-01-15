package handler

import (
	"idel/errors/message"
	"log"
)

type BaseErrorHandler struct {
}

func (h *BaseErrorHandler) BuildAppException(err error, status int) *message.AppErrorMessage {
	return h.BuildAppErrorFromMessage(err.Error(), status)
}

func (h *BaseErrorHandler) BuildAppErrorFromMessage(msg string, status int) *message.AppErrorMessage {
	log.Println("Error:", msg)
	return message.NewAppErrorMessage(msg, status)
}
