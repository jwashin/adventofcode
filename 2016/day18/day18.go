package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("1.", countSafeTiles(string(input), 40))
	fmt.Println("2.", makeLines2(string(input), 400000))
}

func trapRow(s string) string {
	w := fmt.Sprintf(".%s.", s)
	newString := []string{}
	// s := newString[len(newString)-1]
	for k := range w {
		if k > 0 && k < len(w)-1 {
			test3 := fmt.Sprintf("%c%c%c", w[k-1], w[k], w[k+1])
			if isTrap(test3) {
				newString = append(newString, "^")
				continue
			}
			newString = append(newString, ".")
		}
	}
	return strings.Join(newString, "")
}

func isTrap(s string) bool {
	for _, v := range []string{"^^.", ".^^", "^..", "..^"} {
		if s == v {
			return true
		}
	}
	return false
}

func makeLines(start string, n int) []string {
	out := []string{}
	out = append(out, start)
	newLine := start
	for len(out) < n {
		newLine = trapRow(newLine)
		out = append(out, newLine)
	}
	return out
}

func makeLines2(start string, n int) int {

	newLine := start
	safeCount := strings.Count(newLine, ".")
	out := 1
	for out < n {
		newLine = trapRow(newLine)
		safeCount += strings.Count(newLine, ".")
		out += 1
	}
	return safeCount
}

func countSafeTiles(start string, n int) int {
	t := makeLines(start, n)
	count := 0
	for _, v := range t {
		count += strings.Count(v, ".")
	}
	return count

}
