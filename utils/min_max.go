package utils

import "math"

// Todo upgrade to generics
func Min(input []int) int {
	res := math.MaxInt

	for _, n := range input {
		if n < res {
			res = n
		}
	}

	return res
}
