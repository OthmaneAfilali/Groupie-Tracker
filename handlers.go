package main

import (
	"net/http"
	"strconv"
)

type Page struct {
	Header string
	Msg   string
}

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" && req.URL.Path != "/groupie-tracker" {
		errorHandler(w, http.StatusNotFound)
		return
	}

    err := indexTmpl.Execute(w, data)
    if ErrorCheck(err) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	page := &Page{Header: "About Us"}
	aboutTmpl.Execute(w, page)
}

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
