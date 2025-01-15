package resolver

import (
	"encoding/json"
	"idel/errors/message"
	"log"
	"net/http"
)

type RestResponseStatusErrorResolver struct{}

func NewRestResponseStatusErrorResolver() *RestResponseStatusErrorResolver {
	return &RestResponseStatusErrorResolver{}
}

func (r *RestResponseStatusErrorResolver) HandleException(w http.ResponseWriter, err error) {
	log.Printf("Error: %s", err.Error())

	errorMessage := message.NewAppErrorMessage(err.Error(), http.StatusForbidden)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
		log.Printf("Error encoding JSON: %s", err.Error())
	}
}
