package groupie

import (
	"html/template"
	"log"
	"net/http"
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
		log.Fatal(err, "endPoints")
	}
	fetchIndex()
	for i := 0; i <= len(Artistians)-1; i++ {
		if Artistians[i].Id == 21 {
			Artistians[i].Image = "https://media.istockphoto.com/id/157030584/vector/thumb-up-emoticon.jpg?s=612x612&w=0&k=20&c=GGl4NM_6_BzvJxLSl7uCDF4Vlo_zHGZVmmqOBIewgKg="
		}
	}
	err = tpl.Execute(w, Artistians)
	if err != nil {
		log.Fatal(err, "endPoints")
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
		log.Fatal(errtpl, "endPoints")
	}

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
	tmpl, errtpl := template.ParseFiles("templates/result.html")
	if errtpl != nil {
		log.Fatal(errtpl, "endPoints")
	}
	sanitized := sanitiseinput(creatin_date, first_album, members, country)
	filtredData := filterData(sanitized)
	tmpl.Execute(w, filtredData)
}

type fiters struct {
	cd, fa, members [2]string
	country         []string
	err             error
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
	f.cd = b(cd, "1957", "2024")
	f.members = b(members, "1", "8")
	f.fa = b(fa, "1967-08-12", "2018-11-15")

	f.country = func(c []string) []string {
		if len(c) == 0 {
			return countries
		}
		ans := []string{}
		for _, j := range c {
			for _, l := range countries {
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

var countries = []string{"argentina", "australia", "austria", "belarus", "belgium", "brazil", "canada", "chile", "china", "colombia", "costa_rica", "czechia", "denmark", "finland", "france", "french_polynesia", "germany", "greece", "hungary", "india", "indonesia", "ireland", "italy", "japan", "mexico", "netherlands", "netherlands_antilles", "new_caledonia", "new_zealand", "norway", "peru", "philippines", "poland", "portugal", "qatar", "romania", "saudi_arabia", "slovakia", "south_korea", "spain", "sweden", "switzerland", "taiwan", "thailand", "uk", "united_arab_emirates", "usa"}
