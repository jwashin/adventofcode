package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func fitsCriteria(s string, criterion []int) bool {
// 	s = strings.ReplaceAll(strings.TrimSpace(s), ".", " ")
// 	t := strings.Fields(s)
// 	if len(t) != len(criterion) {
// 		return false
// 	}
// 	olist := []int{}
// 	for _, v := range t {
// 		olist = append(olist, len(v))
// 	}
// 	return listEqual(olist, criterion)
// }

// func listEqual(a []int, b []int) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i, v := range a {
// 		if v != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

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
		count += getResult(v)
	}
	return count
}

// func part2(s string) int {
// 	// 7017 too low
// 	s = strings.TrimSpace(s)
// 	data := strings.Split(s, "\n")
// 	count := 0
// 	for _, v := range data {
// 		count += countTheWays2(v)
// 	}
// 	return count
// }

func getResult(s string) int {
	s = strings.TrimSpace(s)
	d2 := strings.Split(s, " ")
	s = d2[0]
	ind := d2[1]
	isplit := strings.Split(ind, ",")
	indicator := []int{}
	for _, v := range isplit {
		length, _ := strconv.Atoi(v)
		indicator = append(indicator, length)
	}

	// we should now be able to
	// countTheWays(s string, indicator []int)
	return countTheWays(s, indicator)
}

// func countTheWays1(s string, indicator []int) int {
// 	s = strings.ReplaceAll(s, ".", " ")
// 	j := strings.Fields(s)
// 	s = strings.Join(j, " ")
// 	qCount := strings.Count(s, "?")
// 	count := 0
// 	binValue := int(math.Pow(2, float64(qCount)))
// 	binHint := strings.Replace("%20b", "20", fmt.Sprint(qCount), 1)
// 	mask := strings.Replace(s, "?", "%s", -1)
// 	for i := 0; i < binValue; i++ {
// 		choices := fmt.Sprintf(binHint, i)
// 		choices = strings.Replace(choices, "0", " ", -1)
// 		choices = strings.Replace(choices, "1", "#", -1)
// 		// for len(choices) > qCount {
// 		// 	choices = choices[1:]
// 		// }
// 		items := []string{}
// 		for _, v := range choices {
// 			items = append(items, string(v))
// 		}
// 		candidate := mask
// 		for _, v := range items {
// 			candidate = strings.Replace(candidate, "%s", v, 1)
// 		}
// 		// candidate := fmt.Sprintf(mask, items)
// 		if fitsCriteria(candidate, indicator) {
// 			count += 1
// 		}
// 	}
// 	return count
// }

var replacements = []string{"#", "."}

func countTheWays(s string, indicator []int) int {
	ls := len(s)
	li := len(indicator)
	if ls == li && li == 0 {
		return 1
	} else if ls != li && (ls == 0 || li == 0) {
		return 0
	}

	firstCharacter := s[0]
	if firstCharacter == '.' {
		return countTheWays(s[1:], indicator)
	} else if firstCharacter == '?' {
		tries := []string{}
		for _, v := range replacements {
			tries = append(tries, strings.Replace(s, "?", v, 1))
		}
		out := 0
		for _, av := range tries {
			out += countTheWays(av, indicator)
		}
		return out
	} else {
		// '#' is the first char
		firstN := s[0:indicator[0]]
		if strings.Count(firstN, "#") == indicator[0] {
			return countTheWays(s[indicator[0]:], indicator[1:])
		}

		items := []string{}
		for _, v := range replacements {
			items = append(items, strings.Replace(s, "?", v, 1))
		}
		count := 0
		for _, v := range items {
			count += countTheWays(v, indicator)
		}
		return count
	}
}
