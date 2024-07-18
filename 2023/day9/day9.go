package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input)))
}

func endState(line []int) bool {
	if len(line) == 0 {
		return false
	}
	for _, v := range line {
		if v != 0 {
			return false
		}
	}
	return true
}

func findNext(history []int) int {
	lines := [][]int{}
	lines = append(lines, history)

	for !endState(lines[len(lines)-1]) {
		newLine := []int{}
		line := lines[len(lines)-1]
		for index := range line[1:] {
			ix := index + 1
			value := line[ix] - line[ix-1]
			newLine = append(newLine, value)
		}
		lines = append(lines, newLine)
	}
	output := 0
	for _, v := range lines {
		output += v[len(v)-1]
	}
	return output
}

func findPrev(history []int) int {
	lines := [][]int{}
	lines = append(lines, history)

	for !endState(lines[len(lines)-1]) {
		newLine := []int{}
		line := lines[len(lines)-1]
		for index := range line[1:] {
			ix := index + 1
			value := line[ix] - line[ix-1]
			newLine = append(newLine, value)
		}
		lines = append(lines, newLine)
	}
	slices.Reverse(lines)
	output := 0
	for _, line := range lines {
		output = line[0] - output
	}
	return output
}

func part1(s string) int {
	data := strings.Split(strings.TrimSpace(s), "\n")
	output := 0
	for _, line := range data {
		history1 := strings.Fields(line)
		history := []int{}
		for _, value := range history1 {
			v, _ := strconv.Atoi(value)
			history = append(history, v)
		}
		output += findNext(history)
	}
	return output
}

func part2(s string) int {
	data := strings.Split(strings.TrimSpace(s), "\n")
	output := 0
	for _, line := range data {
		history1 := strings.Fields(line)
		history := []int{}
		for _, value := range history1 {
			v, _ := strconv.Atoi(value)
			history = append(history, v)
		}
		output += findPrev(history)
	}
	return output
}
