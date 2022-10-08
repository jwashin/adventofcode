package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	input, _ := os.ReadFile("input.txt")
	data := string(input)
	fmt.Println("1.", part1(data, 20))
	fmt.Println("2.", part1(data, 50000000000))
}

func newGeneration(old map[int]string, rules map[string]string) map[int]string {
	// find minimum index. add 2 '.' below the minimum
	minIndex := math.MaxInt
	maxIndex := math.MinInt
	for k := range old {
		if k < minIndex {
			minIndex = k
		}
		if k > maxIndex {
			maxIndex = k
		}
	}

	// for string(b[minIndex])+string(b[minIndex+1]) != ".." {
	// 	b[minIndex-1] = '.'
	// 	minIndex = minIndex - 1
	// }
	new := map[int]string{}
	for potNumber := minIndex - 3; potNumber < maxIndex+3; potNumber++ {
		st := ""
		for idx := potNumber - 2; idx <= potNumber+2; idx++ {
			sn := old[idx]
			if sn == "" {
				sn = "."
			}
			st += sn
		}
		j := rules[st]

		if j == "#" {
			new[potNumber] = j
		}
	}

	// g := view(new)
	// fmt.Println(g)
	return new
}

func view(s map[int]string) string {
	min := math.MaxInt
	max := math.MinInt
	for k := range s {
		if k > max {
			max = k
		}
		if k < min {
			min = k
		}
	}
	out := []string{}
	for i := min; i <= max; i++ {
		t := s[i]
		if t == "" {
			t = "."
		}
		out = append(out, t)
	}
	return strings.Join(out, "")
}

// 4050000008879 too high
// 4050000000879 too high
// 4050000000798 just right

func part1(s string, generations int) int {
	initialState := ""
	rules := map[string]string{}
	for _, v := range strings.Split(s, "\n") {
		if strings.Contains(v, "initial") {
			t := strings.Fields(v)
			initialState = t[2]
		}
		v := strings.TrimSpace(v)
		if len(v) > 0 {
			t := strings.Fields(v)
			rules[t[0]] = t[2]
		}
	}
	state := initialState

	pots := map[int]string{}
	for k, v := range state {
		pots[k] = string(v)
	}
	// oss := view(pots)
	s100 := 0
	// newSum := sumPots(pots)
	for i := 0; i < generations; i++ {
		// oldSum := newSum
		pots = newGeneration(pots, rules)
		// if i < 1000 {
		// newSum = sumPots(pots)
		// nl := []int{}
		// for k := range pots {
		// 	nl = append(nl, k)
		// }
		// sort.Slice(nl, func(i2, j int) bool {
		// 	return nl[i2] < nl[j]
		// })
		// getting the function right
		// fmt.Println(i, newSum, s100+81*(i-100), s100+81*(i-100)-newSum)
		// }
		// sum difference between iterations is 81 after iteration 99
		if i == 100 {
			s100 = sumPots(pots)
		}
		if i > 100 {
			// generations is <, not <=
			return s100 + 81*(generations-1-100)
		}
	}

	return sumPots(pots)
}

func sumPots(pots map[int]string) int {
	sum := 0
	for k, v := range pots {
		if v == "#" {
			sum += k
		}
	}
	return sum
}
