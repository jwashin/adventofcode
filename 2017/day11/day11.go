package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	return hexStepsAway(string(input))
}
func part2() int {
	input, _ := os.ReadFile("input.txt")
	return maxStepsAway(string(input))
}

func maxStepsAway(data string) int {
	path := strings.Split(strings.TrimSpace(data), ",")
	start := NewHex(0, 0)
	newHex := start
	max := 0
	for _, v := range path {
		newHex = HexNeighbor(newHex, getDirection(v))
		dist := HexDistance(newHex, start)
		if dist > max {
			max = dist
		}
	}
	return max

}

// https://github.com/pmcxs/hexgrid/blob/master/hex.go

type hex struct {
	q int // x axis
	r int // y axis
	s int // z axis
}

type direction int

func NewHex(q, r int) hex {

	h := hex{q: q, r: r, s: -q - r}
	return h

}

var directions = []hex{
	NewHex(1, 0),
	NewHex(1, -1),
	NewHex(0, -1),
	NewHex(-1, 0),
	NewHex(-1, +1),
	NewHex(0, +1),
}

func hexStepsAway(data string) int {
	path := strings.Split(strings.TrimSpace(data), ",")
	start := NewHex(0, 0)
	newHex := NewHex(0, 0)
	for _, v := range path {
		newHex = HexNeighbor(newHex, getDirection(v))
	}
	return HexDistance(newHex, start)

}

func getDirection(s string) direction {
	t := []string{"se", "ne", "n", "nw", "sw", "s"}
	idx := 0
	for k, v := range t {
		if v == s {
			idx = k
		}
	}
	return direction(idx)
}

// Subtracts two hexagons
func HexSubtract(a, b hex) hex {
	return NewHex(a.q-b.q, a.r-b.r)
}
func HexLength(hex hex) int {
	return int((math.Abs(float64(hex.q)) + math.Abs(float64(hex.r)) + math.Abs(float64(hex.s))) / 2.)
}

func HexDistance(a, b hex) int {
	sub := HexSubtract(a, b)
	return HexLength(sub)
}

// Returns the neighbor hexagon at a certain direction
func HexNeighbor(h hex, direction direction) hex {
	directionOffset := directions[direction]
	return NewHex(h.q+directionOffset.q, h.r+directionOffset.r)
}
