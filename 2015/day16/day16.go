package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", findSue())
}

func findSue() string {
	data := `children: 3
	cats: 7
	samoyeds: 2
	pomeranians: 3
	akitas: 0
	vizslas: 0
	goldfish: 5
	trees: 3
	cars: 2
	perfumes: 1`

	auntinput, _ := os.ReadFile("input.txt")
	auntsdata := strings.Split(string(auntinput), "\n")
	for _, v := range auntsdata {
		firstColon := strings.Index(v, ":")
		aunt := v[:firstColon]
		cString := strings.TrimSpace(v[firstColon+1:])
		characteristics := strings.Split(cString, ",")
		foundAll := true
		for _, v := range characteristics {
			c := strings.TrimSpace(v)
			if !strings.Contains(data, c) {
				foundAll = false
				break
			}
		}
		if foundAll {
			return aunt
		}

	}
	return "not found"
}

func makeDict(aLine string) map[string]int {
	newDict := map[string]int{}

}

func findSue2() string {
	data := `children: 3
	cats: 7
	samoyeds: 2
	pomeranians: 3
	akitas: 0
	vizslas: 0
	goldfish: 5
	trees: 3
	cars: 2
	perfumes: 1`

	dataDict := map[string]int{}

	for _, v := range strings.Split(data, "\n") {
		t := strings.Split(v, ":")
		key := strings.TrimSpace(t[0])
		val, _ := strconv.Atoi(strings.TrimSpace(t[1]))
		dataDict[key] = val
	}

	auntinput, _ := os.ReadFile("input.txt")
	auntsdata := strings.Split(string(auntinput), "\n")
	for _, v := range auntsdata {
		firstColon := strings.Index(v, ":")
		aunt := v[:firstColon]
		cString := strings.TrimSpace(v[firstColon+1:])
		characteristics := strings.Split(cString, ",")
		foundAll := true
		for _, v := range characteristics {
			c := strings.TrimSpace(v)
			if !strings.Contains(data, c) {
				foundAll = false
				break
			}
		}
		if foundAll {
			return aunt
		}

	}
	return "not found"
}
