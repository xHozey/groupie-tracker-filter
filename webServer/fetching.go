package groupie

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var FilterSearch Final

func FetchIndex() {
	var wg sync.WaitGroup
	var artistians []Artist
	var location LocationIndex
	var dates DateIndex
	var relation RelationIndex
	wg.Add(4)
	go fetchAllData(&artistians, "artists", &wg)
	go fetchAllData(&location, "locations", &wg)
	go fetchAllData(&dates, "dates", &wg)
	go fetchAllData(&relation, "relation", &wg)
	wg.Wait()

	FilterSearch = Final{artistians, location, dates, relation}

}

func fetchAllData(dataVar interface{}, path string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/" + path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&dataVar)
	if err != nil {
		log.Fatal(err)
	}
}

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
