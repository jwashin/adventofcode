package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", follow(string(input), 2))
	fmt.Println("part 2:", follow(string(input), 10))
}

type coordinate struct {
	x int
	y int
}

func follow(s string, knots int) int {
	rope := []*coordinate{}
	for len(rope) < knots {
		rope = append(rope, &coordinate{0, 0})
	}
	tailPositions := map[coordinate]int{{0, 0}: 1}

	// decision map
	// dx, dy: movex, movey
	d2m := map[coordinate]coordinate{}

	d2m[coordinate{0, -2}] = coordinate{0, -1}
	d2m[coordinate{0, 2}] = coordinate{0, 1}
	d2m[coordinate{-2, 0}] = coordinate{-1, 0}
	d2m[coordinate{2, 0}] = coordinate{1, 0}

	d2m[coordinate{2, 1}] = coordinate{1, 1}
	d2m[coordinate{2, -1}] = coordinate{1, -1}
	d2m[coordinate{-2, 1}] = coordinate{-1, 1}
	d2m[coordinate{-2, -1}] = coordinate{-1, -1}

	d2m[coordinate{1, 2}] = coordinate{1, 1}
	d2m[coordinate{1, -2}] = coordinate{1, -1}
	d2m[coordinate{-1, 2}] = coordinate{-1, 1}
	d2m[coordinate{-1, -2}] = coordinate{-1, -1}

	d2m[coordinate{-2, -2}] = coordinate{-1, -1}
	d2m[coordinate{2, 2}] = coordinate{1, 1}
	d2m[coordinate{-2, 2}] = coordinate{-1, 1}
	d2m[coordinate{2, -2}] = coordinate{1, -1}

	data := strings.Split(s, "\n")
	for _, v := range data {
		f := strings.Fields(v)
		direction := f[0]
		dist, _ := strconv.Atoi(f[1])
		moveQ := ""
		for len(moveQ) < dist {
			moveQ += direction
		}
		for len(moveQ) > 0 {
			move := moveQ[0]
			moveQ = moveQ[1:]
			// head move
			head := rope[0]
			switch move {
			case 'U':
				head.y += 1
			case 'D':
				head.y -= 1
			case 'L':
				head.x -= 1
			case 'R':
				head.x += 1
			}
			// tail moves
			for k, v := range rope {
				if k == 0 {
					continue
				}
				dx := rope[k-1].x - v.x
				dy := rope[k-1].y - v.y
				move := d2m[coordinate{dx, dy}]
				v.x += move.x
				v.y += move.y
				if k == len(rope)-1 {
					tailPositions[coordinate{v.x, v.y}] += 1
				}
			}
		}
	}
	return len(tailPositions)
}
