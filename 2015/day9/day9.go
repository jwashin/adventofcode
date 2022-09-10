package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	// fmt.Println("1.", getShortestRoute(string(input)))
	fmt.Println("2.", getLongestRoute(string(input)))
}

func getDistances(s string) map[string]int {
	distances := map[string]int{}
	for _, v := range strings.Split(s, "\n") {
		f := ""
		t := ""
		d := 0
		fmt.Sscanf(v, "%s to %s = %d", &f, &t, &d)
		distances[f+","+t] = d
		distances[t+","+f] = d
	}
	return distances
}

func getShortestRoute(s string) int {
	distances := getDistances(s)
	nm := collectNames(distances)
	nameDict := map[string]string{}
	ts := "abcdefghijklmnopqrstuvwxyz"
	for ky, val := range nm {
		nameDict[string(ts[ky])] = val
	}
	ts = ts[:len(nameDict)]
	perms := permutations(ts)
	// perms are in abc bac, etc
	minDistance := math.MaxInt
	for _, routePerm := range perms {
		circuitDistance := 0
		for idx, item := range strings.Split(routePerm, "")[1:] {
			// a-b becomes Dublin-London
			circuitDistance += distances[nameDict[item]+","+nameDict[string(routePerm[idx])]]
		}
		if circuitDistance < minDistance {
			minDistance = circuitDistance
		}
	}
	return minDistance
}

func getLongestRoute(s string) int {
	distances := getDistances(s)
	nm := collectNames(distances)
	nameDict := map[string]string{}
	ts := "abcdefghijklmnopqrstuvwxyz"
	for ky, val := range nm {
		nameDict[string(ts[ky])] = val
	}
	ts = ts[:len(nameDict)]
	perms := permutations(ts)
	// perms are in abc bac, etc
	maxDistance := 0
	for _, routePerm := range perms {
		circuitDistance := 0
		for idx, item := range strings.Split(routePerm, "")[1:] {
			// a-b becomes Dublin-London
			circuitDistance += distances[nameDict[item]+","+nameDict[string(routePerm[idx])]]
		}
		if circuitDistance > maxDistance {
			maxDistance = circuitDistance
		}
	}
	return maxDistance
}

func collectNames(distances map[string]int) []string {
	namesDict := map[string]int{}
	for k := range distances {
		t := strings.Split(k, ",")
		for _, name := range t {
			namesDict[name] = 1
		}
	}
	names := []string{}
	for ky := range namesDict {
		names = append(names, ky)
	}
	return names
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
