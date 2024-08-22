package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
}

func part1(s string) int {
	// 54128 too high
	// 53206 too high
	// 51697 too high
	// 45958 too high
	lines := strings.Split(s, "\n")
	// lines = append(lines, "")
	count := 0
	// for len(lines[0]) == 0 {
	// 	lines = lines[1:]
	// }
	item := []string{}
	for len(lines) > 0 {
		line := lines[0]
		if len(line) != 0 {
			item = append(item, line)
			lines = lines[1:]
		} else {

			// fmt.Println(len(lines))
			if len(item) > 0 {
				countHoriz := palindromeIndex(item)
				count += countHoriz * 100
				countVert := palindromeIndex(transpose(item))
				count += countVert
				item = []string{}
				lines = lines[1:]
				fmt.Println(countHoriz, countVert)
			}
		}
	}
	return count
}

func transpose(arr []string) []string {
	newArray := []string{}
	for range arr[0] {
		newArray = append(newArray, "")
	}
	for _, row := range arr {
		for i, v := range row {
			newArray[i] += string(v)
		}

	}
	return newArray
}

func fullPalindromeAt(arr []string, left int) bool {
	if left < 2 {
		return false
	}
	leftIndex := left
	rightIndex := left + 1
	ky := 0
	for {
		testa := leftIndex - ky
		testb := rightIndex + ky
		if testa < 0 || testb > len(arr)-1 {
			return false
		}
		if arr[testa] == arr[testb] {
			if testa == 0 || testb == len(arr)-1 {
				return true
			}

		}
		ky += 1
	}
}

func palindromeIndex(s []string) int {
	bestIndex := 0
	for i := 0; i < len(s); i++ {
		// Even-length palindrome
		palindrome := fullPalindromeAt(s, i)
		if palindrome {
			return i + 1
		}
	}
	return bestIndex
}
