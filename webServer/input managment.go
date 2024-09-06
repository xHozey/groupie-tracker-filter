package groupie

import (
	"errors"
	"strings"
)

func sanitiseinput(cd, fa, members, country, locations []string) fiters {
	var f fiters

	f.cd = check_min_max(cd, Static.Date[0], Static.Date[1], &f.err)
	f.fa = check_min_max(fa, Static.Fa[0], Static.Fa[1], &f.err)
	f.country_loc = validateCountry(country, locations, &f.err)
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

func validateCountry(countries, locations []string, err *error) map[string][]string {
	if len(countries) == 0 {
		return nil
	}

	ans := make(map[string][]string)
	chekloc := len(locations) != 0
	for _, c := range countries {
		for _, l := range Static.Countries {
			if c == l {

				if len(ans[l]) == 0 {
					ans[l] = []string{}
				}
				if chekloc {
					for _, l1 := range locations {
						if strings.HasSuffix(l1, l) {
							for _, c1 := range Static.Countloc[l] {
								if l1 == c1 {
									ans[l] = append(ans[l], l1)
								}
							}
						}
					}
				}
				break
			}
		}
	}
	if len(countries) > len(ans) {
		*err = errors.Join(*err, errors.New("\n - invalid input : country\n "))
	}

	return ans
}

func validateMembers(f *fiters, members []string) {
	if len(members) == 0 {
		f.members = []bool{true, true, true, true, true, true, true, true, true}
		return
	}
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
