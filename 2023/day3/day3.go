package main

import (
	"strconv"
	"strings"
)

var symbols = "/@#$%&*+="
var digits = "0123456789"

func getIntegers(s string) ([]string, []int) {
	numbers := []string{}
	endIndexes := []int{}
	candidate := ""
	for index, v := range s {
		if strings.Contains(digits, string(v)) {
			candidate += string(v)
		} else {
			if len(candidate) > 0 {
				numbers = append(numbers, candidate)
				candidate = ""
				endIndexes = append(endIndexes, index)
			}
		}
	}
	return numbers, endIndexes
}

func getValidNumbers(s string) []int {
	output := []int{}
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	for index, line := range lines {
		numbers, endIndexes := getIntegers(line)
		for _, number := range numbers {
			if isSymbolAdjacent(number, endIndexes, s, index) {
				theNumber, _ := strconv.Atoi(number)
				output = append(output, theNumber)
			}
		}
	}
	return output
}

func isSymbolAdjacent(number string, endIndexes []int, s string, index int) bool {
	testZone := []string{}
	lines := strings.Split(s, "\n")
	if index > 0 {
		testZone = append(testZone, lines[index-1])
	}
	testZone = append(testZone, lines[index])
	if index < len(lines)-1 {
		testZone = append(testZone, lines[index+1])
	}
	start := endIndexes[index] - len(number)
	if start < 0 {
		start = 0
	}
	end := endIndexes[index] + 1
	if end >= len(testZone[0]) {
		end -= 1
	}
	for _, v := range testZone {
		testString := v[start:end]
		for _, t := range testString {
			if strings.Contains(symbols, string(t)) {
				return true
			}
		}
	}
	return false
}
