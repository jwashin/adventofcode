package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	// fmt.Println(getTestValue(string(data), 10))
	fmt.Println(makeDistributionCounts(string(data), 40))
}

func getTestValue(data string, count int) int {
	data = doPairInsert(data, count)
	distribution := map[string]int{}
	for _, v := range data {
		distribution[string(v)] += 1
	}
	min := math.MaxInt
	max := 0
	for _, v := range distribution {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

func GetValueFunction(distribution map[string]int) int {
	min := math.MaxInt
	max := 0
	for _, v := range distribution {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

func makeDistributionCounts(data string, steps int) int {
	seed, rules := parse(data)
	polymerDict := map[string]int{}
	first := getPairs(seed)
	for _, v := range first {
		polymerDict[v] += 1
	}
	for i := 0; i < steps; i++ {
		polymerDict = growPolymer(polymerDict, rules)
	}
	atoms := statsPolymer(polymerDict, seed)
	return GetValueFunction(atoms)
}

func growPolymer(aDict map[string]int, rules map[string]string) map[string]int {
	newDict := map[string]int{}
	for k, v := range aDict {
		a := string(k[0]) + rules[k]
		b := rules[k] + string(k[1])
		newDict[a] += v
		newDict[b] += v
	}
	return newDict
}

func statsPolymer(aDict map[string]int, seed string) map[string]int {
	atoms := map[string]int{}
	for k, v := range aDict {
		atoms[string(k[0])] += v
		atoms[string(k[1])] += v
	}
	atoms[string(seed[0])] += 1
	atoms[string(seed[len(seed)-1])] += 1
	for k := range atoms {
		atoms[k] = atoms[k] / 2
	}
	return atoms
}

func makeDistribution(data string, count int) int {
	seed, oldrules := parse(data)
	rules := map[string]string{}
	for k, v := range oldrules {
		rules[k] = string(k[0]) + v + string(k[1])
	}
	distribution := map[string]int{}
	for _, s := range getPairs(seed) {
		for ct := 1; ct <= count; ct++ {
			s = makeInserts(s, rules)
			fmt.Println(ct, s)
		}
	}
	for _, v := range seed {
		distribution[string(v)] += 1
	}
	min := math.MaxInt
	max := 0
	for _, v := range distribution {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
	// for _,k := range seed{
	// 	current
	// }
}

func makeInserts(aString string, rules map[string]string) string {
	if len(aString) == 2 {
		return rules[aString]
	}
	pairs := getPairs(aString)
	newString := string(aString[0])
	for _, val := range pairs {
		newString += makeInserts(val, rules)[1:]
	}
	return newString
}

func getPairs(aString string) []string {
	aList := []string{}
	for idx := range aString {
		if idx == 0 {
			continue
		}
		aList = append(aList, string(aString[idx-1])+string(aString[idx]))
	}
	return aList

}

func doPairInsert(data string, count int) string {

	seed, rules := parse(data)

	for x := 0; x < count; x++ {
		seed = pairInsert(seed, rules)
		fmt.Println(x, len(seed))
	}
	return seed
}

func pairInsert(seed string, rules map[string]string) string {

	newString := ""
	for x := 1; x < len(seed); x++ {
		insert := rules[string(seed[x-1])+string(seed[x])]
		newString = newString + string(seed[x-1]) + insert
	}
	newString = newString + string(seed[len(seed)-1])
	fmt.Println(newString)
	return newString
}

func parse(data string) (string, map[string]string) {
	trimmed := strings.TrimSpace(data)
	split := strings.Split(trimmed, "\n")
	seed := strings.TrimSpace(split[0])
	ruleData := split[2:]
	rules := map[string]string{}
	for _, v := range ruleData {
		v = strings.TrimSpace(v)
		s := strings.Split(v, " -> ")
		rules[s[0]] = s[1]
	}
	return seed, rules
}
