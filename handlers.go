package main

import (
	"net/http"
	"html/template"
)

func Handler(w http.ResponseWriter, req *http.Request) {
    tmpl := template.Must(template.ParseFiles("./templates/index.html"))
    err := tmpl.Execute(w, data)
    if ErrorCheck(err) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}