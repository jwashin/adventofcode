package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	ic := diagnose(string(input), 10000)
	return ic
}

// 2508638 too low

func part2() int {
	input, _ := os.ReadFile("input.txt")
	ic := diagnose2(string(input), 10000000)
	return ic
}

type Grid map[string]string

func (g Grid) clean(y, x int) {
	delete(g, c2s(y, x))
}

func (g Grid) infect(y, x int) {
	g[c2s(y, x)] = "#"
}

func (g Grid) isInfected(y, x int) bool {
	return g[c2s(y, x)] == "#"
}

func (g Grid) weaken(y, x int) {
	g[c2s(y, x)] = "W"
}

func (g Grid) flag(y, x int) {
	g[c2s(y, x)] = "F"
}

func (g Grid) getItem(y, x int) string {
	return g[c2s(y, x)]
}

func c2s(y, x int) string {
	return fmt.Sprintf("%d %d", y, x)
}

type carrier struct {
	x         int
	y         int
	direction string
}

func (c *carrier) turnRight() {
	d := map[string]string{
		"n": "e",
		"e": "s",
		"s": "w",
		"w": "n",
	}
	c.direction = d[c.direction]
}
func (c *carrier) turnLeft() {
	d := map[string]string{
		"n": "w",
		"e": "n",
		"s": "e",
		"w": "s",
	}
	c.direction = d[c.direction]
}
func (c *carrier) reverse() {
	d := map[string]string{
		"n": "s",
		"e": "w",
		"s": "n",
		"w": "e",
	}
	c.direction = d[c.direction]
}
func (c *carrier) move() {
	if c.direction == "n" {
		c.y -= 1
	}
	if c.direction == "s" {
		c.y += 1
	}
	if c.direction == "e" {
		c.x += 1
	}
	if c.direction == "w" {
		c.x -= 1
	}
}

func iter2(g *Grid, c *carrier) int {
	infectionCount := 0
	y, x := c.y, c.x
	item := g.getItem(y, x)
	switch item {
	case "":
		c.turnLeft()
		g.weaken(y, x)
	case "#":
		c.turnRight()
		g.flag(y, x)
	case "F":
		c.reverse()
		g.clean(y, x)
	case "W":
		g.infect(y, x)
		infectionCount += 1
	}
	c.move()
	return infectionCount
}

func iter(g *Grid, c *carrier) int {
	infectionCount := 0
	y, x := c.y, c.x
	if g.isInfected(y, x) {
		c.turnRight()
	} else {
		c.turnLeft()
	}
	if g.isInfected(y, x) {
		g.clean(y, x)
	} else {
		g.infect(y, x)
		infectionCount += 1
	}
	c.move()
	return infectionCount
}

func inits(s string) (carrier, Grid) {
	t := strings.Split(strings.TrimSpace(s), "\n")
	center := len(t[0]) / 2
	c := carrier{y: center, x: center, direction: "n"}
	g := Grid{}
	for y, v := range t {
		for x, i := range v {
			if i == '#' {
				g[c2s(y, x)] = "#"
			}
		}
	}
	return c, g
}

func diagnose2(start string, activityCount int) int {
	c, g := inits(start)
	infections := 0
	ac := 0
	for ac < activityCount {

		i := iter2(&g, &c)
		ac += 1

		infections += i
	}
	return infections
}

func diagnose(start string, activityCount int) int {
	c, g := inits(start)
	infections := 0
	ac := 0
	for ac < activityCount {
		i := iter(&g, &c)
		ac += 1
		infections += i
	}
	return infections
}
