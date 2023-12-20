package main

import (
	"fmt"
	"regexp"

	"github.com/vitorduarte/aoc/cmd/utils"
)

type cordinates struct {
	left  string
	right string
}

func main() {
	input, _ := utils.ReadStrings("input.txt")

	fmt.Println("⭐ 2023 - Day 4 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))

}

func solvePart1(input []string) int {
	instructions, cordinatesMap := parseInput(input)
	currPos := "AAA"
	count := 0

	for {
		if currPos == "ZZZ" {
			break
		}
		// Use module to make sure we'll go back to the first
		// instruction when we reach the end of the instructions
		instIdx := count % len(instructions)
		inst := instructions[instIdx]

		if inst == 'R' {
			currPos = cordinatesMap[currPos].right
		}

		if inst == 'L' {
			currPos = cordinatesMap[currPos].left
		}

		count++
	}

	return count
}

func solvePart2(input []string) int {
	instructions, cordinatesMap := parseInput(input)
	startCords := findStartingCordinates(cordinatesMap)

	rates := []int{}
	for _, cord := range startCords {
		_, rate := findCandidatesRate(cord, instructions, cordinatesMap)
		rates = append(rates, rate)
	}

	lcm, _ := utils.LCM(rates...)
	return lcm

}

func parseInput(input []string) (string, map[string]cordinates) {
	instructions := input[0]
	cordinatesMap := map[string]cordinates{}

	for _, line := range input[2:] {
		r := regexp.MustCompile(`([A-Z]{3})`)
		matches := r.FindAllStringSubmatch(line, -1)

		key := matches[0][0]
		left := matches[1][0]
		right := matches[2][0]
		cordinatesMap[key] = cordinates{left: left, right: right}

	}
	return instructions, cordinatesMap
}

func findStartingCordinates(cordinatesMap map[string]cordinates) (res []string) {
	for key := range cordinatesMap {
		if key[2] == 'A' {
			res = append(res, key)
		}
	}

	return
}

func findCandidatesRate(start string, instructions string, cordinatesMap map[string]cordinates) (candidates []int, rate int) {
	currPos := start
	count := 0

	for {
		if currPos[2] == 'Z' {
			// res := fmt.Sprintf("%v-%v", count, start)
			// fmt.Printf("res:>> %v\n", res)
			maxCandidates := len(candidates)
			if maxCandidates > 3 {
				if candidates[maxCandidates-1]-candidates[maxCandidates-2] == candidates[maxCandidates-2]-candidates[maxCandidates-3] {
					rate = candidates[maxCandidates-1] - candidates[maxCandidates-2]
					return
				}
			}
			candidates = append(candidates, count)
		}

		// Use module to make sure we'll go back to the first
		// instruction when we reach the end of the instructions
		instIdx := count % len(instructions)
		inst := instructions[instIdx]

		if inst == 'R' {
			currPos = cordinatesMap[currPos].right
		}

		if inst == 'L' {
			currPos = cordinatesMap[currPos].left
		}

		if currPos == start {
			rate = count - candidates[len(candidates)-1]
			return
		}

		count++
	}

	return
}
