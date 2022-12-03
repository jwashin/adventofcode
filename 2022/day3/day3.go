package main

import (
	"fmt"
	"os"
	"strings"
)

var alfas = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1", part1(string(input)))
	fmt.Println("Part 2", part2(string(input)))

}

func part1(s string) int {
	data := strings.Split(s, "\n")
	sum := 0
	for _, item := range data {
		sum += priority(commonChar(item))
	}
	return sum
}

func part2(s string) int {
	data := strings.Split(s, "\n")
	sum := 0
	group := []string{}
	for _, elf := range data {
		elf = strings.TrimSpace(elf)
		if len(group) < 3 {
			group = append(group, elf)
		}
		if len(group) == 3 {
			sum += priority(commonChar2(group))
			group = []string{}
		}
	}
	return sum
}

func commonChar2(rucksacks []string) string {
	for _, v := range rucksacks[0] {
		for _, n := range rucksacks[1] {
			for _, m := range rucksacks[2] {
				if v == n && n == m {
					return string(v)
				}
			}
		}
	}
	return ""
}

func commonChar(s string) string {
	s = strings.TrimSpace(s)
	firstHalf := s[:len(s)/2]
	secondHalf := s[len(s)/2:]
	for _, v := range firstHalf {
		for _, n := range secondHalf {
			if v == n {
				return string(v)
			}
		}
	}
	return ""
}

func priority(s string) int {
	return strings.Index(alfas, s) + 1
}
