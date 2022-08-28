package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	t := findShortestPath(1, 1, 31, 39, 1362)
	fmt.Println("1. ", t)
	s := closerThan50(1, 1, 1362)
	fmt.Println("2. ", s)
}

func genChar(x int, y int, factor int) string {

	sum := (x*x + 3*x + 2*x*y + y + y*y) + factor

	// sum := ((x + y) * (x + y)) + 3*x + y + factor

	binrep := strconv.FormatInt(int64(sum), 2)
	oneBits := strings.Count(binrep, "1")
	if oneBits%2 == 0 {
		return "."
	}
	return "#"
}

func gentestSpace(factor int) string {
	t := []string{"  0123456789"}
	for ky, y := range []int{0, 1, 2, 3, 4, 5, 6} {
		newString := fmt.Sprint(ky) + " "
		for _, x := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
			newString = newString + genChar(x, y, factor)
		}
		t = append(t, newString)
	}
	return strings.Join(t, "\n")
}

// https://www.sohamkamani.com/golang/sets/
type vertexSet map[string]struct{}

// Adds an item to the set
func (s vertexSet) add(vertex string) {
	s[vertex] = struct{}{}
}

// // Removes an item from the set
// func (s vertexSet) remove(vertex string) {
// 	delete(s, vertex)
// }

// Returns a boolean value describing if the item exists in the set
func (s vertexSet) has(vertex string) bool {
	_, ok := s[vertex]
	return ok
}

func coords(item string) (x, y int) {
	// gotta be "x,y"
	t := strings.Split(item, ",")
	x, _ = strconv.Atoi(t[0])
	y, _ = strconv.Atoi(t[1])
	return
}

func isaSpace(x, y int, factor int) bool {
	return genChar(x, y, factor) == "."
}

func closerThan50(startx, starty, factor int) int {
	count := 0
	for y := 0; y < 51; y++ {
		for x := 0; x < 51; x++ {
			if isaSpace(x, y, factor) {
				dist := findShortestPath(startx, starty, x, y, factor)
				if dist <= 50 {
					count += 1
				}

			}

		}
	}

	return count
}

func findShortestPath(startx, starty, destx, desty int, factor int) int {
	// each string is an x,y coordinate
	if startx == destx && starty == desty {
		return 0
	}
	used := vertexSet{}
	q := map[string]int{"1,1": 0}
	dest := fmt.Sprintf("%d,%d", destx, desty)
	for len(q) > 0 {
		// find smallest currScore
		minScore := math.MaxInt
		currentItem := ""
		for k, v := range q {
			if v < minScore {
				minScore = v
				currentItem = k
			}
		}
		currentx, currenty := coords(currentItem)
		delete(q, currentItem)
		used.add(currentItem)

		// check the nodes around the current node
		for _, value := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			if isaSpace(currentx+value[0], currenty+value[1], factor) {
				asString := fmt.Sprintf("%d,%d", currentx+value[0], currenty+value[1])
				if !used.has(asString) {
					if q[asString] > minScore+1 {
						q[asString] = minScore + 1
						continue
					}
					q[asString] = minScore + 1
					if asString == dest {
						return minScore + 1
					}
				}

			}
		}
	}
	return math.MaxInt
}
