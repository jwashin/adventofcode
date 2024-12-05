package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", mulconsume(string(input)))
	fmt.Println("Part 2:", condmulconsume(string(input)))

}

func mulconsume(s string) int {
	result := 0

	for len(s) > 0 {
		if strings.HasPrefix(s, "mul(") {
			s = s[4:]
			multiplicand := ""
			for isDigit(s[0]) {
				multiplicand += string(s[0])
				s = s[1:]
			}
			if s[0] == ',' {
				s = s[1:]
			}
			multiplier := ""
			for isDigit(s[0]) {
				multiplier += string(s[0])
				s = s[1:]
			}
			if s[0] == ')' {

				m1, er1 := strconv.Atoi(multiplicand)
				m2, er2 := strconv.Atoi(multiplier)
				if er1 == nil && er2 == nil {
					result += m1 * m2
				}
				s = s[1:]
			}

		} else {
			s = s[1:]
		}
	}
	return result
}

func condmulconsume(s string) int {
	result := 0
	enabled := true

	for len(s) > 0 {
		if strings.HasPrefix(s, "do()") {
			enabled = true
			s = s[4:]
		}
		if strings.HasPrefix(s, "don't()") {
			enabled = false
			s = s[7:]
		}
		if strings.HasPrefix(s, "mul(") {
			s = s[4:]
			multiplicand := ""
			for isDigit(s[0]) {
				multiplicand += string(s[0])
				s = s[1:]
			}
			if s[0] == ',' {
				s = s[1:]
			}
			multiplier := ""
			for isDigit(s[0]) {
				multiplier += string(s[0])
				s = s[1:]
			}
			if s[0] == ')' {

				m1, er1 := strconv.Atoi(multiplicand)
				m2, er2 := strconv.Atoi(multiplier)
				if er1 == nil && er2 == nil && enabled {
					result += m1 * m2
				}
				s = s[1:]
			}

		} else {
			s = s[1:]
		}
	}
	return result
}

func isDigit(s byte) bool {
	return s >= '0' && s <= '9'

}
