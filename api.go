package main

import (
	"io"
	"encoding/json"
	"net/http"
	"log"
)

func fetchAllData() PageData {
    baseURL := "https://groupietrackers.herokuapp.com/api"
    
    var index Index
    json.Unmarshal(FetchApi(baseURL), &index)

    var data PageData

    json.Unmarshal(FetchApi(index.Artists), &data.Artists)
    json.Unmarshal(FetchApi(index.Locations), &data.Locations)
    json.Unmarshal(FetchApi(index.Dates), &data.Dates)
    json.Unmarshal(FetchApi(index.Relation), &data.Relations)

    return data
}

func FetchApi(url string) []byte {
    res, err := http.Get(url)
    ErrorCheck(err)
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    ErrorCheck(err)
    return body
}

func ErrorCheck(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
