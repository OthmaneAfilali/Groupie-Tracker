package main

import (
	"net/http"
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

