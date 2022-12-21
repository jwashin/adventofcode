package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println(part1(string(input)))
	fmt.Println(part2(string(input)))
}

func part2(s string) int {
	data := strings.Split(s, "\n")
	humn := -1
	for {
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
		humn += 1
		values["humn"] = humn
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
				// if currItem == "vslw"{
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "hpnl" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "whwg" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "fhhm" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "zlmt" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "fgvl" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "hstf" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "stzw" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "mzsz" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "vbbc" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "nhvm" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "bfcg" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "vbdq" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "ttgc" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "tztg" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "bbgj" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "gzdf" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "rlwp" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "drdn" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "bsbv" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "fgrn" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "mlvd" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "wtzd" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "ntnz" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "frpz" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "pjbq" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "dzbb" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "frlr" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "glgl" {
				// 	fmt.Println(v1, v2)
				// }
				// if currItem == "sbtm" {
				// 	fmt.Println(v1, v2)
				// }
				if currItem == "root" {
					if v1 == v2 {
						return values["humn"]
					}
				}
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
			} else {
				q = append(q, currItem)
			}
		}
	}
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
