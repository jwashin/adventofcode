package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part1:", part1(string(input), 10))
	fmt.Println("part2:", part2(string(input)))
}

type coordinate struct {
	row    int
	column int
}

func (c coordinate) neighbors() []coordinate {
	out := []coordinate{}
	for row := c.row - 1; row <= c.row+1; row++ {
		for column := c.column - 1; column <= c.column+1; column++ {
			if !(column == c.column && row == c.row) {
				out = append(out, coordinate{row, column})
			}
		}
	}
	return out
}
func (c coordinate) echelon(s string) []coordinate {
	out := []coordinate{}
	if s == "N" {
		row := c.row - 1
		for column := c.column - 1; column <= c.column+1; column++ {
			out = append(out, coordinate{row, column})
		}
		return out
	}
	if s == "S" {
		row := c.row + 1
		for column := c.column - 1; column <= c.column+1; column++ {
			out = append(out, coordinate{row, column})
		}
		return out
	}

	if s == "E" {
		column := c.column + 1
		for row := c.row - 1; row <= c.row+1; row++ {
			out = append(out, coordinate{row, column})
		}
		return out
	}
	if s == "W" {
		column := c.column - 1
		for row := c.row - 1; row <= c.row+1; row++ {
			out = append(out, coordinate{row, column})
		}
		return out
	}

	return out
}

// set implementation
// https://linuxhint.com/golang-set/
type void struct{}

var member void

func existAny(elflocs map[coordinate]void, a []coordinate) bool {
	for _, v := range a {
		if _, ok := elflocs[v]; ok {
			return true
		}
	}
	return false
}

func calc(elfLocs map[coordinate]void) int {
	minx := math.MaxInt
	miny := math.MaxInt
	maxx := math.MinInt
	maxy := math.MinInt
	for v := range elfLocs {
		if v.column < minx {
			minx = v.column
		}
		if v.column > maxx {
			maxx = v.column
		}
		if v.row < miny {
			miny = v.row
		}
		if v.row > maxy {
			maxy = v.row
		}
	}
	area := (maxy - miny + 1) * (maxx - minx + 1)
	return area - len(elfLocs)

}

func part2(s string) int {

	checkOrder := "NSWE"

	elfLocs := map[coordinate]void{}
	for k, v := range strings.Split(s, "\n") {
		v = strings.ReplaceAll(v, "\t", "")
		for m, w := range v {
			if w == '#' {
				loc := coordinate{row: k, column: m}
				elfLocs[loc] = member
			}
		}
	}

	round := 0
	for {
		round += 1
		activeElves := map[coordinate]void{}
		for v := range elfLocs {
			nb := v.neighbors()
			if existAny(elfLocs, nb) {
				activeElves[v] = member
			}
		}

		// first half
		proposedMoves := [][2]coordinate{}
		proposedNewPositions := map[coordinate]int{}
		for v := range activeElves {
			for _, d := range checkOrder {
				direction := string(d)
				nb := v.echelon(direction)
				if !existAny(elfLocs, nb) {
					var pos coordinate
					if direction == "N" {
						pos = coordinate{row: v.row - 1, column: v.column}
					}
					if direction == "S" {
						pos = coordinate{row: v.row + 1, column: v.column}
					}
					if direction == "E" {
						pos = coordinate{row: v.row, column: v.column + 1}
					}
					if direction == "W" {
						pos = coordinate{row: v.row, column: v.column - 1}
					}
					proposedNewPositions[pos] += 1
					move := [2]coordinate{}
					move[0] = v
					move[1] = pos
					proposedMoves = append(proposedMoves, move)
					break
				}
			}
		}
		// second half
		moved := false
		for _, m := range proposedMoves {
			dest := m[1]
			orig := m[0]
			if proposedNewPositions[dest] == 1 {
				elfLocs[dest] = member
				delete(elfLocs, orig)
				moved = true
			}

		}
		d := checkOrder[0]
		checkOrder = checkOrder[1:] + string(d)
		if !moved {
			return round
		}
	}
}

func part1(s string, rounds int) int {

	checkOrder := "NSWE"

	elfLocs := map[coordinate]void{}
	for k, v := range strings.Split(s, "\n") {
		v = strings.ReplaceAll(v, "\t", "")
		for m, w := range v {
			if w == '#' {
				loc := coordinate{row: k, column: m}
				elfLocs[loc] = member
			}
		}
	}

	round := 0
	for round < rounds {
		round += 1
		activeElves := map[coordinate]void{}
		for v := range elfLocs {
			nb := v.neighbors()
			if existAny(elfLocs, nb) {
				activeElves[v] = member
			}
		}
		// first half
		proposedMoves := [][2]coordinate{}
		proposedNewPositions := map[coordinate]int{}
		for v := range activeElves {
			for _, d := range checkOrder {
				direction := string(d)
				nb := v.echelon(direction)
				if !existAny(elfLocs, nb) {
					var pos coordinate
					if direction == "N" {
						pos = coordinate{row: v.row - 1, column: v.column}
					}
					if direction == "S" {
						pos = coordinate{row: v.row + 1, column: v.column}
					}
					if direction == "E" {
						pos = coordinate{row: v.row, column: v.column + 1}
					}
					if direction == "W" {
						pos = coordinate{row: v.row, column: v.column - 1}
					}
					proposedNewPositions[pos] += 1
					move := [2]coordinate{}
					move[0] = v
					move[1] = pos
					proposedMoves = append(proposedMoves, move)
					break
				}
			}
		}
		// second half
		for _, m := range proposedMoves {
			dest := m[1]
			orig := m[0]
			if proposedNewPositions[dest] == 1 {
				elfLocs[dest] = member
				delete(elfLocs, orig)
			}
		}
		d := checkOrder[0]
		checkOrder = checkOrder[1:] + string(d)
	}
	return calc(elfLocs)
}
