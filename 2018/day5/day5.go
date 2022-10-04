package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var alpha = "abcdefghijklmnopqrstuvwxyz"

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	t := react(string(input))
	return len(t)
}

func part2() int {
	input, _ := os.ReadFile("input.txt")

	Alpha := strings.ToUpper(alpha)
	min := math.MaxInt
	for i, v := range alpha {
		s := string(input)
		toRemoveA := string(v)
		toRemoveB := string(Alpha[i])
		s = strings.ReplaceAll(s, toRemoveA, "")
		s = strings.ReplaceAll(s, toRemoveB, "")
		t := react(s)
		z := len(t)
		if z < min {
			min = z
		}
	}
	return min

}

func react(s string) string {
	Alpha := strings.ToUpper(alpha)
	oldLen := 0

	for oldLen != len(s) {
		oldLen = len(s)
		for i, v := range alpha {
			a := string(v) + string(Alpha[i])
			b := string(Alpha[i]) + string(v)
			s = strings.ReplaceAll(s, a, "")
			s = strings.ReplaceAll(s, b, "")
		}
	}
	return s
}
