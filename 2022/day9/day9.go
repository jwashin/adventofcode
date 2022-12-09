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

func follow(s string, count int) int {
	data := strings.Split(s, "\n")
	nTails := count
	rope := []*coordinate{}
	for x := 0; x < nTails; x++ {
		rope = append(rope, &coordinate{0, 0})
	}

	tailPositions := map[string]int{"0 0": 1}
	d := map[coordinate]coordinate{}
	// dx, dy: movex, movey
	d[coordinate{0, -2}] = coordinate{0, 1}
	d[coordinate{0, 2}] = coordinate{0, -1}
	d[coordinate{-2, 0}] = coordinate{1, 0}
	d[coordinate{2, 0}] = coordinate{-1, 0}

	d[coordinate{2, 1}] = coordinate{-1, -1}
	d[coordinate{2, -1}] = coordinate{-1, 1}
	d[coordinate{-2, 1}] = coordinate{1, -1}
	d[coordinate{-2, -1}] = coordinate{1, 1}

	d[coordinate{1, 2}] = coordinate{-1, -1}
	d[coordinate{1, -2}] = coordinate{-1, 1}
	d[coordinate{-1, 2}] = coordinate{1, -1}
	d[coordinate{-1, -2}] = coordinate{1, 1}

	d[coordinate{-2, -2}] = coordinate{1, 1}
	d[coordinate{2, 2}] = coordinate{-1, -1}
	d[coordinate{-2, 2}] = coordinate{1, -1}
	d[coordinate{2, -2}] = coordinate{-1, 1}

	for _, v := range data {
		f := strings.Fields(v)
		direction := f[0]
		dist, _ := strconv.Atoi(f[1])
		moveQ := ""
		for m := 0; m < dist; m++ {
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
				if k > 0 {
					currX := rope[k-1].x
					currY := rope[k-1].y
					tailX := v.x
					tailY := v.y
					isLast := k == len(rope)-1

					move := d[coordinate{currX - tailX, currY - tailY}]
					v.x -= move.x
					v.y -= move.y
					if isLast {
						tailPositions[fmt.Sprintf("%d %d", v.x, v.y)] += 1
					}
				}
			}
		}
	}
	return len(tailPositions)
}
