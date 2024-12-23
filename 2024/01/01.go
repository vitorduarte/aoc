package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/vitorduarte/aoc/utils"
)

func main() {
	leftNums, rightNums, err := utils.ReadPairInts("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("⭐ 2024 - Day 1 ⭐")
	fmt.Printf("Part 1: %.0f\n", solvePart1(leftNums, rightNums))
	fmt.Printf("Part 2: %d\n", solvePart2(leftNums, rightNums))

}

func solvePart1(leftNums, rightNums []int) float64 {
	var sum float64

	sort.IntSlice(leftNums).Sort()
	sort.IntSlice(rightNums).Sort()

	for i := 0; i < len(leftNums); i++ {
		sum += math.Abs(float64(leftNums[i] - rightNums[i]))
	}

	return sum
}

func solvePart2(leftNums, rightNums []int) int {
	var sum int

	rightFreq := map[int]int{}

	for _, num := range rightNums {
		rightFreq[num]++
	}

	for _, num := range leftNums {
		sum += num * rightFreq[num]
	}

	return sum
}
