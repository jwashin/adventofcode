package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", part1(string(input)))
	fmt.Println("part 2:", part2(string(input)))
}

// 134 too low
// 4418 too low

var revop = map[string]string{"+": "-", "-": "+", "*": "/", "/": "*"}

func part2(s string) int {
	data := strings.Split(s, "\n")
	values := map[string]int{}
	fns := map[string][3]string{}
	for _, v := range data {
		j := strings.Fields(v)
		key := j[0][:4]
		if len(j) == 2 {
			value, _ := strconv.Atoi(j[1])
			values[key] = value
		}
	}
	for _, v := range data {
		j := strings.Fields(v)
		key := j[0][:4]
		if len(j) == 4 {
			a := j[1]
			op := j[2]
			b := j[3]
			t := [3]string{}
			t[0] = a
			t[1] = op
			t[2] = b
			fns[key] = t

			// t = [3]string{}
			// t[0] = key
			// t[1] = revop[op]
			// t[2] = b
			// fns[a] = t

			// if strings.Contains("*+", op) {
			// 	t = [3]string{}
			// 	t[0] = key
			// 	t[1] = revop[op]
			// 	t[2] = a
			// 	fns[b] = t
			// }

			// if strings.Contains("-/", op) {
			// 	t = [3]string{}
			// 	t[0] = a
			// 	t[1] = op
			// 	t[2] = key
			// 	fns[b] = t
			// }
		}
	}
	delete(values, "humn")
	// the second one is for test
	// fns["humn"] = [3]string{"zdwl", "+", "bnrn"}
	fns["humn"] = [3]string{"dvpt", "+", "ptdq"}

	q := []string{}
	currItem := ""
	for ky := range fns {
		if _, ok := values[ky]; !ok {
			q = append(q, ky)
		}
	}
	for len(q) > 0 {
		currItem = q[0]
		q = q[1:]

		_, ok0 := values[currItem]

		v1, v2 := 0, 0
		ok1, ok2 := false, false

		v1s := fns[currItem][0]
		v2s := fns[currItem][2]
		op := fns[currItem][1]

		v1, ok1 = values[v1s]
		v2, ok2 = values[v2s]

		if ok1 && ok2 {
			// we have both. perform operation
			value := 0
			if op == "*" {
				value = v1 * v2
			} else if op == "+" {
				value = v1 + v2
			} else if op == "/" {
				value = v1 / v2
			} else if op == "-" {
				value = v1 - v2
			}
			if currItem == "humn" {
				return value
			}
			values[currItem] = value

		} else if (ok1 || ok2) && currItem == "root" {
			if ok1 {
				values[v2s] = v1
				// q = append(q, v1s)
				// q = append(q, v2s)
			}
			if ok2 {
				values[v1s] = v2
				// q = append(q, v2s)
				// q = append(q, v1s)
			}
		}
		if ok0 && ok1 {
			fmt.Println("ok1")

			//  in c = a $ b
			// we have c and a, need b
			t := [3]string{}

			a := v1s
			key := v2s
			b := currItem

			if strings.Contains("*+", op) {
				t[0] = key
				t[1] = revop[op]
				t[2] = a
				fns[b] = t
			}

			if strings.Contains("-/", op) {
				t[0] = a
				t[1] = op
				t[2] = key
				fns[b] = t
			}
			// q = append(q, currItem)

		}
		if ok0 && ok2 {
			fmt.Println("ok2")
			t := [3]string{}
			t[0] = currItem
			t[1] = revop[op]
			t[2] = v2s
			fns[v1s] = t
			// q = append(q, v1s)
			// q = append(q, currItem)
		} else {
			q = append(q, currItem)
		}

		// if ok1 || ok2 {
		// 	if ok1 {
		// 	}
		// 	if ok2 {
		// 	}
		// 	// perform the reverse operation, and we're done
		// }

	}
	// remaining := q[0]
	// r := values[remaining]
	// st := values[fns[remaining][2]]
	// fmt.Println(remaining, fns[remaining])
	// fmt.Println(remaining, "(", r, ")", fns[remaining], st)
	// return r + st
	return 1
}

func part1(s string) int {
	data := strings.Split(s, "\n")
	values := map[string]int{}
	fns := map[string][3]string{}

	for _, v := range data {
		j := strings.Fields(v)
		key := j[0][:4]
		if len(j) == 2 {
			value, _ := strconv.Atoi(j[1])
			values[key] = value
		}
		if len(j) == 4 {
			t := [3]string{}
			t[0] = j[1]
			t[1] = j[2]
			t[2] = j[3]
			fns[key] = t
		}
	}
	q := []string{}
	currItem := ""
	for ky := range fns {
		q = append(q, ky)
	}
	for len(q) > 0 {
		currItem = q[0]
		q = q[1:]

		v1, v2 := 0, 0
		ok1, ok2 := false, false

		v1s := fns[currItem][0]
		v2s := fns[currItem][2]
		op := fns[currItem][1]

		v1, ok1 = values[v1s]
		v2, ok2 = values[v2s]

		if ok1 && ok2 {
			// perform operation
			value := 0
			if op == "*" {
				value = v1 * v2
			} else if op == "+" {
				value = v1 + v2
			} else if op == "/" {
				value = v1 / v2
			} else if op == "-" {
				value = v1 - v2
			}
			// add value to values
			values[currItem] = value
			if currItem == "root" {
				return value
			}
		} else {
			q = append(q, currItem)
		}
	}
	return -1
}
