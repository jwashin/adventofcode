package main

import (
	"fmt"
	"sort"
	"strconv"
)

var puzzleInput = 5093

func main() {
	// fmt.Println("1.", part1(puzzleInput))
	fmt.Println("2.", part2(puzzleInput))
}

func powerLevel(x, y, serialNumber int) int {
	rackId := x + 10
	pl := rackId * y
	pl += serialNumber
	pl *= rackId
	s := []byte(fmt.Sprintf("%d", pl))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	d := 0
	if len(s) > 2 {
		d, _ = strconv.Atoi(string(s[2]))
	}
	return d - 5
}

type row []int

type grid []row

func part2(serialNumber int) string {
	gridsize := 300
	g := grid{}
	for y := 1; y <= gridsize; y++ {
		r := row{}
		for x := 1; x <= gridsize; x++ {
			r = append(r, powerLevel(x, y, serialNumber))
		}
		g = append(g, r)
	}

	sumvalues := map[string]int{}

	squareSize := 1
	gmaxes := map[string]int{}
	for squareSize < gridsize {
		for ydx, row := range g {
			y := ydx + 1
			for idx := range row {
				x := idx + 1
				sum := 0
				// TODO: we're doing too many here. should stop when we can't get a full square
				for sy := ydx; sy < ydx+squareSize && sy < gridsize; sy++ {
					for sx := idx; sx < idx+squareSize && sx < gridsize; sx++ {
						sum += g[sy][sx]
					}
				}
				sumvalues[fmt.Sprintf("%d,%d", x, y)] = sum
			}
		}
		max := 0
		for _, v := range sumvalues {
			if v > max {
				max = v
			}
		}
		maxes := []string{}
		for k, v := range sumvalues {
			if v == max {
				maxes = append(maxes, k)
			}
		}
		sort.Slice(maxes, func(i, j int) bool {
			return i < j
		})
		gmaxes[maxes[0]+","+fmt.Sprint(squareSize)] = max
		squareSize += 1
		if squareSize%30 == 0 {
			fmt.Println(squareSize)
		}
	}
	max := 0
	id := ""
	for k, v := range gmaxes {
		if v > max {
			max = v
			id = k
		}
	}
	return id
}

func part1(serialNumber int) string {
	gridsize := 300
	g := grid{}
	for y := 1; y <= gridsize; y++ {
		r := row{}
		for x := 1; x <= gridsize; x++ {
			r = append(r, powerLevel(x, y, serialNumber))
		}
		g = append(g, r)
	}

	sumvalues := map[string]int{}

	for ydx, row := range g {
		y := ydx + 1
		for idx := range row {
			x := idx + 1
			sum := 0
			for sy := ydx; sy < ydx+3 && sy < gridsize; sy++ {
				for sx := idx; sx < idx+3 && sx < gridsize; sx++ {
					sum += g[sy][sx]
				}
			}
			sumvalues[fmt.Sprintf("%d,%d", x, y)] = sum
		}
	}
	max := 0
	for _, v := range sumvalues {
		if v > max {
			max = v
		}
	}
	maxes := []string{}
	for k, v := range sumvalues {
		if v == max {
			maxes = append(maxes, k)
		}
	}
	sort.Slice(maxes, func(i, j int) bool {
		return i < j
	})
	return maxes[0]
}
