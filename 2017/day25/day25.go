package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1", part1(false))
}

type turingMachine struct {
	state        string
	interval     int
	cursor       int
	tape         map[int]int
	instructions map[string]instruction
}

func (t *turingMachine) run() int {
	for n := 0; n < t.interval; n++ {
		currentValue := fmt.Sprint(t.tape[t.cursor])
		i := t.instructions[t.state+currentValue]
		t.tape[t.cursor] = i.value
		t.cursor += i.move
		t.state = i.nextState
	}
	count := 0
	for _, v := range t.tape {
		count += v
	}
	return count
}

type instruction struct {
	// state     string
	// condition int
	// we'll use state and condition as key.
	value     int
	move      int
	nextState string
}

var testInstructions = `Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.`

func (t *turingMachine) getInstructions(s string) {
	program := strings.Split(s, "\n")
	firstLine := program[0]
	t.state = string(firstLine[len(firstLine)-2])
	secondLine := program[1]
	csAfter := 0
	fmt.Sscanf(secondLine, "Perform a diagnostic checksum after %d steps.", &csAfter)
	t.interval = csAfter

	instr := map[string]instruction{}
	state := ""
	value := ""
	write := 0
	move := 1
	for _, v := range program[3:] {
		v = strings.TrimSpace(strings.ReplaceAll(v, ":", " :"))
		v = strings.ReplaceAll(v, ".", "")
		if strings.Contains(v, "In state") {
			t := strings.Fields(v)
			state = t[2]
		}
		if strings.Contains(v, "current value") {
			t := strings.Fields(v)
			value = t[5]
		}
		if strings.Contains(v, "Write") {
			t := strings.Fields(v)
			write, _ = strconv.Atoi(t[4])
		}
		if strings.Contains(v, "Move") {
			t := strings.Fields(v)
			ts := t[6]
			if ts == "right" {
				move = 1
			}
			if ts == "left" {
				move = -1
			}
		}
		if strings.Contains(v, "Continue") {
			t := strings.Fields(v)
			nextState := t[4]
			item := instruction{move: move, value: write, nextState: nextState}
			instr[state+value] = item
		}
	}
	t.instructions = instr
}

func part1(test bool) int {

	input := testInstructions
	if !test {
		xinput, _ := os.ReadFile("input.txt")
		input = string(xinput)
	}
	t := turingMachine{tape: map[int]int{}}
	t.getInstructions(input)
	s := t.run()
	return s

}
