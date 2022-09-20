package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.TrimSpace(string(input))
	fmt.Println("1.", countMatchesNext(data))
	fmt.Println("2.", countMatchesHalfway(data))
}

func countMatchesNext(s string) int {
	endString := string(s[0])
	testString := s + endString
	count := 0
	ix := 0
	for ix <= len(s)-1 {
		if testString[ix+1] == testString[ix] {
			d, _ := strconv.Atoi(string(testString[ix]))
			count += d
		}
		ix += 1
	}
	return count
}

func countMatchesHalfway(s string) int {
	count := 0
	ix := 0
	length := len(s)
	lastIdx := len(s) - 1
	for ix <= lastIdx {
		compIndex := ix + length/2
		if compIndex > lastIdx {
			compIndex -= length
		}
		if s[compIndex] == s[ix] {
			d, _ := strconv.Atoi(string(s[ix]))
			count += d
		}
		ix += 1
	}
	return count
}
