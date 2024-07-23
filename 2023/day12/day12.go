package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
}

func part1(s string) int {
	// 7017 too low
	s = strings.TrimSpace(s)
	data := strings.Split(s, "\n")
	count := 0
	for _, v := range data {
		z := strings.Split(v, " ")
		itemCounts := []int{}
		ss := strings.Split(z[1], ",")
		for _, d := range ss {
			i, _ := strconv.Atoi(d)
			itemCounts = append(itemCounts, i)
		}
		count += countTheWays(z[0], itemCounts)
	}
	return count
}

// https://old.reddit.com/r/adventofcode/comments/18ghux0/2023_day_12_no_idea_how_to_start_with_this_puzzle/kd0npmi/
// well, you could analyze the string left to right.
// if it starts with a ., discard the . and recursively check again.
// if it starts with a ?, replace the ? with a . and recursively check again,
//    AND replace it with a # and recursively check again.
// if it starts with a #, check if it is long enough for the first group,
//   check if all characters in the first [grouplength] characters are not '.',
//   and then remove the first [grouplength] chars and the first group number, recursively check again.
// at some point you will get to the point of having an empty string and more groups to do - that is a zero.
//   or you have an empty string with zero groups to do - that is a one.
// there are more rules to check than these few, which are up to you to find. but this is a way to work out the solution.

func countTheWays(s string, groupCounts []int) int {
	if len(s) == 0 && len(groupCounts) == 0 {
		return 1
	} else if len(s) == 0 && len(groupCounts) > 0 {
		return 0
	}
	if s[0] == '.' {
		return countTheWays(s[1:], groupCounts)
	}
	if s[0] == '?' {
		return countTheWays("."+s[1:], groupCounts) +
			countTheWays("#"+s[1:], groupCounts)
	}
	if s[0] == '#' {
		if len(s) >= groupCounts[0] {
			z := s[:groupCounts[0]]
			if !strings.Contains(z, ".") {
				return countTheWays(s[groupCounts[0]:], groupCounts[1:])
			} else {
				return 0
			}

		}

	}
	return 0
}
