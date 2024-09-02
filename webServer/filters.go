package groupie

import (
	"strconv"
	"strings"
)

func filterData(m []string, cdMin int, cdMax int, fa []string, country string) Final {
	mint := slicestrtoint(m)

	filterData := filterCD(FilterSearch, cdMin, cdMax)
	filterData = filterFA(filterData, fa)
	filterData = filterM(filterData, mint)
	filterData = filterLocation(filterData, country)

	return filterData
}

func filterCD(data Final, min int, max int) Final {
	if min == 0 && max == 0 {
		return data
	}
	var result Final
	for _, artist := range data.Art {
		if artist.CreationDate >= min && artist.CreationDate <= max {
			result.Art = append(result.Art, artist)
		}
	}

	return result
}

func filterFA(data Final, fa []string) Final {
	var result Final
	if len(fa) == 0 {
		return data
	}
	for _, artist := range data.Art {
		for _, val := range fa {
			if strings.Contains(artist.FirstAlbum, val) {
				result.Art = append(result.Art, artist)
			}
		}
	}
	return result
}

func filterM(data Final, member []int) Final {
	var result Final
	if len(member) == 0 {
		return data
	}
	for _, artist := range data.Art {
		for i := 0; i < len(member); i++ {
			if len(artist.Members) == member[i] {
				result.Art = append(result.Art, artist)
			}
		}
	}
	return result
}

func filterLocation(data Final, c string) Final {
	var result Final
	if c == "" {
		return data
	}
	for _, artist := range data.Art {
		for _, locationArtist := range FilterSearch.Location.Index {
			for _, loc := range locationArtist.Location {
				if strings.Contains(loc, c) {
					if artist.Id == locationArtist.Id {
						result.Art = append(result.Art, artist)
					}
				}
			}
		}
	}
	return result
}

func getRangeInt(cd1 int, cd2 int) []int {
	var result []int
	for i := cd1; i <= cd2; i++ {
		result = append(result, i)
	}
	return result
}

func getRangeStr(min string, max string) []string {
	var result []string

	if min == "" || max == "" {
		return result
	}

	intMin, err := strconv.Atoi(min)
	if err != nil {
		return result
	}

	intMax, err := strconv.Atoi(max)
	if err != nil {
		return result
	}

	intR := getRangeInt(intMin, intMax)
	for i := 0; i < len(intR); i++ {
		result = append(result, strconv.Itoa(intR[i]))
	}
	return result
}

func slicestrtoint(member []string) []int {
	var membresint []int
	for _, vstr := range member {
		vint, _ := strconv.Atoi(vstr)
		membresint = append(membresint, vint)
	}
	return membresint
}
