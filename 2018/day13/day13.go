package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// s := part1(false)
	// fmt.Println("1.", s)
	fmt.Println("2.", part2(false))
}

func part2(test bool) string {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	if test {
		// 	data = strings.Join([]string{"|", "v", "|", "|", "|", "^", "|"}, "\n")
		// }
		data = `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`
	}
	// initialize grid
	grid := stringGrid{}
	data = strings.ReplaceAll(data, "\t", "")
	// data = strings.ReplaceAll(data, " ", ".")
	for _, v := range strings.Split(data, "\n") {
		grid = append(grid, v)
	}
	// find carts
	c := carts{}
	for j, row := range grid {
		for i := range row {
			loc := coordinate{x: i, y: j}
			v := grid.getItem(loc)
			var item cart
			if v == "^" {
				item = cart{location: loc, currentDirection: "north", nextOption: "left", active: true}
				c = append(c, &item)
				grid.setItem(loc, "|")
			}
			if v == "v" {
				item = cart{location: loc, currentDirection: "south", nextOption: "left", active: true}
				c = append(c, &item)
				grid.setItem(loc, "|")
			}
			if v == ">" {
				item = cart{location: loc, currentDirection: "east", nextOption: "left", active: true}
				c = append(c, &item)
				grid.setItem(loc, "-")
			}
			if v == "<" {
				item = cart{location: loc, currentDirection: "west", nextOption: "left", active: true}
				c = append(c, &item)
				grid.setItem(loc, "-")
			}
		}
	}
	// have carts, c and grid, g. Now make them move and detect crashes
	iteration := 0
	for {
		c.sort()
		fmt.Println(iteration)
		// printout(grid, c)
		iteration += 1
		for _, v := range c {
			newloc, dir := grid.nextCoordinate(v)
			v.location = newloc
			v.currentDirection = dir
			if c.detectCrash() {
				for _, n := range c {
					if n.location.String() == v.location.String() {
						v.active = false
						n.active = false
					}
				}
			}
		}
		if c.activeCount() == 1 {
			for _, z := range c {
				if z.active {
					return z.location.String()
				}
			}
		}
	}
}

func part1(test bool) string {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	if test {
		// 	data = strings.Join([]string{"|", "v", "|", "|", "|", "^", "|"}, "\n")
		// }
		data =
			`/->-\
	|   |  /----\
	| /-+--+-\  |
	| | |  | v  |
	\-+-/  \-+--/
	  \------/  `
	}
	// initialize grid
	grid := stringGrid{}
	data = strings.ReplaceAll(data, "\t", "")
	// data = strings.ReplaceAll(data, " ", ".")
	for _, v := range strings.Split(data, "\n") {
		grid = append(grid, v)
	}
	// find carts
	c := carts{}
	for j, row := range grid {
		for i := range row {
			loc := coordinate{x: i, y: j}
			v := grid.getItem(loc)
			var item cart
			if v == "^" {
				item = cart{location: loc, currentDirection: "north", nextOption: "left", active: true}
				c = append(c, &item)
				grid.setItem(loc, "|")
			}
			if v == "v" {
				item = cart{location: loc, currentDirection: "south", nextOption: "left", active: true}
				c = append(c, &item)
				grid.setItem(loc, "|")
			}
			if v == ">" {
				item = cart{location: loc, currentDirection: "east", nextOption: "left", active: true}
				c = append(c, &item)
				grid.setItem(loc, "-")
			}
			if v == "<" {
				item = cart{location: loc, currentDirection: "west", nextOption: "left", active: true}
				c = append(c, &item)
				grid.setItem(loc, "-")
			}
		}
	}
	// have carts, c and grid, g. Now make them move and detect crashes
	iteration := 0
	for {
		c.sort()
		fmt.Println(iteration)
		printout(grid, c)
		iteration += 1
		for _, v := range c {
			if v.active {
				newloc, dir := grid.nextCoordinate(v)
				v.location = newloc
				v.currentDirection = dir
				if c.detectCrash() {
					return v.location.String()
				}
			}
		}
	}
}

