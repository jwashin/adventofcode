package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func initData(aString string) []string {
	tableau := []string{}
	for _, line := range strings.Split(aString, "\n") {
		stripped := strings.TrimSpace(line)
		if len(stripped) > 0 {
			tableau = append(tableau, stripped)
		}
	}
	return tableau
}

func same(t1 []string, t2 []string) bool {
	for idx, val := range t1 {
		if t2[idx] != val {
			return false
		}
	}
	return true
}

func move(tableau []string, char string) []string {
	newTableau := []string{}
	f := char + "."
	t := "." + char
	for _, line := range tableau {
		oldFirstChar := string(line[0])
		oldLastChar := string(line[len(line)-1])
		newLine := strings.ReplaceAll(line, f, t)
		if oldLastChar == char && oldFirstChar == "." {
			newLine = oldLastChar + newLine[1:len(newLine)-1] + oldFirstChar
		}
		newTableau = append(newTableau, newLine)
	}
	return newTableau
}

func moveRight(tableau []string) []string {
	critter := ">"
	return move(tableau, critter)
}

func moveDown(tableau []string) []string {
	critter := "v"
	t := transpose(tableau)
	s := move(t, critter)
	return transpose(s)
}

func transpose(aTableau []string) []string {
	newTableau := []string{}
	maps := map[string]string{}
	for y, row := range aTableau {
		for x, char := range row {
			maps[fmt.Sprintf("%d,%d", x, y)] = string(char)
		}
	}
	lenx := len(aTableau[0])
	leny := len(aTableau)
	for x := 0; x < lenx; x++ {
		aString := ""
		for y := 0; y < leny; y++ {
			aString += maps[fmt.Sprintf("%d,%d", x, y)]
		}
		newTableau = append(newTableau, aString)
	}
	return newTableau
}

func doStep(aTableau []string) []string {
	t1 := moveRight(aTableau)
	t2 := moveDown(t1)
	return t2
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	fmt.Println(stepsToStop(string(data)))
}

func stepsToStop(aString string) int {
	tableau := initData(aString)
	steps := 0
	equal := false
	for !equal {
		steps += 1
		tableau2 := doStep(tableau)
		equal = same(tableau, tableau2)
		tableau = tableau2
	}
	return steps
}
