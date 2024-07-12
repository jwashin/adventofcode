package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var symbols = "/@#$%&*+="
var digits = "0123456789"

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", part1(string(input)))
}

func part1(s string) int {
	data := getValidNumbers(s)
	total := 0
	for _, v := range data {
		total += v
	}
	return total
}

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
		for ky, number := range numbers {
			if isSymbolAdjacent(number, endIndexes, s, index, ky) {
				theNumber, _ := strconv.Atoi(number)
				output = append(output, theNumber)
			}
		}
	}
	return output
}

func isSymbolAdjacent(number string, endIndexes []int, s string, index int, ky int) bool {
	testZone := []string{}
	lines := strings.Split(s, "\n")
	if index > 0 {
		testZone = append(testZone, lines[index-1])
	}
	testZone = append(testZone, lines[index])
	if index < len(lines)-1 {
		testZone = append(testZone, lines[index+1])
	}
	start := endIndexes[ky] - len(number) - 1
	if start < 0 {
		start = 0
	}
	end := endIndexes[ky] + 1
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
