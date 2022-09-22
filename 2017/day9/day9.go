package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	return scoreGroups(string(input))
}
func part2() int {
	input, _ := os.ReadFile("input.txt")
	return countGarbage(string(input))
}

func removeGarbage(s string) string {
	inGarbage := false
	ignoreNext := false
	newS := []byte{}
	for _, v := range s {
		if ignoreNext {
			ignoreNext = false
			continue
		}
		c := v
		if inGarbage && c == '!' {
			ignoreNext = true
			continue
		}
		if c == '<' {
			inGarbage = true
		}
		if c == '>' {
			inGarbage = false
			continue
		}
		if !inGarbage {
			newS = append(newS, byte(c))
		}
	}
	return string(newS)
}

func countGarbage(s string) int {
	inGarbage := false
	ignoreNext := false
	count := 0
	for _, v := range s {
		if ignoreNext {
			ignoreNext = false
			continue
		}
		c := v
		if inGarbage && c == '!' {
			ignoreNext = true
			continue
		}
		if c == '<' && !inGarbage {
			inGarbage = true
			continue
		}
		if c == '>' {
			inGarbage = false
			continue
		}
		if inGarbage {
			count += 1
		}
	}
	return count
}

func scoreGroups(s string) int {
	s = removeGarbage(s)
	score := 0
	scoreIncrement := 0
	for _, v := range s {
		if v == '{' {
			scoreIncrement += 1
		}
		if v == '}' {
			score += scoreIncrement
			scoreIncrement -= 1
		}
	}
	return score
}
