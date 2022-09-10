package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("1.", nextPassword("hepxcrrq"))
	fmt.Println("2.", nextPassword("hepxxyzz"))
}

func nextPassword(s string) string {
	t := increment(s)
	for !isValid(t) {
		t = increment(t)
	}
	return t
}

func increment(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	item := s[len(s)-1]
	alphaIndex := strings.Index(alphabet, string(item))
	if item != 'z' {
		return s[:len(s)-1] + string(alphabet[alphaIndex+1])
	}
	return increment(s[:len(s)-1]) + "a"
}

func isValid(s string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	if strings.ContainsAny(s, "iol") {
		return false
	}
	// need a pair aa,bb,cc, etc
	found := hasaPair(s, alphabet)
	if found == "" {
		return false
	}
	// need a second pair, not the same
	ix := strings.Index(alphabet, found)
	sa := alphabet[:ix] + alphabet[ix+1:]
	found = hasaPair(s, sa)
	if found == "" {
		return false
	}
	for ky, val := range alphabet[2:] {
		t := string(alphabet[ky+0]) + string(alphabet[ky+1]) + string(val)
		if strings.Contains(s, t) {
			return true
		}
	}
	return false
}

func hasaPair(s string, alphabet string) string {
	for _, val := range alphabet {
		if strings.Contains(s, string(val)+string(val)) {
			return string(val)
		}
	}
	return ""
}
