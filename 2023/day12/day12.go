package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func fitsCriteria(s string, criterion string) bool {
	// s = strings.TrimSpace(s)
	t := strings.Fields(s)
	if len(t) != len(strings.Split(criterion, ",")) {
		return false
	}
	olist := []string{}
	for _, v := range t {
		olist = append(olist, fmt.Sprint(len(v)))
	}
	return strings.Join(olist, ",") == criterion
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

func part2(s string) int {
	// 7017 too low
	s = strings.TrimSpace(s)
	data := strings.Split(s, "\n")
	count := 0
	for _, v := range data {
		count += countTheWays2(v)
	}
	return count
}

func countTheWays(s string) int {
	// s = strings.TrimSpace(s)

	d2 := strings.Split(s, " ")
	s = d2[0]
	indicator := d2[1]
	s = strings.ReplaceAll(s, ".", " ")
	j := strings.Fields(s)
	s = strings.Join(j, " ")

	qCount := strings.Count(s, "?")
	count := 0
	binValue := int(math.Pow(2, float64(qCount)))
	binHint := strings.Replace("%20b", "20", fmt.Sprint(qCount), 1)
	mask := strings.Replace(s, "?", "%s", -1)
	for i := 0; i < binValue; i++ {
		choices := fmt.Sprintf(binHint, i)
		choices = strings.Replace(choices, "0", ".", -1)
		choices = strings.Replace(choices, "1", "#", -1)
		// for len(choices) > qCount {
		// 	choices = choices[1:]
		// }
		items := []string{}
		for _, v := range choices {
			items = append(items, string(v))
		}
		candidate := mask
		for _, v := range items {
			candidate = strings.Replace(candidate, "%s", v, 1)
		}
		// candidate := fmt.Sprintf(mask, items)
		if fitsCriteria(candidate, indicator) {
			count += 1
		}
	}
	return count
}

func countTheWays2(s string) int {
	// s = strings.TrimSpace(s)

	d2 := strings.Split(s, " ")
	s = d2[0]
	indicator := d2[1]
	s = strings.ReplaceAll(s, ".", " ")
	j := strings.Fields(s)
	s = strings.Join(j, " ")
	d3 := []string{}
	d4 := []string{}
	for i := 0; i <= 5; i++ {
		d3 = append(d3, s)
		d4 = append(d4, indicator)
	}
	s = strings.Join(d3, "?")
	indicator = strings.Join(d4, ",")

	qCount := strings.Count(s, "?")
	count := 0
	binValue := int(math.Pow(2, float64(qCount)))
	binHint := strings.Replace("%20b", "20", fmt.Sprint(qCount), 1)
	mask := strings.Replace(s, "?", "%s", -1)
	for i := 0; i < binValue; i++ {

		choices := fmt.Sprintf(binHint, i)
		choices = strings.Replace(choices, "0", ".", -1)
		choices = strings.Replace(choices, "1", "#", -1)
		// for len(choices) > qCount {
		// 	choices = choices[1:]
		// }
		items := []string{}
		for _, v := range choices {
			items = append(items, string(v))
		}
		candidate := mask
		for _, v := range items {
			candidate = strings.Replace(candidate, "%s", v, 1)
		}
		// candidate := fmt.Sprintf(mask, items)
		if fitsCriteria(candidate, indicator) {
			count += 1
		}
	}
	return count
}
