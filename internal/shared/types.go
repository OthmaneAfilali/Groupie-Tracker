package shared

type Index struct {
    Artists   string
    Locations string
    Dates     string
    Relation  string
}

type People struct {
    ID           int64    `json:"id"`
    Image        string   `json:"image"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int64    `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
    Locations    string   `json:"locations"`
    ConcertDates string   `json:"concertDates"`
    Relations    string   `json:"relations"`
}

type Location struct {
    Index []struct {
        ID        int64    `json:"id"`
        Locations []string `json:"locations"`
        Dates     string   `json:"dates"`
    }
}

type Date struct {
    Index []struct {
        ID    int64    `json:"id"`
        Dates []string `json:"dates"`
    }
}

type Relation struct {
    Index []struct {
        ID             int64               `json:"id"`
        DatesLocations map[string][]string `json:"datesLocations"`
    }
}

type PageData struct {
    Artists   []People
    Locations Location
    Dates     Date
    Relations Relation
}

type Page struct {
    Header string
    Msg    string
}