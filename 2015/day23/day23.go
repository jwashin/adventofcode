package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	c := puter{}
	c.registers = map[string]uint{"a": 0, "b": 0}
	c.getInstructions(string(data))
	c.run()
	fmt.Println("1.", c.registers["b"])
	// c.getInstructions(string(data))
	// c.registers = map[string]int{"a": 12, "b": 0, "c": 0, "d": 0}
	// c.run()
	// fmt.Println("2.", c.registers["a"])
}

// part2 6096 too low

func test(s string) uint {
	c := puter{}
	c.registers = map[string]uint{}
	c.getInstructions(s)
	c.run()
	return c.registers["a"]
}

type puter struct {
	registers map[string]uint
	program   []Instruction
}

type Instruction struct {
	name string
	args []string
}

func (c puter) run() {
	inst := 0
	for inst < len(c.program) && inst > 0 {
		inst = c.doItem(inst)
	}
}

func (c *puter) getInstructions(s string) {
	p := strings.Split(s, "\n")
	t := []Instruction{}
	for _, v := range p {
		z := strings.TrimSpace(v)
		if len(z) > 0 {
			x := strings.Split(z, " ")
			i := Instruction{name: strings.TrimSpace(x[0]), args: x[1:]}
			t = append(t, i)
		}
	}

	c.program = t
}

func isEven(x uint) bool {
	return x%2 == 0
}

func (c *puter) doItem(idx int) int {
	instruction := c.program[idx]
	ins := instruction.name

	if ins == "hlf" {
		register := instruction.args[0]
		c.registers[register] /= 2
	}

	if ins == "tpl" {
		register := instruction.args[0]
		c.registers[register] *= 3
	}

	if ins == "inc" {
		register := instruction.args[0]
		c.registers[register] += 1
	}

	if ins == "jmp" {
		jmp, _ := strconv.Atoi(instruction.args[0])
		return idx + jmp
	}

	if ins == "jie" {
		j := instruction.args[1]
		st, jmx := j[0], j[1:]
		factor := 1
		if st == '-' {
			factor = -1
		}
		j1, _ := strconv.Atoi(jmx)
		jmp := factor * j1

		register := strings.Replace(instruction.args[0], ",", "", 1)
		d := c.registers[register]

		if isEven(d) {
			return idx + jmp
		}
	}

	if ins == "jio" {
		j := instruction.args[1]
		st, jmx := j[0], j[1:]
		factor := 1
		if st == '-' {
			factor = -1
		}
		j1, _ := strconv.Atoi(jmx)
		jmp := factor * j1
		register := strings.Replace(instruction.args[0], ",", "", 1)

		d := c.registers[register]

		if !isEven(d) {
			return idx + jmp
		}
	}

	return idx + 1
}
