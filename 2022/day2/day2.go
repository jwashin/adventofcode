package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	fmt.Println("Part 1", part1(data))
	fmt.Println("Part 2", part2(data))
}

func part1(data string) int {
	return playGame(data)
}
func part2(data string) int {
	return playGame2(data)
}

func gamePlay(opponent string, player string) int {
	score := 0
	if opponent == "A" {
		// opponent rock
		if player == "X" {
			// player rock
			score += 1
			// draw
			score += 3
		} else if player == "Y" {
			// player paper
			score += 2
			// win
			score += 6
		} else if player == "Z" {
			// player scissors
			score += 3
			//lose
			score += 0
		}
	} else if opponent == "B" {
		// opponent paper
		if player == "Y" {
			// player paper
			score += 2
			// draw
			score += 3
		} else if player == "Z" {
			// player scissors
			score += 3
			// win
			score += 6
		} else if player == "X" {
			// player rock
			score += 1
			//lose
			score += 0
		}
	} else if opponent == "C" {
		// opponent scissors
		if player == "Z" {
			// player scissors
			score += 3
			// draw
			score += 3
		} else if player == "X" {
			// player rock
			score += 1
			// win
			score += 6
		} else if player == "Y" {
			// player paper
			score += 2
			//lose
			score += 0
		}
	}
	return score
}

func gamePlay2(opponent string, player string) int {
	if player == "Z" {
		// should win
		if opponent == "A" {
			// opponent rock
			return gamePlay(opponent, "Y")
		} else if opponent == "B" {
			// opponent paper
			return gamePlay(opponent, "Z")

		} else if opponent == "C" {
			// opponent scissors
			return gamePlay(opponent, "X")
		}
	} else if player == "X" {
		// should lose
		if opponent == "A" {
			// opponent rock
			return gamePlay(opponent, "Z")
		} else if opponent == "B" {
			// opponent paper
			return gamePlay(opponent, "X")

		} else if opponent == "C" {
			// opponent scissors
			return gamePlay(opponent, "Y")
		}
	} else if player == "Y" {
		// should draw
		if opponent == "A" {
			// opponent rock
			return gamePlay(opponent, "X")
		} else if opponent == "B" {
			// opponent paper
			return gamePlay(opponent, "Y")

		} else if opponent == "C" {
			// opponent scissors
			return gamePlay(opponent, "Z")
		}
	}
	return 0
}

func playGame(s string) int {
	t := strings.Split(s, "\n")
	total := 0
	for _, v := range t {
		j := strings.Fields(v)
		score := gamePlay(j[0], j[1])
		total += score
	}
	return total
}

func playGame2(s string) int {
	t := strings.Split(s, "\n")
	total := 0
	for _, v := range t {
		j := strings.Fields(v)
		score := gamePlay2(j[0], j[1])
		total += score
	}
	return total
}
