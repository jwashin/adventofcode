package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("1.", getBestHappiness(string(input)))
	fmt.Println("2.", getBestHappiness2(string(input)))
}

func makeData(s string) map[string]int {
	data := map[string]int{}
	t := strings.Split(s, "\n")
	for _, v := range t {
		var a, b string
		var h int
		n, _ := fmt.Sscanf(v, "%s would gain %d happiness units by sitting next to %s", &a, &h, &b)
		if n == 3 {
			// get rid of the dot sometimes
			if strings.Contains(b, ".") {
				b = b[:len(b)-1]
			}
			data[fmt.Sprintf("%s,%s", a, b)] = h
			continue
		}
		fmt.Sscanf(v, "%s would lose %d happiness units by sitting next to %s", &a, &h, &b)
		// get rid of the dot sometimes
		if strings.Contains(b, ".") {
			b = b[:len(b)-1]
		}
		data[fmt.Sprintf("%s,%s", a, b)] = -h
	}
	return data
}

func makeKeys(data map[string]int) map[string]string {
	// keys is {Alice:a, Bob:b ...}
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	newList := map[string]string{}
	for k := range data {
		s := strings.Split(k, ",")
		for _, v := range s {
			newList[v] = ""
		}
	}
	bareList := []string{}
	for v := range newList {
		bareList = append(bareList, v)
	}
	for ky, val := range bareList {
		newList[val] = string(alphabet[ky])
	}
	return newList
}

func getBestHappiness(s string) int {
	data := makeData(s)
	keys := makeKeys(data)

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	max := 0
	alphas := alphabet[:len(keys)]
	j := permutations(alphas)
	for _, perm := range j {
		out := 0
		circle := perm + string(perm[0])
		for k, v := range data {
			names := strings.Split(k, ",")
			a := keys[names[0]]
			b := keys[names[1]]
			if strings.Contains(circle, a+b) || strings.Contains(circle, b+a) {
				out += v
			}
		}
		if out > max {
			max = out
		}

	}
	return max
}

func getBestHappiness2(s string) int {
	// just add one more letter to the permutations
	data := makeData(s)
	keys := makeKeys(data)

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	max := 0
	alphas := alphabet[:len(keys)+1]
	j := permutations(alphas)
	for _, perm := range j {
		out := 0
		circle := perm + string(perm[0])
		for k, v := range data {
			names := strings.Split(k, ",")
			a := keys[names[0]]
			b := keys[names[1]]
			if strings.Contains(circle, a+b) || strings.Contains(circle, b+a) {
				out += v
			}
		}
		if out > max {
			max = out
		}

	}
	return max
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
