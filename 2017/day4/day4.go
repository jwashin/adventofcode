package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("1.", countValid(string(input)))
	fmt.Println("2.", countValid2(string(input)))
}

func countValid(passphrases string) int {
	p := strings.Split(passphrases, "\n")
	count := 0
	for _, v := range p {
		if isValid(v) {
			count += 1
		}
	}
	return count
}
func countValid2(passphrases string) int {
	p := strings.Split(passphrases, "\n")
	count := 0
	for _, v := range p {
		if isValid2(v) {
			count += 1
		}
	}
	return count
}

func alphabetizeWord(s string) string {
	t := []string{}
	for _, v := range s {
		t = append(t, string(v))
	}
	sort.Strings(t)
	return strings.Join(t, "")
}

func alphabetizeIndividualWords(passphrase string) string {
	d := strings.Fields(passphrase)
	out := []string{}
	for _, v := range d {
		s := alphabetizeWord(v)
		out = append(out, s)
	}
	return strings.Join(out, " ")
}

func isValid2(passphrase string) bool {
	s := alphabetizeIndividualWords(passphrase)
	return isValid(s)
}

func isValid(passphrase string) bool {
	d := strings.Fields(passphrase)
	td := map[string]int{}
	for _, item := range d {
		td[item] += 1
	}
	for _, v := range td {
		if v > 1 {
			return false
		}
	}
	return true
}
