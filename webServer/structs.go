package groupie

type Result struct {
	Art          Artist
	Location     Locations
	Date         Dates
	DateLocation Relation
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Loc struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
