package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input)))
}

func part1(s string) int {

	ttrim := strings.TrimSpace(s)
	ssplit := strings.Split(ttrim, "\n")
	result := 0
	for _, v := range ssplit {
		levels := []int{}
		report := strings.Fields(v)
		for _, v1 := range report {
			level, _ := strconv.Atoi(v1)
			levels = append(levels, level)
		}
		if issafe(levels) {
			result += 1
		}
	}
	return result
}

func part2(s string) int {

	ttrim := strings.TrimSpace(s)
	ssplit := strings.Split(ttrim, "\n")
	result := 0
	for _, v := range ssplit {
		levels := []int{}
		report := strings.Fields(v)
		for _, v1 := range report {
			level, _ := strconv.Atoi(v1)
			levels = append(levels, level)
		}
		if issafe(levels) || anysafe(levels) {
			result += 1
		}
	}
	return result
}

func anysafe(a []int) bool {

	for i := range a {
		templist := []int{}

		for j, v := range a {
			if j != i {
				templist = append(templist, v)
			}
		}
		if issafe(templist) {
			return true
		}
	}
	return false
}

func issafe(a []int) bool {
	currdirection := ""
	if a[1] < a[0] {
		currdirection = "-"
	} else {
		currdirection = "+"
	}
	for k, v := range a[1:] {
		index := a[k]
		next := v
		diff, direction := compare(index, next)
		if currdirection != direction {
			return false
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func compare(a, b int) (diff int, direction string) {
	if a < b {
		direction = "+"
	} else {
		direction = "-"
	}
	diff = abs(a - b)
	return diff, direction
}
