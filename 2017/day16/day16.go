package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pgms = "abcdefghijklmnop"

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

type instruction struct {
	name string
	sa   string
	sb   string
	ia   int
	ib   int
}

func (i instruction) do(s string) string {
	if i.name == "s" {
		return spin(s, i.ia)
	}
	if i.name == "x" {
		return exchange(s, i.ia, i.ib)
	}
	return partner(s, i.sa, i.sb)
}

func part2() string {
	s := pgms
	input, _ := os.ReadFile("input.txt")
	instrs := []instruction{}
	dance := strings.Split(string(input), ",")
	for _, v := range dance {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		if v[0] == 's' {
			n, _ := strconv.Atoi(v[1:])
			i := instruction{name: "s", ia: n}
			instrs = append(instrs, i)
			// s = spin(s, n)
		}
		if v[0] == 'x' {
			d := strings.Split(v[1:], "/")
			a, _ := strconv.Atoi(d[0])
			b, _ := strconv.Atoi(d[1])
			i := instruction{name: "x", ia: a, ib: b}
			instrs = append(instrs, i)
			// s = exchange(s, a, b)
		}
		if v[0] == 'p' {
			d := strings.Split(v[1:], "/")
			a := d[0]
			b := d[1]
			i := instruction{name: "p", sa: a, sb: b}
			instrs = append(instrs, i)
			// s = partner(s, a, b)
		}
	}

	// aha. the sequence of permutations repeated in 36 dances. We optimized for nothing...
	for i := 0; i < 1000000000%36; i++ {
		// if i%1000000 == 0 {
		// fmt.Println(i)
		// }
		for _, v := range instrs {
			s = v.do(s)
		}
		if s == pgms {
			fmt.Println("rpt:", i)
		}

	}

	return s
}

func part1() string {
	s := pgms
	input, _ := os.ReadFile("input.txt")
	dance := strings.Split(string(input), ",")
	for _, v := range dance {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		if v[0] == 's' {
			n, _ := strconv.Atoi(v[1:])
			s = spin(s, n)
		}
		if v[0] == 'x' {
			d := strings.Split(v[1:], "/")
			a, _ := strconv.Atoi(d[0])
			b, _ := strconv.Atoi(d[1])
			s = exchange(s, a, b)
		}
		if v[0] == 'p' {
			d := strings.Split(v[1:], "/")
			a := d[0]
			b := d[1]
			s = partner(s, a, b)
		}
	}
	return s
}

func spin(s string, n int) string {
	// move n items from rear to front
	split := len(s) - n
	lastN := s[split:]
	first := s[:split]
	return lastN + first
}

func exchange(s string, a, b int) string {
	f := []byte(s)
	A, B := f[a], f[b]
	f[b] = A
	f[a] = B
	return string(f)
}

func partner(s string, a, b string) string {
	A := strings.Index(s, a)
	B := strings.Index(s, b)
	return exchange(s, A, B)
}
