package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
}

type rule struct {
	from string
	to   string
}

func moleculeCount(s string, rules []rule) int {
	originalMolecule := s
	molecules := map[string]bool{}
	for _, r := range rules {
		t := singleReplacement(originalMolecule, r.from, r.to)
		for _, k := range t {
			molecules[k] = true
		}
	}
	return len(molecules)
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")
	rules := []rule{}
	originalMolecule := ""
	for _, t := range data {
		item := strings.TrimSpace(t)
		if strings.Contains(item, "=>") {
			d := strings.Split(item, " ")
			newRule := rule{strings.TrimSpace(d[0]), strings.TrimSpace(d[2])}
			rules = append(rules, newRule)
			continue
		}
		if len(item) == 0 {
			continue
		}
		originalMolecule = strings.TrimSpace(item)
	}
	return moleculeCount(originalMolecule, rules)

}

func singleReplacement(s string, fromString string, toString string) []string {
	s2 := strings.ReplaceAll(s, fromString, "ZZ#")
	howmany := strings.Count(s2, "ZZ")
	if howmany == 1 {
		return []string{strings.ReplaceAll(s, fromString, toString)}
	}
	s3 := strings.Split(s2, "#")
	out := []string{}
	for act := 0; act < howmany; act++ {
		nl := []string{}
		for i, v := range s3 {
			if len(v) > 0 {
				rep := fromString
				if i == act {
					rep = toString
				}
				nl = append(nl, strings.ReplaceAll(s3[i], "ZZ", rep))
			}
		}
		out = append(out, strings.Join(nl, ""))
	}
	return out
}

// https://code-maven.com/slides/golang/solution-permutations
func permutations(word string) []string {
	if word == "" {
		return []string{""}
	}
	perms := []string{}
	for i, rn := range word {
		rest := word[:i] + word[i+1:]
		//fmt.Println(rest)
		for _, result := range permutations(rest) {
			perms = append(perms, fmt.Sprintf("%c", rn)+result)
		}
		//perms = append(perms, fmt.Sprintf("%c\n", rn))
	}
	return perms
}
