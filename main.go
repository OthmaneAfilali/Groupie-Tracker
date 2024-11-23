package main

import (
	"groupie-tracker/internal/api"
	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/shared"
	"html/template"
	"log"
	"net/http"
)

type Index struct {
	Artists   string
	Locations string
	Dates     string
	Relation  string
}

type People struct {
	ID           int64
	Image        string
	Name         string
	Members      []string
	CreationDate int64
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Location struct {
	Index []struct {
		ID        int64
		Locations []string
		Dates     string
	}
}

type Date struct {
	Index []struct {
		ID    int64
		Dates []string
	}
}

type Relation struct {
	Index []struct {
		ID             int64
		DatesLocations map[string][]string
	}
}

type PageData struct {
	Artists   []People
	Locations Location
	Dates     Date
	Relations Relation
}

func init() {
	shared.IndexTmpl = template.Must(template.ParseFiles("./templates/index.html"))
	shared.AboutTmpl = template.Must(template.ParseFiles("./templates/about.html"))
	shared.BioTmpl = template.Must(template.ParseFiles("./templates/bio.html"))
	shared.ErrTmpl = template.Must(template.ParseFiles("./templates/error.html"))
}

func main() {
	shared.Data = api.FetchAllData()
	err := api.ValidateData(shared.Data)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/groupie-tracker/about", handlers.AboutHandler)
	http.HandleFunc("/groupie-tracker/bio", handlers.BioHandler)
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
