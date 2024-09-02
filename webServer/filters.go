package groupie

import (
	"log"
	"strconv"
	"strings"
)

func filterData(f fiters, data []Artist) []Artist {
	var filtredArtists []Artist
	g := true
	for a, v := range data {
		if v.CreationDate < atoi(f.cd[0]) || v.CreationDate > atoi(f.cd[1]) {
			continue
		}

		if f.fa[1] < date(v.FirstAlbum) || date(v.FirstAlbum) > f.fa[1] {
			continue
		}

		if !f.members[len(v.Members)] {
			continue
		}
		g = true
		for _, c := range f.country {
			t := loc.Index[a].Location
			for i, loc := range t {
				if strings.HasSuffix(loc, c) {
					break
				}
				if i == len(t)-1 {
					g = false
				}
			}

		}
		if g {
			filtredArtists = append(filtredArtists, v)
		}
	}
	return filtredArtists
}

func date(fa string) string {
	s := strings.Split(fa, "-")
	day := s[2]
	year := s[0]
	month := "-" + s[1] + "-"
	date := day + month + year
	return date
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("atoi: %s | %v", s, err)
	}
	return i
}
