package lines

import (
	"strconv"
	"strings"
)

type Line struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func (l Line) points() [][]int {

	points := [][]int{}

	startX := l.X1
	endX := l.X2
	incX := 0
	if l.X1 < l.X2 {
		incX = 1
	}
	if l.X1 > l.X2 {
		incX = -1
	}

	startY := l.Y1
	endY := l.Y2
	incY := 0
	if l.Y1 > l.Y2 {
		incY = -1
	}
	if l.Y1 < l.Y2 {
		incY = 1
	}

	point := []int{startX, startY}
	points = append(points, point)
	for {
		point = []int{point[0] + incX, point[1] + incY}
		points = append(points, point)
		if point[0] == endX && point[1] == endY {
			return points
		}
	}

}

func HVFilter(input []Line) []Line {
	filtered := []Line{}
	for _, value := range input {
		if (value.X1 == value.X2) || (value.Y1 == value.Y2) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func MakeWorldArray(input []Line) [][]int {
	xMax := 0
	yMax := 0
	xMin := 2000000000
	yMin := 2000000000

	for _, val := range input {
		if val.X1 > xMax {
			xMax = val.X1 + 1
		}
		if val.X1 < xMin {
			xMin = val.X1
		}
		if val.X2 > xMax {
			xMax = val.X2 + 1
		}
		if val.X2 < xMin {
			xMin = val.X2
		}
		if val.Y1 > yMax {
			yMax = val.Y1 + 1
		}
		if val.Y1 < yMin {
			yMin = val.Y1
		}

		if val.Y2 > yMax {
			yMax = val.Y2 + 1
		}
		if val.Y2 < yMin {
			yMin = val.Y2
		}

	}
	matrix := make([][]int, xMax+1)
	for i := 0; i <= xMax; i++ {
		col := make([]int, yMax+1)
		for idx := range col {
			col[idx] = 0
		}
		matrix[i] = col
	}
	return matrix
}

func CountIntersections(input []string) int {
	lines := LinesFromData(input)
	// lines = HVFilter(lines)
	matrix := MakeWorldArray(lines)
	for _, line := range lines {
		points := line.points()
		for _, point := range points {
			matrix[point[0]][point[1]] += 1
		}
	}
	intersections := 0
	for _, row := range matrix {
		for _, col := range row {
			if col >= 2 {
				intersections += 1
			}
		}
	}
	return intersections
}

func LinesFromData(data []string) []Line {
	lines := []Line{}
	for _, value := range data {
		cs := strings.Split(value, " -> ")
		c1 := strings.Split(strings.TrimSpace(cs[0]), ",")
		c2 := strings.Split(strings.TrimSpace(cs[1]), ",")
		X1, _ := strconv.Atoi(c1[0])
		Y1, _ := strconv.Atoi(c1[1])
		X2, _ := strconv.Atoi(c2[0])
		Y2, _ := strconv.Atoi(c2[1])
		lines = append(lines, Line{X1, Y1, X2, Y2})

	}
	return lines
}
