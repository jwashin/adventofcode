package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	s := string(input)
	fmt.Println("1.", part1(s))
	fmt.Println("2.", part2(s, 10000))

}

type coordinate struct {
	x, y int
}

func (a coordinate) distance(b *coordinate) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func (a coordinate) String() string {
	return fmt.Sprintf("%d %d", a.x, a.y)
}

func (a *coordinate) calcClosest(c map[string]*coordinate) (string, int) {
	closest := ""
	distances := map[string]int{}

	totalDistance := 0
	for i, v := range c {
		x := a.distance(v)
		totalDistance += x
		distances[i] = x
	}
	min := math.MaxInt

	for _, v := range distances {

		if v < min {
			min = v
		}
	}
	count := 0
	for k, v := range distances {
		if v == min {
			count += 1
			closest = k
		}
	}
	if count > 1 {
		closest = "."
	}
	return closest, totalDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func makeOrigins(s string) map[string]*coordinate {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	alpha += strings.ToUpper(alpha)
	coordinates := map[string]*coordinate{}
	for i, v := range strings.Split(s, "\n") {
		v = strings.ReplaceAll(v, ",", "")
		t := strings.Fields(v)
		x, _ := strconv.Atoi(t[0])
		y, _ := strconv.Atoi(t[1])
		c := coordinate{x: x, y: y}
		coordinates[string(alpha[i])] = &c
	}
	return coordinates
}

type stringGrid []string

// 3433 too high

func part1(s string) int {
	origins := makeOrigins(s)
	xmax, ymax := 0, 0
	xmin := math.MaxInt
	ymin := math.MaxInt

	for _, v := range origins {
		if v.x > xmax {
			xmax = v.x
		}
		if v.x < xmin {
			xmin = v.x
		}
		if v.y > ymax {
			ymax = v.y
		}
		if v.y < ymin {
			ymin = v.y
		}
	}

	t := stringGrid{}
	for j := ymin; j < ymax; j++ {
		line := ""
		for i := xmin; i < xmax; i++ {
			c := coordinate{x: i, y: j}
			d, _ := c.calcClosest(origins)
			line += d
		}
		t = append(t, line)
	}
	stage1Counts := map[string]int{}
	for _, line := range t {
		for _, v := range line {
			stage1Counts[string(v)] += 1
		}
	}

	m := stringGrid{}
	for j := ymin - 10; j < ymax+10; j++ {
		line := ""
		for i := xmin - 10; i < xmax+10; i++ {
			c := coordinate{x: i, y: j}
			d, _ := c.calcClosest(origins)
			line += d
		}
		m = append(m, line)
	}
	stage2Counts := map[string]int{}

	for _, line := range m {
		for _, v := range line {
			stage2Counts[string(v)] += 1
		}
	}
	newList := map[string]int{}
	for k, v := range stage1Counts {
		if v == stage2Counts[k] {
			newList[k] = v
		}
	}
	max := 0
	for _, v := range newList {
		if v > max {
			max = v
		}
	}
	return max

}
func part2(s string, maxDistance int) int {
	origins := makeOrigins(s)
	xmax, ymax := 0, 0
	xmin := math.MaxInt
	ymin := math.MaxInt

	for _, v := range origins {
		if v.x > xmax {
			xmax = v.x
		}
		if v.x < xmin {
			xmin = v.x
		}
		if v.y > ymax {
			ymax = v.y
		}
		if v.y < ymin {
			ymin = v.y
		}
	}

	t := stringGrid{}
	for j := ymin; j < ymax; j++ {
		line := ""
		for i := xmin; i < xmax; i++ {
			c := coordinate{x: i, y: j}
			_, totalDistance := c.calcClosest(origins)
			if totalDistance < maxDistance {
				line += "#"
			} else {
				line += "."
			}

		}
		t = append(t, line)
	}
	stage1Counts := map[string]int{}
	for _, line := range t {
		for _, v := range line {
			stage1Counts[string(v)] += 1
		}
	}
	return stage1Counts["#"]
}
