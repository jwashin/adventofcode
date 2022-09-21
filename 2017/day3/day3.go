package main

import "fmt"

var puzzleInput = 361527

func main() {
	fmt.Println("1.", spiralDistance(puzzleInput))
	fmt.Println("2.", part2())
}

func neighboringCoordinates(y, x int) []string {
	t := []string{}
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if !(i == y && j == x) {
				t = append(t, fmt.Sprintf("%d %d", i, j))
			}
		}
	}
	return t
}

func part2() int {
	// spiral["y x"] = value
	spiral := map[string]int{"0 0": 1}
	n := 1
	for {
		n += 1
		y, x := coordfromX(n)
		value := 0
		for _, v := range neighboringCoordinates(y, x) {
			value += spiral[v]
		}
		if value > puzzleInput {
			return value
		}
		spiral[fmt.Sprintf("%d %d", y, x)] = value
		// if n < 8 {
		// 	fmt.Println(value)
		// }
	}
}

func spiralDistance(s int) int {
	x, y := coordfromX(s)
	return abs(x) + abs(y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func coordfromX(n int) (int, int) {
	t := leastSquareGE(n)
	currentValue := t * t
	valuesPerRow := t
	startx := t / 2
	starty := t / 2

	// go counterclockwise from the bottom right
	x := startx
	y := starty

	// bottom row, going left
	for x > startx-(valuesPerRow-1) {
		if currentValue == n {
			return y, x
		}
		x -= 1
		currentValue -= 1
	}
	// left side, going up
	for y > starty-(valuesPerRow-1) {
		if currentValue == n {
			return y, x
		}
		y -= 1
		currentValue -= 1
	}
	// top row, going right
	for x < startx {
		if currentValue == n {
			return y, x
		}
		x += 1
		currentValue -= 1
	}
	// right side, going down
	for y <= starty {
		if currentValue == n {
			return y, x
		}
		y += 1
		currentValue -= 1
	}

	return y, x
}

func leastSquareGE(n int) int {
	// get the square root of the number at the
	// lower right corner of the square holding n
	if n == 1 {
		return 1
	}
	x := 1
	sq := x
	for sq < n {
		x += 2
		sq = x * x
	}
	return x
}
