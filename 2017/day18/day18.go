package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var testData = `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`

var part2Test = `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`

func main() {
	// fmt.Println(part1(false))
	fmt.Println(part2(false))
}

// 127 too low

func part2(test bool) int {
	input, _ := os.ReadFile("input.txt")
	if test {
		input = []byte(part2Test)
	}
	data := string(input)
	ch1 := make(chan int, 5000)
	ch2 := make(chan int, 5000)
	c := puter{send: ch1, receive: ch2, id: 0, registers: map[string]int{"p": 0}, program: getInstructions(data)}

	d := puter{send: ch2, receive: ch1, id: 1, registers: map[string]int{"p": 1}, program: getInstructions(data)}
	out1 := make(chan int)
	out2 := make(chan int)
	go c.run(out1)
	go d.run(out2)

	a := <-out2
	fmt.Println(a)
	return a
}

func part1(test bool) int {
	input, _ := os.ReadFile("input.txt")
	if test {
		input = []byte(testData)
	}
	data := string(input)
	c := puter1{registers: map[string]int{"part1": 0}, program: []Instruction{}, frequencies: []int{}}
	c.program = getInstructions(data)
	c.run()
	return c.registers["part1"]

}

type Instruction struct {
	name string
	args []string
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

type puter struct {
	registers map[string]int
	program   []Instruction
	id        int
	send      chan int
	receive   chan int
	sendCount int
}

func (c puter) run(n chan int) {
	idx := 0
	for idx >= 0 && idx < len(c.program) {
		idx = c.doItem(idx)
	}
	n <- c.sendCount
}

func (c *puter) doItem(idx int) int {
	instruction := c.program[idx]
	ins := instruction.name

	if ins == "snd" {
		// time.Sleep(0)
		x := 0
		v0 := instruction.args[0]
		if isNum(v0) {
			x, _ = strconv.Atoi(v0)
		} else {
			x = c.registers[v0]
		}
		if c.id == 1 {
			c.sendCount += 1
			fmt.Println(c.sendCount)
		}
		// go func() {
		c.send <- x
		// }()
	}

	if ins == "rcv" {
		v0 := instruction.args[0]
		// time.Sleep(0)
		// go func() {
		b := <-c.receive
		c.registers[v0] = b
		// }()
	}

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
	if ins == "add" {

		v0 := instruction.args[0]
		y := 0
		v1 := instruction.args[1]
		if isNum(v1) {
			y, _ = strconv.Atoi(v1)
		} else {
			y = c.registers[v1]
		}
		c.registers[v0] += y
	}
	if ins == "mul" {
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
		c.registers[instruction.args[0]] = x * y
	}
	if ins == "mod" {
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
		c.registers[v0] = x % y
	}
	if ins == "jgz" {
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
		if x > 0 {
			return idx + y
		}
	}
	return idx + 1
}

type puter1 struct {
	registers   map[string]int
	program     []Instruction
	frequencies []int
}

func (c puter1) run() {
	idx := 0
	for idx >= 0 && idx < len(c.program) && c.registers["part1"] == 0 {
		idx = c.doItem(idx)
	}
}

func (c *puter1) doItem(idx int) int {
	instruction := c.program[idx]
	ins := instruction.name

	if ins == "snd" {
		x := 0
		v0 := instruction.args[0]
		if isNum(v0) {
			x, _ = strconv.Atoi(v0)
		} else {
			x = c.registers[v0]
		}
		c.frequencies = append(c.frequencies, x)
	}
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
	if ins == "add" {

		v0 := instruction.args[0]
		y := 0
		v1 := instruction.args[1]
		if isNum(v1) {
			y, _ = strconv.Atoi(v1)
		} else {
			y = c.registers[v1]
		}
		c.registers[v0] += y
	}
	if ins == "mul" {
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
		c.registers[instruction.args[0]] = x * y
	}
	if ins == "mod" {
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
		c.registers[v0] = x % y
	}
	if ins == "rcv" {
		v0 := instruction.args[0]
		x := c.registers[v0]
		if x != 0 {
			c.registers["part1"] = c.frequencies[len(c.frequencies)-1]
		}
	}
	if ins == "jgz" {
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
		if x > 0 {
			return idx + y
		}
	}
	return idx + 1
}
