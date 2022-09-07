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
	c.getInstructions(string(data))
	n := 0
	ok := false
	for !ok {
		n += 1
		c.registers = map[string]int{"a": n, "b": 0, "c": 0, "d": 0}
		ok = c.run()
	}
	fmt.Println(n)
}

type puter struct {
	registers map[string]int
	program   []Instruction
	output    string
}

type Instruction struct {
	name string
	args []string
}

func (c puter) run() bool {
	inst := 0
	test := "010101010101010101010101010"
	for inst < len(c.program) {
		inst = c.doItem(inst)
		if len(c.output) == len(test) {
			return c.output == test
		}
	}
	return true
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

func (c *puter) doItem(idx int) int {
	instruction := c.program[idx]
	ins := instruction.name

	if ins == "out" {
		val, err := strconv.Atoi(instruction.args[0])
		if err == nil {
			c.output += fmt.Sprint(val)
		} else {
			c.output += fmt.Sprint(c.registers[instruction.args[0]])
		}
	}

	if ins == "inc" {
		register := instruction.args[0]
		c.registers[register] += 1
	}

	if ins == "dec" {
		register := instruction.args[0]
		c.registers[register] -= 1
	}

	if ins == "cpy" {
		dest := instruction.args[1]
		val, err := strconv.Atoi(instruction.args[0])
		// first arg is not a number
		if err != nil {
			val = c.registers[instruction.args[0]]
		}
		_, err1 := strconv.Atoi(dest)
		// second arg is a number. invalid instruction
		if err1 == nil {
			return idx + 1
		}
		c.registers[dest] = val
	}

	if ins == "jnz" {
		arg1, arg2 := instruction.args[0], instruction.args[1]
		val, err := strconv.Atoi(arg1)
		// first arg is not a number
		if err != nil {
			val = c.registers[arg1]
		}
		if val != 0 {
			y, err := strconv.Atoi(arg2)
			if err == nil {
				return idx + y
			}
			return idx + c.registers[arg2]
		}
	}
	return idx + 1
}
