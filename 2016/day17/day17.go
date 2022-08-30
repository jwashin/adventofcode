package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := "pxxbnzuo"
	fmt.Println("1.", findShortestPath(input))
	fmt.Println("2. ", findLongestPathLength(input))
}

func possibleNexts(path string) []string {
	s := calcMD5(path)
	out := []string{}
	for ky, val := range []string{"U", "D", "L", "R"} {
		if strings.Contains("bcdef", string(s[ky])) {
			out = append(out, val)
		}
	}
	return out

}

func findShortestPath(passcode string) string {
	startPath := ""
	endx := 3
	endy := 3
	// UDDD:4
	q := map[string]int{startPath: 0}
	currentPath := ""
	for len(q) > 0 {
		// find shortest path in q and use it
		min := math.MaxInt
		for ky, v := range q {
			if v < min {
				min = v
				currentPath = ky
			}
		}
		delete(q, currentPath)
		for _, v := range possibleNexts(passcode + currentPath) {
			x, y := location(currentPath + v)
			if x == endx && y == endy {
				return currentPath + v
			}
			if x >= 0 && y >= 0 {
				newpath := currentPath + v
				q[newpath] = len(newpath)
			}
		}

	}
	return ""
}

func findLongestPathLength(passcode string) int {
	startPath := ""
	endx := 3
	endy := 3
	// done := map[string]int{}
	doneLength := 0
	// donePath := ""

	// UDDD:4
	q := map[string]int{startPath: 0}
	currentPath := ""
	for len(q) > 0 {
		// find shortest path in q and use it
		min := math.MaxInt
		for ky, v := range q {
			if v < min {
				min = v
				currentPath = ky
			}
		}
		delete(q, currentPath)
		for _, v := range possibleNexts(passcode + currentPath) {
			testpath := currentPath + v
			x, y := location(testpath)
			if x == endx && y == endy {
				// we got to the end!
				// done[testpath] = len(testpath)
				tl := len(testpath)
				if tl > doneLength {
					// donePath = testpath
					doneLength = tl
				}

			}
			if (x >= 0 && y >= 0) && !(x == 3 && y == 3) {
				q[testpath] = len(testpath)
			}
		}

	}
	return doneLength
}

func calcMD5(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func location(path string) (int, int) {
	x := 0
	y := 0
	for _, v := range path {
		if string(v) == "U" {
			y -= 1
			if y < 0 {
				return -1, -1
			}
		}
		if string(v) == "D" {
			y += 1
			if y > 3 {
				return -1, -1
			}
		}
		if string(v) == "L" {
			x -= 1
			if x < 0 {
				return -1, -1
			}
		}
		if string(v) == "R" {
			x += 1
			if x > 3 {
				return -1, -1
			}
		}
	}
	return x, y
}
