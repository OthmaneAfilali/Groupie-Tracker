package main

import (
	"net/http"
)

type Page struct {
	Header string
	Msg    string
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" && req.URL.Path != "/groupie-tracker" {
		errorHandler(w, http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	indexTmpl.Execute(w, data)
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	page := &Page{Header: "About Us"}
	aboutTmpl.Execute(w, page)
}

func bioHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		handleBioPost(w, req)
	case "GET":
		handleBioGet(w, req)
	default:
		errorHandler(w, http.StatusMethodNotAllowed)
	}
}

func handleBioPost(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if ErrorCheck(err) {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	artistName := req.FormValue("name")
	if artistName == "" {
		http.Error(w, "Artist name is required", http.StatusBadRequest)
		return
	}

	http.Redirect(w, req, "/groupie-tracker/bio?name="+artistName, http.StatusSeeOther)
}

func handleBioGet(w http.ResponseWriter, req *http.Request) {
	artistName := req.URL.Query().Get("name")
	if artistName == "" {
		http.Error(w, "Artist name is required", http.StatusBadRequest)
		return
	}

	var artist *People
	for _, a := range data.Artists {
		if a.Name == artistName {
			artist = &a
			break
		}
	}

	if artist == nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	bioTmpl.Execute(w, artist)
}
