package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.ReadFile("input.txt")
	data := strings.TrimSpace(string(in))
	parsed := strings.Split(data, "\n")
	triangleCount := 0
	for _, v := range parsed {
		p2 := parse(v)
		if possibleTriangle(p2[0], p2[1], p2[2]) {
			triangleCount += 1
		}
	}
	fmt.Printf("1: %v\n", triangleCount)
	part2Count := 0
	newList := [][]string{}
	for _, v := range parsed {
		p2 := parse(v)
		newList = append(newList, p2)
		if len(newList) < 3 {
			continue
		}
		for _, y := range []int{0, 1, 2} {
			if possibleTriangle(newList[0][y], newList[1][y], newList[2][y]) {
				part2Count += 1
			}
		}
		newList = [][]string{}
	}
	fmt.Printf("2: %v\n", part2Count)
}

func parse(a string) []string {
	r, _ := regexp.Compile(`\s*([0-9]{1,3})\s*`)
	return r.FindAllString(a, 3)
}

func possibleTriangle(a string, b string, c string) bool {
	sidea, _ := strconv.Atoi(strings.TrimSpace(a))
	sideb, _ := strconv.Atoi(strings.TrimSpace(b))
	sidec, _ := strconv.Atoi(strings.TrimSpace(c))
	tl := []int{sidea, sideb, sidec}
	mx := max(tl)
	ns := []int{}
	// make new list without max
	for _, val := range tl {
		if val != mx {
			ns = append(ns, val)
		}
	}
	// must have two remaining
	for len(ns) < 2 {
		ns = append(ns, mx)
	}
	return sum(ns) > mx
}

func sum(d []int) int {
	c := 0
	for _, v := range d {
		c += v
	}
	return c
}

func max(d []int) int {
	m := d[0]
	for _, v := range d {
		if v > m {
			m = v
		}
	}
	return m
}
