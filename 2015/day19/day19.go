package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

type rule struct {
	from string
	to   string
}

type rules []rule

func (r rules) Len() int {
	return len(r)
}
func (r rules) Less(i, j int) bool {
	return len(r[i].to) < len(r[j].to)
}
func (r rules) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
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

func part2() int {
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
	return greedyLeast(originalMolecule, rules, "e")
}

func fewestSteps(target string, rules []rule, start string) int {

	// do Dijkstra, maybe

	currentValue := 0
	currentItem := start
	targetLen := len(target)

	q := map[string]int{currentItem: 0}

	for len(q) > 0 {
		// find least of q
		min := math.MaxInt
		for k, val := range q {
			if val < min {
				min = val
				currentItem = k
				currentValue = val
			}
		}
		delete(q, currentItem)
		for _, rule := range rules {
			t := singleReplacement(currentItem, rule.from, rule.to)
			for _, newString := range t {
				if len(newString) <= targetLen && len(newString) > 0 {
					if newString == target {
						return currentValue + 1
					}
					q[newString] = currentValue + 1
				}
			}
		}
	}
	return -1
}

func fewestStepsA(target string, rules []rule, start string) int {

	// do Dijkstra, maybe

	currentValue := 0
	currentItem := start

	q := map[string]int{currentItem: 0}

	for len(q) > 0 {
		// find least of q
		min := math.MaxInt
		for k, val := range q {
			if len(k) < min {
				min = val
				currentItem = k
				currentValue = val
			}
		}
		delete(q, currentItem)
		for _, rule := range rules {
			t := singleUnReplacement(currentItem, rule.to, rule.from)
			for _, newString := range t {
				if len(newString) > 0 {
					if newString == target {
						return currentValue + 1
					}
					q[newString] = currentValue + 1
				}
			}
		}
	}
	return -1
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

func singleUnReplacement(s string, fromString string, toString string) []string {
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
				rep := toString
				if i == act {
					rep = fromString
				}
				nl = append(nl, strings.ReplaceAll(s3[i], "ZZ", rep))
			}
		}
		out = append(out, strings.Join(nl, ""))
	}
	return out
}

func greedyLeast(molecule string, rules rules, target string) int {
	//  Hmmm. this repeats NRnBSiRnCaRnFArYFArFArF 210
	// but the python app works
	steps := 0
	// sort.Sort(rules)
	for molecule != "e" {
		for _, v := range rules {
			for {
				new := strings.Replace(molecule, v.to, v.from, 1)
				if molecule != new {
					steps += 1
					molecule = new
				} else {
					break
				}
			}
		}
		fmt.Println(molecule, steps)
	}
	return steps

}
