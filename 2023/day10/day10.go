package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
}

type Coordinate struct {
	row int
	col int
}

func findStart(tableau []string) (Coordinate, error) {
	for rowindex, row := range tableau {
		for columnIndex, character := range row {
			if character == 'S' {
				return Coordinate{rowindex, columnIndex}, nil
			}
		}
	}
	return Coordinate{}, errors.New("not found")
}

func currentCharacter(c Coordinate, tableau []string) byte {
	return tableau[c.row][c.col]
}

func candidates(currentLocation Coordinate, tableau []string) []Coordinate {
	currentChar := currentCharacter(currentLocation, tableau)
	out := []Coordinate{}
	a, b := Coordinate{0, 0}, Coordinate{0, 0}
	if currentChar == '7' {
		a = Coordinate{currentLocation.row, currentLocation.col - 1}
		b = Coordinate{currentLocation.row + 1, currentLocation.col}
	} else if currentChar == '|' {
		a = Coordinate{currentLocation.row - 1, currentLocation.col}
		b = Coordinate{currentLocation.row + 1, currentLocation.col}
	} else if currentChar == '-' {
		a = Coordinate{currentLocation.row, currentLocation.col - 1}
		b = Coordinate{currentLocation.row, currentLocation.col + 1}
	} else if currentChar == 'F' {
		a = Coordinate{currentLocation.row, currentLocation.col + 1}
		b = Coordinate{currentLocation.row + 1, currentLocation.col}
	} else if currentChar == 'J' {
		a = Coordinate{currentLocation.row - 1, currentLocation.col}
		b = Coordinate{currentLocation.row, currentLocation.col - 1}
	} else if currentChar == 'L' {
		a = Coordinate{currentLocation.row - 1, currentLocation.col}
		b = Coordinate{currentLocation.row, currentLocation.col + 1}
	} else if currentChar == 'S' {
		s_candidates := []Coordinate{}
		if currentLocation.row != 0 {
			s_candidates = append(s_candidates, Coordinate{currentLocation.row - 1, currentLocation.col})
		}
		if currentLocation.row < len(tableau) {
			s_candidates = append(s_candidates, Coordinate{currentLocation.row + 1, currentLocation.col})
		}
		if currentLocation.col < len(tableau[0]) {
			s_candidates = append(s_candidates, Coordinate{currentLocation.row, currentLocation.col + 1})
		}
		if currentLocation.col > 0 {
			s_candidates = append(s_candidates, Coordinate{currentLocation.row, currentLocation.col - 1})
		}
		for index, cand := range s_candidates {
			if index == 0 && strings.Contains("7|F", string(currentCharacter(cand, tableau))) {
				out = append(out, cand)
			}
			if index == 1 && strings.Contains("L|J", string(currentCharacter(cand, tableau))) {
				out = append(out, cand)
			}
			if index == 2 && strings.Contains("7-J", string(currentCharacter(cand, tableau))) {
				out = append(out, cand)
			}
			if index == 3 && strings.Contains("F-L", string(currentCharacter(cand, tableau))) {
				out = append(out, cand)
			}
		}
		return out
	}
	// if !(a.row == currentLocation.row && a.col == currentLocation.col) {
	// 	out = append(out, a)
	// }
	// if !(b.row == currentLocation.row && b.col == currentLocation.col) {
	// 	out = append(out, b)
	// }
	out = append(out, a)
	out = append(out, b)
	return out
}

func part1(s string) int {
	tableau := strings.Split(strings.TrimSpace(s), "\n")
	// find S
	start, err := findStart(tableau)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	// count := 0
	distances := map[Coordinate]int{}
	distances[start] = 0
	// currentLocation := start
	queue := []Coordinate{start}
	used := map[Coordinate]int{}
	for len(queue) > 0 {
		cand := queue[0]
		used[cand] = 1
		queue = queue[1:]
		// count += 1

		locs := candidates(cand, tableau)
		for _, v := range locs {
			if used[v] != 1 {
				distances[v] = distances[cand] + 1
				queue = append(queue, v)
			}
		}

	}
	max := 0
	for _, v := range distances {
		if v > max {
			max = v
		}
	}
	return max
}
func part2(s string) int {
	tableau := strings.Split(strings.TrimSpace(s), "\n")
	// find S
	start, err := findStart(tableau)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	// count := 0
	distances := map[Coordinate]int{}
	distances[start] = 0
	// currentLocation := start
	queue := []Coordinate{start}
	used := map[Coordinate]int{}
	for len(queue) > 0 {
		cand := queue[0]
		used[cand] = 1
		queue = queue[1:]
		// count += 1
		locs := candidates(cand, tableau)
		for _, v := range locs {
			if used[v] != 1 {
				distances[v] = distances[cand] + 1
				queue = append(queue, v)
			}
		}
	}
	insideCount := 0
	insideItems := []Coordinate{}
	for row, line := range tableau {
		outside := used[Coordinate{row, 0}] != 1
		for col := range line[1:] {
			currLoc := Coordinate{row, col + 1}
			if used[currLoc] == 1 {
				if strings.Contains("|", string(currentCharacter(Coordinate{row, col + 1}, tableau))) {
					outside = !outside
				}
			} else if !outside {
				insideCount += 1
				insideItems = append(insideItems, currLoc)
			}
		}

	}

	return len(insideItems)
}
