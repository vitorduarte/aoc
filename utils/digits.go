package utils

import (
	"strconv"
	"unicode"
)

var (
	digitsSpelled = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9}
)

func GetFirstDigit(input string, shouldSpell bool) int {
	for i := 0; i < len(input); i++ {
		char := input[i]
		isDigit := unicode.IsDigit(rune(char))

		if isDigit {
			digit, _ := strconv.Atoi(string(char))
			return digit
		}

		if !isDigit && shouldSpell {
			ok, digit, _ := hasSpelledDigit(input, i, false)
			if ok {
				return digit
			}
		}
	}

	return 0
}

func GetLastDigit(input string, shouldSpell bool) int {
	for i := len(input) - 1; i >= 0; i-- {
		char := input[i]
		isDigit := unicode.IsDigit(rune(char))

		if isDigit {
			digit, _ := strconv.Atoi(string(char))
			return digit
		}

		if !isDigit && shouldSpell {
			ok, digit, _ := hasSpelledDigit(input, i, true)
			if ok {
				return digit
			}
		}
	}

	return 0
}

func GetDigitsFromString(input string, shouldSpell bool) []int {
	var output []int

	for i := 0; i < len(input); i++ {
		char := input[i]
		isDigit := unicode.IsDigit(rune(char))

		if isDigit {
			digit, _ := strconv.Atoi(string(char))
			output = append(output, digit)
		}

		if !isDigit && shouldSpell {
			ok, digit, step := hasSpelledDigit(input, i, false)
			if ok {
				output = append(output, digit)
				i += step
			}
		}
	}

	return output
}

func hasSpelledDigit(input string, currIndex int, reverse bool) (valid bool, digit int, step int) {
	for _, step := range []int{3, 4, 5} {
		var minIdx, maxIdx int

		if reverse {
			minIdx = currIndex + 1 - step
			maxIdx = currIndex + 1
		} else {
			minIdx = currIndex
			maxIdx = currIndex + step
		}

		if minIdx < 0 || maxIdx > len(input) {
			return false, 0, 0
		}

		word := input[minIdx:maxIdx]
		digit, ok := digitsSpelled[word]
		if ok {
			return true, digit, step - 1
		}
	}

	return false, 0, 0
}
