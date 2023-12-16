package utils

import (
	"regexp"
	"strconv"
)

func GetNumsFromString(input string) (res []int) {
	r := regexp.MustCompile(`(\d+)`)
	matches := r.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		num, _ := strconv.Atoi(match[1])
		res = append(res, num)
	}
	return
}
