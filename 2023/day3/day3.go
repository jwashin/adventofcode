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
	fmt.Println("part 2:", part2(string(input)))
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

func part2(s string) int {
	// 13270325 too low
	s = strings.TrimSpace(s)
	total := 0
	xes := getAllStarCoordinates(s)
	tableau := strings.Split(s, "\n")
	for _, x := range xes {
		j := getIntegersAdjacentToCoordinate(x, tableau)
		if len(j) == 2 {
			total += j[0] * j[1]
		}
	}
	return total
}

func getIntegersAdjacentToCoordinate(c Coordinate, tableau []string) []int {
	items := []int{}
	testRows := []string{}
	coords := map[IntId]int{}
	if c.row > 0 {
		testRows = append(testRows, tableau[c.row-1])
	}
	testRows = append(testRows, tableau[c.row])
	if c.row < len(tableau)-1 {
		testRows = append(testRows, tableau[c.row+1])
	}
	for _, row := range testRows {
		if c.column > 0 {
			i, err := getIntegerAtIndex(row, c.column-1)
			if err == nil {
				coords[i] = 1
			}
		}
		i, err := getIntegerAtIndex(row, c.column)
		if err == nil {
			coords[i] = 1
		}
		if c.column < len(row)-1 {
			i, err := getIntegerAtIndex(row, c.column+1)
			if err == nil {
				coords[i] = 1
			}
		}
	}
	if len(coords) == 2 {
		for k := range coords {
			items = append(items, k.value)
		}
	}

	return items
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

type Coordinate struct {
	row    int
	column int
}

func getAllStarCoordinates(s string) []Coordinate {
	tableau := strings.Split(s, "\n")
	stars := []Coordinate{}
	for line := range tableau {
		for column, char := range tableau[line] {
			if char == '*' {
				stars = append(stars, Coordinate{line, column})
			}
		}
	}
	return stars
}

type IntId struct {
	value int
	start int
	end   int
}

func getIntegerAtIndex(s string, index int) (IntId, error) {
	candidate := string(s[index])
	_, err := strconv.Atoi(candidate)
	if err != nil {
		return IntId{}, err
	}
	leftDone := false
	leftidx := index
	for !leftDone && leftidx > 0 {
		leftidx -= 1
		if strings.Contains(digits, string(s[leftidx])) {
			candidate = string(s[leftidx]) + candidate
		} else {
			leftDone = true
		}
	}
	rightDone := false
	rightidx := index
	for !rightDone && rightidx <= len(s)-2 {
		rightidx += 1
		if strings.Contains(digits, string(s[rightidx])) {
			candidate = candidate + string(s[rightidx])
		} else {
			rightDone = true
		}
	}
	t1, err1 := strconv.Atoi(candidate)
	if err1 != nil {
		return IntId{}, err
	}
	return IntId{t1, leftidx, rightidx}, nil
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
