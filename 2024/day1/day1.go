package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1", getDiffSum(string(input)))
	fmt.Println("Part 2", similarity(string(input)))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getDiffSum(s string) int {

	lines := strings.Split(s, "\n")
	lista := []int{}
	listb := []int{}
	for _, v := range lines {
		items := strings.Fields(v)
		if len(items) == 2 {
			a, _ := strconv.Atoi(items[0])
			lista = append(lista, a)
			b, _ := strconv.Atoi(items[1])
			listb = append(listb, b)
		}
	}
	sort.Ints(lista)
	sort.Ints(listb)
	result := 0
	for k := range lista {
		result += abs(lista[k] - listb[k])
	}
	return result
}
func similarity(s string) int {

	lines := strings.Split(s, "\n")
	lista := []int{}
	listb := []int{}
	for _, v := range lines {
		items := strings.Fields(v)
		if len(items) == 2 {
			a, _ := strconv.Atoi(items[0])
			lista = append(lista, a)
			b, _ := strconv.Atoi(items[1])
			listb = append(listb, b)
		}
	}
	result := 0
	for _, v := range lista {
		result += v * countinlist(v, listb)
	}
	return result
}

func countinlist(a int, b []int) int {
	count := 0
	for _, v := range b {
		if v == a {
			count += 1
		}
	}
	return count
}
