package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input)))
}

var directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

func letterAt(row, col int, arr []string) byte {
	if row < 0 || col < 0 {
		return ' '
	}
	if row > len(arr)-1 || col > len(arr[0])-1 {
		return ' '
	}

	return arr[row][col]
}

func part1(input string) int {

	data := strings.Split(strings.TrimSpace(input), "\n")
	count := 0
	for row, s := range data {
		for col, item := range s {
			if item == 'X' {
				for _, direction := range directions {
					aString := stringAt(data, row, col, 4, direction)
					if aString == "XMAS" {
						count += 1
					}
				}

			}
		}
	}
	return count
}

func part2(input string) int {
	// 2035 too high
	//1027 too low
	data := strings.Split(strings.TrimSpace(input), "\n")
	count := 0
	for row, s := range data {
		for col, item := range s {
			if item == 'A' {
				if hasCrossMAS(data, row, col) {
					count += 1
				}
			}
		}
	}
	return count
}

func stringAt(arr []string, row, col, length int, direction []int) string {
	s := string(letterAt(row, col, arr))
	for len(s) < length {
		row += direction[0]
		col += direction[1]
		s += string(letterAt(row, col, arr))
	}
	return s
}

func hasCrossMAS(arr []string, row, col int) bool {
	// directions1 = [][]int{{-1, -1}, {1, 1}}}
	// directions2 = [][]int{{{-1, 1}, {1, -1}}

	tl := string(letterAt(row-1, col+1, arr))
	br := string(letterAt(row+1, col-1, arr))

	tr := string(letterAt(row-1, col-1, arr))
	bl := string(letterAt(row+1, col+1, arr))

	string1 := tl + "A" + br
	string2 := tr + "A" + bl

	return (string2 == "MAS" || string2 == "SAM") && (string1 == "MAS" || string1 == "SAM")
}
