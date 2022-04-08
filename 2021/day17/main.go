package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

type Coord struct {
	x int
	y int
}

func (c Coord) hitsTarget(tl Coord, lr Coord, velocity Coord) (bool, int) {
	maxY := math.MinInt
	currLoc := Coord{c.x, c.y}
	hitTarget := false
	cv := velocity
	for {
		newLoc, newVel := next(currLoc, cv)
		if newLoc.y > maxY {
			maxY = currLoc.y
		}
		if newLoc.isInTarget(tl, lr) {
			hitTarget = true
			break
		}
		if newLoc.y < lr.y {
			break
		}
		currLoc = newLoc
		cv = newVel
	}
	return hitTarget, maxY
}

func String(c Coord) string {
	return fmt.Sprintf("[%d,%d]", c.x, c.y)
}

func (c Coord) isInTarget(tl Coord, lr Coord) bool {
	return c.x >= tl.x && c.x <= lr.x && c.y <= tl.y && c.y >= lr.y
}

func main() {
	s, _ := ioutil.ReadFile("input.txt")
	fmt.Println(highestY(string(s)))
	fmt.Println(countVelocityChoices(string(s)))
}

func highestY(aString string) int {
	tl, lr := parseTarget(aString)
	start := Coord{0, 0}
	highestY := math.MinInt

	mx := findMx(tl, lr)
	for xvel := -mx; xvel < mx; xvel++ {
		for yvel := -mx; yvel < mx; yvel++ {
			hitsTarget, maxY := start.hitsTarget(tl, lr, Coord{xvel, yvel})
			if maxY > highestY && hitsTarget {
				highestY = maxY
			}
		}
	}
	return highestY
}

func findMx(tl Coord, lr Coord) int {
	xl := math.Abs(float64(lr.x - tl.x))
	yl := math.Abs(float64(tl.y - lr.y))
	var ret int
	if xl > yl {
		ret = int(xl)
	} else {
		ret = int(yl)
	}
	return ret * 5
}

func countVelocityChoices(aString string) int {
	tl, lr := parseTarget(aString)
	choices := []Coord{}
	start := Coord{0, 0}
	highestY := math.MinInt
	mx := findMx(tl, lr)
	for xvel := -mx; xvel < mx; xvel++ {
		for yvel := -mx; yvel < mx; yvel++ {
			hitsTarget, maxY := start.hitsTarget(tl, lr, Coord{xvel, yvel})
			if hitsTarget {
				choices = append(choices, Coord{xvel, yvel})
			}
			if maxY > highestY && hitsTarget {
				highestY = maxY
			}
		}
	}
	return len(choices)
}

func parseTarget(aString string) (Coord, Coord) {
	var x1, x2, y1, y2 int
	fmt.Sscanf(aString, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 < y2 {
		y1, y2 = y2, y1
	}

	return Coord{x1, y1}, Coord{x2, y2}
}

func next(position Coord, velocity Coord) (Coord, Coord) {
	newX := position.x + velocity.x
	newY := position.y + velocity.y

	newDx := velocity.x

	if velocity.x < 0 {
		newDx = velocity.x + 1
	}
	if velocity.x > 0 {
		newDx = velocity.x - 1
	}

	newDy := velocity.y - 1

	return Coord{newX, newY}, Coord{newDx, newDy}
}
