package main

import (
	"fmt"
	"log"

	"github.com/vitorduarte/aoc/cmd/utils"
)

func main() {
	input, err := utils.ReadInts("../../inputs/2018/01.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1 Solution:", solvePart1(input))
	fmt.Println("Part 2 Solutiion:", solvePart2(input))

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

	for true {
		freqUsed[currFreq] = true
		currFreq += input[idx%len(input)]

		if _, ok := freqUsed[currFreq]; ok {
			return
		}
		idx++
	}

	return
}
