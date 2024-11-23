package main

import (
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
    Artists    []People
    Locations  Location
    Dates      Date
    Relations  Relation
}

var data PageData

func main() {
    data = fetchAllData()
	http.Handle("/assets/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/", Handler)
	log.Println("Listening on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
