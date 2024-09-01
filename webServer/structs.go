package groupie

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func init() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	err = json.Unmarshal(file, &Static)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
	}

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err1 := json.NewDecoder(resp.Body).Decode(&Artistians)
	if err1 != nil {
		log.Fatal(err)
	}

	loca, eror := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if eror != nil {
		log.Fatal(eror)
	}
	defer loca.Body.Close()
	err = json.NewDecoder(loca.Body).Decode(&loc)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	Static.index, err = template.ParseFiles("templates/index.html")
	if err != nil || Static.index == nil {
		log.Fatal(err, "endPoints")
	}

	Static.artist, err = template.ParseFiles("templates/artist.html")
	if err != nil || Static.artist == nil {
		log.Fatal(err, "endPoints")
	}
	Static.result, err = template.ParseFiles("templates/result.html")
	if err != nil || Static.result == nil {
		log.Fatal(err, "endPoints")
	}
}

type static struct {
	Countries []string `json:"countries"`
	Date      []string `json:"date"`
	Fa        []string `json:"fa"`
	Member    []string `json:"members"`
	index     *template.Template
	artist    *template.Template
	result    *template.Template
}

var Static static

type fiters struct {
	cd, fa, members [2]string
	country         []string
	err             error
}

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
