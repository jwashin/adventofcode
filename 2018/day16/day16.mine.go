package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

// func main() {
// 	fmt.Println("1.", part1())
// 	fmt.Println(part2())
// }

func part1() int {
	samples, _ := parseInput()
	count := 0
	for _, v := range samples {
		if len(candidateOpcodes(v.before, v.a, v.b, v.c, v.after)) >= 3 {
			count += 1
		}
	}
	return count
}

type void struct{}
type intSet map[int]void

var member void

func (s intSet) add(i int) {
	s[i] = member
}
func (s intSet) remove(i int) {
	delete(s, i)
}
func (s intSet) items() []int {
	items := []int{}
	for k := range s {
		items = append(items, k)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	return items
}
func (s intSet) contains(i int) bool {
	for v := range s {
		if v == i {
			return true
		}
	}
	return false
}

// using the other guy's code, we see the 
// answer to part2 is 610, but his part1 was wrong...

// TODO: get the right answer

// 562 not right
// 581 not right
// 556 not right
// 553 not right
// 567 not right
// 528 too low
// 534 too low
// 550 too low
func part2() int {
	samples, program := parseInput()
	ops := map[int]string{}

	possibles := map[string]intSet{}
	for _, op := range opcodes {
		possibles[op] = intSet{}
	}
	for _, v := range samples {
		t := candidateOpcodes(v.before, v.a, v.b, v.c, v.after)
		s := strings.Join(t, " ")
		for _, k := range opcodes {
			if strings.Contains(s, k) {
				// possibles[k][v.opcode] = true
				possibles[k].add(v.opcode)
			}
		}
	}
	// ns := map[string][]int{}
	// for k, v := range possibles {
	// 	ns[k] = v.items()
	// }
	// fmt.Println(ns)
	possibles["gtir"].remove(0)
	possibles["gtrr"].remove(8)
	// possibles["eqri"].remove()

	for len(ops) < 16 {
		for oc := 0; oc < 16; oc++ {
			count := 0
			for _, pset := range possibles {
				if pset.contains(oc) {
					count += 1
				}
			}
			if count == 1 {
				ky := ""
				for opcode, pset := range possibles {
					if pset.contains(oc) {
						ops[oc] = opcode
						ky = opcode
						break
					}
				}
				delete(possibles, ky)
				for _, v := range possibles {
					v.remove(oc)
				}
			}
		}
	}

	// ops[2] = "bori"

	// ops[13] = "muli"
	// ops[11] = "mulr"
	// ops[9] = "borr"
	// ops[12] = "bani"
	// ops[10] = "addr"
	// ops[15] = "addi"
	// ops[6] = "seti"
	// ops[14] = "banr"
	// ops[1] = "setr"

	// ops[3] = "gtir"
	// ops[0] = "eqrr"
	// ops[8] = "gtrr"

	// ops[4] = "eqri"
	// ops[5] = "eqir"
	// ops[4] = "gtri"
	// ops[8] = "gtrr"

	// ops[7] = "eqir"
	// ops[1] = "setr"
	// ops[7] = "gtrr"

	// prev try
	// ops[7] = "eqri"
	// ops[4] = "gtri"
	// ops[3] = "gtrr"
	// ops[14] = "banr"
	// ops[1] = "setr"
	// ops[12] = "bani"
	// ops[15] = "addi"
	// ops[10] = "addr"
	// ops[9] = "borr"
	// ops[2] = "bori"
	// ops[11] = "mulr"
	// ops[13] = "muli"
	// ops[8] = "gtir"

	// for len(possibles) > 0 {
	// 	for k, v := range possibles {
	// 		if len(v) == 1 {
	// 			for x := range v {
	// 				currentValue = x
	// 				ops[currentValue] = k
	// 				// delete(v, x)
	// 				for _, d := range possibles {
	// 					d.remove(currentValue)
	// 					if len(v) == 0 {
	// 						delete(possibles, k)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	start := []int{0, 0, 0, 0}
	for _, v := range program {
		start = operation(start, ops[v[0]], v[1], v[2], v[3])
	}
	return start[0]

}

type sample struct {
	before          []int
	opcode, a, b, c int
	after           []int
}

func parseInput() ([]sample, [][]int) {
	samples := []sample{}
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")
	var item sample
	sampleStart := 0
	blanks := 0
	program := [][]int{}
	for k, v := range data {
		if strings.Contains(v, "Before") {
			sampleStart = k
			item = sample{}
			v = strings.ReplaceAll(v, ",", "")
			v = strings.ReplaceAll(v, "[", "")
			v = strings.ReplaceAll(v, "]", "")
			before := []int{}
			j := strings.Fields(v)
			for i := 1; i <= 4; i++ {
				d, _ := strconv.Atoi(j[i])
				before = append(before, d)
			}
			item.before = before
		}
		if k == sampleStart+1 {
			j := strings.Fields(v)
			opcode, _ := strconv.Atoi(j[0])
			a, _ := strconv.Atoi(j[1])
			b, _ := strconv.Atoi(j[2])
			c, _ := strconv.Atoi(j[3])
			item.opcode = opcode
			item.a = a
			item.b = b
			item.c = c
		}
		if k == sampleStart+2 {
			v = strings.ReplaceAll(v, ",", "")
			v = strings.ReplaceAll(v, "[", "")
			v = strings.ReplaceAll(v, "]", "")
			after := []int{}
			j := strings.Fields(v)
			for i := 1; i <= 4; i++ {
				d, _ := strconv.Atoi(j[i])
				after = append(after, d)
			}
			item.after = after
			samples = append(samples, item)
			blanks = 0
		}
		if strings.TrimSpace(v) == "" {
			blanks += 1
		}
		if blanks > 2 && len(strings.Fields(v)) > 0 {
			theList := []int{}
			j := strings.Fields(v)
			for i := 0; i < 4; i++ {
				d, _ := strconv.Atoi(j[i])
				theList = append(theList, d)
			}
			program = append(program, theList)
		}
	}
	return samples, program
}

var opcodes = []string{"addr", "addi", "mulr", "muli",
	"banr", "bani", "borr", "bori",
	"setr", "seti", "gtir", "gtri",
	"gtrr", "eqir", "eqri", "eqrr"}

func candidateOpcodes(before []int, a, b, c int, after []int) []string {
	candidates := []string{}
	for _, v := range opcodes {
		testItem := []int{}
		testItem = append(testItem, before...)
		test := operation(testItem, v, a, b, c)
		equal := true
		for i := range after {
			if test[i] != after[i] {
				equal = false
				break
			}
		}
		if equal {
			candidates = append(candidates, v)
		}
	}
	return candidates
}

func operation(register []int, opcode string, a int, b int, c int) []int {
	switch opcode {
	case "addr":
		register[c] = register[a] + register[b]
	case "addi":
		register[c] = register[a] + b
	case "mulr":
		register[c] = register[a] * register[b]
	case "muli":
		register[c] = register[a] * b
	case "banr":
		register[c] = register[a] & register[b]
	case "bani":
		register[c] = register[a] & b
	case "borr":
		register[c] = register[a] | register[b]
	case "bori":
		register[c] = register[a] | b
	case "setr":
		register[c] = register[a]
	case "seti":
		register[c] = a
	case "gtir":
		register[c] = 0
		if a > register[b] {
			register[c] = 1
		}
	case "gtri":
		register[c] = 0
		if register[a] > b {
			register[c] = 1
		}
	case "gtrr":
		register[c] = 0
		if register[a] > register[b] {
			register[c] = 1
		}
	case "eqir":
		register[c] = 0
		if a == register[b] {
			register[c] = 1
		}
	case "eqri":
		register[c] = 0
		if register[a] == b {
			register[c] = 1
		}
	case "eqrr":
		register[c] = 0
		if register[a] == register[b] {
			register[c] = 1
		}
	}
	newRegisters := []int{}
	newRegisters = append(newRegisters, register...)

	return newRegisters
}
