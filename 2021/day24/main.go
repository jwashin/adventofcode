package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type ALU struct {
	input          []int
	currInputIndex int
	register       map[string]int
	instructions   []string
}

func (a *ALU) setInstructions(aString string) {
	a.instructions = []string{}
	data := strings.Split(aString, "\n")
	for _, k := range data {
		instr := strings.TrimSpace(k)
		if len(instr) > 0 {
			a.instructions = append(a.instructions, instr)
		}
	}
}

func (a *ALU) parseInput(i int) bool {
	a.currInputIndex = 0
	t := fmt.Sprint(i)
	out := []int{}
	for _, v := range t {
		d, _ := strconv.Atoi(string(v))
		if d == 0 {
			return false
		}
		out = append(out, d)
	}
	a.input = out
	return true
}

func (a *ALU) validate(input int, instructions string) bool {
	if a.parseInput(input) {
		a.setInstructions(instructions)
		a.doInstructions()
		fmt.Println("done")
		return a.register["z"] == 0
	}
	return false
}

func (a *ALU) doInstructions() {
	a.register = map[string]int{}
	for idx, k := range a.instructions {
		a.doInstruction(k, idx)
	}
}

func isAlpha(aString string) bool {
	return strings.Contains("abcdefghijklmnopqrstuvwxyz", aString)
}

func (a ALU) getInt(aString string) int {
	if isAlpha(aString) {
		return a.register[aString]
	}
	v, _ := strconv.Atoi(aString)
	return v
}
func (a *ALU) validateInt(anInt int) bool {
	if a.parseInput(anInt) {
		a.register = map[string]int{}
		a.doInstructions()
		fmt.Println("done")
		return a.register["z"] == 0
	}
	return false
}

func (a *ALU) doInstruction(aString string, idx int) bool {
	split := strings.Split(aString, " ")
	verb := split[0]
	first := split[1]
	second := split[len(split)-1]
	if verb == "inp" {
		a.register[first] = a.input[a.currInputIndex]
		a.currInputIndex += 1
	} else if verb == "add" {
		a.register[first] += a.getInt(second)
	} else if verb == "mul" {
		a.register[first] *= a.getInt(second)
	} else if verb == "div" {
		num := float64(a.register[first])
		sub := float64(a.getInt(second))
		if sub == 0.0 {
			fmt.Println("Division by zero on line " + fmt.Sprint(idx))
			return false
		}
		quo := int(math.Round(num / sub))
		a.register[first] = quo
	} else if verb == "mod" {
		if a.register[first] < 0 {
			fmt.Println("Undefined mod on line " + fmt.Sprint(idx))
			return false
		}
		if a.getInt(second) <= 0 {
			fmt.Println("Undefined mod on line " + fmt.Sprint(idx))
			return false
		}
		a.register[first] = a.register[first] % a.getInt(second)
	} else if verb == "eql" {
		t := a.register[first] == a.getInt(second)
		if t {
			a.register[first] = 1
		} else {
			a.register[first] = 0
		}
	}
	return true
}

func validate(a int, instructions string) bool {
	alu := ALU{}
	return alu.validate(a, instructions)
}

func main() {
	start := 99999999999999 + 1
	data, _ := ioutil.ReadFile("input.txt")
	alu := ALU{}
	alu.setInstructions(string(data))
	validates := false
	for !validates {
		start -= 1
		validates = alu.validateInt(start)
		fmt.Println(start)
	}
}
