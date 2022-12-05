package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var upperAlphas = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))

}

func part2(s string) string {

	tableau, instructions := parseInput(s)

	for _, v := range instructions {
		t := strings.Fields(v)
		count, _ := strconv.Atoi(t[1])
		source, _ := strconv.Atoi(t[3])
		destination, _ := strconv.Atoi(t[5])

		last := len(tableau[source]) - 1
		first := last - count + 1

		toMove := tableau[source][first:]

		tableau[source] = tableau[source][:first]
		tableau[destination] = append(tableau[destination], toMove...)
	}

	out := ""
	for c := 1; c <= len(tableau); c++ {
		out += tableau[c][len(tableau[c])-1]
	}
	return out

}

func part1(s string) string {

	tableau, instructions := parseInput(s)
	for _, v := range instructions {
		t := strings.Fields(v)
		count, _ := strconv.Atoi(t[1])
		source, _ := strconv.Atoi(t[3])
		destination, _ := strconv.Atoi(t[5])

		for c := 0; c < count; c++ {
			last := len(tableau[source]) - 1
			item := tableau[source][last]
			tableau[source] = tableau[source][:last]
			tableau[destination] = append(tableau[destination], item)
		}
	}
	out := ""
	for c := 1; c <= len(tableau); c++ {
		out += tableau[c][len(tableau[c])-1]
	}
	return out

}

func parseInput(s string) (map[int][]string, []string) {
	input := strings.Split(s, "\n")
	tableau := map[int][]string{}
	instructions := []string{}

	for _, v := range input {
		// v = strings.TrimSpace(v) heh. nope! we need the leading space
		if strings.Contains(v, "[") {
			for idx, m := range v {
				if strings.Contains(upperAlphas, string(m)) {
					crate := string(m)
					bin := idx/4 + 1
					// hmm. no need to initialize slices...
					tableau[bin] = append(tableau[bin], crate)
				}
			}
		} else if strings.Contains(v, "move") {
			instructions = append(instructions, v)
		}
	}
	// reverse the stacks to put bottom item at zero index
	for k := range tableau {
		for i, j := 0, len(tableau[k])-1; i < j; i, j = i+1, j-1 {
			tableau[k][i], tableau[k][j] = tableau[k][j], tableau[k][i]
		}
	}
	return tableau, instructions
}
