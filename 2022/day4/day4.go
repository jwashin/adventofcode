package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	fmt.Println("part 1.", part1(data))
	fmt.Println("part 2.", part2(data))
}

func part1(s string) int {
	data := strings.Split(s, "\n")
	count := 0
	for _, pairs := range data {
		pairs = strings.TrimSpace(pairs)
		p := strings.Split(pairs, ",")
		if containsx(p[0], p[1]) {
			count += 1
		}
	}
	return count
}

func part2(s string) int {
	data := strings.Split(s, "\n")
	count := 0
	for _, pairs := range data {
		pairs = strings.TrimSpace(pairs)
		p := strings.Split(pairs, ",")
		if containsy(p[0], p[1]) {
			count += 1
		}
	}
	return count
}

func containsx(s1 string, s2 string) bool {
	one := strings.Split(s1, "-")
	two := strings.Split(s2, "-")

	min1, _ := strconv.Atoi(one[0])
	max1, _ := strconv.Atoi(one[1])

	min2, _ := strconv.Atoi(two[0])
	max2, _ := strconv.Atoi(two[1])

	if min2 >= min1 && max2 <= max1 {
		return true
	}

	if min1 >= min2 && max1 <= max2 {
		return true
	}
	return false
}

// 850 too low

func containsy(s1 string, s2 string) bool {
	one := strings.Split(s1, "-")
	two := strings.Split(s2, "-")

	min1, _ := strconv.Atoi(one[0])
	max1, _ := strconv.Atoi(one[1])

	min2, _ := strconv.Atoi(two[0])
	max2, _ := strconv.Atoi(two[1])

	if min2 >= min1 && min2 <= max1 {
		return true
	}

	if max2 >= min1 && max2 <= max1 {
		return true
	}

	if min1 >= min2 && min1 <= max2 {
		return true
	}

	if max1 >= min2 && max1 <= max2 {
		return true
	}

	return false
}
