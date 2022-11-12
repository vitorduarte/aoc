package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := readInts("../../inputs/2018/01.txt")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(solvePart1(input))

	fmt.Println(solvePart2(input))

}

func readInts(filename string) ([]int, error) {
	var result []int

	file, err := os.Open(filename)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		result = append(result, x)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
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
