package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", follow(string(input)))
}

type coordinate struct {
	x int
	y int
}

func follow10(s string, count int) int {
	data := strings.Split(s, "\n")
	nTails := 10
	rope := []*coordinate{}
	for x := 0; x < nTails; x++ {
		rope = append(rope, &coordinate{0, 0})
	}

	tailPositions := map[string]int{"0 0": 1}

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
			// tail move
			for k, v := range rope[1:] {
				currX := rope[k-1].x
				currY := rope[k-1].y
				tailX := v.x
				tailY := v.y
				isLast := v == rope[len(rope)-1]
				if currX == tailX && currY-tailY == -2 {
					v.y -= 1
					if isLast {
						tailPositions[fmt.Sprintf("%d %d", v.x, v.y)] += 1
					}
					continue
				}
				if currX == tailX && currY-tailY == 2 {
					v.y += 1
					if isLast {
						tailPositions[fmt.Sprintf("%d %d", v.x, v.y)] += 1
					}
					continue
				}

				if currY == tailY && currX-tailX == -2 {
					v.x -= 1
					if isLast {
						tailPositions[fmt.Sprintf("%d %d", v.x, v.y)] += 1
					}
					continue
				}
				if currY == tailY && currX-tailX == 2 {
					v.x += 1
					if isLast {
						tailPositions[fmt.Sprintf("%d %d", v.x, v.y)] += 1
					}
					continue
				}

				if currX-tailX == 2 && (currY-tailY == 1 || currY-tailY == -1) {
					v.x += 1
					v.y = currY
					if isLast {
						tailPositions[fmt.Sprintf("%d %d", v.x, v.y)] += 1
					}
					continue
				}

				if currX-tailX == -2 && (currY-tailY == 1 || currY-tailY == -1) {
					v.x -= 1
					v.y = currY
					if isLast {
						tailPositions[fmt.Sprintf("%d %d", v.x, v.y)] += 1
					}
					continue
				}

				if currY-tailY == -2 && (currX-tailX == 1 || currX-tailX == -1) {
					v.y -= 1
					v.x = currX
					if isLast {
						tailPositions[fmt.Sprintf("%d %d", v.x, v.y)] += 1
					}
					continue
				}

				if currY-tailY == 2 && (currX-tailX == 1 || currX-tailX == -1) {
					v.y += 1
					v.x = currX
					if isLast {
						tailPositions[fmt.Sprintf("%d %d", v.x, v.y)] += 1
					}
					continue
				}
			}
		}
	}
	return len(tailPositions)
}

func follow(s string) int {
	data := strings.Split(s, "\n")

	currX, currY := 0, 0
	tailX, tailY := 0, 0

	tailPositions := map[string]int{"0 0": 1}

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
			switch move {
			case 'U':
				currY += 1
			case 'D':
				currY -= 1
			case 'L':
				currX -= 1
			case 'R':
				currX += 1
			}
			// tail move
			if currX == tailX && currY-tailY == -2 {
				tailY -= 1
				tailPositions[fmt.Sprintf("%d %d", tailX, tailY)] += 1
				continue
			}
			if currX == tailX && currY-tailY == 2 {
				tailY += 1
				tailPositions[fmt.Sprintf("%d %d", tailX, tailY)] += 1
				continue
			}

			if currY == tailY && currX-tailX == -2 {
				tailX -= 1
				tailPositions[fmt.Sprintf("%d %d", tailX, tailY)] += 1
				continue
			}
			if currY == tailY && currX-tailX == 2 {
				tailX += 1
				tailPositions[fmt.Sprintf("%d %d", tailX, tailY)] += 1
				continue
			}

			if currX-tailX == 2 && (currY-tailY == 1 || currY-tailY == -1) {
				tailX += 1
				tailY = currY
				tailPositions[fmt.Sprintf("%d %d", tailX, tailY)] += 1
				continue
			}

			if currX-tailX == -2 && (currY-tailY == 1 || currY-tailY == -1) {
				tailX -= 1
				tailY = currY
				tailPositions[fmt.Sprintf("%d %d", tailX, tailY)] += 1
				continue
			}

			if currY-tailY == -2 && (currX-tailX == 1 || currX-tailX == -1) {
				tailY -= 1
				tailX = currX
				tailPositions[fmt.Sprintf("%d %d", tailX, tailY)] += 1
				continue
			}

			if currY-tailY == 2 && (currX-tailX == 1 || currX-tailX == -1) {
				tailY += 1
				tailX = currX
				tailPositions[fmt.Sprintf("%d %d", tailX, tailY)] += 1
				continue
			}

		}
	}
	return len(tailPositions)
}
