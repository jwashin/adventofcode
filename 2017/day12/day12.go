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
	input, _ := os.ReadFile("input.txt")
	x := connectsToZero(string(input))
	return x
}

func part2() int {
	input, _ := os.ReadFile("input.txt")
	x := countGroups(string(input))
	return x
}

type program struct {
	id          int
	connections []int
}

func parseLine(s string) program {
	z := strings.Fields(s)
	id, _ := strconv.Atoi(z[0])
	dsts := z[1:]
	dests := []int{}
	for _, v := range dsts {
		v = strings.Replace(v, ",", "", 1)
		v2, _ := strconv.Atoi(v)
		dests = append(dests, v2)
	}
	return program{id: id, connections: dests}

}

func countGroups(s string) int {
	data := strings.Split(s, "\n")
	c := map[int]program{}
	seeds := []int{}
	for _, d := range data {
		j := parseLine(d)
		c[j.id] = j
		seeds = append(seeds, j.id)
	}
	count := 0
	used := map[int]bool{}
	for len(seeds) > 0 {
		count += 1
		q := []int{seeds[0]}
		for len(q) > 0 {
			currentNodeId := q[0]
			used[currentNodeId] = true
			q = q[1:]
			currentNode := c[currentNodeId]
			for _, d := range currentNode.connections {
				if !used[d] {
					q = append(q, d)
				}
			}
		}
		seeds = []int{}
		for k := range c {
			if !used[k] {
				seeds = append(seeds, k)
			}
		}
	}
	return count
}

func connectsToZero(s string) int {
	data := strings.Split(s, "\n")
	used := map[int]bool{}
	c := map[int]program{}
	for _, d := range data {
		j := parseLine(d)
		c[j.id] = j
	}
	q := []int{0}
	for len(q) > 0 {
		currentNodeId := q[0]
		used[currentNodeId] = true
		q = q[1:]
		currentNode := c[currentNodeId]
		for _, d := range currentNode.connections {
			if !used[d] {
				q = append(q, d)
			}
		}
	}
	return len(used)
}
