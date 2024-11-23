package handlers

import (
	"groupie-tracker/internal/utils"
	"groupie-tracker/internal/shared"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" && req.URL.Path != "/groupie-tracker" {
		utils.ErrorHandler(w, http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	shared.IndexTmpl.Execute(w, shared.Data)
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	page := &shared.Page{Header: "About Us"}
	shared.AboutTmpl.Execute(w, page)
}

func BioHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		handleBioPost(w, req)
	case "GET":
		handleBioGet(w, req)
	default:
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
	}
}

func handleBioPost(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if utils.ErrorCheck(err) {
		utils.ErrorHandler(w, http.StatusBadRequest)
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

	var artist *shared.People
	for _, a := range shared.Data.Artists {
		if a.Name == artistName {
			artist = &a
			break
		}
	}

	if artist == nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	shared.BioTmpl.Execute(w, artist)
}
