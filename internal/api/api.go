package api

import (
	"groupie-tracker/internal/utils"
	"groupie-tracker/internal/shared"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchAllData() shared.PageData {
	baseURL := "https://groupietrackers.herokuapp.com/api"

	var index shared.Index
	if err := fetchAndUnmarshal(baseURL, &index); err != nil {
		utils.LogError("Failed to fetch index data", err)
		return shared.PageData{}
	}

	var data shared.PageData
	if err := fetchAndUnmarshal(index.Artists, &data.Artists); err != nil {
		utils.LogError("Failed to fetch artists data", err)
		return shared.PageData{}
	}

	if err := fetchAndUnmarshal(index.Locations, &data.Locations); err != nil {
		utils.LogError("Failed to fetch locations data", err)
		return shared.PageData{}
	}

	if err := fetchAndUnmarshal(index.Dates, &data.Dates); err != nil {
		utils.LogError("Failed to fetch dates data", err)
		return shared.PageData{}
	}

	if err := fetchAndUnmarshal(index.Relation, &data.Relations); err != nil {
		utils.LogError("Failed to fetch relations data", err)
		return shared.PageData{}
	}

	return data
}

func fetchAndUnmarshal(url string, v interface{}) error {
	body := FetchApi(url)
	if body == nil {
		return fmt.Errorf("failed to fetch data from %s", url)
	}
	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("failed to unmarshal data from %s: %w", url, err)
	}
	return nil
}

func FetchApi(url string) []byte {
	res, err := http.Get(url)
	if utils.ErrorCheck(err) {
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if utils.ErrorCheck(err) {
		return nil
	}
	return body
}

func ValidateData(data shared.PageData) error {
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
