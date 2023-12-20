package utils

import "fmt"

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) (int, error) {
	if len(integers) < 2 {
		return 0, fmt.Errorf("input should have at least 2 elements")
	}

	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result, _ = LCM([]int{result, integers[i]}...)
	}

	return result, nil
}
