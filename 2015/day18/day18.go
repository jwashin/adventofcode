package main

import (
	"fmt"
	"os"
	"strings"
)

// Conway's Life

func main() {
	fmt.Println(part1())
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	tableau := string(input)
	t2 := life(tableau, 100)
	return strings.Count(t2, "#")
}

func life(s string, iter int) string {
	out := []string{}
	tableau := initCells(s)
	i := 0
	for i < iter {
		tableau.iter()
		i += 1
	}
	for _, line := range tableau {
		s = ""
		for _, cell := range line {
			if cell.state {
				s = s + "#"
				continue
			}
			s = s + "."
		}
		out = append(out, s)
	}
	return strings.Join(out, "\n")
}

type cell struct {
	state     bool
	nextState bool
}

type cells [][]*cell

func (c cells) isLive(y, x int) bool {
	if y >= 0 && y <= len(c)-1 {
		if x >= 0 && x <= len(c[0])-1 {
			return c[y][x].state
		}
	}
	return false
}

func (c cells) countLiveNeighbors(y0, x0 int) int {
	count := 0
	for y := y0 - 1; y <= y0+1; y++ {
		for x := x0 - 1; x <= x0+1; x++ {
			// the below line gave me some heartburn. Logic!
			if !(x == x0 && y == y0) {
				if c.isLive(y, x) {
					count += 1
				}
			}
		}
	}

	return count
}

func (c cells) fix4Cells() {
	// this is just for part 2.
	// it's used in initCells() and iter()
	c[0][0].state = true
	c[0][len(c[0])-1].state = true
	c[len(c)-1][0].state = true
	c[len(c)-1][len(c[0])-1].state = true
}

func (c cells) iter() {
	for y, line := range c {
		for x, cell := range line {
			t := c.countLiveNeighbors(y, x)
			if cell.state {
				if t == 2 || t == 3 {
					cell.nextState = true
				} else {
					cell.nextState = false
				}
			}
			if !cell.state {
				if t == 3 {
					cell.nextState = true
				} else {
					cell.nextState = false
				}
			}
		}
	}
	for _, line := range c {
		for _, cell := range line {
			cell.state = cell.nextState
		}
	}
	// uncomment the next line for part2
	// c.fix4Cells()
}

func initCells(s string) cells {
	tableau := cells{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			tableauLine := []*cell{}
			for _, character := range line {
				c := cell{state: false}
				if character == '#' {
					c.state = true
				}
				tableauLine = append(tableauLine, &c)
			}
			tableau = append(tableau, tableauLine)
		}
	}
	// uncomment the next line for part2
	// tableau.fix4Cells()
	return tableau
}
