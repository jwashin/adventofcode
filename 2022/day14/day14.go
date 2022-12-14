package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", dropSand(string(input)))
	fmt.Println("part 2:", dropSand2(string(input)))
}

type coordinate struct {
	x int
	y int
}

func makePoint(s string) coordinate {
	p := strings.Split(s, ",")
	x, _ := strconv.Atoi(p[0])
	y, _ := strconv.Atoi(p[1])
	return coordinate{x, y}
}

func countStaticSand(m map[coordinate]int) int {
	count := 0
	for _, v := range m {
		if v > 1 {
			count += 1
		}
	}
	return count
}

func addLine(m map[coordinate]int, start coordinate, end coordinate) map[coordinate]int {
	if start.x == end.x {
		// vertical line
		s, e := start.y, end.y
		if start.y > end.y {
			s, e = end.y, start.y
		}
		yNew := s
		for yNew <= e {
			coord := coordinate{start.x, yNew}
			m[coord] = 1
			yNew += 1
		}
	} else if start.y == end.y {
		// horizontal line
		s, e := start.x, end.x
		if start.x > end.x {
			s, e = end.x, start.x
		}
		xNew := s
		for xNew <= e {
			coord := coordinate{xNew, start.y}
			m[coord] = 1
			xNew += 1
		}
	}
	return m
}

func parse(s string) map[coordinate]int {
	m := map[coordinate]int{}
	for _, v := range strings.Split(s, "\n") {
		v = strings.TrimSpace(v)
		f := strings.Fields(v)
		for len(f) >= 3 {
			startPoint := makePoint(f[0])
			endPoint := makePoint(f[2])
			m = addLine(m, startPoint, endPoint)
			f = f[2:]
		}
	}
	return m
}

// dropSand2 was copy-pasted from dropSand and modified. Could be refactored.

func dropSand2(s string) int {
	matrix := parse(s)
	maxY := 0
	for k := range matrix {
		if k.y > maxY {
			maxY = k.y
		}
	}
	startPoint := coordinate{500, 0}
	stop := false
	for !stop {
		currLoc := startPoint
		for {
			if matrix[coordinate{currLoc.x, currLoc.y + 1}] < 1 {
				currLoc = coordinate{currLoc.x, currLoc.y + 1}
			} else if matrix[coordinate{currLoc.x - 1, currLoc.y + 1}] < 1 {
				currLoc = coordinate{currLoc.x - 1, currLoc.y + 1}
			} else if matrix[coordinate{currLoc.x + 1, currLoc.y + 1}] < 1 {
				currLoc = coordinate{currLoc.x + 1, currLoc.y + 1}
			} else {
				matrix[currLoc] = 2
				if currLoc == startPoint {
					stop = true
				}
				break
			}
			// floor
			if currLoc.y == maxY+1 {
				matrix[currLoc] = 2
				break
			}
		}
	}
	return countStaticSand(matrix)
}

func dropSand(s string) int {
	matrix := parse(s)
	maxY := 0
	for k := range matrix {
		if k.y > maxY {
			maxY = k.y
		}
	}
	startPoint := coordinate{500, 0}
	stop := false
	for !stop {
		currLoc := startPoint
		for {
			if matrix[coordinate{currLoc.x, currLoc.y + 1}] < 1 {
				currLoc = coordinate{currLoc.x, currLoc.y + 1}
			} else if matrix[coordinate{currLoc.x - 1, currLoc.y + 1}] < 1 {
				currLoc = coordinate{currLoc.x - 1, currLoc.y + 1}
			} else if matrix[coordinate{currLoc.x + 1, currLoc.y + 1}] < 1 {
				currLoc = coordinate{currLoc.x + 1, currLoc.y + 1}
			} else {
				matrix[currLoc] = 2
				break
			}
			if currLoc.y > maxY {
				stop = true
				break
			}
		}
	}
	return countStaticSand(matrix)
}
