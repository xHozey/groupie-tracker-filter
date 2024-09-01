package groupie

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var (
	Artistians []Artist
	loc        Loc
)

func FetchData(dataVar interface{}, path string, id string, wg *sync.WaitGroup) {
	defer wg.Done()
	artist, errArt := http.Get("https://groupietrackers.herokuapp.com/api/" + path + id)
	if errArt != nil {
		log.Fatal(errArt)
	}
	defer artist.Body.Close()
	json.NewDecoder(artist.Body).Decode(&dataVar)
}

func fetchArtist(s string) Result {
	var wg sync.WaitGroup
	var art Artist
	var loc Locations
	var dat Dates
	var rel Relation

	wg.Add(4)
	go FetchData(&art, "artists/", s, &wg)
	go FetchData(&loc, "locations/", s, &wg)
	go FetchData(&dat, "dates/", s, &wg)
	go FetchData(&rel, "relation/", s, &wg)
	wg.Wait()
	result := Result{art, loc, dat, rel}

	return result
}
