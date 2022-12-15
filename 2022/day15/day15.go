package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", part1(string(input), 2000000))
	fmt.Println("part 2:", part2(string(input), 4000000))
}

type sensor struct {
	loc    coord
	beacon coord
	sweep  int
}

func (s sensor) outerLines() [][]int {
	// list of [-1 or +1 , y-intercept]
	r := [][]int{}
	a := []int{1, s.loc.y - s.loc.x + s.sweep + 1}
	b := []int{1, s.loc.y - s.loc.x - s.sweep - 1}
	c := []int{-1, s.loc.y + s.loc.x + s.sweep + 1}
	d := []int{-1, s.loc.y + s.loc.x - s.sweep - 1}
	r = append(r, a)
	r = append(r, b)
	r = append(r, c)
	r = append(r, d)
	return r
}

func (d sensor) lineIntersections(g sensor) []coord {
	intersections := []coord{}
	for k, a := range d.outerLines() {
		for el, b := range g.outerLines() {
			if k > el {
				if a[0] == b[0] {
					continue
				}
				t := coord{abs(b[1]-a[1]) / 2, abs(a[1]+b[1]) / 2}
				intersections = append(intersections, t)
			}
		}
	}
	return intersections
}

type coord struct {
	x int
	y int
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func (a coord) manhattanDistance(b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func getSensors(s string) []sensor {
	sensors := []sensor{}
	for _, v := range strings.Split(s, "\n") {
		v = strings.ReplaceAll(v, ",", "")
		v = strings.ReplaceAll(v, ":", "")
		v = strings.ReplaceAll(v, "=", " ")
		f := strings.Fields(v)
		x1, _ := strconv.Atoi(f[3])
		y1, _ := strconv.Atoi(f[5])

		x2, _ := strconv.Atoi(f[11])
		y2, _ := strconv.Atoi(f[13])

		loc := coord{x1, y1}
		beacon := coord{x2, y2}
		sweep := loc.manhattanDistance(beacon)

		sensor := sensor{loc: loc, beacon: beacon, sweep: sweep}
		sensors = append(sensors, sensor)
	}
	return sensors
}

func part2(s string, max int) int {
	// I got help from reddit with this algorithm
	// https://www.reddit.com/r/adventofcode/comments/zmjzu7/2022_day_15_part_2_no_search_formula/
	// especially
	// https://www.reddit.com/r/adventofcode/comments/zmcn64/comment/j0b90nr/

	sensors := getSensors(s)
	candidates := []sensor{}
	for k, v := range sensors {
		for j, m := range sensors {
			if j > k {
				if v.loc.manhattanDistance(m.loc) == v.sweep+m.sweep+2 {
					candidates = append(candidates, v)
					candidates = append(candidates, m)
				}
			}
		}
	}
	for k, v := range candidates {
		for el, w := range candidates {
			if el > k {
				ins := w.lineIntersections(v)
				for _, cand := range ins {
					if cand.x >= 0 && cand.y >= 0 && cand.x <= max && cand.y <= max {
						if !isInSensorRange(sensors, cand) {
							return cand.x*4000000 + cand.y
						}
					}

				}
			}
		}
	}

	d := candidates[0].loc
	// just to return something...
	return d.x*4000000 + d.y

}

// func isBetweenBeacons(sensors []sensor, c coord) bool {
// 	maxX, minX := sensors[0].beacon.x, sensors[0].beacon.x
// 	maxY, minY := sensors[0].beacon.y, sensors[0].beacon.y

// 	for _, s := range sensors {
// 		if s.beacon.x > maxX {
// 			maxX = s.beacon.x
// 		}
// 		if s.beacon.y > maxY {
// 			maxY = s.beacon.y
// 		}
// 		if s.beacon.x < minX {
// 			minX = s.beacon.x
// 		}
// 		if s.beacon.y < minY {
// 			minY = s.beacon.y
// 		}
// 	}
// 	if c.x < maxX && c.x > minX && c.y < maxY && c.y > minY {
// 		return true
// 	}
// 	return false
// }

// func isBetweenSensors(sensors []sensor, c coord) bool {
// 	maxX, minX := sensors[0].loc.x, sensors[0].loc.x
// 	maxY, minY := sensors[0].loc.y, sensors[0].loc.y

// 	for _, s := range sensors {
// 		if s.loc.x > maxX {
// 			maxX = s.loc.x
// 		}
// 		if s.loc.y > maxY {
// 			maxY = s.loc.y
// 		}
// 		if s.loc.x < minX {
// 			minX = s.loc.x
// 		}
// 		if s.loc.y < minY {
// 			minY = s.loc.y
// 		}
// 	}
// 	if c.x < maxX && c.x > minX && c.y < maxY && c.y > minY {
// 		return true
// 	}
// 	return false
// }

func isInSensorRange(sensors []sensor, c coord) bool {
	for _, s := range sensors {
		if c == s.beacon {
			break
		}
		if c.manhattanDistance(s.loc) <= s.sweep {
			return true
		}
	}
	return false
}

// func part1tableau(s string, row int) int {
// 	sensors := getSensors(s)

// 	tableau := map[coord]int{}

// 	for _, v := range sensors {
// 		xMin := v.loc.x - v.sweep()
// 		xMax := v.loc.x + v.sweep()

// 		for x := xMin; x <= xMax; x++ {
// 			d := coord{x, row}
// 			if d.manhattanDistance(v.loc) <= v.sweep() {
// 				tableau[d] = 1
// 			}
// 			if d == v.beacon {
// 				tableau[d] = 2
// 			}
// 			if d == v.loc {
// 				tableau[d] = 3
// 			}
// 		}
// 	}
// 	count := 0
// 	for _, v := range tableau {
// 		if v == 1 {
// 			count += 1
// 		}
// 	}
// 	return count
// }

// 4184160 too low
// 4228766 too low
func part1(s string, row int) int {
	sensors := getSensors(s)
	xMax := sensors[0].beacon.x
	xMin := sensors[0].beacon.x
	for _, v := range sensors {
		if v.loc.x+v.sweep > xMax {
			xMax = v.loc.x + v.sweep
		}
		if v.loc.x-v.sweep < xMin {
			xMin = v.loc.x - v.sweep
		}
	}
	count := 0
	for x := xMin; x <= xMax; x++ {
		test := coord{x, row}
		if isInSensorRange(sensors, test) {
			count += 1
		}
	}
	return count
}
