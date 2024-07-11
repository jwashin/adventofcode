package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part2: ", part1(string(input)))

}

func part1(s string) int {
	total := 0
	t := strings.Split(s, "\n")
	for _, v := range t {
		first, last := firstAndLastDigits(v)
		lineValue := recombineNumber(first, last)
		total += lineValue
	}
	return total

}

func recombineNumber(a int, b int) int {
	return 10*a + b
}

var findDigits = map[string]int{
	"1":     1,
	"one":   1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"5":     5,
	"five":  5,
	"six":   6,
	"6":     6,
	"7":     7,
	"seven": 7,
	"eight": 8,
	"8":     8,
	"nine":  9,
	"9":     9,
}

func firstAndLastDigits(s string) (int, int) {
	hasFirstDigit := false
	firstDigit := 0
	lastDigit := 0
	for {
		for ky := range findDigits {
			if strings.Index(s, ky) == 0 {
				foundDigit := findDigits[ky]
				lastDigit = foundDigit
				if !hasFirstDigit {
					firstDigit = lastDigit
					hasFirstDigit = true
				}
				s = s[len(ky)-1:]
				break
			}

		}
		if len(s) > 0 {
			s = s[1:]
		} else {
			return firstDigit, lastDigit
		}
	}

}
