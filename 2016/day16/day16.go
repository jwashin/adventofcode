package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("1. ", filldisk("01110110101001000", 272))
	fmt.Println("2. ", filldisk("01110110101001000", 35651584))
}

func dragonCurve2(s string) string {
	a := s
	bprime := []string{}
	ln := len(s)
	for i := range s {
		if s[ln-i-1] == '0' {
			bprime = append(bprime, "1")
			continue
		}
		bprime = append(bprime, "0")
	}
	return a + "0" + strings.Join(bprime, "")

}

func checkSum2(s string) string {
	checksum := []string{}
	pair := ""
	for _, v := range s {
		pair += string(v)
		if len(pair) == 2 {
			if pair[0] == pair[1] {
				checksum = append(checksum, "1")
			} else {
				checksum = append(checksum, "0")
			}
			pair = ""
		}
	}
	t := strings.Join(checksum, "")
	if len(checksum)%2 == 1 {
		return t
	}
	return checkSum2(t)
}

func filldisk(init string, length int) string {
	c := init
	// fmt.Println("calc dragons")
	for len(c) < length {
		c = dragonCurve2(c)
		// fmt.Println(len(c))
	}
	// fmt.Println("calc checksum")
	c = c[:length]
	return checkSum2(c)
}
