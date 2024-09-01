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
	tpl := Static.index
	for i := 0; i <= len(Artistians)-1; i++ {
		if Artistians[i].Id == 21 {
			Artistians[i].Image = "https://media.istockphoto.com/id/157030584/vector/thumb-up-emoticon.jpg?s=612x612&w=0&k=20&c=GGl4NM_6_BzvJxLSl7uCDF4Vlo_zHGZVmmqOBIewgKg="
		}
	}
	err := tpl.Execute(w, struct {
		Artists []Artist
		Loc     []Locations
	}{Artistians, loc.Index})
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
	if data.Art.Id == 21 {
		data.Art.Image = "https://media.istockphoto.com/id/157030584/vector/thumb-up-emoticon.jpg?s=612x612&w=0&k=20&c=GGl4NM_6_BzvJxLSl7uCDF4Vlo_zHGZVmmqOBIewgKg="
	}
	tmpl := Static.artist

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err, "endPoints")
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
	tmpl := Static.result
	sanitized := sanitiseinput(creatin_date, first_album, members, country)
	filtredData := filterData(sanitized, Artistians)
	tmpl.Execute(w, filtredData)
}

func sanitiseinput(cd, fa, members, country []string) fiters {
	var f fiters
	b := func(data []string, a, b string) [2]string {
		if len(data) == 0 || len(data) > 2 || data[0] == "" && data[1] == "" {
			return [2]string{a, b}
		} else if data[0] == "" {
			data[0] = a
		}
		if len(data) == 1 || data[1] == "" {
			return [2]string{data[0], data[0]}
		} else {
			if data[0] >= data[1] {
				data[0], data[1] = data[1], data[0]
			}
			if data[0] < a {
				data[0] = a
			}
			if data[1] > b {
				data[1] = b
			}
			return [2]string{data[0], data[1]}
		}
	}
	f.cd = b(cd, Static.Date[0], Static.Date[1])
	f.members = b(members, Static.Member[0], Static.Member[1])
	f.fa = b(fa, Static.Fa[0], Static.Fa[1])

	f.country = func(c []string) []string {
		if len(c) == 0 {
			return Static.Countries
		}
		ans := []string{}
		for _, j := range c {
			for _, l := range Static.Countries {
				if j[1:len(j)-1] == l {
					ans = append(ans, l)
					break
				}
			}
		}
		return ans
	}(country)
	return f
}

func Search(w http.ResponseWriter, r *http.Request) {
	var result []Artist
	search := r.FormValue("search")
	seen := make([]bool, 53)

	for i, artist := range Artistians {
		for _, member := range artist.Members {
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
	f := sanitiseinput(r.Form["year"], r.Form["first-album"], r.Form["members"], r.Form["countries"])
	result = filterData(f, result)
	tml, err := template.ParseFiles("templates/result.html")
	if err != nil {
		log.Fatal(err)
	}
	tml.Execute(w, result)
}
