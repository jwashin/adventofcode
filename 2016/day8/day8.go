package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")
	r := initArray(6, 50)
	ct := 0
	for _, v := range data {
		cmd, a, b := parse(v)
		if cmd == "rect" {
			r = rect(r, a, b)
			// ct += a * b
		}
		if cmd == "rotcol" {
			r = rotCol(r, a, b)
		}
		if cmd == "rotrow" {
			r = rotRow(r, a, b)
		}
		printArray(r)
	}
	// ct = 0
	for _, v := range r {
		ct += strings.Count(v, "#")
	}
	fmt.Println(ct)
	printArray(r)
}

func parse(astring string) (cmd string, a int, b int) {
	if strings.Contains(astring, "rect") {
		fmt.Sscanf(astring, "rect %dx%d", &a, &b)
		cmd = "rect"
		fmt.Println(cmd, a, b)
		return
	}
	if strings.Contains(astring, "column") {
		fmt.Sscanf(astring, "rotate column x=%d by %d", &a, &b)
		cmd = "rotcol"
		fmt.Println(cmd, a, b)
		return
	}
	if strings.Contains(astring, "row") {
		fmt.Sscanf(astring, "rotate row y=%d by %d", &a, &b)
		cmd = "rotrow"
		fmt.Println(cmd, a, b)
		return
	}
	fmt.Println(cmd, a, b)
	return
}

func initArray(rows int, cols int) []string {
	t := []string{}
	for len(t) < rows {
		t = append(t, strings.Repeat(".", cols))
	}
	return t
}

func printArray(r []string) {
	for _, k := range r {
		fmt.Println(k)
	}
}

func rect(r []string, a int, b int) []string {
	newRect := []string{}
	for idx, row := range r {
		if idx < b {
			insert := strings.Repeat("#", a)
			ns := insert + row[a:]
			newRect = append(newRect, ns)
			continue
		}
		newRect = append(newRect, row)
	}
	// printArray(newRect)
	return newRect
}

func rotateString(s string, n int) string {
	// {"3", args{".##", 1}, "#.#"},
	// 0123
	// 3012
	// 2301

	spl := len(s) - n
	return s[spl:] + s[:spl]
}

func rotCol(r []string, x int, n int) []string {
	newRect := []string{}
	ns := ""
	for _, v := range r {
		ns += string(v[x])
	}
	gs := rotateString(ns, n)
	fmt.Println(n, ns, "to", gs)
	for idx, v := range r {
		newRect = append(newRect, v[:x]+string(gs[idx])+v[x+1:])
	}
	return newRect
}

func rotRow(r []string, y int, n int) []string {
	r[y] = rotateString(r[y], n)
	return r
}
