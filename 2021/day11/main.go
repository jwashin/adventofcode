package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data := getInputFile()
	fmt.Println(flashOctopi(data, 100))
	fmt.Println(firstStepToUnisonFlash(data))
}

func getInputFile() string {
	d, _ := ioutil.ReadFile("input.txt")
	return string(d)
}

type Coordinate struct {
	x int
	y int
}

func firstStepToUnisonFlash(input string) int {
	inputData := strings.Split(input, "\n")
	arr := [][]int{}
	// initialize setup
	for _, val := range inputData {
		row := []int{}
		val = strings.TrimSpace(val)
		if len(val) > 0 {
			for _, e := range val {
				energy, _ := strconv.Atoi(string(e))
				row = append(row, energy)
			}
			arr = append(arr, row)
		}
	}
	maxX := len(arr[0])
	maxY := len(arr)
	target := maxX * maxY
	// add 1 energy to each octopus
	step := 1
	for {
		flashingOctopi := []Coordinate{}
		for x := 0; x < maxX; x++ {
			for y := 0; y < maxY; y++ {
				arr[x][y] += 1
			}
		}
		// flash octopi
		doAgain := true
		for doAgain {
			doAgain = false
			for x := 0; x < maxX; x++ {
				for y := 0; y < maxY; y++ {
					if arr[x][y] >= 10 && !listContainsCoordinate(flashingOctopi, Coordinate{x, y}) {
						doAgain = true
						flashingOctopi = append(flashingOctopi, Coordinate{x, y})
						incrementNeighbors(x, y, maxX, maxY, arr)
					}
				}
			}
		}
		count := len(flashingOctopi)
		if count == target {
			return step
		}
		// zero all flashing octopi
		for _, value := range flashingOctopi {
			arr[value.x][value.y] = 0
		}
		step += 1
	}
}

func flashOctopi(input string, steps int) int {
	inputData := strings.Split(input, "\n")
	totalFlashes := 0
	arr := [][]int{}
	// initialize setup
	for _, val := range inputData {
		row := []int{}
		val = strings.TrimSpace(val)
		if len(val) > 0 {
			for _, e := range val {
				energy, _ := strconv.Atoi(string(e))
				row = append(row, energy)
			}
			arr = append(arr, row)
		}
	}
	maxX := len(arr[0])
	maxY := len(arr)
	// add 1 energy to each octopus
	for step := 0; step < steps; step++ {
		flashingOctopi := []Coordinate{}
		for x := 0; x < maxX; x++ {
			for y := 0; y < maxY; y++ {
				arr[x][y] += 1
			}
		}
		// flash octopi
		doAgain := true
		for doAgain {
			doAgain = false
			for x := 0; x < maxX; x++ {
				for y := 0; y < maxY; y++ {
					if arr[x][y] >= 10 && !listContainsCoordinate(flashingOctopi, Coordinate{x, y}) {
						doAgain = true
						flashingOctopi = append(flashingOctopi, Coordinate{x, y})
						incrementNeighbors(x, y, maxX, maxY, arr)
					}
				}
			}
		}
		totalFlashes += len(flashingOctopi)
		// zero all flashing octopi
		for _, value := range flashingOctopi {
			arr[value.x][value.y] = 0
		}
	}
	return totalFlashes
}

func incrementNeighbors(x int, y int, maxX int, maxY int, arr [][]int) {
	topX := x + 1
	topY := y + 1
	smallX := x - 1
	smallY := y - 1
	listOfNeighbors := []Coordinate{}
	for x1 := smallX; x1 <= topX; x1++ {
		for y1 := smallY; y1 <= topY; y1++ {
			listOfNeighbors = append(listOfNeighbors, Coordinate{x1, y1})
		}
	}

	for _, val := range listOfNeighbors {
		if val.x >= 0 && val.x < maxX && val.y >= 0 && val.y < maxY {
			arr[val.x][val.y] += 1
		}
	}
}

func listContainsCoordinate(coordinates []Coordinate, c Coordinate) bool {
	for _, val := range coordinates {
		if c.x == val.x && c.y == val.y {
			return true
		}
	}
	return false
}
