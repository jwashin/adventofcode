package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	fmt.Println(scoreIllegals(string(data)))
	fmt.Println(scoreIncompletes(string(data)))
}

func scoreIllegals(aString string) int {
	data := strings.Split(aString, "\n")
	syntaxScore := 0

	for _, val := range data {
		t := findFirstIllegalCharacter(val)
		if t == ")" {
			syntaxScore += 3
		}
		if t == "]" {
			syntaxScore += 57
		}
		if t == "}" {
			syntaxScore += 1197
		}
		if t == ">" {
			syntaxScore += 25137
		}

	}
	return syntaxScore
}

func scoreIncompletes(aString string) int {
	data := strings.Split(aString, "\n")
	completionScores := []int{}

	for _, val := range data {
		CValue := 0
		res := findFirstIllegalCharacter(val)
		if len(res) == 1 {
			continue
		}
		for _, z := range res {
			t := string(z)
			if t == "(" {
				CValue = CValue*5 + 1
			}
			if t == "[" {
				CValue = CValue*5 + 2
			}
			if t == "{" {
				CValue = CValue*5 + 3
			}
			if t == "<" {
				CValue = CValue*5 + 4
			}
		}
		completionScores = append(completionScores, CValue)

	}
	sort.Ints(completionScores)
	loc := int(len(completionScores) / 2)
	v := completionScores[loc]
	return v
}

func findFirstIllegalCharacter(aString string) string {
	// openers := "([{<"
	closers := ")]}>"
	for idx, val := range aString {
		v := string(val)
		if strings.Contains(closers, v) {
			if idx == 0 {
				return v
			}
			opener := openerFor(v)
			if string(aString[idx-1]) == opener {
				rm := opener + v
				newString := strings.Replace(aString, rm, "", 1)
				return findFirstIllegalCharacter(newString)
			} else {
				return v
			}

		}
	}
	l := len(aString) - 1
	ns := ""
	for idx := range aString {
		ns += string(aString[l-idx])
	}
	return ns
}

func openerFor(closer string) string {
	closers := ")]}>"
	openers := "([{<"
	for idx, val := range closers {
		if string(val) == closer {
			return string(openers[idx])
		}
	}
	return ""
}
