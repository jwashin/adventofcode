package main

import (
	"fmt"
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
	c.getInstructions(string(input))
	c.run()
	return c.steps
}

func part2() int {
	input, _ := os.ReadFile("input.txt")
	c := puter{}
	c.getInstructions(string(input))
	c.run2()
	return c.steps
}

type puter struct {
	program []int
	steps   int
}

func test() int {
	c := puter{}
	c.program = []int{0, 3, 0, 1, -3}
	c.run()
	return c.steps
}
func test2() int {
	c := puter{}
	c.program = []int{0, 3, 0, 1, -3}
	c.run2()
	return c.steps
}

func (c *puter) run() {
	inst := 0
	c.steps = 0
	for inst < len(c.program) && inst >= 0 {
		c.steps += 1
		inst = c.doItem(inst)
	}
}

func (c *puter) run2() {
	inst := 0
	c.steps = 0
	for inst < len(c.program) && inst >= 0 {
		c.steps += 1
		inst = c.doItem2(inst)
	}
}

func (c *puter) getInstructions(s string) {
	p := strings.Split(s, "\n")
	t := []int{}
	for _, v := range p {
		z := strings.TrimSpace(v)
		if len(z) > 0 {
			d, _ := strconv.Atoi(z)
			t = append(t, d)
		}
	}
	c.program = t
}

func (c *puter) doItem(idx int) int {
	j := c.program[idx]
	c.program[idx] += 1
	return idx + j
}

func (c *puter) doItem2(idx int) int {
	j := c.program[idx]
	if j >= 3 {
		c.program[idx] -= 1
	} else {
		c.program[idx] += 1
	}
	return idx + j
}
