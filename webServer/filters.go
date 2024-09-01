package groupie

import (
	"log"
	"strconv"
	"strings"
)

func filterData(f fiters, data []Artist) []Artist {
	var g = make([]bool, len(data))
	var filtredArtists []Artist
	for I, v := range data {
		g[I] = true
		if v.CreationDate < atoi(f.cd[0]) || v.CreationDate > atoi(f.cd[1]) {
			g[I] = false
			continue
		}

		if f.fa[1] < date(v.FirstAlbum) || date(v.FirstAlbum) > f.fa[1] {
			g[I] = false
			continue
		}

		if atoi(f.members[0]) < len(v.Members) && len(v.Members) > atoi(f.members[1]) {
			g[I] = false
			continue
		}
		for _, c := range f.country {
			t := loc.Index[I].Location
			for i, loc := range t {
				if strings.HasSuffix(loc, c) {
					break
				}
				if i == len(t)-1 {
					g[I] = false
					break
					break
				}
			}

		}
	}
	for i, j := range g {
		if j {
			filtredArtists = append(filtredArtists, data[i])
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
