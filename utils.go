package main

import (
	"net/http"
	"strconv"
	"log"
)

func errorHandler(w http.ResponseWriter, statusCode int) {
	page := &Page{Header: strconv.Itoa(statusCode)}
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
	errTmpl.Execute(w, page)
}

func ErrorCheck(err error) bool {
    if err != nil {
        log.Println(err)
        return true
    }
    return false
}
