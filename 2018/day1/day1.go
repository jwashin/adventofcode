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

func part1() int {
	inpt, _ := os.ReadFile("input.txt")
	data := string(inpt)
	sum := 0
	for _, v := range strings.Split(data, "\n") {
		e := strings.ReplaceAll(v, "+", "")
		x, _ := strconv.Atoi(e)
		sum += x
	}
	return sum
}

func part2() int {
	inpt, _ := os.ReadFile("input.txt")
	data := string(inpt)
	sum := 0
	counts := map[int]int{}
	for {
		for _, v := range strings.Split(data, "\n") {
			e := strings.ReplaceAll(v, "+", "")
			x, _ := strconv.Atoi(e)
			sum += x
			counts[sum] += 1
			if counts[sum] == 2 {
				return sum
			}
		}
	}
}
