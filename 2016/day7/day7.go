package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	supporters := 0
	input, _ := os.ReadFile("input.txt")
	for _, line := range strings.Split(string(input), "\n") {
		data := strings.TrimSpace(line)
		if supportsTLS(data) {
			supporters += 1
		}
	}
	fmt.Println("1. ", supporters)
	supporters = 0
	for _, line := range strings.Split(string(input), "\n") {
		data := strings.TrimSpace(line)
		t := bracketParse(data)
		if t.supportsSSL() {
			supporters += 1
		}
	}
	fmt.Println("2. ", supporters)
}

func has4palindrome(s string) bool {
	for idx := range s {
		if idx+3 < len(s) {
			if s[idx] == s[idx+3] &&
				s[idx+1] == s[idx+2] &&
				s[idx] != s[idx+2] {
				return true
			}
		}
	}
	return false
}

func getABAs(s string) []string {
	out := []string{}
	for idx := range s {
		if idx+2 < len(s) {
			if s[idx] == s[idx+2] &&
				s[idx] != s[idx+1] {
				out = append(out, s[idx:idx+3])
			}
		}
	}
	return out
}

func aba2bab(s string) string {
	return string(s[1]) + string(s[0]) + string(s[1])
}

func supportsTLS(s string) bool {
	data := bracketParse(s)
	for _, item := range data.inside {
		if has4palindrome(item) {
			return false
		}
	}
	for _, item := range data.outside {
		if has4palindrome(item) {
			return true
		}
	}
	return false
}

type BracketParsed struct {
	inside  []string
	outside []string
}

func (t BracketParsed) supportsSSL() bool {
	for _, s := range t.outside {
		abas := getABAs(s)
		for _, aba := range abas {
			for _, bab := range t.inside {
				if strings.Contains(bab, aba2bab(aba)) {
					return true
				}
			}
		}
	}
	return false
}

func bracketParse(s string) BracketParsed {
	var out BracketParsed
	newString := ""
	inBracket := false
	for _, val := range s {
		tst := string(val)
		if tst != "[" && tst != "]" && !inBracket {
			newString += tst
			continue
		}
		if tst == "[" {
			inBracket = true
			out.outside = append(out.outside, newString)
			newString = ""
			continue
		}
		if tst == "]" {
			inBracket = false
			out.inside = append(out.inside, newString)
			newString = ""
			continue
		}
		newString += tst
	}
	if len(newString) > 0 {
		out.outside = append(out.outside, newString)
	}
	return out
}
