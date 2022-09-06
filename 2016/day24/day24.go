package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	s, _ := os.ReadFile("input.txt")
	grid := makeGrid(string(s))
	// fmt.Println("1.", part1(grid))
	fmt.Println("2.", part2(grid))
}

func part1(grid charGrid) int {
	p := grid.getPointsOfInterest()
	distances := map[string]int{}
	items := ""
	for ky1, v1 := range p {
		// getting the string for permutations while we are here
		if ky1 != 0 {
			items += fmt.Sprintf("%d", ky1)
		}
		for ky2, v2 := range p {
			if ky1 >= ky2 {
				continue
			}
			d := grid.getShortestPath(v1, v2)
			distances[fmt.Sprintf("%d,%d", ky1, ky2)] = d
			distances[fmt.Sprintf("%d,%d", ky2, ky1)] = d
		}
	}
	j := permutations(items)
	minCost := math.MaxInt
	for _, permutation := range j {
		testCost := 0
		path := "0" + permutation
		for len(path) >= 2 {
			step := fmt.Sprintf("%s,%s", string(path[0]), string(path[1]))
			testCost += distances[step]
			path = path[1:]
		}
		if testCost < minCost {
			minCost = testCost
		}
	}
	return minCost
}

func part2(grid charGrid) int {
	p := grid.getPointsOfInterest()
	distances := map[string]int{}
	items := ""
	for ky1, v1 := range p {
		if ky1 != 0 {
			items += fmt.Sprintf("%d", ky1)
		}
		for ky2, v2 := range p {
			if ky1 >= ky2 {
				continue
			}
			d := grid.getShortestPath(v1, v2)
			distances[fmt.Sprintf("%d,%d", ky1, ky2)] = d
			distances[fmt.Sprintf("%d,%d", ky2, ky1)] = d
		}
	}
	j := permutations(items)
	minCost := math.MaxInt
	for _, permutation := range j {
		testCost := 0
		// the only difference is here
		path := "0" + permutation + "0"
		for len(path) >= 2 {
			step := fmt.Sprintf("%s,%s", string(path[0]), string(path[1]))
			testCost += distances[step]
			path = path[1:]
		}
		if testCost < minCost {
			minCost = testCost
		}
	}
	return minCost
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

func makeGrid(s string) charGrid {
	t := strings.Split(s, "\n")
	if len(t[len(t)-1]) == 0 {
		return t[:len(t)-1]
	}
	return t
}

func s2xy(s string) (int, int) {
	x, y := 0, 0
	fmt.Sscanf(s, "%d,%d", &x, &y)
	return x, y
}

type charGrid []string

// func (c charGrid) getChar(x, y int) byte {
// 	return c[y][x]
// }

func (c charGrid) isFree(x, y int) bool {
	return c[y][x] != '#'
}

func (c charGrid) getPointsOfInterest() map[int]string {
	points := map[int]string{}
	for y, ln := range c {
		for x, character := range strings.Split(ln, "") {
			if strings.Contains("0123456789", string(character)) {
				p, _ := strconv.Atoi(string(character))
				points[p] = fmt.Sprintf("%d,%d", x, y)
			}
		}
	}
	return points
}

func (c charGrid) getShortestPath(start string, dest string) int {
	q := map[string]int{start: 0}
	currentNode := start
	currentDistance := 0
	for len(q) > 0 {
		min := math.MaxInt
		for ky, val := range q {
			if val < min {
				min = val
				currentNode = ky
				currentDistance = q[currentNode]
			}
		}
		delete(q, currentNode)
		for _, cand := range c.destinations(currentNode) {
			newDist := currentDistance + 1
			if cand == dest {
				return newDist
			}
			if q[cand] > 0 {
				if newDist < q[cand] {
					q[cand] = newDist
				}
			}
			q[cand] = newDist
		}
	}
	return currentDistance
}

func (c charGrid) destinations(point string) []string {
	x, y := s2xy(point)
	dests := []string{}

	if c.isFree(x, y-1) {
		dests = append(dests, fmt.Sprintf("%d,%d", x, y-1))
	}
	if c.isFree(x, y+1) {
		dests = append(dests, fmt.Sprintf("%d,%d", x, y+1))
	}
	if c.isFree(x+1, y) {
		dests = append(dests, fmt.Sprintf("%d,%d", x+1, y))
	}
	if c.isFree(x-1, y) {
		dests = append(dests, fmt.Sprintf("%d,%d", x-1, y))
	}
	return dests
}
