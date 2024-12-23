package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadPairInts(filename string) (leftNums []int, rightNums []int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return leftNums, rightNums, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var leftNum, rightNum int
		fields := strings.Fields(scanner.Text())

		leftNum, err = strconv.Atoi(fields[0])
		if err != nil {
			return
		}

		rightNum, err = strconv.Atoi(fields[1])
		if err != nil {
			return
		}

		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)
	}

	return
}
