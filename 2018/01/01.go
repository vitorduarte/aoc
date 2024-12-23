package main

import (
	"fmt"
	"log"

	"github.com/vitorduarte/aoc/utils"
)

func main() {
	input, err := utils.ReadInts("./01.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 1 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))
}

func solvePart1(input []int) int {
	result := 0
	for _, freq := range input {
		result += freq
	}

	return result
}

func solvePart2(input []int) (currFreq int) {
	freqUsed := make(map[int]bool)
	idx := 0

	for {
		freqUsed[currFreq] = true
		currFreq += input[idx%len(input)]

		if _, ok := freqUsed[currFreq]; ok {
			return
		}
		idx++
	}

	return
}
