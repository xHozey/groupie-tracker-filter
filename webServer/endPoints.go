package groupie

import (
	"html/template"
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
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	FilterSearch.Art[21].Image = "https://media.istockphoto.com/id/157030584/vector/thumb-up-emoticon.jpg?s=612x612&w=0&k=20&c=GGl4NM_6_BzvJxLSl7uCDF4Vlo_zHGZVmmqOBIewgKg="

	err = tpl.Execute(w, FilterSearch)
	if err != nil {
		log.Fatal(err)
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
	if data.Art.Id == 21 {
		data.Art.Image = "https://media.istockphoto.com/id/157030584/vector/thumb-up-emoticon.jpg?s=612x612&w=0&k=20&c=GGl4NM_6_BzvJxLSl7uCDF4Vlo_zHGZVmmqOBIewgKg="
	}
	tmpl, errtpl := template.ParseFiles("templates/artist.html")
	if errtpl != nil {
		log.Fatal(errtpl)
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func Filter(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	members := r.Form["members"]
	country := r.FormValue("countries")
	cdMin, _ := strconv.Atoi(r.FormValue("year1"))
	cdMax, _ := strconv.Atoi(r.FormValue("year2"))
	faMin := r.FormValue("faMin")
	faMax := r.FormValue("faMax")

	faR := getRangeStr(faMin, faMax)

	tmpl, errtpl := template.ParseFiles("templates/result.html")
	if errtpl != nil {
		log.Fatal(errtpl)
	}
	filtredData := filterData(members, cdMin, cdMax, faR, country)
	tmpl.Execute(w, filtredData)
}

func Search(w http.ResponseWriter, r *http.Request) {
	var result Final
	search := r.FormValue("search")
	seen := make([]bool, 53)

	for _, artist := range FilterSearch.Art {
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(artist.Name), search) {
				if !seen[artist.Id] {
					result.Art = append(result.Art, artist)
					seen[artist.Id] = true
				}
			}
			if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), search) {
				if !seen[artist.Id] {
					result.Art = append(result.Art, artist)
					seen[artist.Id] = true
				}
			}
			if strings.Contains(strings.ToLower(artist.FirstAlbum), search) {
				if !seen[artist.Id] {
					result.Art = append(result.Art, artist)
					seen[artist.Id] = true
				}
			}
			if strings.Contains(strings.ToLower(member), search) {
				if !seen[artist.Id] {
					result.Art = append(result.Art, artist)
					seen[artist.Id] = true
				}
			}
		}
		for _, location := range FilterSearch.Location.Index {
			for _, loc := range location.Location {
				if strings.Contains(strings.ToLower(loc), search) {
					if location.Id == artist.Id && !seen[artist.Id] {
						result.Art = append(result.Art, artist)
						seen[artist.Id] = true
					}
				}
			}
		}
	}

	tml, err := template.ParseFiles("templates/result.html")
	if err != nil {
		log.Fatal(err)
	}
	tml.Execute(w, result)
}
