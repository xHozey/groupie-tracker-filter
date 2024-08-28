package groupie

import (
	"strconv"
	"strings"
)

func filterData(m []string, cd int, fa string, country string) []Artist {
	var filtredArtists []Artist
	for I, v := range Artistians {
		addart := true
		membrbool := true
		locabool := true
		if cd != 1957 && v.CreationDate != cd {
			addart = false
		}
		if fa != "" && v.FirstAlbum != fa {
			addart = false
		}
		if len(m) != 0 {
			for _, T := range m {
				Tint, _ := strconv.Atoi(T)
				if len(v.Members) != Tint {
					membrbool = false
				}else if addart{
					membrbool = true
					break
				}
			}
		}
		if !membrbool {
			addart = false
		}
		for _,v := range loc.Index[I].Location {
			if !strings.Contains(v, country) {
				locabool = false
			}else if addart{
				locabool = true
				break
			}
		}
		if !locabool {
			addart = false
		} 
		if addart {
			filtredArtists = append(filtredArtists, v)
		}

	}
	return filtredArtists
}

func spltdate(fa string) string {
	s := strings.Split(fa, "-")
	day := s[2]
	year := s[0]
	month := "-" + s[1] + "-"
	date := day + month + year
	return date
}
