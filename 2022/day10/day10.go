package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", part1(string(input)))
	fmt.Println("part 2.")
	fmt.Println(drawScreen(string(input)))
}

func part1(s string) int {
	gt := 20
	inc := 40
	total := 0
	max := 220
	for gt <= max {
		j := signalStrength(s, gt)
		total += j
		gt += inc
	}
	return total
}

// not proud of this.
// getX and signalStrength should work identically,
// but they don't

func drawScreen(s string) string {
	cycle := -1
	out := []string{}
	ns := ""
	for cycle <= 240 {
		cycle += 1
		crs := getX(s, cycle)
		px := cycle % 40
		if crs == px-1 || crs == px || crs == px+1 {
			ns += "#"
		} else {
			ns += "."
		}
		if len(ns) == 40 {
			out = append(out, ns)
			ns = ""
		}
	}
	return strings.Join(out, "\n")

}

func getX(s string, c int) int {
	data := strings.Split(s, "\n")
	x := 1
	cycle := 1
	for _, v := range data {
		f := strings.Fields(v)
		if f[0] == "noop" {
			cycle += 1
			if cycle > c {
				return x
			}
		}
		if f[0] == "addx" {
			cycle += 1
			if cycle > c {
				return x
			}
			cycle += 1
			a, _ := strconv.Atoi(f[1])
			x += a
			if cycle > c {
				return x
			}
		}
	}
	return 0
}

func signalStrength(s string, c int) int {
	data := strings.Split(s, "\n")
	x := 1
	cycle := 1
	for _, v := range data {
		f := strings.Fields(v)
		if f[0] == "noop" {
			cycle += 1
			if cycle == c {
				return c * x
			}
		}
		if f[0] == "addx" {
			cycle += 1
			if cycle == c {
				return c * x
			}
			cycle += 1
			a, _ := strconv.Atoi(f[1])
			x += a
			if cycle == c {
				return c * x
			}
		}
	}
	return 0
}
