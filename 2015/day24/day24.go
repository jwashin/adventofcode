package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1", part1())
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	return makeLists(string(input))
}

func listSum(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

func qe(data []int) int {
	product := 1
	for _, v := range data {
		product *= v
	}
	return product
}

func makeLists(input string) int {
	d := strings.Split(input, "\n")
	data := []int{}
	for _, v := range d {
		a, _ := strconv.Atoi(v)
		data = append(data, a)
	}
	weight := listSum(data)
	// use 3 for part 1, 4 for part 2
	target := weight / 3
	min := 0
	ct := 1
	for ct < len(data) {
		ct += 1
		pairs := findCombos(data, target, ct)
		if len(pairs) > 0 {
			min = math.MaxInt
			for _, v := range pairs {
				if qe(v) < min {
					min = qe(v)
				}
			}
			return min
		}
	}
	return 0
}

func findCombos(data []int, total int, count int) [][]int {
	out := [][]int{}
	comb(len(data), count, func(c []int) {
		newList := []int{}
		for _, v := range c {
			newList = append(newList, data[v])
		}
		if listSum(newList) == total {
			out = append(out, newList)
		}
	})
	return out
}

// https://rosettacode.org/wiki/Combinations#Go
func comb(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
	}
	rc(0, 0)
}
