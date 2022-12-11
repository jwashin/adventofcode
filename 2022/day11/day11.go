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

func monkeyBusiness(s string, rounds int, part2 bool) int {
	monkeys := parseMonkeys(s)

	for round := 1; round <= rounds; round++ {
		for _, v := range monkeys {
			throws := v.doTurn(part2)
			for _, w := range throws {
				d := monkeys[w.destination]
				d.items = append(d.items, w.item)
			}
		}
		// if round == 0 || round == 1 || round == 20 || round%1000 == 0 {
		// 	fmt.Println("== after round", round, "==")
		// 	for _, v := range monkeys {
		// 		fmt.Println("Monkey", v.id, "inspected items", v.inspections, "times.")
		// 	}
		// }
	}
	ins := []int{}
	for _, v := range monkeys {
		ins = append(ins, v.inspections)
	}
	sort.Ints(ins)
	return int(ins[len(ins)-1]) * int(ins[len(ins)-2])
}

type Monkey struct {
	id          int
	items       []int
	op          string
	test        int
	t           int
	f           int
	inspections int
	// needed least common multiple for part 2.
	// It's product of all the monkeys' test factors
	// We mod the thrown values by this to keep them from
	// getting too big
	lcm int
}

type throw struct {
	item        int
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
		d := m.test
		if nv%d == 0 {
			throws = append(throws, throw{nv, m.t})
		} else {
			throws = append(throws, throw{nv, m.f})
		}
	}
	m.items = []int{}
	return throws
}

func (m *Monkey) doOp(item int) int {
	f := strings.Fields(m.op)
	var first, second, val int
	if f[0] == "old" {
		first = item
	} else {
		first, _ = strconv.Atoi(f[0])
	}
	if f[2] == "old" {
		second = item
	} else {
		second, _ = strconv.Atoi(f[2])
	}
	if f[1] == "*" {
		val = first * second
	}
	if f[1] == "+" {
		val = first + second
	}
	return val % m.lcm
}

func parseMonkeys(s string) []*Monkey {

	monkeys := []*Monkey{}

	id, t, f, test := 0, 0, 0, 0
	items := []int{}
	op := ""
	lcm := 1
	for _, v := range strings.Split(s, "\n") {
		v = strings.TrimSpace(v)

		if strings.Contains(v, "Monkey") {
			id, _ = strconv.Atoi(string(v[7]))
		}
		if strings.Contains(v, "Starting") {
			items = []int{}
			ix := strings.Index(v, ":")
			z := v[ix+1:]
			z = strings.ReplaceAll(z, ",", "")
			j := strings.Fields(z)
			for _, w := range j {
				d, _ := strconv.Atoi(w)
				items = append(items, int(d))
			}
		}
		if strings.Contains(v, "Operation") {
			ix := strings.Index(v, "=")
			op = v[ix+1:]
		}
		if strings.Contains(v, "Test") {
			f1 := strings.Fields(v)
			test, _ = strconv.Atoi(f1[3])
			lcm *= test
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
	for _, v := range monkeys {
		v.lcm = lcm
	}
	return monkeys
}
