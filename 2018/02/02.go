package main

import (
	"fmt"
	"log"

	"github.com/vitorduarte/aoc/utils"
)

func main() {
	input, err := utils.ReadStrings("./02.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 2 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))

}

func countRepetitions(word string) (hasDoubleChar, hasTripleChar bool) {
	charCounter := make(map[rune]int)
	for _, ch := range word {
		charCounter[ch] += 1
	}

	for _, value := range charCounter {
		if value == 3 {
			hasTripleChar = true
		}

		if value == 2 {
			hasDoubleChar = true
		}

	}

	return
}

func solvePart1(input []string) int {

	countDouble := 0
	countTriple := 0

	for _, boxId := range input {
		hasDouble, hasTriple := countRepetitions(boxId)

		if hasDouble {
			countDouble++
		}

		if hasTriple {
			countTriple++
		}
	}

	return countDouble * countTriple
}

func solvePart2(input []string) string {
	for i, boxId := range input {
		for _, boxIdToCompare := range input[i+1:] {
			if boxIdMatches(boxId, boxIdToCompare) {
				return buildFixedBoxId(boxId, boxIdToCompare)
			}
		}
	}

	return ""
}

func boxIdMatches(first string, second string) bool {
	charDiff := 0

	for idx, char := range first {
		if byte(char) != second[idx] {
			charDiff++
		}
	}

	return charDiff == 1
}

func buildFixedBoxId(first string, second string) string {
	var fixedBoxId []rune

	for idx, char := range first {
		if byte(char) == second[idx] {
			fixedBoxId = append(fixedBoxId, char)
		}
	}

	return string(fixedBoxId)
}
