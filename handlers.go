package main

import (
	"net/http"
	"html/template"
)

func Handler(w http.ResponseWriter, req *http.Request) {
    tmpl := template.Must(template.ParseFiles("./templates/index.html"))
    tmpl.Execute(w, data)
}