package groupie

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	tpl := Static.index
	err := tpl.Execute(w, struct {
		Artists   []Artist
		Loc       []Locations
		Countries []string
		Locations map[string][]string
	}{Artistians, loc.Index, Static.Countries, Static.Countloc})
	// fmt.Println(Static.Countloc)
	if err != nil {
		log.Print(err, "endPoints")
	}
}

func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("Id")
	data := fetchArtist(id)
	if data.Art.Id == 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tmpl := Static.artist

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "internal server eror", http.StatusInternalServerError)
		return
	}
}

func Filter(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	creatin_date := r.Form["year"]
	first_album := r.Form["first-album"]
	members := r.Form["members"]
	country := r.Form["countries"]
	locations := r.Form["location"]
	fmt.Println(locations)
	tmpl := Static.result
	sanitized := sanitiseinput(creatin_date, first_album, members, country, locations)
	filtredData := filterData(sanitized, Artistians)
	err = tmpl.Execute(w, filtredData)
	if err != nil {
		http.Error(w, "internal server eror", http.StatusInternalServerError)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	var result []Artist
	search := strings.ToLower(r.FormValue("search"))
	seen := make([]bool, 53)
	for i, artist := range Artistians {
		if strings.Contains(strings.ToLower(artist.Name), search) {
			if !seen[artist.Id] {
				result = append(result, artist)
				seen[artist.Id] = true
			}
		}
		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), search) {
			if !seen[artist.Id] {
				result = append(result, artist)
				seen[artist.Id] = true
			}
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), search) {
			if !seen[artist.Id] {
				result = append(result, artist)
				seen[artist.Id] = true
			}
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), search) {
				if !seen[artist.Id] {
					result = append(result, artist)
					seen[artist.Id] = true
				}
			}
		}
		for _, location := range loc.Index[i].Location {
			if strings.Contains(strings.ToLower(location), search) {
				if !seen[artist.Id] {
					result = append(result, artist)
					seen[artist.Id] = true
				}
			}
		}
	}
	fmt.Println(r.Form["location"], r.Form["countries"])
	f := sanitiseinput(r.Form["year"], r.Form["first-album"], r.Form["members"], r.Form["countries"], r.Form["location"])
	result = filterData(f, result)
	tml := Static.result
	err := tml.Execute(w, result)
	if err != nil {
		http.Error(w, "internal server eror", http.StatusInternalServerError)
		return
	}
}
