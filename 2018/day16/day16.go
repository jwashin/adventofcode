package main

import (
	"os"
	"strconv"
	"strings"
)

func part1() {

}

type sample struct {
	before          []int
	opcode, a, b, c int
	after           []int
}

func parseInput() {
	samples := []sample{}
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")
	var item sample
	for _, v := range data {
		if strings.Contains(v, "Before") {
			item = sample{}
			v = strings.ReplaceAll(v, ",", "")
			v = strings.ReplaceAll(v, "[", "")
			v = strings.ReplaceAll(v, "]", "")
			item.before = []int{}
			j := strings.Fields(v)
			for i := 1; i <= 4; i++ {
				d, _ := strconv.Atoi(j[i])
				item.before = append(item.before, d)
			}

		}
	}
}

var opcodes = []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti",
	"gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}

func candidateOpcodes(before []int, a, b, c int, after []int) []string {
	candidates := []string{}
	for _, v := range opcodes {
		test := operation(before, v, a, b, c)
		equal := true
		for i := range after {
			if test[i] != after[i] {
				equal = false
				break
			}
		}
		if equal {
			candidates = append(candidates, v)
		}
	}
	return candidates
}

func operation(registers []int, opcode string, a int, b int, c int) []int {
	switch opcode {
	case "addr":
		registers[c] = registers[a] + registers[b]
	case "addi":
		registers[c] = registers[a] + b
	case "mulr":
		registers[c] = registers[a] * registers[b]
	case "muli":
		registers[c] = registers[a] * b
	case "banr":
		registers[c] = registers[a] & registers[b]
	case "bani":
		registers[c] = registers[a] & b
	case "borr":
		registers[c] = registers[a] | registers[b]
	case "bori":
		registers[c] = registers[a] | b
	case "setr":
		registers[c] = registers[a]
	case "seti":
		registers[c] = a
	case "gtir":
		registers[c] = 0
		if a > registers[b] {
			registers[c] = 1
		}
	case "gtri":
		registers[c] = 0
		if registers[a] > b {
			registers[c] = 1
		}
	case "gtrr":
		registers[c] = 0
		if registers[a] > registers[b] {
			registers[c] = 1
		}
	case "eqir":
		registers[c] = 0
		if a == registers[b] {
			registers[c] = 1
		}
	case "eqri":
		registers[c] = 0
		if registers[a] == b {
			registers[c] = 1
		}
	case "eqrr":
		registers[c] = 0
		if registers[a] == registers[b] {
			registers[c] = 1
		}
	}
	return registers
}
