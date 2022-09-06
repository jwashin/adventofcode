package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNumber(s string) {

}

func main() {
	data, _ := os.ReadFile("input.txt")
	c := puter{}
	c.registers = map[string]int{"a": 7, "b": 0, "c": 0, "d": 0}
	c.getInstructions(string(data))
	c.run()
	fmt.Println("1.", c.registers["a"])
	c.getInstructions(string(data))
	c.registers = map[string]int{"a": 12, "b": 0, "c": 0, "d": 0}
	c.run()
	fmt.Println("2.", c.registers["a"])
}

// part2 6096 too low

func test(s string) int {
	c := puter{}
	c.registers = map[string]int{}
	c.getInstructions(s)
	c.run()
	return c.registers["a"]
}

type puter struct {
	registers map[string]int
	program   []Instruction
}

type Instruction struct {
	name string
	args []string
}

func (c puter) run() {
	inst := 0
	for inst < len(c.program) {
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

func (c *puter) doItem(idx int) int {
	instruction := c.program[idx]
	ins := instruction.name
	if ins == "tgl" {
		opregister := instruction.args[0]
		affectedIndex := idx + c.registers[opregister]
		if affectedIndex >= 0 && affectedIndex < len(c.program) {
			opinst := c.program[affectedIndex]
			tins := opinst.name
			if tins == "inc" {
				c.program[affectedIndex].name = "dec"
			}
			if tins == "dec" || tins == "tgl" {
				c.program[affectedIndex].name = "inc"
			}
			if tins == "jnz" {
				c.program[affectedIndex].name = "cpy"
			}
			if tins == "cpy" {
				c.program[affectedIndex].name = "jnz"
			}
		}
	}

	// if ins == "mul" {
	// 	// TODO skip this if invalid instruction
	// 	args := instruction.args
	// 	dest := args[2]
	// 	multiplicand, err := strconv.Atoi(args[0])
	// 	if err != nil {
	// 		return idx + 1
	// 	}

	// 	multiplier, err := strconv.Atoi(args[1])
	// 	if err != nil {
	// 		return idx + 1
	// 	}
	// 	c.registers[dest] = multiplicand * multiplier
	// }

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
