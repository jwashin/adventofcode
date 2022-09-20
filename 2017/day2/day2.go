package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := makeSpreadsheet(string(input))
	fmt.Println("1.", checksum(data))
	fmt.Println("2.", checksum2(data))
}

func min(a []int) int {
	min := math.MaxInt
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func max(a []int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func difference(r []int) int {
	max := max(r)
	min := min(r)
	return max - min
}

func makeSpreadsheet(s string) [][]int {
	n := strings.Split(s, "\n")
	spreadsheet := [][]int{}
	for _, line := range n {
		newLine := []int{}
		for _, field := range strings.Split(line, "\t") {
			d, _ := strconv.Atoi(field)
			newLine = append(newLine, d)
		}
		spreadsheet = append(spreadsheet, newLine)
	}
	return spreadsheet
}

func checksum(s [][]int) int {
	cs := 0
	for _, v := range s {
		cs += difference(v)
	}
	return cs
}

func checksum2(s [][]int) int {
	cs := 0
	for _, v := range s {
		cs += xdivision(v)
	}
	return cs
}

func xdivision(r []int) int {
	for _, v := range r {
		for _, x := range r {
			if x != v {
				f := x % v
				if f == 0 {
					if x/v > 0 {
						return x / v
					}
				}
			}
		}
	}
	return 0
}
