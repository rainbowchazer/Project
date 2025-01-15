package handler

import (
	"encoding/json"
	"idel/errors/message"
	"net/http"
)

type RequestErrorHandler struct{}

func (h *RequestErrorHandler) HandleMissingParameterError(w http.ResponseWriter, errMessage string) {
	h.writeResponse(w, message.NewAppErrorMessage(errMessage, http.StatusBadRequest))
}

func (h *RequestErrorHandler) HandleUnsupportedMediaTypeError(w http.ResponseWriter, errMessage string) {
	h.writeResponse(w, message.NewAppErrorMessage(errMessage, http.StatusUnsupportedMediaType))
}

func (h *RequestErrorHandler) HandleMethodNotAllowedError(w http.ResponseWriter, errMessage string) {
	h.writeResponse(w, message.NewAppErrorMessage(errMessage, http.StatusMethodNotAllowed))
}

func (h *RequestErrorHandler) writeResponse(w http.ResponseWriter, appError *message.AppErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appError.Code)
	_ = json.NewEncoder(w).Encode(appError)
}
