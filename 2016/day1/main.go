package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.TrimSpace(string(input))
	out1, out2 := getDistance(string(data))
	fmt.Printf("1: %v\n", out1)
	fmt.Printf("2: %v\n", out2)
}

type vector struct {
	direction string
	distance  int
}

func getDistance(input string) (int, int) {
	directions := strings.Split(input, ", ")
	currDirection := "North"
	currentLocationN := 0
	currentLocationE := 0
	locations := map[string]int{}
	loc2dist := 0
	for _, value := range directions {
		v := parseDirections(value)
		nd := newDirection(currDirection, v.direction)
		if nd == "North" {
			for v.distance > 0 {
				currentLocationN += 1
				tl := fmt.Sprintf("%d,%d", currentLocationN, currentLocationE)
				locations[tl] += 1
				if locations[tl] == 2 && loc2dist == 0 {
					loc2dist = abs(currentLocationE) + abs(currentLocationN)
				}
				v.distance -= 1
			}
		}
		if nd == "East" {
			for v.distance > 0 {
				currentLocationE += 1
				tl := fmt.Sprintf("%d,%d", currentLocationN, currentLocationE)
				locations[tl] += 1
				if locations[tl] == 2 && loc2dist == 0 {
					loc2dist = abs(currentLocationE) + abs(currentLocationN)
				}
				v.distance -= 1
			}
		}
		if nd == "South" {

			for v.distance > 0 {
				currentLocationN -= 1
				tl := fmt.Sprintf("%d,%d", currentLocationN, currentLocationE)
				locations[tl] += 1
				if locations[tl] == 2 && loc2dist == 0 {
					loc2dist = abs(currentLocationE) + abs(currentLocationN)
				}
				v.distance -= 1
			}
			// currentLocationN -= v.distance
		}
		if nd == "West" {
			for v.distance > 0 {
				currentLocationE -= 1
				tl := fmt.Sprintf("%d,%d", currentLocationN, currentLocationE)
				locations[tl] += 1
				if locations[tl] == 2 && loc2dist == 0 {
					loc2dist = abs(currentLocationE) + abs(currentLocationN)
				}
				v.distance -= 1
			}

			// currentLocationE -= v.distance
		}
		currDirection = nd
	}

	return abs(currentLocationE) + abs(currentLocationN), loc2dist
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func parseDirections(s string) vector {

	direction := string(s[0])
	distance, _ := strconv.Atoi(s[1:])
	return vector{direction: direction, distance: distance}
}

func newDirection(currentDirection string, turn string) string {
	if currentDirection == "North" {
		if turn == "R" {
			return "East"
		}
		if turn == "L" {
			return "West"
		}
	}
	if currentDirection == "South" {
		if turn == "L" {
			return "East"
		}
		if turn == "R" {
			return "West"
		}
	}
	if currentDirection == "East" {
		if turn == "L" {
			return "North"
		}
		if turn == "R" {
			return "South"
		}
	}
	if currentDirection == "West" {
		if turn == "L" {
			return "South"
		}
		if turn == "R" {
			return "North"
		}
	}
	return "None"
}
