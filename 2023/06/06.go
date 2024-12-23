package main

import (
	"fmt"
	"strings"

	"github.com/vitorduarte/aoc/utils"
)

func main() {
	input, _ := utils.ReadStrings("input.txt")

	fmt.Println("⭐ 2023 - Day 4 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))

}

func solvePart1(input []string) int {
	res := 1

	time, distance := parseTimeDistance(input)

	for i := 0; i < len(time); i++ {
		res = res * getNumWins(time[i], distance[i])
	}

	return res
}

func solvePart2(input []string) int {
	res := 1

	clearInput := removeSpaces(input)
	time, distance := parseTimeDistance(clearInput)

	for i := 0; i < len(time); i++ {
		res = res * getNumWins(time[i], distance[i])
	}

	return res
}

func parseTimeDistance(input []string) ([]int, []int) {
	time := utils.GetNumsFromString(input[0])
	distance := utils.GetNumsFromString(input[1])

	return time, distance
}

func getNumWins(time int, distanceToWin int) int {
	winsCount := 0
	halfDists := getDistancesFirstHalf(time)

	for _, dist := range halfDists {
		if dist > distanceToWin {
			winsCount++
		}
	}

	// Doubling the wins because we only looked at first half
	res := winsCount * 2

	// If time is even, the number of peak wouldn't repeat
	// we should remove it from result
	if time%2 == 0 {
		res--
	}

	return res
}

func getDistancesFirstHalf(time int) []int {
	res := []int{}

	// Numbers generated will act as a parable
	// we just need to get the first part of parable
	for i := 1; i <= time/2; i++ {
		res = append(res, time*i-i*i)
	}

	return res
}

func removeSpaces(input []string) []string {
	res := []string{}

	for _, line := range input {
		res = append(res, strings.ReplaceAll(line, " ", ""))
	}

	return res
}
