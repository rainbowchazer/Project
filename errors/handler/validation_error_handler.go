package handler

import (
	"encoding/json"
	"idel/errors/message"
	"log"
	"net/http"
)

type ValidationErrorHandler struct{}

func (h *ValidationErrorHandler) HandleValidationError(w http.ResponseWriter, violations map[string]string) {
	log.Printf("Validation errors: %v", violations)

	validationErrorMessage := message.NewValidationErrorMessage(violations)

	h.writeResponse(w, validationErrorMessage)
}

func (h *ValidationErrorHandler) writeResponse(w http.ResponseWriter, validationError *message.ValidationErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(validationError.Code)
	_ = json.NewEncoder(w).Encode(validationError)
}
