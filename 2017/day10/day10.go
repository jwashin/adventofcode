package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	input, _ := os.ReadFile("input.txt")
	fmt.Println("2.", part2(string(input)))
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), ",")
	lengths := []int{}
	for _, v := range data {
		d, _ := strconv.Atoi(v)
		lengths = append(lengths, d)
	}
	j := knotHash(lengths, false)
	return j[0] * j[1]
}

func part2(data string) string {
	lengths := []int{}
	for _, v := range data {
		lengths = append(lengths, int(v))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)
	z := knotHash64(lengths, false)
	newS := []int{}
	hxs := []int{}
	for _, v := range z {
		newS = append(newS, v)
		if len(newS) == 16 {
			hxs = append(hxs, xorList(newS))
			newS = []int{}
		}
	}
	d := ""
	for _, v := range hxs {
		r := fmt.Sprintf("%x", v)
		if len(r) == 1 {
			r = "0" + r
		}
		d += r
	}
	return d
}

func makeList(test bool) []int {
	out := []int{}
	n := 256
	if test {
		n = 5
	}
	x := 0
	for x < n {
		out = append(out, x)
		x += 1
	}
	return out
}

func xorList(s []int) int {
	x := s[0]
	for _, v := range s[1:] {
		x = x ^ v
	}
	return x
}

func knotHash64(lengths []int, test bool) []int {
	s := makeList(test)
	currentPosition := 0
	skipSize := 0
	for n := 0; n < 64; n++ {
		for _, length := range lengths {
			var selection []int
			if currentPosition+length > len(s)-1 {
				selection = s[currentPosition:]
				selection = append(selection, s[0:length-len(selection)]...)
			} else {
				selection = s[currentPosition : currentPosition+length]
			}
			for i, j := 0, len(selection)-1; i < j; i, j = i+1, j-1 {
				selection[i], selection[j] = selection[j], selection[i]
			}
			for i, v := range selection {
				idx := i + currentPosition
				if idx > len(s)-1 {
					idx -= len(s)
				}
				s[idx] = v
			}
			currentPosition += length + skipSize
			for currentPosition > len(s)-1 {
				currentPosition -= len(s)
			}
			skipSize += 1
		}
	}
	return s
}

func knotHash(lengths []int, test bool) []int {
	s := makeList(test)
	currentPosition := 0
	skipSize := 0
	for _, length := range lengths {
		var selection []int
		if currentPosition+length > len(s)-1 {
			selection = s[currentPosition:]
			selection = append(selection, s[0:length-len(selection)]...)
		} else {
			selection = s[currentPosition : currentPosition+length]
		}
		for i, j := 0, len(selection)-1; i < j; i, j = i+1, j-1 {
			selection[i], selection[j] = selection[j], selection[i]
		}
		for i, v := range selection {
			idx := i + currentPosition
			if idx > len(s)-1 {
				idx -= len(s)
			}
			s[idx] = v
		}
		currentPosition += length + skipSize
		for currentPosition > len(s)-1 {
			currentPosition -= len(s)
		}
		skipSize += 1
	}
	return s
}
