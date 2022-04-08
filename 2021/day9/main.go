package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := GetInputData("input.txt")
	fmt.Println(sumRisk(data))
	fmt.Println(doBasinTask(data))
}

func GetInputData(aString string) string {
	dataFile, _ := ioutil.ReadFile(aString)
	return string(dataFile)
}

func sumRisk(aString string) int {
	input := makeArray(aString)
	value := 0
	xMax := len(input)
	yMax := len(input[0])
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			if isaLowPoint(x, y, input) {
				value += input[x][y] + 1
			}
		}
	}
	return value
}

type Coordinate struct {
	x int
	y int
}

func doBasinTask(aString string) int {
	input := makeArray(aString)
	basinCounts := []int{}
	// basins := [][]Coordinate{}
	lowPoints := findLowPoints(input)
	for _, val := range lowPoints {
		newBasin := MakeBasin(input, val)
		// basins = append(basins, newBasin)
		basinCounts = append(basinCounts, len(newBasin))
	}
	sort.Ints(basinCounts)
	value1 := basinCounts[len(basinCounts)-1]
	value2 := basinCounts[len(basinCounts)-2]
	value3 := basinCounts[len(basinCounts)-3]
	return value1 * value2 * value3

}

func MakeBasin(arr [][]int, lowPoint Coordinate) []Coordinate {
	basin := []Coordinate{}
	finishedPoints := []Coordinate{}
	basin = append(basin, lowPoint)
	for {
		if len(finishedPoints) == len(basin) {
			return basin
		}
		for _, currCell := range basin {
			if cellIsIn(finishedPoints, currCell) {
				continue
			}
			newCells := GetInBasinNeighbors(currCell, arr)
			for _, v := range newCells {
				if !cellIsIn(basin, v) {
					basin = append(basin, v)
				}
			}
			finishedPoints = append(finishedPoints, currCell)
		}
	}
}

func GetInBasinNeighbors(cell Coordinate, arr [][]int) []Coordinate {
	locations := []Coordinate{}
	x := cell.x
	y := cell.y
	if (y - 1) >= 0 {
		value := arr[x][y-1]
		if value < 9 {
			locations = append(locations, Coordinate{x, y - 1})
		}
	}
	if (y + 1) < len(arr[0]) {
		value := arr[x][y+1]
		if value < 9 {
			locations = append(locations, Coordinate{x, y + 1})
		}
	}
	if (x - 1) >= 0 {
		value := arr[x-1][y]
		if value < 9 {
			locations = append(locations, Coordinate{x - 1, y})
		}
	}
	if (x + 1) < len(arr) {
		value := arr[x+1][y]
		if value < 9 {
			locations = append(locations, Coordinate{x + 1, y})
		}
	}
	return locations
}

func cellIsIn(list []Coordinate, cell Coordinate) bool {
	for _, val := range list {
		if val.x == cell.x && val.y == cell.y {
			return true
		}
	}
	return false
}

func findLowPoints(input [][]int) []Coordinate {
	lowPoints := []Coordinate{}
	// input := makeArray(aString)
	xMax := len(input)
	yMax := len(input[0])
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			if isaLowPoint(x, y, input) {
				lowPoints = append(lowPoints, Coordinate{x, y})
			}
		}
	}
	return lowPoints
}

func isaLowPoint(x int, y int, arr [][]int) bool {
	locations := []int{}
	if (y - 1) >= 0 {
		locations = append(locations, arr[x][y-1])
	}
	if (y + 1) < len(arr[0]) {
		locations = append(locations, arr[x][y+1])
	}
	if (x - 1) >= 0 {
		locations = append(locations, arr[x-1][y])
	}
	if (x + 1) < len(arr) {
		locations = append(locations, arr[x+1][y])
	}
	if len(locations) < 2 {
		return false
	}
	currVal := arr[x][y]
	for _, val := range locations {
		if val <= currVal {
			// fmt.Println(x,y,currVal, false)
			return false
		}
	}
	return true
}

func makeArray(input string) [][]int {
	inputData := strings.Split(input, "\n")
	arr := [][]int{}
	for idx := range inputData {
		cleaned := strings.TrimSpace(inputData[idx])
		if len(cleaned) > 0 {
			arr = append(arr, []int{})
			for _, v := range cleaned {
				vi, _ := strconv.Atoi(string(v))
				arr[idx] = append(arr[idx], vi)
			}
		}
	}
	return arr
}
