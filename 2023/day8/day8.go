package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
	fmt.Println("Part2:", part2(string(input)))
}

func getData(s string) (string, map[string][]string) {
	data := strings.Split(strings.TrimSpace(s), "\n")
	vectors := map[string][]string{}
	directions := data[0]
	for _, v := range data[2:] {
		parsed := strings.Fields(v)
		left := parsed[2]
		left = left[1:4]
		right := parsed[3]
		right = right[:3]
		vectors[parsed[0]] = []string{left, right}
	}
	return directions, vectors
}

func part1(s string) int {
	directions, vectors := getData(s)
	start := "AAA"
	end := "ZZZ"
	count := 0
	currLocation := start
	for currLocation != end {
		count += 1
		currInstruction := string(directions[0])
		directions = directions[1:] + currInstruction
		if currInstruction == "L" {
			currLocation = vectors[currLocation][0]
		} else if currInstruction == "R" {
			currLocation = vectors[currLocation][1]
		}
	}
	return count
}

func endState(currentLocation string, count int) bool {
	if currentLocation[2] == 'A' && count > 0 {
		return true
	}
	return false
}

func part2(s string) int {
	directions, vectors := getData(s)
	startList := []string{}
	counts := map[string][]int{}
	for k := range vectors {
		if k[2] == 'A' {
			startList = append(startList, k)
		}
	}

	for _, start := range startList {

		count := 0
		currLocation := start
		for {
			count += 1
			currInstruction := string(directions[0])
			directions = directions[1:] + currInstruction
			if currInstruction == "L" {
				currLocation = vectors[currLocation][0]
			} else if currInstruction == "R" {
				currLocation = vectors[currLocation][1]
			}
			if currLocation[2] == 'Z' {
				counts[start] = append(counts[start], count)
			}
			if len(counts[start]) > 2 {
				// counts[start] = append(counts[start], count)
				// if count > 0 {
				count = 0
				break
				// }
			}
		}
		// counts[start] = append(counts[start], count)
	}
	d := []int{}
	for _, v := range counts {
		d = append(d, v[0])
	}

	product := lcmm(d)
	return product
}

// Euclidean algorithm for finding the greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Returns the least common multiple of two numbers
func lcm(m, n int) int {
	return m / gcd(m, n) * n
}

func lcmm(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

// func part2(s string) int {
// 	// this takes forever
// 	directions, vectors := getData(s)
// 	directionLen := len(directions)
// 	currLocations := []string{}
// 	for k := range vectors {
// 		if k[2] == 'A' {
// 			currLocations = append(currLocations, k)
// 		}
// 	}
// 	count := 0
// 	dirIndex := 0
// 	for !isEndCondition(currLocations) {
// 		count += 1
// 		if count%10000000 == 0 {
// 			fmt.Println(count)
// 		}
// 		if dirIndex == directionLen {
// 			dirIndex = 0
// 		}
// 		currInstruction := directions[dirIndex]
// 		dirIndex += 1
// 		newLocations := []string{}
// 		if currInstruction == 'L' {
// 			for _, v := range currLocations {
// 				newLocations = append(newLocations, vectors[v][0])
// 			}
// 		} else if currInstruction == 'R' {
// 			for _, v := range currLocations {
// 				newLocations = append(newLocations, vectors[v][1])
// 			}
// 		}
// 		currLocations = newLocations
// 	}
// 	return count
// }
//
// func isEndCondition(currentCondition []string) bool {
// 	for _, k := range currentCondition {
// 		if k[2] != 'Z' {
// 			return false
// 		}
// 	}
// 	return true
// }
