package main

import (
	"log"
	"net/http"

	web "groupie/webServer"
)

func main() {
	web.FetchIndex()
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/artist", web.ArtistInfo)
	http.HandleFunc("/result", web.Filter)
	http.HandleFunc("/search", web.Search)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
