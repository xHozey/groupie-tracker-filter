package groupie

import (
	"errors"
	"fmt"
	"strings"
)

func sanitiseinput(cd, fa, members, country, locations []string) fiters {
	var f fiters

	f.cd = check_min_max(cd, Static.Date[0], Static.Date[1], &f.err)
	f.fa = check_min_max(fa, Static.Fa[0], Static.Fa[1], &f.err)
	f.country, f.locations = validateCountry(country, locations, &f.err)
	fmt.Println(locations, f.locations)
	validateMembers(&f, members)
	return f
}

func check_min_max(data []string, a, b string, f *error) [2]string {
	if len(data) == 0 || len(data) > 2 || data[0] == "" && data[1] == "" {
		return [2]string{a, b}
	}
	if data[0] == "" {
		data[0] = a
	}
	if len(data) == 1 || data[1] == "" {
		return [2]string{data[0], data[0]}
	}
	if data[0] >= data[1] {
		data[0], data[1] = data[1], data[0]
		*f = errors.Join(*f, errors.New("\n - carefull min and max values swaped\n "))
	}
	if data[0] < a {
		data[0] = a
	}
	if data[1] > b {
		data[1] = b
	}
	return [2]string{data[0], data[1]}
}

func validateCountry(countries, locations []string, err *error) ([]string, []string) {
	if len(countries) == 0 {
		return nil, nil
	}
	validCountries := []string{}
	validlocations := []string{}
	chekloc := len(locations) != 0
	for _, c := range countries {
		for _, l := range Static.Countries {
			if c == l {
				fmt.Println("daz", c, chekloc)
				validCountries = append(validCountries, l)
				if chekloc {
					for _, l1 := range locations {
						if strings.HasSuffix(l1, l) {
							fmt.Println("daz2", l1)
							for _, c1 := range Static.Countloc[l] {
								if l1 == c1 {
									validlocations = append(validlocations, l1)
									fmt.Println("daz3", c1)
								}
							}
						}
					}
				}
				break
			}
		}
	}
	if len(countries) > len(validCountries) {
		*err = errors.Join(*err, errors.New("\n - invalid input : country\n "))
	}
	fmt.Println(validCountries, validlocations)
	return validCountries, validlocations
}

func validateMembers(f *fiters, members []string) {
	ans := make([]bool, 9)
	for _, v := range members {
		if i := atoi(v); i < 9 && i > 0 {
			ans[i] = true
		} else {
			f.err = errors.Join(f.err, errors.New("\n - invalid input : members\n "))
		}
	}
	f.members = ans
}
