package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input)))
}

type coordinate struct {
	row int
	col int
}

func makeTableau(s string) (tableau map[coordinate]rune, height, width int) {
	t := strings.Split(strings.TrimSpace(s), "\n")
	tableau = map[coordinate]rune{}
	for row, v := range t {
		for col, char := range v {
			if char != '.' {
				c := coordinate{row, col}
				tableau[c] = char
			}
		}
	}
	return tableau, len(t), len(t[0])
}

func turnRight(currGlyph rune) rune {

	rmap := map[rune]rune{'>': 'v', 'v': '<', '<': '^', '^': '>'}

	return rmap[currGlyph]

}

func getStart(tableau map[coordinate]rune) (rune, coordinate) {
	cands := "<>^v"
	var start = coordinate{}
	var glyph = ' '

	for k, v := range tableau {
		if strings.ContainsRune(cands, v) {
			glyph = v
			start = k
			break
		}
	}
	return glyph, start
}

func part1(s string) int {
	// 5533 too low
	tableau, height, width := makeTableau(s)
	currGlyph, currLoc := getStart(tableau)
	delete(tableau, currLoc)
	usedLocations := map[coordinate]bool{currLoc: true}
	moveMap := map[rune]coordinate{'^': {-1, 0},
		'v': {1, 0},
		'>': {0, 1},
		'<': {0, -1},
	}
	for currLoc.col > 0 && currLoc.row > 0 && currLoc.col < width-1 && currLoc.row < height-1 {
		candNextMove := moveMap[currGlyph]
		candNextLoc := coordinate{currLoc.row + candNextMove.row, currLoc.col + candNextMove.col}
		_, ok := tableau[candNextLoc]
		if ok {
			currGlyph = turnRight(currGlyph)
		} else {

			currLoc = candNextLoc
			usedLocations[currLoc] = true
		}

	}
	return len(usedLocations)
}
func part2(s string) int {

	tableau, height, width := makeTableau(s)
	startGlyph, startLoc := getStart(tableau)
	delete(tableau, startLoc)

	moveMap := map[rune]coordinate{'^': {-1, 0},
		'v': {1, 0},
		'>': {0, 1},
		'<': {0, -1},
	}
	goodObstacleCount := 0

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			candObstacleLoc := coordinate{row, col}
			_, ok := tableau[candObstacleLoc]
			if ok || candObstacleLoc == startLoc {
				continue
			}
			tableau[candObstacleLoc] = '#'

			currGlyph, currLoc := startGlyph, startLoc
			usedLocations := map[coordinate]rune{currLoc: currGlyph}

			for currLoc.col > 0 && currLoc.row > 0 && currLoc.col < width-1 && currLoc.row < height-1 {
				candNextMove := moveMap[currGlyph]
				candNextLoc := coordinate{currLoc.row + candNextMove.row, currLoc.col + candNextMove.col}
				_, ok := tableau[candNextLoc]
				if ok {
					currGlyph = turnRight(currGlyph)
				} else {

					currLoc = candNextLoc
					val, ok := usedLocations[currLoc]
					if ok {
						if val == currGlyph {
							goodObstacleCount += 1
							break
						}
					}
					usedLocations[currLoc] = currGlyph
				}

			}
			delete(tableau, candObstacleLoc)
		}
	}
	return goodObstacleCount
}
