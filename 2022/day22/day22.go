package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
}

func getData(s string) ([]string, string) {
	tableau := []string{}
	instruct := ""
	gotTableau := false
	maxlen := 0
	nt := []string{}
	for _, v := range strings.Split(s, "\n") {
		v = strings.ReplaceAll(v, "\t", "")
		if v == "" {
			gotTableau = true
			continue
		}
		if gotTableau {
			instruct = v
			continue
		}

		tableau = append(tableau, v)
		if len(v) > maxlen {
			maxlen = len(v)
		}
	}
	for _, v := range tableau {
		for len(v) < maxlen {
			v = v + " "
		}
		nt = append(nt, v)
	}
	return nt, instruct
}

type coordinate struct {
	row    int
	column int
}

func turn(currDir string, instruction string) string {
	if instruction == "R" {
		if currDir == ">" {
			return "v"
		}
		if currDir == "v" {
			return "<"
		}
		if currDir == "<" {
			return "^"
		}
		if currDir == "^" {
			return ">"
		}
	} else if instruction == "L" {
		if currDir == ">" {
			return "^"
		}
		if currDir == "v" {
			return ">"
		}
		if currDir == "<" {
			return "v"
		}
		if currDir == "^" {
			return "<"
		}
	}
	return ""
}

func move(t []string, c coordinate, direction string, distance string) coordinate {
	dist, _ := strconv.Atoi(distance)
	currLoc := c
	for dist > 0 {
		dist -= 1
		newLoc, err := moveOne(t, currLoc, direction)
		if err != nil {
			break
		}
		currLoc = newLoc
	}
	return currLoc
}

func moveOne(t []string, c coordinate, direction string) (coordinate, error) {
	var newLoc coordinate
	var err error
	if direction == ">" {
		newLoc, err = moveTo(t, coordinate{c.row, c.column + 1}, direction)
	}
	if direction == "<" {
		newLoc, err = moveTo(t, coordinate{c.row, c.column - 1}, direction)
	}
	if direction == "^" {
		newLoc, err = moveTo(t, coordinate{c.row - 1, c.column}, direction)
	}
	if direction == "v" {
		newLoc, err = moveTo(t, coordinate{c.row + 1, c.column}, direction)
	}
	return newLoc, err
}

func moveTo(t []string, c coordinate, direction string) (coordinate, error) {
	if direction == ">" && c.column > len(t[c.row])-1 {
		c.column = 0
	}
	if direction == "<" && c.column < 0 {
		c.column = len(t[c.row]) - 1
	}
	if direction == "v" && c.row > len(t)-1 {
		c.row = 0
	}
	if direction == "^" && c.row < 0 {
		c.row = len(t) - 1
	}
	glyph := string(t[c.row][c.column])
	if glyph == "." {
		return c, nil
	}
	if glyph == "#" {
		return c, errors.New("hit a wall")
	}
	if direction == ">" {
		row := t[c.row]
		currIndex := c.column
		for k := currIndex + 1; k < len(row); k++ {
			if row[k] == '.' {
				return coordinate{c.row, k}, nil
			}
			if row[k] == '#' {
				return c, errors.New("hit a wall")
			}
		}
		for k := 0; k < currIndex; k++ {
			if row[k] == '#' {
				return c, errors.New("hit a wall")
			}
			if row[k] == '.' {
				return coordinate{c.row, k}, nil
			}
		}
	}
	if direction == "<" {
		row := t[c.row]
		currIndex := c.column
		for k := currIndex - 1; k >= 0; k-- {
			if row[k] == ' ' {
				continue
			}
			if row[k] == '.' {
				return coordinate{c.row, k}, nil
			}
			if row[k] == '#' {
				return c, errors.New("hit a wall")
			}
		}
		for k := len(row) - 1; k > currIndex; k-- {
			if row[k] == ' ' {
				continue
			}
			if row[k] == '#' {
				return c, errors.New("hit a wall")
			}
			if row[k] == '.' {
				return coordinate{c.row, k}, nil
			}
		}
	}

	if direction == "v" {
		row := ""
		for _, v := range t {
			row += string(v[c.column])
		}
		currIndex := c.column
		for k := currIndex + 1; k < len(row); k++ {
			if row[k] == ' ' {
				continue
			}
			if row[k] == '#' {
				return c, errors.New("hit a wall")
			}
			if row[k] == '.' {
				return coordinate{k, c.column}, nil
			}
		}
		for k := 0; k < currIndex; k++ {
			if row[k] == ' ' {
				continue
			}
			if row[k] == '#' {
				return c, errors.New("hit a wall")
			}
			if row[k] == '.' {
				return coordinate{k, c.column}, nil
			}
		}
	}

	if direction == "^" {
		row := ""
		for _, v := range t {
			row += string(v[c.column])
		}
		currIndex := c.column
		for k := currIndex - 1; k >= 0; k-- {
			if row[k] == ' ' {
				continue
			}
			if row[k] == '#' {
				return c, errors.New("hit a wall")
			}
			if row[k] == '.' {
				return coordinate{k, c.column}, nil
			}
		}
		for k := len(row) - 1; k > currIndex; k-- {
			if row[k] == ' ' {
				continue
			}
			if row[k] == '#' {
				return c, errors.New("hit a wall")
			}
			if row[k] == '.' {
				return coordinate{k, c.column}, nil
			}
		}
	}

	return c, errors.New("invalid direction")
}

// 116156 too high
// 78360 too high

func part1(s string) int {
	tableau, instructions := getData(s)
	start := coordinate{0, 0}
	for k, v := range tableau[0] {
		if v == '.' {
			start = coordinate{0, k}
			break
		}
	}
	currLoc := start
	direction := ">"
	dist := ""
	for len(instructions) > 0 {
		t := instructions[0]
		instructions = instructions[1:]
		if t == 'R' || t == 'L' {
			if len(dist) > 0 {
				// move
				currLoc = move(tableau, currLoc, direction, dist)
				dist = ""
			}
			direction = turn(direction, string(t))
		} else {
			dist += string(t)
		}
	}
	// final move, if any
	if len(dist) > 0 {
		currLoc = move(tableau, currLoc, direction, dist)
	}

	facing := map[string]int{">": 0, "v": 1, "<": 2, "^": 3}
	return (currLoc.row+1)*1000 + (currLoc.column+1)*4 + facing[direction]
}
