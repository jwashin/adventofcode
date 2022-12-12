package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
	fmt.Println("Part2:", part2(string(input)))
}

func part1(s string) int {
	darr := strings.Split(s, "\n")
	arr := []string{}
	for _, v := range darr {
		arr = append(arr, strings.TrimSpace(v))
	}
	start := find(arr, "S")
	return hillClimb(arr, start)
}
func part2(s string) int {
	darr := strings.Split(s, "\n")
	arr := []string{}
	for _, v := range darr {
		arr = append(arr, strings.TrimSpace(v))
	}
	allAs := findAll(arr, "a")
	allAs = append(allAs, find(arr, "S"))
	min := math.MaxInt
	for _, v := range allAs {
		t := hillClimb(arr, v)
		if t < min {
			min = t
		}
	}
	return min
}

var alphas = "SabcdefghijklmnopqrstuvwxyzE"

type coordinate struct {
	x int
	y int
}

func findAll(a []string, d string) []coordinate {
	s := []coordinate{}
	for y, v := range a {
		for x, c := range v {
			if string(c) == d {
				s = append(s, coordinate{x, y})
			}
		}
	}
	return s
}

func find(a []string, d string) coordinate {
	var s coordinate
	for y, v := range a {
		for x, c := range v {
			if string(c) == d {
				return coordinate{x, y}
			}
		}
	}
	return s
}

func coordFromPath(arr []string, start coordinate, path string) (coordinate, string, error) {
	loc := start
	for _, v := range path {
		if v == 'v' {
			loc.y += 1
		}
		if v == '^' {
			loc.y -= 1
		}
		if v == '>' {
			loc.x += 1
		}
		if v == '<' {
			loc.x -= 1
		}
	}
	if loc.y < 0 || loc.y > len(arr)-1 ||
		loc.x < 0 || loc.x > len(arr[0])-1 {
		return loc, "", errors.New("new location out of bounds")
	}
	val := string(arr[loc.y][loc.x])
	return loc, val, nil
}

func hillClimb(arr []string, start coordinate) int {
	// dijkstra, sorta
	end := find(arr, "E")
	currLoc := start
	currVal := "S"
	checked := map[coordinate]bool{start: true}
	queue := []string{""}
	for len(queue) > 0 {
		// shortest path goes first
		sort.Slice(queue, func(i, j int) bool {
			return len(queue[i]) < len(queue[j])
		})
		currPath := queue[0]
		queue = queue[1:]
		currLoc, currVal, _ = coordFromPath(arr, start, currPath)

		if currLoc == end {
			return len(currPath)
		}
		for _, direction := range []string{"^", "v", "<", ">"} {
			loc, newVal, err := coordFromPath(arr, currLoc, direction)
			if err == nil {
				a := strings.Index(alphas, currVal)
				b := strings.Index(alphas, newVal)
				// if a == b || b-a == 1 || (currVal == "o" && newVal == "m") {
				if a == b || b-a == 1 || b < a {
					if !checked[loc] {
						queue = append(queue, currPath+direction)
						checked[loc] = true
					}
				}
			}
		}
	}
	return math.MaxInt
}
