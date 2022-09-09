package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")
	c := 0
	m := 0

	for _, v := range data {
		cc, mc := strCounts(v)
		c += cc
		m += mc
	}
	fmt.Println("1.", c-m)
	c = 0
	m = 0

	for _, v := range data {
		cc, mc := strunCounts(v)
		c += cc
		m += mc
	}
	fmt.Println("2.", m-c)

}

// 1.
// 1206 too low
// 1159 too low

func strCounts(s string) (int, int) {

	codelen := len(s)
	stringLength := len(unescape(s))

	return codelen, stringLength
}

func strunCounts(s string) (int, int) {

	codelen := len(s)
	stringLength := len(escape(s))

	return codelen, stringLength
}

func escape(s string) string {

	r := s

	r = strings.ReplaceAll(r, string('\\'), "\\"+"\\")

	// replace " with \"
	r = strings.ReplaceAll(r, string('"'), "\\"+string('"'))

	// put surrounding quotes back on
	r = string('"') + r + string('"')

	fmt.Println(s)
	fmt.Println(r)

	return r
}

func unescape(s string) string {

	reg, _ := regexp.Compile(`\\x[0123456789abcdef]{2}`)

	r := s
	// \xdd to char
	y := reg.FindAllString(r, -1)
	if len(y) > 0 {
		// charcode := 0
		for _, v := range y {

			r = strings.Replace(r, v, "_", 1)
		}
	}

	// fix surrounding quotes
	if (r[0]) == '"' && r[len(r)-1] == '"' {
		r = r[1 : len(r)-1]
	}
	// replace \\ with \
	r = strings.ReplaceAll(r, "\\"+"\\", string('\\'))

	// replace \" with "
	r = strings.ReplaceAll(r, "\\"+string('"'), string('"'))
	fmt.Println(s)
	fmt.Println(r)
	return r
}
