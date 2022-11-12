package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadStrings(filename string) ([]string, error) {
	var result []string

	file, err := os.Open(filename)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}
