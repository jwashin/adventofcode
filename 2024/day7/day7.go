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

func part1(s string) int {

	data := strings.Split(strings.TrimSpace(s), "\n")
	output := 0
	for _, v := range data {
		output += getValue(v)
	}
	return output
}

func getValue(s string) int {
	input := strings.Fields(s)
	value, _ := strconv.Atoi(strings.Replace(input[0], ":", "", 1))
	operands := []int{}
	for _, v := range input[1:] {
		operand, _ := strconv.Atoi(v)
		operands = append(operands, operand)
	}
	operators := []string{"+", "*"}
	if isPossible(value, operands, operators) {
		return value
	}

	return 0
}

func isPossible(value int, operands []int, operators []string) bool {
	oldVals := []int{operands[0]}
	operands = operands[1:]

	for len(operands) > 0 {
		newValues := []int{}
		curr := operands[0]
		operands = operands[1:]
		for _, op := range operators {
			for _, v2 := range oldVals {
				if op == "*" {
					newValues = append(newValues, v2*curr)
				}
				if op == "+" {
					newValues = append(newValues, curr+v2)
				}
			}
		}
		oldVals = newValues

	}
	if slices.Contains(oldVals, value) {
		return true
	}

	return false
}
func part2(s string) int {

	data := strings.Split(strings.TrimSpace(s), "\n")
	output := 0
	for _, v := range data {
		output += getValue2(v)
	}
	return output
}

func getValue2(s string) int {
	input := strings.Fields(s)
	value, _ := strconv.Atoi(strings.Replace(input[0], ":", "", 1))
	operands := []int{}
	for _, v := range input[1:] {
		operand, _ := strconv.Atoi(v)
		operands = append(operands, operand)
	}
	operators := []string{"+", "*", "||"}
	if isPossible2(value, operands, operators) {
		return value
	}

	return 0
}

func isPossible2(value int, operands []int, operators []string) bool {
	oldVals := []int{operands[0]}
	operands = operands[1:]

	for len(operands) > 0 {
		newValues := []int{}
		curr := operands[0]
		operands = operands[1:]
		for _, op := range operators {
			for _, v2 := range oldVals {
				if op == "*" {
					newValues = append(newValues, v2*curr)
				}
				if op == "+" {
					newValues = append(newValues, curr+v2)
				}
				if op == "||" {
					currs := fmt.Sprint(curr)
					v2s := fmt.Sprint(v2)
					s2 := v2s + currs
					out, _ := strconv.Atoi(s2)
					newValues = append(newValues, out)
				}
			}
		}
		oldVals = newValues

	}
	if slices.Contains(oldVals, value) {
		return true
	}

	return false
}
