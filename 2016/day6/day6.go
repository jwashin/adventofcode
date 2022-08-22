package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println(maxRepeatsByColumn(string(input)))
	fmt.Println(leastRepeatsByColumn(string(input)))
}

func maxRepeatsByColumn(a string) string {
	out := ""
	data := strings.Split(a, "\n")

	for idx := range data[0] {
		td := map[string]int{}
		for _, row := range data {
			td[string(row[idx])] += 1
		}
		max := 0
		maxChar := ""
		for k, v := range td {
			if v > max {
				max = v
				maxChar = k
			}
		}
		out = out + maxChar
	}
	return out
}
func leastRepeatsByColumn(a string) string {
	out := ""
	data := strings.Split(a, "\n")

	for idx := range data[0] {
		td := map[string]int{}
		for _, row := range data {
			td[string(row[idx])] += 1
		}
		min := 99
		minChar := ""
		for k, v := range td {
			if v < min {
				min = v
				minChar = k
			}
		}
		out = out + minChar
	}
	return out
}
