package groupie

import "errors"

func sanitiseinput(cd, fa, members, country []string) fiters {
	var f fiters

	f.cd = check_min_max(cd, Static.Date[0], Static.Date[1], &f.err)
	f.fa = check_min_max(fa, Static.Fa[0], Static.Fa[1], &f.err)
	f.country = validateCountry(country, &f.err)
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

func validateCountry(countries []string, err *error) []string {
	if len(countries) == 0 {
		return nil
	}
	validCountries := []string{}
	for _, c := range countries {
		for _, l := range Static.Countries {
			if c == l {
				validCountries = append(validCountries, l)
				break
			}
		}
	}
	if len(countries) > len(validCountries) {
		*err = errors.Join(*err, errors.New("\n - invalid input : country\n "))
	}
	return validCountries
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
	*f.members = ans
}
