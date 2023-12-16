package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vitorduarte/aoc/cmd/utils"
)

func main() {
	input, _ := utils.ReadStrings("input.txt")

	fmt.Println("⭐ 2023 - Day 4 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))

}

func solvePart1(input []string) int {
	return solve(input, false)
}

func solvePart2(input []string) int {
	return solve(input, true)
}

func solve(input []string, joker bool) int {
	res := 0
	hands := parseHands(input, joker)
	sort.Sort(hands)

	currRank := 1
	for _, hand := range hands {
		res += hand.bid * currRank
		currRank++
	}

	return res
}

func parseHands(input []string, joker bool) (hands Hands) {
	for _, line := range input {
		lineSplit := strings.Split(line, " ")
		cards := lineSplit[0]
		bid, _ := strconv.Atoi(lineSplit[1])

		hands = append(hands, Hand{cards: cards, bid: bid, joker: joker})
	}

	return
}
