package piloting

import (
	"strconv"
	"strings"
)

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (p Position) SetPosition(instruction string) Position {

	data := strings.Split(instruction, " ")
	instr := data[0]
	amount, _ := strconv.Atoi(data[1])

	if instr == "forward" {
		return Position{p.Horizontal + amount, p.Depth + p.Aim*amount, p.Aim}
	}
	if instr == "down" {
		return Position{p.Horizontal, p.Depth, p.Aim + amount}
	}
	if instr == "up" {
		return Position{p.Horizontal, p.Depth, p.Aim - amount}
	}

	return Position{}
}

func Navigate(instructions []string) int {
	start := Position{0, 0, 0}
	for _, value := range instructions {
		start = start.SetPosition(value)
	}
	return start.Depth * start.Horizontal
}
