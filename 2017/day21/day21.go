package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("test", test())
	fmt.Println("Part1.", part1())
	fmt.Println("Part2.", part2())
}

var start = ".#./..#/###"

func part1() int {
	input, _ := os.ReadFile("input.txt")
	rules := makeRules(string(input))
	j := stringGrid(strings.Split(start, "/"))
	for i := 0; i < 5; i++ {
		j = j.iter(rules)
	}
	count := 0
	for _, v := range j {
		count += strings.Count(v, "#")
	}
	return count
}
func part2() int {
	input, _ := os.ReadFile("input.txt")
	rules := makeRules(string(input))
	j := stringGrid(strings.Split(start, "/"))
	for i := 0; i < 18; i++ {
		j = j.iter(rules)
	}
	count := 0
	for _, v := range j {
		count += strings.Count(v, "#")
	}
	return count
}

func test() int {
	rules := makeRules(`../.# => ##./#../...
	.#./..#/### => #..#/..../..../#..#`)
	j := stringGrid(strings.Split(start, "/"))
	for i := 0; i < 2; i++ {
		j = j.iter(rules)
	}
	count := 0
	for _, v := range j {
		count += strings.Count(v, "#")
	}
	return count

}

func makeRules(s string) map[string]string {
	t := strings.Split(s, "\n")
	rules := map[string]string{}
	for _, v := range t {
		f := strings.Fields(v)
		i := f[0]
		x := f[2]
		// dunno if some of these are redundant or if we have them all
		rules[i] = x
		rules[rotate90(i)] = x
		rules[rotate90(rotate90(i))] = x
		// rules[rotate90(rotate90(rotate90(i)))] = x
		rules[flipEW(i)] = x
		rules[flipNS(i)] = x
		rules[flipNS(rotate90(i))] = x
		rules[flipEW(rotate90(i))] = x
		// rules[flipNS(rotate90(rotate90(i)))] = x
		// rules[flipEW(rotate90(rotate90(i)))] = x
	}
	return rules
}

func rotate90(s string) string {
	var newS []byte
	if len(s) == 5 {
		s = strings.ReplaceAll(s, "/", "")
		newS = []byte{s[2], s[0], '/', s[3], s[1]}
	}
	if len(s) == 11 {
		s = strings.ReplaceAll(s, "/", "")
		newS = []byte{s[6], s[3], s[0], '/', s[7], s[4], s[1], '/', s[8], s[5], s[2]}
	}
	return string(newS)
}

func flipNS(s string) string {
	var newS []byte
	s = strings.ReplaceAll(s, "/", "")
	if len(s) == 4 {
		newS = []byte{s[2], s[3], '/', s[0], s[1]}
	}
	if len(s) == 9 {
		s = strings.ReplaceAll(s, "/", "")
		newS = []byte{s[6], s[7], s[8], '/', s[3], s[4], s[5], '/', s[0], s[1], s[2]}
	}
	return string(newS)
}

func flipEW(s string) string {
	var newS []byte
	s = strings.ReplaceAll(s, "/", "")
	if len(s) == 4 {
		newS = []byte{s[1], s[0], '/', s[3], s[2]}
	}
	if len(s) == 9 {
		newS = []byte{s[2], s[1], s[0], '/', s[5], s[4], s[3], '/', s[8], s[7], s[6]}
	}
	return string(newS)
}

type stringGrid []string

func (s stringGrid) size() int {
	return len(s)
}

func (s stringGrid) makeKeys() []string {
	n := len(s)
	// n rows with D*n items
	// make nxn arrays going across, then keyize them
	keys := []string{}

	for len(s[0]) > 0 {
		square := stringGrid{}
		for j := range s {
			line, t := s[j][:n], s[j][n:]
			square = append(square, line)
			s[j] = t
		}
		keys = append(keys, strings.Join(square, "/"))
	}
	return keys
}

func (s stringGrid) iter(rules map[string]string) stringGrid {
	size := s.size()
	var n int
	if size%2 == 0 {
		n = 2
	} else {
		n = 3
	}
	out := stringGrid{}
	for len(s) >= n {
		// n rows at a time
		rk := stringGrid{}
		for row := 0; len(rk) < n; row++ {
			rk = append(rk, s[row])
		}
		s = s[n:]
		// I have a string grid of n rows
		kys := rk.makeKeys()
		// kys are "/" delimited squares
		rows := makeNewSquares(kys, rules)
		out = append(out, rows...)
	}
	return out
}

func makeNewSquares(kys []string, rules map[string]string) []string {
	squares := []string{}
	for _, v := range kys {
		// apply rule to get new list of patterns
		// for _, f := range strings.Split(v, "/") {
		squares = append(squares, matchAndGet(v, rules))
		// }
	}

	lines := []stringGrid{}
	for _, v := range squares {
		f := strings.Split(v, "/")
		j := stringGrid(f)
		lines = append(lines, j)
	}
	items := []string{}
	for i := 0; i < len(lines[0]); i++ {
		newString := ""
		for _, v := range lines {
			newString += v[i]
		}
		items = append(items, newString)
	}
	return items
}

func matchAndGet(input string, rules map[string]string) string {
	t := rules[input]
	if t != "" {
		return t
	}
	return fmt.Sprint("Can't match", input)
}
