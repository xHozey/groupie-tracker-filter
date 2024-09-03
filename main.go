package main

import (
	"log"
	"net/http"

	web "groupie/webServer"
)

func main() {
	fs := http.FileServer(http.Dir("./templates/assests"))
	http.Handle("/templates/assests/", http.StripPrefix("/templates/assests/", fs))
	http.HandleFunc("/search", web.Search)
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/artist", web.ArtistInfo)
	http.HandleFunc("/result", web.Filter)
	log.Fatal(http.ListenAndServe(":8080", nil), "Listen and Serve")
}
