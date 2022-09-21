package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	c := puter{}
	c.program = strings.Split(string(input), "\n")
	c.run()
	return c.highestRegister()
}

func part2() int {
	input, _ := os.ReadFile("input.txt")
	c := puter{}
	c.program = strings.Split(string(input), "\n")
	c.run()
	return c.maxVal
}

type puter struct {
	register map[string]int
	program  []string
	maxVal   int
}

func test() int {
	c := puter{}
	c.program = []string{"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10"}
	c.run()
	return c.highestRegister()
}
func test2() int {
	c := puter{}
	c.program = []string{"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10"}
	c.run()
	return c.maxVal
}

func (c *puter) run() {
	c.register = map[string]int{}
	c.maxVal = -math.MaxInt
	for _, instruction := range c.program {
		a := strings.Fields(instruction)
		// condition
		p1 := c.register[a[4]]
		p2, _ := strconv.Atoi(a[6])
		cond := false
		if a[5] == ">" && p1 > p2 {
			cond = true
		}
		if a[5] == "<" && p1 < p2 {
			cond = true
		}
		if a[5] == ">=" && p1 >= p2 {
			cond = true
		}
		if a[5] == "<=" && p1 <= p2 {
			cond = true
		}
		if a[5] == "==" && p1 == p2 {
			cond = true
		}
		if a[5] == "!=" && p1 != p2 {
			cond = true
		}
		if cond {
			p3, _ := strconv.Atoi(a[2])
			if a[1] == "inc" {
				c.register[a[0]] += p3
			}
			if a[1] == "dec" {
				c.register[a[0]] -= p3
			}
			if c.register[a[0]] > c.maxVal {
				c.maxVal = c.register[a[0]]
			}
		}

	}
}

func (c *puter) highestRegister() int {
	max := -math.MaxInt
	for _, v := range c.register {
		if v > max {
			max = v
		}
	}
	return max
}
