package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("1.", part1(string(input)))
	fmt.Println("2.", part2(string(input)))
}

func part1(s string) int {
	used := map[string]int{}
	t := strings.Fields(s)
	data := []int{}
	for _, v := range t {
		d, _ := strconv.Atoi(v)
		data = append(data, d)
	}
	used[strings.Join(asStrings(data), ",")] += 1
	n := 0
	for {
		n += 1
		data := cycle(data)
		ky := strings.Join(asStrings(data), ",")
		used[ky] += 1
		if used[ky] > 1 {
			break
		}
	}
	return n
}

func part2(s string) int {
	used := map[string]int{}
	t := strings.Fields(s)
	data := []int{}
	for _, v := range t {
		d, _ := strconv.Atoi(v)
		data = append(data, d)
	}
	used[strings.Join(asStrings(data), ",")] += 1
	n := 0
	for {
		n += 1
		data := cycle(data)
		ky := strings.Join(asStrings(data), ",")
		used[ky] += 1
		x := 0
		if used[ky] > 1 {
			for {
				x += 1
				data := cycle(data)
				ky2 := strings.Join(asStrings(data), ",")
				if ky2 == ky {
					return x
				}
			}
		}
	}
}

func asStrings(d []int) []string {
	nl := []string{}
	for _, v := range d {
		d := fmt.Sprint(v)
		nl = append(nl, d)
	}
	return nl
}

func maxIdx(t []int) (int, int) {
	max := t[0]
	maxIdx := 0
	for k, v := range t {
		if v > max {
			max = v
			maxIdx = k
		}
	}
	return maxIdx, max
}

func cycle(bank []int) []int {
	idx, max := maxIdx(bank)
	bank[idx] = 0
	for max > 0 {
		idx += 1
		if idx == len(bank) {
			idx = 0
		}
		bank[idx] += 1
		max -= 1
	}
	return bank
}
