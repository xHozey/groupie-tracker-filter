package groupie

import (
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

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
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

	f := sanitiseinput(r.Form["year"], r.Form["first-album"], r.Form["members"], r.Form["country"], r.Form["location"])
	result = filterData(f, result)
	tml := Static.result
	err := tml.Execute(w, result)
	if err != nil {
		http.Error(w, "internal server eror", http.StatusInternalServerError)
		return
	}
}
