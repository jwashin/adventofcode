package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1", part1(string(input)))
	fmt.Println("Part 2", getPowerSum(string(input)))
}

func part1(s string) int {
	r := 12
	g := 13
	b := 14
	return sumOfPossibleIds(r, g, b, s)

}

func isPossibleSample(s string, r int, g int, b int) bool {
	s = strings.TrimSpace(s)
	cubes := strings.Split(s, ",")
	for _, countAndColor := range cubes {
		var count = 0
		var color = ""
		_, err := fmt.Sscanf(countAndColor, "%d %s", &count, &color)
		if err == nil {
			if color == "red" && count > r {
				return false
			}
			if color == "green" && count > g {
				return false
			}
			if color == "blue" && count > b {
				return false
			}
		}
	}
	return true
}

func isPossibleGame(gameString string, r int, g int, b int) bool {

	samples := strings.Split(gameString, ";")
	for _, v := range samples {
		if !isPossibleSample(v, r, g, b) {
			return false
		}
	}
	return true
}

func sumOfPossibleIds(r int, g int, b int, s string) int {
	// for each sample in a game reds must be
	// less than r, blues must be less than b,
	// and greens must be less than g.
	// if true, add the number of the game
	// example:
	// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	idSum := 0
	s = strings.TrimSpace(s)
	gameSplit := strings.Split(s, "\n")
	for _, gm := range gameSplit {
		t := strings.Split(gm, ":")
		// game number is in t[0]
		gameString := t[1]
		if isPossibleGame(gameString, r, g, b) {
			split := strings.Split(t[0], " ")
			gameStr := split[1]
			gameVal, _ := strconv.Atoi(gameStr)
			idSum += gameVal
		}
	}
	return idSum
}

func getPowerSum(s string) int {
	powerSum := 0
	s = strings.TrimSpace(s)
	gameSplit := strings.Split(s, "\n")
	for _, gm := range gameSplit {
		t := strings.Split(gm, ":")
		// game number is in t[0]
		gameString := t[1]
		powerSum += power(gameString)
	}
	return powerSum
}

func power(s string) int {
	r, g, b := 0, 0, 0

	samples := strings.Split(s, ";")
	// split on ; for samples
	for _, v := range samples {
		// split on "," for individual colors and counts
		ccs := strings.Split(v, ",")
		for _, countAndColor := range ccs {
			var count = 0
			var color = ""
			_, err := fmt.Sscanf(countAndColor, "%d %s", &count, &color)
			if err == nil {
				if color == "red" && count > r {
					r = count
				}
				if color == "green" && count > g {
					g = count
				}
				if color == "blue" && count > b {
					b = count
				}
			}

		}

	}
	return r * g * b
}
