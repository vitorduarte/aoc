package main

import (
	"fmt"
	"regexp"
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

func solvePart1(input []string) (result int) {
	for _, line := range input {
		wins := countWins(line)

		if wins > 0 {
			result += 1 << (wins - 1)
		}
	}
	return
}

type CardCount map[int]int

func (c CardCount) getCount(cardId int) int {
	count, ok := c[cardId]
	if !ok {
		c[cardId] = 1
		return 1
	}

	return count
}

func (c CardCount) incrementCount(cardId int, increment int) {
	count, ok := c[cardId]
	if !ok {
		c[cardId] = 1 + increment
		return
	}

	c[cardId] = count + increment
}

func solvePart2(input []string) (result int) {
	cards := CardCount{}

	for i, line := range input {
		currCardId := i + 1
		nextCardId := currCardId + 1
		currCardCount := cards.getCount(currCardId)

		wins := countWins(line)
		for i := nextCardId; i < nextCardId+wins; i++ {
			cards.incrementCount(i, currCardCount)
		}
	}

	for _, card := range cards {
		result += card
	}
	return
}

func countWins(input string) int {
	winNums, myNums := parseCardInput(input)

	wins := 0
	for _, myNum := range myNums {
		_, ok := winNums[myNum]
		if ok {
			wins++
		}
	}

	return wins
}

func parseCardInput(input string) (winningNums map[int]bool, myNums []int) {
	r, _ := regexp.Compile(`Card\s*\d*: ([\d*\s*]*) \| ([\d*\s*]*)`)
	matches := r.FindAllStringSubmatch(input, -1)

	winningNums = convertStringToNumbersMap(matches[0][1])
	myNums = convertStringToNumbers(matches[0][2])
	return
}

func convertStringToNumbers(input string) (result []int) {
	for _, num := range strings.Split(input, " ") {
		convNum, err := strconv.Atoi(num)
		if err == nil {
			result = append(result, convNum)
		}
	}

	return
}

func convertStringToNumbersMap(input string) map[int]bool {
	result := map[int]bool{}

	for _, num := range strings.Split(input, " ") {
		convNum, err := strconv.Atoi(num)
		if err == nil {
			result[convNum] = true
		}
	}

	return result
}
