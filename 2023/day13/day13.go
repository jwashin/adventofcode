package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
}

func part2(s string) int {
	// 54128 too high
	// 53206 too high
	// 51697 too high
	// 45958 too high
	// 45138 wrong
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
				horiz := item
				vert := transpose(item)

				countHoriz := palindromeIndex(horiz)
				count += countHoriz * 100
				countVert := palindromeIndex(vert)
				count += countVert

				lines = lines[1:]

				if countHoriz == 0 && countVert == 0 {
					fmt.Println("boom")
				}
				item = []string{}
				fmt.Println(countHoriz, countVert)
			}
		}
	}
	return count
}

func part1(s string) int {
	// 54128 too high
	// 53206 too high
	// 51697 too high
	// 45958 too high
	// 45138 wrong
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
				horiz := item
				vert := transpose(item)

				countHoriz := palindromeIndex(horiz)
				count += countHoriz * 100
				countVert := palindromeIndex(vert)
				count += countVert

				lines = lines[1:]

				if countHoriz == 0 && countVert == 0 {
					fmt.Println("boom")
				}
				item = []string{}
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

func oneCharDifferent(a string, b string) bool {
	if a == b {
		return false
	}
	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff += 1
			if diff > 1 {
				return false
			}
		}
	}
	return diff > 0
}

func fullPalindromeAt(arr []string, left int) bool {
	leftStart := left
	// we are only working with "even" palindromes here
	rightStart := left + 1
	increment := 0
	leftIndex, rightIndex := leftStart, rightStart
	if leftIndex < 0 || rightIndex > len(arr)-1 {
		return false
	}
	for arr[leftIndex] == arr[rightIndex] {
		// If we have arrived at a left or right boundary, good.
		if leftIndex == 0 || rightIndex == len(arr)-1 {
			return true
		}
		increment += 1
		leftIndex = leftStart - increment
		rightIndex = rightStart + increment

	}
	return false
}

func palindromeIndex(s []string) int {
	bestIndex := 0
	for i := 0; i < len(s); i++ {
		palindrome := fullPalindromeAt(s, i)
		if palindrome {
			return i + 1
		}
	}
	return bestIndex
}

func toNumbers(item []string) []int {
	newList := []int{}
	for _, v := range item {
		v = strings.ReplaceAll(v, "#", "1")
		v = strings.ReplaceAll(v, ".", "0")
		i, err := strconv.ParseInt(v, 2, 64)
		if err == nil {
			newList = append(newList, int(i))
		}
	}
	return newList
}
