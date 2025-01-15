package handler

import (
	"encoding/json"
	"idel/errors/data"
	"idel/errors/message"
	"net/http"
)

type GlobalErrorHandler struct{}

func (h *GlobalErrorHandler) HandleException(w http.ResponseWriter, err error) {
	h.writeResponse(w, message.NewAppErrorMessage(err.Error(), http.StatusBadRequest))
}

func (h *GlobalErrorHandler) HandleNoSuchFileException(w http.ResponseWriter, err *data.PermissionError) {
	h.writeResponse(w, message.NewAppErrorMessage(err.Error(), http.StatusNotFound))
}

func (h *GlobalErrorHandler) HandlePermissionException(w http.ResponseWriter, err *data.PermissionError) {
	h.writeResponse(w, message.NewAppErrorMessage(err.Error(), http.StatusForbidden))
}

func (h *GlobalErrorHandler) writeResponse(w http.ResponseWriter, appError *message.AppErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appError.Code)
	_ = json.NewEncoder(w).Encode(appError)
}

func MiddlewareErrorHandler(h *GlobalErrorHandler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					err, ok := rec.(error)
					if !ok {
						err = data.NewPermissionError("Something went wrong")
					}
					h.HandleException(w, err)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
