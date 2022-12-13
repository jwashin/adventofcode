package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
	fmt.Println("Part2:", part2(string(input)))
}

type pair struct {
	left  string
	right string
}

func part2(s string) int {
	input := strings.Split(s, "\n")
	data := []string{}
	for _, v := range input {
		s := strings.TrimSpace(v)
		if len(s) > 0 {
			data = append(data, s)
		}
	}
	p1 := "[[2]]"
	p2 := "[[6]]"
	data = append(data, p1)
	data = append(data, p2)
	sort.Slice(data, func(i, j int) bool {
		tst, _ := correctOrder(data[i], data[j])
		return tst
	})
	k1, k2 := 0, 0
	for k, v := range data {
		if v == p1 {
			k1 = k + 1
		}
		if v == p2 {
			k2 = k + 1
		}
	}
	return k1 * k2
}

// 6963 too high
func part1(s string) int {
	sum := 0
	data := strings.Split(s, "\n")
	pairs := []pair{}
	acc := []string{}
	for _, v := range data {
		v = strings.TrimSpace(v)
		if len(v) > 0 {
			acc = append(acc, v)
			if len(acc) == 2 {
				pairs = append(pairs, pair{acc[0], acc[1]})
				acc = []string{}
			}
		}
	}
	for k, v := range pairs {
		s, err := correctOrder(v.left, v.right)
		if err == nil && s {
			sum += k + 1
		}
	}
	return sum
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isList(s string) bool {
	return strings.Contains(s, "[") && s[0] == '[' && s[len(s)-1] == ']'
}

func parseList(s string) []string {
	items := []string{}
	acc := ""
	list := 0
	st := s[1 : len(s)-1]
	for len(st) > 0 {
		t := st[0]
		st = st[1:]
		if t == '[' {
			list += 1
			acc += string(t)
			continue
		}
		if t == ']' {
			list -= 1
			acc += string(t)
			continue
		}
		if t == ',' {
			if list == 0 {
				if len(acc) > 0 {
					items = append(items, acc)
				}
				acc = ""
			} else {
				acc += string(t)
			}
			continue
		}
		acc += string(t)
	}
	if len(acc) > 0 {
		items = append(items, acc)
	}
	return items
}

func correctOrder(left string, right string) (bool, error) {
	if isInt(left) && isInt(right) {
		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)
		if l > r {
			return false, nil
		}
		if l < r {
			return true, nil
		}
	}
	if isList(left) && isList(right) {
		l := parseList(left)
		r := parseList(right)

		for k, v := range l {
			if k == len(r) {
				return false, nil
			}
			t, err := correctOrder(v, r[k])
			if err == nil {
				return t, nil
			}
		}
		if len(l) < len(r) {
			return true, nil
		}
	}
	if isInt(left) && !isInt(right) {
		return correctOrder("["+left+"]", right)
	}
	if isInt(right) && !isInt(left) {
		return correctOrder(left, "["+right+"]")
	}
	return false, errors.New("undecided")
}
