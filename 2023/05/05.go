package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/vitorduarte/aoc/utils"
)

type Sheet map[string]string
type Almanac map[string]Sheet

func createAlmanac(input []string) Almanac {
	currentSheetId := ""
	res := Almanac{}

	for _, line := range input {
		sheetId, isSheetId := getSheetIdFromInput(line)
		if isSheetId {
			currentSheetId = sheetId
			continue
		}

		seeds := getNumsFromLine(line)
		if len(seeds) > 0 {
			// The input representation is reverse
			res.addValueToSheet(currentSheetId, seeds[1], seeds[0], seeds[2])
		}
	}

	return res
}

func (a Almanac) getSheet(id string) Sheet {
	sheet, ok := a[id]
	if ok {
		return sheet
	}

	return Sheet{}
}

func (a Almanac) addValueToSheet(id string, startKey, startValue, increment int) {
	sheet := a.getSheet(id)
	key := fmt.Sprintf("%v-%v", startKey, startKey+increment-1)
	value := fmt.Sprintf("%v-%v", startValue, startValue+increment-1)

	sheet[key] = value
	a[id] = sheet
}

func (s Sheet) getValue(num int) int {
	for key, value := range s {
		keyMin, keyMax := getMinMaxFromId(key)
		valueMin, _ := getMinMaxFromId(value)
		if keyMin <= num && keyMax >= num {
			increment := num - keyMin
			return valueMin + increment
		}
	}

	return num
}

func (s Sheet) getValueByRange(input []string) []string {
	response := []string{}

	for _, valueRange := range input {
		response = append(response, s.mapIdRange(valueRange)...)
	}

	return response
}

func (s Sheet) mapIdRange(idRange string) []string {
	response := []string{}

	if idRange == "" {
		return response
	}

	idMin, idMax := getMinMaxFromId(idRange)
	idIncrement := idMax - idMin

	for sheetKey, sheetValue := range s {
		keyMin, keyMax := getMinMaxFromId(sheetKey)
		valueMin, valueMax := getMinMaxFromId(sheetValue)

		// The minimum value of id is inside this key
		if idMin >= keyMin && idMin <= keyMax {
			mapIncrement := idMin - keyMin

			// In case every values are inside the same range
			if idMax <= keyMax {
				mappedValue := createRangeIdentifier(valueMin+mapIncrement, valueMin+idIncrement+mapIncrement)
				return append(response, mappedValue)
			}

			mappedId := createRangeIdentifier(valueMin+mapIncrement, valueMax)
			response := []string{mappedId}

			upperId := createRangeIdentifier(keyMax+1, idMax)
			response = append(response, s.mapIdRange(upperId)...)

			return response
		}

		// The maximum value of id is inside this key
		if idMax >= keyMin && idMax <= keyMax {
			mapDecrement := keyMax - idMax
			// If all the values are inside the same range,
			// the previous if would catch it, so we can asume that
			// only the maximum value is inside this key
			mappedId := createRangeIdentifier(valueMin, valueMax-mapDecrement)
			response := []string{mappedId}

			lowerId := createRangeIdentifier(idMin, keyMin-1)
			response = append(response, s.mapIdRange(lowerId)...)
			return response

		}

		// The key is a subset of id
		if keyMin >= idMin && keyMax <= idMax {
			mappedId := createRangeIdentifier(valueMin, valueMax)
			response := []string{mappedId}

			lowerId := createRangeIdentifier(idMin, keyMin-1)
			upperId := createRangeIdentifier(keyMax+1, idMax)

			response = append(response, s.mapIdRange(lowerId)...)
			response = append(response, s.mapIdRange(upperId)...)
			return response
		}
	}

	// If couldn't map, return itself
	return []string{idRange}
}

func createRangeIdentifier(min, max int) string {
	return fmt.Sprintf(
		"%d-%d",
		min,
		max,
	)
}

func getMinMaxFromId(id string) (min, max int) {
	idSplitted := strings.Split(id, "-")
	max, _ = strconv.Atoi(idSplitted[1])
	min, _ = strconv.Atoi(idSplitted[0])
	return
}

func main() {
	input, _ := utils.ReadStrings("input.txt")

	fmt.Println("⭐ 2023 - Day 5 ⭐")
	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))
}

func solvePart1(input []string) int {
	seeds := getNumsFromLine(input[0])
	almanac := createAlmanac(input[1:])

	seedsLocation := getSeedsLocation(seeds, almanac)

	return utils.Min(seedsLocation)
}

func solvePart2(input []string) int {
	seedsInput := getNumsFromLine(input[0])
	seedRangeIds := createSeedRangeIds(seedsInput)
	almanac := createAlmanac(input[1:])

	seedsLocation := getSeedsLocationByRange(seedRangeIds, almanac)

	minLocation := math.MaxInt
	for _, location := range seedsLocation {
		min, _ := getMinMaxFromId(location)
		if min < minLocation {
			minLocation = min
		}
	}

	return minLocation
}

func createSeedRangeIds(input []int) []string {
	response := []string{}

	for i := 0; i < len(input)-1; i += 2 {
		id := createRangeIdentifier(input[i], input[i]+input[i+1])
		response = append(response, id)
	}

	return response
}

func getNumsFromLine(input string) (seeds []int) {
	r := regexp.MustCompile(`(\d+)`)
	matches := r.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		num, _ := strconv.Atoi(match[1])
		seeds = append(seeds, num)
	}
	return
}

func getSheetIdFromInput(input string) (string, bool) {
	r := regexp.MustCompile(`(.*) map:`)

	matches := r.FindAllStringSubmatch(input, -1)
	if len(matches) > 0 {
		return matches[0][1], true
	}

	return "", false
}

func getSeedsLocation(seeds []int, almanac Almanac) (location []int) {
	sheetFlow := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	for _, seed := range seeds {
		currentId := seed
		for _, sheetId := range sheetFlow {
			currentId = almanac[sheetId].getValue(currentId)
		}
		location = append(location, currentId)
	}

	return
}

func getSeedsLocationByRange(rangeIds []string, almanac Almanac) (location []string) {
	sheetFlow := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	currentRangeIds := rangeIds

	for _, sheetId := range sheetFlow {
		currentRangeIds = almanac.getSheet(sheetId).getValueByRange(currentRangeIds)
	}

	return currentRangeIds
}
