package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", monkeyBusiness(string(input), 20, false))
	fmt.Println("Part2:", monkeyBusiness(string(input), 10000, true))
}

func monkeyBusiness(s string, rounds int, part2 bool) uint {
	monkeys := parseMonkeys(s)

	for round := 1; round <= rounds; round++ {
		for _, v := range monkeys {
			throws := v.doTurn(part2)
			for _, w := range throws {
				d := monkeys[w.destination]
				d.items = append(d.items, w.item)
			}
		}
		if round == 1 || round == 20 || round%1000 == 0 {
			fmt.Println("== after round", round, "==")
			for _, v := range monkeys {
				fmt.Println("Monkey", v.id, "inspected items", v.inspections, "times.")
			}
		}
	}
	ins := []int{}
	for _, v := range monkeys {
		ins = append(ins, v.inspections)
	}
	sort.Ints(ins)
	return uint(ins[len(ins)-1]) * uint(ins[len(ins)-2])
}

type Monkey struct {
	id          int
	items       []uint
	op          string
	test        string
	t           int
	f           int
	inspections int
}

type throw struct {
	item        uint
	destination int
}

func (m *Monkey) doTurn(part2 bool) []throw {
	throws := []throw{}
	for _, v := range m.items {
		m.inspections += 1
		nv := m.doOp(v)
		if !part2 {
			nv = nv / 3
		}
		f := strings.Fields(m.test)
		d1, _ := strconv.Atoi(f[len(f)-1])
		d := uint(d1)
		if nv%d == 0 {
			throws = append(throws, throw{nv, m.t})
		} else {
			throws = append(throws, throw{nv, m.f})
		}
	}
	m.items = []uint{}
	return throws
}

func (m *Monkey) doOp(item uint) uint {
	f := strings.Fields(m.op)
	first, second, val := uint(0), uint(0), uint(0)
	if f[0] == "old" {
		first = item
	} else {
		irst, _ := strconv.Atoi(f[0])
		first = uint(irst)
	}
	if f[2] == "old" {
		second = item
	} else {
		econd, _ := strconv.Atoi(f[2])
		second = uint(econd)
	}
	if f[1] == "*" {
		val = first * second
	}
	if f[1] == "+" {
		val = first + second
	}
	return val
}

func parseMonkeys(s string) []*Monkey {

	monkeys := []*Monkey{}

	id, t, f := 0, 0, 0
	items := []uint{}
	op, test := "", ""

	for _, v := range strings.Split(s, "\n") {
		v = strings.TrimSpace(v)

		if strings.Contains(v, "Monkey") {
			id, _ = strconv.Atoi(string(v[7]))
		}
		if strings.Contains(v, "Starting") {
			items = []uint{}
			ix := strings.Index(v, ":")
			z := v[ix+1:]
			z = strings.ReplaceAll(z, ",", "")
			j := strings.Fields(z)
			for _, w := range j {
				d, _ := strconv.Atoi(w)
				items = append(items, uint(d))
			}
		}
		if strings.Contains(v, "Operation") {
			ix := strings.Index(v, "=")
			op = v[ix+1:]
		}
		if strings.Contains(v, "Test") {
			ix := strings.Index(v, ":")
			test = v[ix+1:]
		}
		if strings.Contains(v, "true") {
			ix := string(v[len(v)-1])
			t, _ = strconv.Atoi(ix)
		}
		if strings.Contains(v, "false") {
			ix := string(v[len(v)-1])
			f, _ = strconv.Atoi(ix)
			monkey := Monkey{id: id, items: items, op: op, test: test, t: t, f: f, inspections: 0}
			monkeys = append(monkeys, &monkey)
		}
	}
	return monkeys
}
