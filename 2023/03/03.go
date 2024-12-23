package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/vitorduarte/aoc/utils"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func main() {
	input, _ := utils.ReadStrings("input.txt")

	fmt.Println("⭐ 2023 - Day 3 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))
}

func solvePart1(input []string) (result int) {
	numbersLines := foundIndexNumbers(input)

	for _, line := range numbersLines {
		isAdjacent, _ := isAdjacentToSymbol(input, line)
		if isAdjacent {
			number, _ := strconv.Atoi(input[line.start.x][line.start.y:line.end.y])
			result += number
		}
	}
	return
}

func solvePart2(input []string) (result int) {
	starMap := map[Point][]int{}

	numbersLines := foundIndexNumbers(input)
	for _, line := range numbersLines {
		isAdjacent, points := isAdjacentToStar(input, line)
		if isAdjacent {
			number, _ := strconv.Atoi(input[line.start.x][line.start.y:line.end.y])

			for _, point := range points {
				starLocations, ok := starMap[point]
				if ok {
					starMap[point] = append(starLocations, number)
				} else {
					starMap[point] = []int{number}
				}
			}
		}
	}

	for _, numbers := range starMap {
		if len(numbers) > 1 {
			gearRatio := 1
			for _, num := range numbers {
				gearRatio *= num
			}
			result += gearRatio
		}
	}
	return
}

func foundIndexNumbers(input []string) (numbersCord []Line) {
	re := regexp.MustCompile(`(\d+)`)

	for x, input := range input {
		matches := re.FindAllStringIndex(input, -1)
		for _, match := range matches {
			numbersCord = append(
				numbersCord,
				Line{
					start: Point{x: x, y: match[0]},
					end:   Point{x: x, y: match[1]},
				})
		}
	}

	return
}

func isAdjacentToSymbol(input []string, line Line) (bool, []Point) {
	return isAdjacentToExpression(input, line, `[^0-9.\s]`)
}

func isAdjacentToStar(input []string, line Line) (bool, []Point) {
	return isAdjacentToExpression(input, line, `[*]`)
}

func isAdjacentToExpression(input []string, line Line, regex string) (isAdjacent bool, points []Point) {
	minX := line.start.x - 1
	maxX := line.end.x + 1
	minY := line.start.y - 1
	maxY := line.end.y + 1

	adjacentPoints := []Point{}

	if minY > 0 {
		adjacentPoints = append(adjacentPoints, Point{x: line.start.x, y: line.start.y - 1})
	} else {
		minY = 0
	}

	if line.end.y < len(input[0]) {
		adjacentPoints = append(adjacentPoints, Point{x: line.end.x, y: line.end.y})
	} else {
		maxY = len(input[0])
	}

	if minX > 0 {
		for y := minY; y < maxY; y++ {
			adjacentPoints = append(adjacentPoints, Point{x: minX, y: y})
		}
	}

	if maxX < len(input) {
		for y := minY; y < maxY; y++ {
			adjacentPoints = append(adjacentPoints, Point{x: maxX, y: y})
		}
	}

	r, _ := regexp.Compile(regex)
	for _, point := range adjacentPoints {
		if r.MatchString(string(input[point.x][point.y])) {
			isAdjacent = true
			points = append(points, point)
		}
	}

	return isAdjacent, points
}