func printout(g stringGrid, c carts) {
	out := stringGrid{}
	out = append(out, g...)

	for _, v := range c {
		if v.active {
			ch := ""
			if v.currentDirection == "north" {
				ch = "^"
			} else if v.currentDirection == "south" {
				ch = "v"

			} else if v.currentDirection == "east" {
				ch = ">"
			} else if v.currentDirection == "west" {
				ch = "<"
			}
			out.setItem(v.location, ch)
		}
	}
	for _, v := range out {
		fmt.Println(v)
	}

}

type coordinate struct {
	y int
	x int
}

func (c coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

type cart struct {
	location         coordinate
	currentDirection string
	nextOption       string
	active           bool
	// lastChar         string
}

type carts []*cart

func (c carts) sort() {
	sort.Slice(c, func(i, j int) bool {
		if c[i].location.y == c[j].location.y {
			return c[i].location.x < c[j].location.x
		} else {
			return c[i].location.y < c[j].location.y
		}
	})
}

func (c carts) detectCrash() bool {
	for k, i := range c {
		for m, j := range c {
			if i.active && j.active && k != m && i.location.x == j.location.x && i.location.y == j.location.y {
				return true
			}
		}
	}
	return false
}
func (c carts) activeCount() int {
	count := 0
	for _, i := range c {
		for _, j := range c {
			if i.active && j.active {
				count += 1
			}
		}
	}
	return count
}

type stringGrid []string

func (g stringGrid) nextCoordinate(c *cart) (coordinate, string) {
	neighbors := g.neighbors(c.location)
	reverses := map[string]string{"north": "south", "west": "east", "south": "north", "east": "west"}

	// remove reverses and blanks from neighbors
	for k, v := range neighbors {
		if string(g.getItem(v)) == " " || k == reverses[c.currentDirection] {
			delete(neighbors, k)
		}
	}
	newDirection := ""
	currCharacter := string(g.getItem(c.location))

	if currCharacter == "+" {
		left := map[string]string{"north": "west", "west": "south", "south": "east", "east": "north"}
		right := map[string]string{"north": "east", "west": "north", "south": "west", "east": "south"}
		if c.nextOption == "left" {
			newDirection = left[c.currentDirection]
			c.nextOption = "straight"
			return neighbors[newDirection], newDirection
		} else if c.nextOption == "right" {
			newDirection = right[c.currentDirection]
			c.nextOption = "left"
			return neighbors[newDirection], newDirection
		} else {
			newDirection = c.currentDirection
			c.nextOption = "right"
			return neighbors[newDirection], newDirection
		}
	}
	if currCharacter == "\\" {
		newdirs := map[string]string{"east": "south", "north": "west", "west": "north", "south": "east"}
		newDirection = newdirs[c.currentDirection]
		return neighbors[newDirection], newDirection
	}
	if currCharacter == "/" {
		newdirs := map[string]string{"east": "north", "north": "east", "west": "south", "south": "west"}
		newDirection = newdirs[c.currentDirection]
		return neighbors[newDirection], newDirection
	}

	for k := range neighbors {
		if c.currentDirection == k {
			return neighbors[k], k
		}
	}

	var out1 coordinate
	var out2 string
	for k, v := range neighbors {
		out1, out2 = v, k
	}
	return out1, out2
}

func (g stringGrid) getItem(c coordinate) string {
	if c.y < len(g) && c.x < len(g[c.y]) {
		return string(g[c.y][c.x])
	}
	return " "
}
func (g stringGrid) setItem(c coordinate, s string) {
	t := g[c.y]
	x := c.x
	n := []string{}
	for k, v := range t {
		if k == x {
			n = append(n, s)
		} else {
			n = append(n, string(v))
		}
	}
	g[c.y] = strings.Join(n, "")
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
	if x < len(g[y])-1 {
		nb["east"] = coordinate{y: y, x: x + 1}
	}
	return nb
}
