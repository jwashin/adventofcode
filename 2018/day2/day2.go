package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

func hasSameLetters(s string, n int) bool {
	for _, v1 := range s {
		if strings.Count(s, string(v1)) == n {
			return true
		}
	}
	return false
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	count2 := 0
	count3 := 0
	for _, s := range strings.Split(data, "\n") {
		if hasSameLetters(s, 2) {
			count2 += 1
		}
		if hasSameLetters(s, 3) {
			count3 += 1
		}
	}
	return count2 * count3
}

func part2() string {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	q := strings.Split(data, "\n")
	for {
		currentItem := q[0]
		q = q[1:]
		for _, v := range q {
			for i := 1; i < len(currentItem)-1; i++ {
				x := currentItem[:i] + currentItem[i+1:]
				y := v[:i] + v[i+1:]
				if x == y {
					return y
				}
			}
		}
	}
}
