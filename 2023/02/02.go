package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/vitorduarte/aoc/utils"
)

type GameSet map[string]int
type Game []GameSet

func main() {
	input, _ := utils.ReadStrings("./input.txt")

	fmt.Println("⭐ 2023 - Day 1 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))

}

func solvePart1(input []string) int {
	sum := 0

	for gameId, record := range input {
		game := parseGame(record)
		isValid := isValidGame(game, GameSet{"red": 12, "green": 13, "blue": 14})
		if isValid {
			sum += gameId + 1
		}
	}

	return sum
}

func solvePart2(input []string) int {
	sum := 0

	for _, record := range input {
		game := parseGame(record)
		sum += getFewestSetPower(game)
	}

	return sum
}

func parseGame(record string) (game Game) {
	r, _ := regexp.Compile(`(\d*) (blue|red|green)`)

	sets := strings.Split(record, ";")
	for _, set := range sets {
		gameSet := GameSet{}
		colors := r.FindAllStringSubmatch(set, -1)

		for _, color := range colors {
			id := color[2]
			value, _ := strconv.Atoi(color[1])
			gameSet[id] = value
		}

		game = append(game, gameSet)
	}

	return
}

func isValidGame(game Game, reference GameSet) bool {
	for _, set := range game {
		for color, value := range set {
			if value > reference[color] {
				return false
			}
		}
	}
	return true
}

func getFewestSetPower(game Game) int {
	result := 1
	fewestSet := GameSet{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range game {
		for color, val := range set {
			if val > fewestSet[color] {
				fewestSet[color] = val
			}
		}
	}

	for _, val := range fewestSet {
		result *= val
	}

	return result
}
