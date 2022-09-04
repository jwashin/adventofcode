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
	c.registers = map[string]int{"a": 7, "b": 0, "c": 0, "d": 0}
	c.getInstructions(string(data))
	c.run()
	fmt.Println("1.", c.registers["a"])
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
	program   []string
}

func (c puter) run() {
	inst := 0
	for inst < len(c.program) {
		inst = c.doItem(inst)
	}
}

func (c *puter) getInstructions(s string) {
	p := strings.Split(s, "\n")
	t := []string{}
	for _, v := range p {
		z := strings.TrimSpace(v)
		if len(z) > 0 {
			t = append(t, z)
		}
	}
	c.program = t
}

func (c puter) doItem(idx int) int {
	instruction := c.program[idx]
	ins := instruction[:3]
	length := len(instruction)
	if ins == "tgl" {
		opregister := string(instruction[length-1])
		affectedIndex := idx + c.registers[opregister]
		if affectedIndex >= 0 && affectedIndex < len(c.program) {
			opinst := c.program[affectedIndex]
			tins := opinst[:3]
			if tins == "inc" {
				c.program[affectedIndex] = "dec" + opinst[3:]
			}
			if tins == "dec" || tins == "tgl" {
				c.program[affectedIndex] = "inc " + opinst[3:]
			}
			if tins == "jnz" {
				c.program[affectedIndex] = "cpy" + opinst[3:]
			}
			if tins == "cpy" {
				c.program[affectedIndex] = "jnz" + opinst[3:]
			}
		}
	}

	if ins == "inc" {
		register := string(instruction[length-1])
		c.registers[register] += 1
	}

	if ins == "dec" {
		register := string(instruction[length-1])
		c.registers[register] -= 1
	}

	if ins == "cpy" {
		parsed := strings.Split(instruction, " ")
		dest := parsed[2]
		val, err := strconv.Atoi(parsed[1])
		if err != nil {
			val = c.registers[parsed[1]]
		}
		c.registers[dest] = val
	}

	if ins == "jnz" {
		parsed := strings.Split(instruction, " ")
		val, err := strconv.Atoi(parsed[1])
		if err != nil {
			val = c.registers[parsed[1]]
		}
		if val != 0 {
			y, err := strconv.Atoi(parsed[2])
			if err == nil {
				return idx + y
			}
			return idx + c.registers[parsed[2]]
		}
	}
	return idx + 1
}
