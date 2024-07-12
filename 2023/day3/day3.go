package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var symbols = "/@#$%&*+=-"
var digits = "0123456789"

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", part1(string(input)))
}

func part1(s string) int {
	// 552835 too low
	// 535104 too low
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
	// if the end of a line is a number...
	// omitting this was my error in part 1
	if len(candidate) > 0 {
		numbers = append(numbers, candidate)
		candidate = ""
		endIndexes = append(endIndexes, len(s)-1)
	}
	return numbers, endIndexes
}

func getIntegerAtIndex(s string, index int) int {
	candidate := string(s[index])
	leftDone := false
	leftidx := index
	for !leftDone && leftidx > 0 {
		leftidx -= 1
		if strings.Contains(digits, string(s[leftidx])) {
			candidate = string(s[leftidx]) + candidate
		} else {
			leftDone = false
		}
	}
	rightDone := false
	rightidx := index
	for !rightDone && rightidx <= len(s)-1 {
		rightidx += 1
		if strings.Contains(digits, string(s[rightidx])) {
			candidate = candidate + string(s[rightidx])
		} else {
			rightDone = false
		}
	}
	t, _ := strconv.Atoi(candidate)
	return t

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
	// index is the current line number

	// make a rectangular blob one char bigger than number
	// on all sides
	testZone := []string{}

	lines := strings.Split(s, "\n")

	// if we are not at the top, add the line above
	if index > 0 {
		testZone = append(testZone, lines[index-1])
	}
	// add self line
	testZone = append(testZone, lines[index])

	// if we are not at the bottom, add the line below
	if index < len(lines)-1 {
		testZone = append(testZone, lines[index+1])
	}

	start := endIndexes[ky] - len(number) - 1
	if start < 0 {
		start = 0
	}

	end := endIndexes[ky] + 1
	if end > len(testZone[0])-1 {
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
