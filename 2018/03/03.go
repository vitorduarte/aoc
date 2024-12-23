package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("./03.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println("Day 2 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))
}

func solvePart1(input string) (overlapCount int) {
	re := regexp.MustCompile(`#(\d*) @ (\d*),(\d*): (\d*)x(\d*)`)
	match := re.FindAllStringSubmatch(input, -1)

	fabricMatrix := [1000][1000]int{}

	for _, claim := range match {
		claimData := convertToIntegers(claim)
		claimId, startX, startY, width, height := claimData[1], claimData[2], claimData[3], claimData[4], claimData[5]

		for i := startX; i < startX+width; i++ {
			for j := startY; j < startY+height; j++ {
				if fabricMatrix[i][j] == 0 {
					fabricMatrix[i][j] = claimId
					continue
				}

				if fabricMatrix[i][j] > 0 {
					overlapCount++
					fabricMatrix[i][j] = -1
				}
			}
		}
	}

	return
}

func solvePart2(input string) (claimId int) {
	re := regexp.MustCompile(`#(\d*) @ (\d*),(\d*): (\d*)x(\d*)`)
	match := re.FindAllStringSubmatch(input, -1)

	fabricMatrix := [1000][1000]int{}
	overlapMap := make(map[int]bool)

	for _, claim := range match {
		claimData := convertToIntegers(claim)
		claimId, startX, startY, width, height := claimData[1], claimData[2], claimData[3], claimData[4], claimData[5]
		overlapMap[claimId] = false

		for i := startX; i < startX+width; i++ {
			for j := startY; j < startY+height; j++ {
				currFabric := fabricMatrix[i][j]

				if currFabric == 0 {
					fabricMatrix[i][j] = claimId
					continue
				}

				if currFabric > 0 {
					overlapMap[currFabric] = true
					overlapMap[claimId] = true
					fabricMatrix[i][j] = claimId
				}
			}
		}
	}

	for key, val := range overlapMap {
		if !val {
			return key
		}
	}

	return
}

func convertToIntegers(input []string) (result []int) {
	for _, attr := range input {
		val, _ := strconv.Atoi(attr)
		result = append(result, val)
	}
	return result
}
