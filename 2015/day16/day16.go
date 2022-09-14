package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", findSue())
	fmt.Println("2.", findSue2())
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
	auntsdata := strings.Split(strings.TrimSpace(string(auntinput)), "\n")
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

func hasKey(aDict map[string]int, key string) bool {
	for k := range aDict {
		if k == key {
			return true
		}
	}
	return false
}

// 296 too high

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

	readings := map[string]int{}

	for _, v := range strings.Split(data, "\n") {
		t := strings.Split(v, ":")
		key := strings.TrimSpace(t[0])
		val, _ := strconv.Atoi(strings.TrimSpace(t[1]))
		readings[key] = val
	}
	aunts := []map[string]int{}
	auntinput, _ := os.ReadFile("input.txt")
	auntsdata := strings.Split(strings.TrimSpace(string(auntinput)), "\n")
	sueNumber, char1, char2, char3 := 0, 0, 0, 0
	char1Name, char2Name, char3Name := "", "", ""
	// aunt := map[string]int{}
	for _, v := range auntsdata {
		// parse this
		// Sue 8: cars: 10, pomeranians: 7, goldfish: 8
		// colons work better if you separate them from their string
		d := strings.ReplaceAll(v, ":", " :")
		fmt.Sscanf(d, "Sue %d : %s : %d, %s : %d, %s : %d", &sueNumber, &char1Name, &char1, &char2Name, &char2, &char3Name, &char3)
		aunt := map[string]int{"sue": sueNumber, char1Name: char1, char2Name: char2, char3Name: char3}
		aunts = append(aunts, aunt)
	}
	auntNumber := 0
	for _, aunt := range aunts {
		auntNumber = aunt["sue"]

		foundAll := true
		for k, value := range aunt {
			if k == "sue" {
				continue
			}
			if !hasKey(readings, k) {
				fmt.Println(k, "not found!")
			}
			if k == "cats" || k == "trees" {
				if value <= readings[k] {
					foundAll = false
				}
				continue
			}
			if k == "pomeranians" || k == "goldfish" {
				if value >= readings[k] {
					foundAll = false
				}
				continue
			}
			if readings[k] != value {
				foundAll = false
			}
		}
		if foundAll {
			return fmt.Sprint(auntNumber)
		}
	}
	return "not found"
}
