package groupie

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
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
	fetchIndex()
	for i := 0; i <= len(Artistians)-1; i++ {
		if Artistians[i].Id == 21 {
			Artistians[i].Image = "https://media.istockphoto.com/id/157030584/vector/thumb-up-emoticon.jpg?s=612x612&w=0&k=20&c=GGl4NM_6_BzvJxLSl7uCDF4Vlo_zHGZVmmqOBIewgKg="
		}
	}
	err = tpl.Execute(w, Artistians)
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
	var date string
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	members := r.Form["members"]
	country := r.FormValue("countries")
	cd, _ := strconv.Atoi(r.FormValue("CreationDate"))
	fa := r.FormValue("first-album")
	if fa != "" {
		date = spltdate(fa)
	}
	tmpl, errtpl := template.ParseFiles("templates/result.html")
	if errtpl != nil {
		log.Fatal(errtpl)
	}
	filtredData := filterData(members, cd, date, country)
	tmpl.Execute(w, filtredData)
}
