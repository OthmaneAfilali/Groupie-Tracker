package utils

import (
	"groupie-tracker/internal/shared"
	"log"
	"net/http"
	"strconv"
)

func ErrorHandler(w http.ResponseWriter, statusCode int) {
	page := &shared.Page{Header: strconv.Itoa(statusCode)}
	w.WriteHeader(statusCode)
	switch page.Header {
	case "400":
		page.Msg = "400 bad request"
	case "404":
		page.Msg = "404 page not found"
	case "405":
		page.Msg = "405 method not allowed"
	default:
		page.Msg = "500 internal server error"
	}
	shared.ErrTmpl.Execute(w, page)
}

func ErrorCheck(err error) bool {
	if err != nil {
		LogError("Error occurred", err)
		return true
	}
	return false
}

func LogError(message string, err error) {
	log.Printf("%s: %v", message, err)
}
