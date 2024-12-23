package main

import (
	"fmt"

	"github.com/vitorduarte/aoc/utils"
)

func main() {
	input, _ := utils.ReadStrings("./input.txt")

	fmt.Println("⭐ 2023 - Day 1 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))
	// fmt.Println("Part 2:", utils.GetLastDigit(input[6], true))
}

func solvePart1(input []string) int {
	result := 0
	for _, word := range input {
		first, last := getFirstAndLastDigit(word, false)
		result += first*10 + last
	}

	return result
}

func solvePart2(input []string) int {
	result := 0
	for _, word := range input {
		first, last := getFirstAndLastDigit(word, true)
		result += first*10 + last
		// fmt.Printf("%s = %d\n", word, first*10+last)
	}

	return result
}

func getFirstAndLastDigit(word string, shouldSpell bool) (first, last int) {
	return utils.GetFirstDigit(word, shouldSpell), utils.GetLastDigit(word, shouldSpell)

	// digits := utils.GetDigitsFromString(word, shouldSpell)

	// if len(digits) > 0 {
	// 	return digits[0], digits[len(digits)-1]
	// }

	// return 0, 0
}
