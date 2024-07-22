package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func breakAndCount(s string) string {
	// s = strings.TrimSpace(s)
	t := strings.Fields(strings.Replace(s, ".", " ", -1))
	olist := []string{}
	for _, v := range t {
		olist = append(olist, fmt.Sprint(len(v)))
	}
	return strings.Join(olist, ",")
}

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
		count += countTheWays(v)
	}
	return count
}

func countTheWays(s string) int {
	s = strings.TrimSpace(s)

	d2 := strings.Split(s, " ")
	s = d2[0]
	indicator := d2[1]

	qCount := strings.Count(s, "?")
	count := 0
	binValue := int(math.Pow(2, float64(qCount)))
	mask := strings.Replace(s, "?", "%s", -1)
	for i := 0; i < binValue; i++ {
		choices := fmt.Sprintf("%20b", i)
		choices = strings.Replace(choices, "0", ".", -1)
		choices = strings.Replace(choices, "1", "#", -1)
		for len(choices) > qCount {
			choices = choices[1:]
		}
		items := []string{}
		for _, v := range choices {
			items = append(items, string(v))
		}
		candidate := mask
		for _, v := range items {
			candidate = strings.Replace(candidate, "%s", v, 1)
		}
		// candidate := fmt.Sprintf(mask, items)
		ind := breakAndCount(candidate)
		if ind == indicator {
			count += 1
		}
	}
	return count
}
