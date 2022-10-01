package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", sinst())
}

func part2() int {
	c := puter{}
	c.registers = map[string]int{"a": 1}
	input, _ := os.ReadFile("input.txt")
	c.program = getInstructions(string(input))
	c.run1()
	return c.registers["h"]
}

func part1() int {
	c := puter{}
	c.registers = map[string]int{}
	input, _ := os.ReadFile("input.txt")
	c.program = getInstructions(string(input))
	t := c.run1()
	return t
}

type Instruction struct {
	name string
	args []string
}

func getInstructions(s string) []Instruction {
	p := strings.Split(s, "\n")
	t := []Instruction{}
	for _, z := range p {
		if len(z) > 0 {
			x := strings.Fields(z)
			i := Instruction{name: x[0], args: x[1:]}
			t = append(t, i)
		}
	}
	return t
}

func isNum(s string) bool {
	if s[0] == '-' {
		s = s[1:]
	}
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

type puter struct {
	registers map[string]int
	program   []Instruction
}

func (c puter) run1() int {
	idx := 0
	mul := 0
	var isMul bool
	for idx >= 0 && idx < len(c.program) {
		idx, isMul = c.doItem(idx)
		if isMul {
			mul += 1
		}
	}
	return mul
}

func (c *puter) doItem(idx int) (int, bool) {
	instruction := c.program[idx]
	ins := instruction.name
	mul := false
	if ins == "set" {
		b := instruction.args[1]
		val := 0
		if isNum(b) {
			val, _ = strconv.Atoi(b)
		} else {
			val = c.registers[b]
		}
		c.registers[instruction.args[0]] = val
	}
	if ins == "sub" {
		b := instruction.args[1]
		val := 0
		if isNum(b) {
			val, _ = strconv.Atoi(b)
		} else {
			val = c.registers[b]
		}
		c.registers[instruction.args[0]] -= val
	}
	if ins == "inc" {
		b := instruction.args[1]
		val := 0
		if isNum(b) {
			val, _ = strconv.Atoi(b)
		} else {
			val = c.registers[b]
		}
		c.registers[instruction.args[0]] += val
	}
	if ins == "mul" {
		x := 0
		mul = true
		v0 := instruction.args[0]
		if isNum(v0) {
			x, _ = strconv.Atoi(v0)
		} else {
			x = c.registers[v0]
		}
		y := 0
		v1 := instruction.args[1]
		if isNum(v1) {
			y, _ = strconv.Atoi(v1)
		} else {
			y = c.registers[v1]
		}
		c.registers[instruction.args[0]] = x * y
	}
	if ins == "jnz" {
		x := 0
		v0 := instruction.args[0]
		if isNum(v0) {
			x, _ = strconv.Atoi(v0)
		} else {
			x = c.registers[v0]
		}

		y := 0
		v1 := instruction.args[1]
		if isNum(v1) {
			y, _ = strconv.Atoi(v1)
		} else {
			y = c.registers[v1]
		}
		if x != 0 {
			return idx + y, mul
		}
	}
	return idx + 1, mul
}

// ok. I got help at
// https://www.reddit.com/r/adventofcode/comments/7lms6p/2017_day_23_solutions/

// TODO: there was an excellent tutorial on how to
// assem to code at the above link.

func sinst() int {
	var a, b, c, f, h int
	a = 1
	b = 57
	c = b
	if a != 0 {
		b = b*100 + 100000
		c = b + 17000
		// set g to something; check if it is zero later
		for b-c != 0 {
			f = 1
			// somehow, I was supposed to figure
			// that this was a prime sieve
			for d := 2; d*d <= b; d++ {
				// composite
				if b%d == 0 {
					f = 0
					break
				}
			}
			if f == 0 {
				h++
			}
			// g = b - c
			b += 17
		}
	}
	return h
}
