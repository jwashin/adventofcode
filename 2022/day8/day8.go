package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", countVisible(string(input)))
	fmt.Println("part 2:", maxScenicScore(string(input)))
}

// 1700 too low
func countVisible(s string) int {
	input := strings.Split(s, "\n")
	data := []string{}
	for _, v := range input {
		data = append(data, strings.TrimSpace(v))
	}
	count := 0
	for y, v := range data {
		for x := range v {
			visible := isVisible(x, y, data)
			if visible {
				count += 1
			}
		}
	}
	return count
}

func maxScenicScore(s string) int {
	input := strings.Split(s, "\n")
	data := []string{}
	for _, v := range input {
		data = append(data, strings.TrimSpace(v))
	}
	score := 0
	for y, v := range data {
		for x := range v {
			sscore := scenicScore(y, x, data)
			if sscore > score {
				score = sscore
			}
		}
	}
	return score
}

func scenicScore(y int, x int, forest []string) int {
	treeHeight, _ := strconv.Atoi(string(forest[y][x]))
	top, bottom, left, right := 0, 0, 0, 0
	t, b, l, r := "", "", "", ""

	for i, v := range forest[y] {
		c := string(v)
		if i < x {
			l = c + l
		}
		if i > x {
			r += c
		}
	}
	for j, line := range forest {
		c := string(line[x])
		if j < y {
			t = c + t
		}
		if j > y {
			b += c
		}
	}

	for _, v := range t {
		ht, _ := strconv.Atoi(string(v))
		if ht < treeHeight {
			top += 1
		}
		if ht >= treeHeight {
			top += 1
			break
		}
	}
	for _, v := range b {
		ht, _ := strconv.Atoi(string(v))
		if ht < treeHeight {
			bottom += 1
		}
		if ht >= treeHeight {
			bottom += 1
			break
		}
	}
	for _, v := range l {
		ht, _ := strconv.Atoi(string(v))
		if ht < treeHeight {
			left += 1
		}
		if ht >= treeHeight {
			left += 1
			break
		}
	}

	for _, v := range r {
		ht, _ := strconv.Atoi(string(v))
		if ht < treeHeight {
			right += 1
		}
		if ht >= treeHeight {
			right += 1
			break
		}
	}
	out := top * bottom * left * right
	return out
}

func isVisible(x int, y int, forest []string) bool {
	treeHeight, _ := strconv.Atoi(string(forest[y][x]))

	top, bottom, left, right := true, true, true, true
	t, b, l, r := "", "", "", ""

	for i, v := range forest[y] {
		c := string(v)
		if i < x {
			l = c + l
		}
		if i > x {
			r += c
		}
	}
	for j, line := range forest {
		c := string(line[x])
		if j < y {
			t = c + t
		}
		if j > y {
			b += c
		}
	}

	for _, v := range t {
		ht, _ := strconv.Atoi(string(v))
		if ht >= treeHeight {
			top = false
			break
		}
	}

	for _, v := range b {
		ht, _ := strconv.Atoi(string(v))
		if ht >= treeHeight {
			bottom = false
			break
		}
	}
	for _, v := range l {
		ht, _ := strconv.Atoi(string(v))
		if ht >= treeHeight {
			left = false
			break
		}
	}
	for _, v := range r {
		ht, _ := strconv.Atoi(string(v))
		if ht >= treeHeight {
			right = false
			break
		}
	}
	return left || right || top || bottom
}
