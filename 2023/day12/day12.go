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

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
	fmt.Println("Part2:", part2(string(input)))
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

func part2(s string) int {

	s = strings.TrimSpace(s)
	data := strings.Split(s, "\n")
	count := 0
	for _, v := range data {
		count += getResult2(v)
	}
	return count
}

func getResult2(s string) int {
	s = strings.TrimSpace(s)
	d2 := strings.Split(s, " ")
	s = d2[0]
	ind := d2[1]
	os := []string{}
	bs := []string{}

	for len(os) < 5 {
		os = append(os, s)
		bs = append(bs, ind)
	}
	s = strings.Join(os, "?")
	ind = strings.Join(bs, ",")
	// s = strings.ReplaceAll(s,"."," ")
	return countArrangements(s, ind)
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
	// s = strings.ReplaceAll(s,"."," ")
	return countArrangements(s, ind)
}

func asCommaDelimited(d []int) string {
	s := []string{}
	for _, v := range d {
		s = append(s, fmt.Sprint(v))
	}
	return strings.Join(s, ",")
}

func asInts(d string) []int {
	f := strings.Split(d, ",")
	g := []int{}
	for _, v := range f {
		w, _ := strconv.Atoi(v)
		g = append(g, w)
	}
	return g
}

var cached = map[string]int{}

func countArrangements(s string, counts string) int {
	for len(s) > 0 && s[0] == '.' {
		s = s[1:]
	}
	stringLength := len(s)
	// consCounts := asInts(counts)
	hasHashes := strings.Contains(s, "#")
	if len(counts) == 0 {
		if hasHashes {
			return 0
		} else {
			return 1
		}
	}

	cacheKey := s + " " + counts
	value, hasKey := cached[cacheKey]
	if hasKey {
		return value
	}

	// value = 0
	targetLength := asInts(counts)[0]

	if stringLength < targetLength {
		return 0
	}

	testStr := strings.Repeat("#", targetLength)
	operativeString := s[:targetLength]

	if s == testStr && !strings.Contains(counts, ",") {
		cached[cacheKey] = 1
		return 1
	}

	if stringLength > targetLength && operativeString == testStr && (s[targetLength] == '.' || s[targetLength] == '?') {
		newcounts := asCommaDelimited(asInts(counts)[1:])
		value += countArrangements(s[targetLength+1:], newcounts)
		cached[cacheKey] = value
		return value
	}

	if strings.Contains(operativeString, "?") {
		z1 := strings.Replace(s, "?", ".", 1)
		z2 := strings.Replace(s, "?", "#", 1)
		value := countArrangements(z2, counts) + countArrangements(z1, counts)
		cached[cacheKey] = value
		return value
	}
	return 0
}
