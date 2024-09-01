package groupie

import (
	"strconv"
	"strings"
)

func filterData(m []string, cd []int, fa []string, country string) Final {

	return Final{}
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
	strMin := strings.Split(min, "-")
	strMax := strings.Split(max, "-")
	if len(strMin) == 3 && len(strMax) == 3 {
		intMin, _ := strconv.Atoi(strMin[2])
		intMax, _ := strconv.Atoi(strMax[2])
		intR := getRangeInt(intMin, intMax)
		for i := 0; i < len(intR); i++ {
			result = append(result, strconv.Itoa(intR[i]))
		}
	}

	return result
}
