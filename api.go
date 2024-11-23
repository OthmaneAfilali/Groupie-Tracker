package main

import (
	"io"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)

func fetchAllData() PageData {
    baseURL := "https://groupietrackers.herokuapp.com/api"
    
    var index Index
    body := FetchApi(baseURL)
    if body == nil || ErrorCheck(json.Unmarshal(body, &index)) {
        return PageData{}
    }

    var data PageData

    body = FetchApi(index.Artists)
    if body == nil || ErrorCheck(json.Unmarshal(body, &data.Artists)) {
        return PageData{}
    }

    body = FetchApi(index.Locations)
    if body == nil || ErrorCheck(json.Unmarshal(body, &data.Locations)) {
        return PageData{}
    }

    body = FetchApi(index.Dates)
    if body == nil || ErrorCheck(json.Unmarshal(body, &data.Dates)) {
        return PageData{}
    }

    body = FetchApi(index.Relation)
    if body == nil || ErrorCheck(json.Unmarshal(body, &data.Relations)) {
        return PageData{}
    }

    return data
}

func FetchApi(url string) []byte {
    res, err := http.Get(url)
    if ErrorCheck(err) {
        return nil
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if ErrorCheck(err) {
        return nil
    }
    return body
}

func validateData(data PageData) error {
    if len(data.Artists) == 0 {
        return fmt.Errorf("artists data is empty")
    }
    for _, artist := range data.Artists {
        if artist.Name == "" {
            return fmt.Errorf("artist with ID %d has no name", artist.ID)
        }
    }

    if len(data.Locations.Index) == 0 {
        return fmt.Errorf("locations data is empty")
    }
    for _, location := range data.Locations.Index {
        if len(location.Locations) == 0 {
            return fmt.Errorf("location with ID %d has no locations", location.ID)
        }
    }

    if len(data.Dates.Index) == 0 {
        return fmt.Errorf("dates data is empty")
    }
    for _, date := range data.Dates.Index {
        if len(date.Dates) == 0 {
            return fmt.Errorf("date with ID %d has no dates", date.ID)
        }
    }

    if len(data.Relations.Index) == 0 {
        return fmt.Errorf("relations data is empty")
    }
    for _, relation := range data.Relations.Index {
        if len(relation.DatesLocations) == 0 {
            return fmt.Errorf("relation with ID %d has no dates/locations", relation.ID)
        }
    }

    return nil
}

func ErrorCheck(err error) bool {
    if err != nil {
        log.Println(err)
        return true
    }
    return false
}
