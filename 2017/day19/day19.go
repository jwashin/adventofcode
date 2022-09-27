package main

import (
	"fmt"
	"os"
	"strings"
)

var alpha = "ABCDEFGHIJKLMNOPQRSTUVWYYZ"

func main() {
	s, d := part1(false)
	fmt.Println("1.", s)
	fmt.Println("2.", d)
}

func part1(test bool) (string, int) {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	if test {
		data = `        |          
		|  +--+    
		A  |  C    
	F---|----E|--+ 
		|  |  |  D 
		+B-+  +--+ `
	}
	outdata := ""

	grid := stringGrid{}
	data = strings.ReplaceAll(data, "\t", "    ")
	data = strings.ReplaceAll(data, " ", ".")
	for _, v := range strings.Split(data, "\n") {
		grid = append(grid, v)
	}
	currentCoordinate := coordinate{0, strings.Index(grid[0], "|")}
	direction := "south"
	n, d := grid.nextCoordinate(currentCoordinate, direction)
	steps := 1
	for d != "" {
		steps += 1
		n, d = grid.nextCoordinate(n, d)
		i := string(grid.getItem(n))
		if strings.Contains(alpha, i) {
			outdata += i
		}
	}
	return outdata, steps
}

type coordinate struct {
	y int
	x int
}

type stringGrid []string

func (g stringGrid) nextCoordinate(c coordinate, currDirection string) (coordinate, string) {
	nb := g.neighbors(c)
	for k, v := range nb {
		if string(g.getItem(v)) == "." {
			delete(nb, k)
		}
	}
	currCharacter := string(g.getItem(c))
	opps := [][]string{{"north", "south"}, {"east", "west"}}
	for _, v := range opps {
		if currDirection == v[0] {
			delete(nb, v[1])
		}
		if currDirection == v[1] {
			delete(nb, v[0])
		}
	}
	if len(nb) == 0 {
		return coordinate{}, ""
	}

	if currCharacter == "+" {
		for k, v := range nb {
			return v, k
		}
	}

	// default go the same direction

	return nb[currDirection], currDirection

}

func (g stringGrid) getItem(c coordinate) string {
	return string(g[c.y][c.x])
}

func (g stringGrid) neighbors(c coordinate) map[string]coordinate {
	nb := map[string]coordinate{}
	y := c.y
	x := c.x
	// north
	if y > 0 {
		nb["north"] = coordinate{y: y - 1, x: x}
	}
	// west
	if x > 0 {
		nb["west"] = coordinate{y: y, x: x - 1}
	}
	// south
	if y < len(g)-1 {
		nb["south"] = coordinate{y: y + 1, x: x}
	}
	// east
	if x < len(g[0])-1 {
		nb["east"] = coordinate{y: y, x: x + 1}
	}
	return nb
}
