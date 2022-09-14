package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

func part2() int {
	input, _ := os.ReadFile("input.txt")
	containers := strings.Split(string(input), "\n")
	return minCount(containers, 150)
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	containers := strings.Split(string(input), "\n")
	return countCombinations(containers, 150)
}

func test17() int {
	return countCombinations([]string{"20", "15", "10", "5", "5"}, 25)
}
func testa17() int {
	return minCount([]string{"20", "15", "10", "5", "5"}, 25)
}

func listSum(items []string) int {
	total := 0
	for _, v := range items {
		amt, _ := strconv.Atoi(v)
		total += amt
	}
	return total
}

func minCount(containers []string, total int) int {
	count := 0
	nContainers := 0
	found := false
	for !found {
		nContainers += 1
		comb := Combinations(containers, nContainers)
		for _, v := range comb {
			t := listSum(v)
			if t == total {
				fmt.Println(v)
				found = true
				count += 1
			}
		}
	}
	return count
}

func countCombinations(containers []string, total int) int {
	count := 0
	nContainers := 0
	for nContainers < len(containers) {
		nContainers += 1
		comb := Combinations(containers, nContainers)
		for _, v := range comb {
			t := listSum(v)
			if t == total {
				count += 1
			}
		}
	}
	return count
}

// https://www.sobyte.net/post/2022-01/go-slice/
func Combinations(iterable []string, r int) (rt [][]string) {
	pool := iterable
	n := len(pool)

	if r > n {
		return
	}

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	result := make([]string, r)
	for i, el := range indices {
		result[i] = pool[el]
	}
	s2 := make([]string, r)
	copy(s2, result)
	rt = append(rt, s2)

	for {
		i := r - 1
		for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
		}

		if i < 0 {
			return
		}

		indices[i] += 1
		for j := i + 1; j < r; j += 1 {
			indices[j] = indices[j-1] + 1
		}

		for ; i < len(indices); i += 1 {
			result[i] = pool[indices[i]]
		}
		s2 = make([]string, r)
		copy(s2, result)
		rt = append(rt, s2)
	}

}
