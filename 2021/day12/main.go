package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileData, _ := ioutil.ReadFile("input.txt")
	fmt.Println(countPaths(string(fileData)))
	fmt.Println(countImprovedPaths(string(fileData)))
}

type Connection struct {
	caveA string
	caveB string
}

func countImprovedPaths(aString string) int {
	connections := initData(aString)
	paths := initPaths(connections, "start")
	paths = doPart2Paths(connections, paths)
	return len(paths)
}

func countPaths(aString string) int {
	connections := initData(aString)
	paths := initPaths(connections, "start")
	paths = doPart1Paths(connections, paths)
	return len(paths)
}

func initData(aString string) []Connection {
	connections := []Connection{}
	firstSplits := strings.Split(strings.TrimSpace(aString), "\n")
	for _, val := range firstSplits {
		secondSplit := strings.Split(val, "-")
		a := strings.TrimSpace(secondSplit[0])
		b := strings.TrimSpace(secondSplit[1])
		if a != "end" && b != "start" {
			connections = append(connections, Connection{a, b})
		}
		if a != "start" && b != "end" {
			connections = append(connections, Connection{b, a})
		}
	}
	return connections
}

func doPart2Paths(data []Connection, paths []string) []string {
	// oldPathCount := len(paths)
	for {
		newPaths := []string{}
		for _, oldPath := range paths {
			splitPath := strings.Split(oldPath, ",")
			lastItem := splitPath[len(splitPath)-1]
			if lastItem == "end" {
				newPaths = appendUnique(newPaths, oldPath)
			} else {
				for _, connection := range allThatStartWith(data, lastItem) {
					newPath := oldPath + "," + connection.caveB
					if !failsPart2Filter(newPath) {
						newPaths = append(newPaths, newPath)
					}
				}
			}
		}
		if allPathsEnd(newPaths) {
			return newPaths
		} else {
			return doPart2Paths(data, newPaths)
		}
	}
}

func stringIsIn(aList []string, aString string) bool {
	for _, val := range aList {
		if val == aString {
			return true
		}
	}
	return false
}

func failsPart2Filter(s string) bool {
	lowerCaseCaveNames := []string{}
	path := strings.Split(s, ",")
	for _, val := range path {
		if strings.ToLower(val) == val && val != "start" && val != "end" {
			if !stringIsIn(lowerCaseCaveNames, val) {
				lowerCaseCaveNames = append(lowerCaseCaveNames, val)
			}
		}
	}
	visitCount := map[string]int{}
	for _, item := range lowerCaseCaveNames {
		for _, cave := range path {
			if item == cave {
				visitCount[cave] += 1
			}
		}
	}
	lowerCaseCavesWith2Visits := 0
	for _, value := range visitCount {
		if value > 2 {
			return true
		}
		if value == 2 {
			lowerCaseCavesWith2Visits += 1
		}
	}
	return lowerCaseCavesWith2Visits > 1
}

func initPaths(data []Connection, start string) []string {

	paths := []string{}
	starts := allThatStartWith(data, start)
	for _, val := range starts {
		path := "start," + val.caveB
		paths = append(paths, path)
	}
	return paths
}

func doPart1Paths(data []Connection, paths []string) []string {
	// oldPathCount := len(paths)
	for {
		newPaths := []string{}
		for _, oldPath := range paths {
			splitPath := strings.Split(oldPath, ",")
			lastItem := splitPath[len(splitPath)-1]
			if lastItem == "end" {
				newPaths = appendUnique(newPaths, oldPath)
			} else {
				for _, connection := range allThatStartWith(data, lastItem) {
					newPath := oldPath + "," + connection.caveB
					if !failsPart1Filter(newPath) {
						newPaths = append(newPaths, newPath)
					}
				}
			}
		}
		if allPathsEnd(newPaths) {
			return newPaths
		} else {
			return doPart1Paths(data, newPaths)
		}
	}
}

func allPathsEnd(paths []string) bool {
	for _, path := range paths {
		if !strings.HasSuffix(path, "end") {
			return false
		}
	}
	return true
}

func appendUnique(aList []string, anItem string) []string {
	tempList := aList
	tempTest := anItem
	for _, val := range tempList {
		if val == tempTest {
			return aList
		}
	}
	return append(aList, anItem)
}

func failsPart1Filter(s string) bool {
	lcs := []string{}
	aList := strings.Split(s, ",")
	for _, val := range aList {
		if strings.ToLower(val) == val {
			lcs = append(lcs, val)
		}
	}
	for _, item := range lcs {
		count := 0
		for _, cave := range aList {
			if item == cave {
				count += 1
			}
			if count >= 2 {
				return true
			}
		}
	}
	return false
}

func allThatStartWith(aList []Connection, aString string) []Connection {
	out := []Connection{}
	for _, val := range aList {
		if val.caveA == aString {
			out = append(out, val)
		}
	}
	return out
}
