package handler

import (
	"idel/errors/data"
	"idel/errors/message"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var appError *message.AppErrorMessage

	switch e := err.(type) {
	case *data.DataError:
		log.Printf("Data error: %s", e.Error())
		appError = message.NewAppErrorMessage(e.Error(), http.StatusForbidden)
	case *data.InvalidDataError:
		log.Printf("Invalid data error: %s", e.Error())
		appError = message.NewAppErrorMessage(e.Error(), http.StatusBadRequest)
	case *data.RelationError:
		log.Printf("Relation error: %s", e.Error())
		appError = message.NewAppErrorMessage(e.Error(), http.StatusBadRequest)
	case *data.EntityNotFoundError:
		log.Printf("Entity not found: %s", e.Error())
		appError = message.NewAppErrorMessage(e.Error(), http.StatusNotFound)
	default:
		log.Printf("Unknown error: %s", err.Error())
		appError = message.NewAppErrorMessage("Internal server error", http.StatusInternalServerError)
	}

	c.JSON(appError.Code, appError)
}
