package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInts(filename string) ([]int, error) {
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
