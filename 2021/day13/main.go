package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dataFile, _ := ioutil.ReadFile("input.txt")
	// paper, instructions := getData(string(dataFile))
	fmt.Println(doFirstInstruction(string(dataFile)))

	paper := doAllInstructions(string(dataFile))
	// paper := doAllInstructions(string(`6,10
	// 0,14
	// 9,10
	// 0,3
	// 10,4
	// 4,11
	// 6,0
	// 6,12
	// 4,1
	// 0,13
	// 10,12
	// 3,4
	// 3,0
	// 8,4
	// 1,10
	// 2,14
	// 8,10
	// 9,0

	// fold along y=7`))
	display(paper)
}

func display(paper map[string]int) {
	maxX := 0
	maxY := 0
	for k := range paper {
		xy := strings.Split(k, ",")
		Y, _ := strconv.Atoi(xy[1])
		X, _ := strconv.Atoi(xy[0])
		if Y > maxY {
			maxY = Y
		}
		if X > maxX {
			maxX = X
		}
	}
	output := []string{}
	for y := 0; y < maxY+1; y++ {
		s := []string{}
		for x := 0; x < maxX+1; x++ {
			if paper[fmt.Sprintf("%d,%d", x, y)] == 1 {
				s = append(s, "#")
			} else {
				s = append(s, " ")
			}
		}
		output = append(output, strings.Join(s, ""))

	}
	for _, val := range output {
		fmt.Println(val)
	}
}

func doAllInstructions(dataFile string) map[string]int {
	paperData, instructionsData := getData(dataFile)
	paper := map[string]int{}
	for _, value := range paperData {
		paper[value] = 1
	}
	for _, instruction := range instructionsData {
		paper = doInstruction(paper, instruction)
	}
	return paper
}

func doInstruction(paper map[string]int, instruction string) map[string]int {
	// display(paper)
	iSplit := strings.Split(instruction, "=")
	xory := string(iSplit[0][len(iSplit[0])-1])
	instructionValue, _ := strconv.Atoi(iSplit[1])
	res := map[string]int{}
	if xory == "y" {
		res = yFold(paper, instructionValue)
	}
	if xory == "x" {
		res = xFold(paper, instructionValue)
	}
	return res
}

func doFirstInstruction(dataFile string) int {
	paperData, instructionsData := getData(dataFile)
	paper := map[string]int{}
	for _, value := range paperData {
		paper[value] = 1
	}
	instruction := instructionsData[0]
	iSplit := strings.Split(instruction, "=")
	xory := string(iSplit[0][len(iSplit[0])-1])
	instructionValue, _ := strconv.Atoi(iSplit[1])
	if xory == "y" {
		res := yFold(paper, instructionValue)
		return len(res)
	}
	if xory == "x" {
		res := xFold(paper, instructionValue)
		return len(res)
	}
	return 0
}

func getData(dataFile string) ([]string, []string) {
	sf := strings.Split(string(dataFile), "\n")
	paper := []string{}
	instructions := []string{}
	for _, value := range sf {
		value := strings.TrimSpace(value)
		if len(value) == 0 {
			continue
		}
		if strings.HasPrefix(value, "fold") {
			instructions = append(instructions, value)
		} else {
			paper = append(paper, value)
		}

	}
	return paper, instructions
}

func yFold(data map[string]int, yFold int) map[string]int {
	// fold the bottom half up
	dst := map[string]int{}
	for k, v := range data {
		xy := strings.Split(k, ",")
		oldY, _ := strconv.Atoi(xy[1])
		oldX, _ := strconv.Atoi(xy[0])
		if oldY < yFold && v == 1 {
			dst[k] = v
		} else {
			newX := oldX
			newY := 2*yFold - oldY
			dst[fmt.Sprintf("%d,%d", newX, newY)] = 1
		}
	}
	return dst
}

func xFold(data map[string]int, xFold int) map[string]int {
	// fold the bottom half up
	dst := map[string]int{}
	for k, v := range data {
		xy := strings.Split(k, ",")
		oldY, _ := strconv.Atoi(xy[1])
		oldX, _ := strconv.Atoi(xy[0])
		if oldX < xFold && v == 1 {
			dst[k] = v
		} else {
			newY := oldY
			newX := 2*xFold - oldX
			dst[fmt.Sprintf("%d,%d", newX, newY)] = 1
		}
	}
	return dst
}
