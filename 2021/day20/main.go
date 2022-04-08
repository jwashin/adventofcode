package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	fmt.Println(enhance(string(input), 50))

}

func getDimensions(aMap map[string]string) (int, int, int, int) {
	minRow := math.MaxInt
	maxRow := math.MinInt
	minColumn := math.MaxInt
	maxColumn := math.MinInt
	for key := range aMap {
		var row, column int
		fmt.Sscanf(key, "%d,%d", &row, &column)
		if row < minRow {
			minRow = row
		}
		if column < minColumn {
			minColumn = column
		}
		if row > maxRow {
			maxRow = row
		}
		if column > maxColumn {
			maxColumn = column
		}
	}
	return minRow, maxRow, minColumn, maxColumn
}

func parseData(aString string) (string, map[string]string) {
	input := strings.Split(aString, "\n")
	var ea string
	data := map[string]string{}
	for row, v := range input {
		if row == 0 {
			ea = strings.TrimSpace(v)
		} else if row >= 2 {
			theString := strings.TrimSpace(v)
			if len(theString) > 0 {
				for col, k := range theString {
					ix := fmt.Sprintf("%d,%d", row-2, col)
					data[ix] = string(k)
				}
			}
		}
	}
	return ea, data
}

func getNine(aString string) []string {
	var frow, fcol int
	fmt.Sscanf(aString, "%d,%d", &frow, &fcol)
	theList := []string{}
	for row := frow - 1; row <= frow+1; row++ {
		for col := fcol - 1; col <= fcol+1; col++ {
			theList = append(theList, fmt.Sprintf("%d,%d", row, col))
		}
	}
	return theList
}

// func pad(theMap map[string]string, aValue int) map[string]string {
// 	minX, maxX, minY, maxY := getDimensions(theMap)
// 	newMap := map[string]string{}
// 	for y := minY - aValue; y <= maxY+aValue; y++ {
// 		for x := minX - aValue; x <= maxX+aValue; x++ {
// 			ky := fmt.Sprintf("%d,%d", x, y)
// 			if theMap[ky] != "#" {
// 				newMap[ky] = "."
// 			} else {
// 				newMap[ky] = "#"
// 			}
// 		}
// 	}
// 	return newMap
// }

func enhance(aString string, count int) int {
	ea, data := parseData(aString)
	print(data)
	for enhancement := 0; enhancement < count; enhancement++ {
		newMap := map[string]string{}

		// evenIsDark := enhancement%2 == 0 && string(ea[0]) != "."
		// minRow, maxRow, minColumn, maxColumn
		minRow, maxRow, minColumn, maxColumn := getDimensions(data)
		outsidePadding := "."
		if string(ea[0]) == "#" {
			if enhancement%2 == 1 {
				outsidePadding = "#"
			}
		}
		for row := minRow - 1; row <= maxRow+1; row++ {
			for col := minColumn - 1; col <= maxColumn+1; col++ {
				key := fmt.Sprintf("%d,%d", row, col)
				binaryString := ""
				pixelKeys := getNine(key)
				for _, pixelKey := range pixelKeys {
					d := data[pixelKey]
					if d == "" {
						d = outsidePadding
					}
					if d == "#" {
						binaryString = binaryString + "1"
					} else if d == "." {
						binaryString = binaryString + "0"
					}
				}
				enhancementIndex, _ := strconv.ParseInt(binaryString, 2, 0)
				outputCharacter := ea[enhancementIndex]
				g := string(outputCharacter)
				if g == "#" {
					newMap[key] = "#"
				} else {
					newMap[key] = "."
				}

			}
			// print(newMap)
		}
		data = newMap
	}
	print(data)
	ct := 0
	for _, val := range data {
		if val == "#" {
			ct += 1
		}
	}

	return ct
}

func print(aMap map[string]string) {
	t := asString(aMap)
	fmt.Println(t)
}

func asString(aMap map[string]string) string {
	minRow, maxRow, minColumn, maxColumn := getDimensions(aMap)
	outStr := ""
	for row := minRow; row <= maxRow; row++ {
		newStr := ""
		for col := minColumn; col <= maxColumn; col++ {
			s := aMap[fmt.Sprintf("%d,%d", row, col)]
			if len(s) > 0 {
				newStr += s
			} else {
				newStr += "."
			}
		}
		outStr += newStr + "\n"
	}
	return outStr
}
