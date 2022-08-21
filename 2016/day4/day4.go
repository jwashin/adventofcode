package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := getInputTxt()
	ssum := 0
	for _, v := range data {
		if len(v) > 0 {
			_, isReal, sectorID := isRealRoom(v)
			if isReal {
				ssum += sectorID
			}
		}
	}
	fmt.Println(ssum)
	for _, v := range data {
		if len(v) > 0 {
			name, isReal, sectorID := isRealRoom(v)
			if isReal {
				trueName := shift(name, sectorID)
				fmt.Println(sectorID, trueName)
			}
		}

	}
}

func shift(name string, id int) string {
	n := id
	newName := ""
	rg := "abcdefghijklmnopqrstuvwxyz"
	for _, v := range name {
		if string(v) == "-" {
			newName += " "
			continue
		}
		i := strings.IndexRune(rg, v)
		i += n
		for i >= 26 {
			i -= 26
		}
		newName += string(rg[i])
	}
	return newName
}

func getInputTxt() []string {
	g, _ := os.ReadFile("input.txt")
	s := strings.Split(string(g), "\n")
	nl := []string{}
	for _, v := range s {
		nl = append(nl, strings.TrimSpace(v))
	}
	return nl
}

func firstParse(s string) (name string, sectorID int, checksum string) {
	r, _ := regexp.Compile(`([a-z-]*)-(\d*)\[(.*)\]`)
	z := r.FindStringSubmatch(s)
	s2, _ := strconv.Atoi(z[2])
	return z[1], s2, z[3]
}

func isRealRoom(s string) (string, bool, int) {
	name, sectorID, checksum := firstParse(s)
	letterCounts := map[string]int{}
	for _, v := range name {
		letterCounts[string(v)] += 1
	}
	sortable := []string{}
	for k, v := range letterCounts {
		if k != "-" {
			sortable = append(sortable, fmt.Sprintf("%03d%v", 100-v, k))
		}
	}
	sort.Strings(sortable)
	check := ""
	for len(check) < 5 {
		item := sortable[len(check)]
		check += string(item[3])
	}
	if check == checksum {
		return name, true, sectorID
	}
	return "", false, 0
}
