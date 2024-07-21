package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1:", sumDistances(string(input), 2))
	fmt.Println("Part 2:", sumDistances(string(input), 1000000))

}

type Point struct {
	row int
	col int
}

func sumDistances(s string, factor int) int {
	factor -= 1
	data := strings.Split(strings.TrimSpace(s), "\n")
	originalStars := []Point{}
	// find empty rows and columns
	emptyRows := []int{}
	emptyCols := []int{}
	colCounts := map[int]int{}
	for row, line := range data {
		if strings.Count(line, ".") == len(line) {
			emptyRows = append(emptyRows, row)
		}
		for col, v := range line {
			if v == '.' {
				colCounts[col] += 1
			}
			if v == '#' {
				originalStars = append(originalStars, Point{row, col})
			}
		}
	}
	for k, v := range colCounts {
		if v == len(data) {
			emptyCols = append(emptyCols, k)
		}
	}
	// we need to use these in order
	sort.Ints(emptyCols)
	sort.Ints(emptyRows)

	// increment star coordinates based on empty rows and cols
	stars := map[Point]Point{}
	for _, v := range originalStars {
		stars[v] = v
	}

	for _, emptyRow := range emptyRows {
		for _, location := range originalStars {
			if location.row > emptyRow {
				stars[location] = Point{stars[location].row + factor, stars[location].col}
			} else {
				stars[location] = Point{stars[location].row, stars[location].col}
			}
		}
	}

	for _, emptyCol := range emptyCols {
		for _, location := range originalStars {
			if location.col > emptyCol {
				stars[location] = Point{stars[location].row, stars[location].col + factor}
			} else {
				stars[location] = Point{stars[location].row, stars[location].col}
			}
		}
	}

	starList := []Point{}
	for _, v := range stars {
		starList = append(starList, v)
	}

	// now, get distances
	tot := 0
	for i, a := range starList {
		for _, b := range starList[i+1:] {
			tot += ManhattanDistance(a, b)
		}
	}
	return tot
}

func ManhattanDistance(a Point, b Point) int {
	y := a.row - b.row
	if y < 0 {
		y = -y
	}
	x := a.col - b.col
	if x < 0 {
		x = -x
	}
	return x + y
}
