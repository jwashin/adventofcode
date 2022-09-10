package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.json")
	fmt.Println("1.", addAllNumbers(string(input)))
	fmt.Println("2.", addAllNumbers(removeObjectsWithRedValues(string(input))))
}

func addAllNumbers(s string) int {
	reg, _ := regexp.Compile(`(-?\d+)`)
	total := 0
	for _, v := range reg.FindAllString(s, -1) {
		item, _ := strconv.Atoi(v)
		total += item
	}
	return total
}

func removeObjectsWithRedValues(s string) string {
	findstr := ":\"red\""
	t := s
	for strings.Contains(t, findstr) {
		// go backwards
		ix := strings.Index(t, findstr)
		b := 1
		for b > 0 {
			ix -= 1
			if t[ix] == '}' {
				b += 1
			}
			if t[ix] == '{' {
				b -= 1
			}
		}
		left := ix
		// now, go forwards
		ix = strings.Index(t, findstr)
		b = 1
		for b > 0 {
			ix += 1
			if t[ix] == '{' {
				b += 1
			}
			if t[ix] == '}' {
				b -= 1
			}
		}
		right := ix
		t = t[:left] + t[right+1:]

	}
	return t
}
