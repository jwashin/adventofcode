package main1

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data := GetInputData("input.txt")
	fmt.Println(makeAPath(data))
	// fmt.Println(doBasinTask(data))
}

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func coordinateFromString(aString string) Coordinate {
	var x, y int
	fmt.Sscanf(aString, "%d,%d", &x, &y)
	return Coordinate{x, y}
}

func GetInputData(aString string) string {
	dataFile, _ := ioutil.ReadFile(aString)
	return string(dataFile)
}

func makeAPath(aString string) int {
	riskLevels := makeBigArray(aString)
	totalRisk := 0

	// array bounds
	xMax := 0
	yMax := 0
	for k := range riskLevels {
		testC := coordinateFromString(k)
		if testC.x > xMax {
			xMax = testC.x
		}
		if testC.y > yMax {
			yMax = testC.y
		}
	}
	start := Coordinate{0, 0}
	destination := Coordinate{xMax, yMax}
	// visited := map[string]int{}
	unvisited := map[string]int{}
	prevs := map[string]string{}

	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			unvisited[Coordinate{x, y}.String()] = math.MaxInt
		}
	}

	// riskLevels[start.String()] = 0
	currentNode := start
	unvisited[start.String()] = 0
	// // done initialization
	// dijkstra
	for {
		// check the nodes around the current node
		currentValue := unvisited[currentNode.String()]
		neighbors := []Coordinate{
			{currentNode.x, currentNode.y + 1},
			{currentNode.x, currentNode.y - 1},
			{currentNode.x + 1, currentNode.y},
			{currentNode.x - 1, currentNode.y}}
		for _, node := range neighbors {
			x := node.x
			y := node.y
			if x >= 0 && y >= 0 && x <= xMax && y <= yMax {
				testKey := node.String()
				if unvisited[testKey] > 0 {
					risk := riskLevels[fmt.Sprintf("%d,%d", x, y)]
					tRisk := currentValue + risk
					if tRisk < unvisited[testKey] {
						unvisited[testKey] = tRisk
						prevs[testKey] = currentNode.String()
					}
				}
			}
		}
		delete(unvisited, currentNode.String())
		min := math.MaxInt
		if len(unvisited) == 0 {
			break
		}
		for key, val := range unvisited {
			if val < min {
				min = val
				currentNode = coordinateFromString(key)
			}
		}
		// s := []string{}
		u := destination.String()
		// s = append(s, u)
		for {
			if u != start.String() {
				totalRisk += riskLevels[u]
			}
			p := prevs[u]
			// s = append(s, p)
			if len(p) > 0 {
				u = p
			} else {
				break
			}
		}
	}
	return totalRisk
}

func makeBigArray(input string) map[string]int {
	inData := strings.Split(input, "\n")
	inputData := []string{}
	for _, k := range inData {
		t := strings.TrimSpace(k)
		if len(t) > 0 {
			inputData = append(inputData, t)
		}
	}
	cellWidth := len(strings.TrimSpace(inputData[0]))
	cellHeight := len(inputData)
	arr := map[string]int{}
	for iy, ly := range inputData {
		for lx, v := range ly {
			vi, _ := strconv.Atoi(string(v))
			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					value := clock9(vi + x + y)
					// arr[Coordinate{lx, iy}.String()] = vi
					arr[Coordinate{cellWidth*x + lx, cellHeight*y + iy}.String()] = value
				}
			}
		}
	}

	return arr
}

func clock9(aValue int) int {
	if aValue > 9 {
		return aValue - 9
	}
	return aValue
}
