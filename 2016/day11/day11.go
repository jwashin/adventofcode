package main

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
	"strings"
)

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

type Graph struct {
	vertices map[string]*Vertex
}

type Vertex struct {
	dist int
	prev string
}

func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	fmt.Println(doGame(false))
}

func getData(test bool) string {
	if test {
		return "E,HM,LM\nHG\nLG\n"
	}
	// return "E,PG,TG,TM,XG,RG,RM,WG,WM\nPM,XM\n\n"
	return "E,PG,TG,TM,XG,RG,RM,WG,WM,YG,YM,ZG,ZM\nPM,XM\n\n"
}

func getSolution(test bool) string {
	if test {
		return normalize("\n\n\nE,HM,LM,HG,LG")
	}
	return normalize("\n\n\nE,PG,TG,TM,XG,RG,RM,WG,WM,PM,XM,YG,YM,ZG,ZM")
}

func contains(ls []string, s string) bool {
	for _, v := range ls {
		if v == s {
			return true
		}
	}
	return false
}

func removeItems(f string, r []string) string {
	s := strings.Split(f, ",")
	newList := []string{}
	for _, v := range s {
		if !contains(r, v) {
			newList = append(newList, v)
		}
	}
	return strings.Join(newList, ",")
}

func addItems(f string, r []string) string {
	s := []string{f}
	if strings.Contains(f, ",") {
		s = strings.Split(f, ",")
	}
	for _, v := range r {
		if !contains(s, v) {
			s = append(s, v)
		}
	}
	if len(s) > 1 {
		return strings.Join(s, ",")
	}
	return s[0]
}

func doGame(test bool) int {
	data := getData(test)
	solution := getSolution(test)
	s := normalize(data)
	used := vertexSet{}
	currentSolution := math.MaxInt
	// initialize the graph
	// string is the string representation of any item in the graph
	graph := Graph{vertices: map[string]*Vertex{s: {dist: 0, prev: ""}}}
	q := graph.vertices
	oldc := 0
	for len(q) > 0 {
		// fmt.Println(currentSolution, len(q), len(used))
		// find min solution so far in unvisited
		min := math.MaxInt
		currentItem := ""
		// currentSolution := math.MaxInt
		for itemRepresentation, v := range q {
			if v.dist < min {
				min = v.dist
				currentItem = itemRepresentation
			}
		}
		// remove currentItem
		currentSolution = min
		if oldc != currentSolution {
			fmt.Println(currentSolution, len(q), len(used))
			oldc = currentSolution
		}
		delete(q, currentItem)
		used.add(currentItem)
		// used = append(used, currentItem)
		// fmt.Println(currentSolution, currentItem)
		// debugprint(currentSolution, currentItem)

		// item is a string. enumerate next possible moves
		floors := fromTableau(currentItem)
		// get the one with the elevator
		floorIdx := currentFloor(floors)
		floorItemsx := strings.Split(floors[floorIdx], ",")
		floorItems := []string{}
		for _, v := range floorItemsx {
			if v != "E" && v != "" {
				floorItems = append(floorItems, v)
			}
		}
		// pull out the elevator for now
		floors[floorIdx] = strings.Join(floorItems, ",")
		toAdd := [][]string{}
		for _, v := range floorItems {
			toAdd = append(toAdd, []string{v})
		}
		doubles := Combinations2(floorItems, 2)
		toAdd = append(toAdd, doubles...)
		neighborRepresentations := []string{}
		newFloor := floors[floorIdx]
		for _, item := range toAdd {
			floorT := removeItems(newFloor, item)
			if isSafeFloor(floorT) {
				for _, ix := range possibleFloors(floorIdx) {
					floorx := addItems(floors[ix], item)
					if isSafeFloor(floorx) {
						newTableau := []string{}
						for kx, v := range floors {
							if kx == ix {
								// put the elevator back on the destination floor
								floorx := addItems(floorx, []string{"E"})
								newTableau = append(newTableau, floorx)
								continue
							}
							if kx == floorIdx {
								newTableau = append(newTableau, floorT)
								continue
							}
							newTableau = append(newTableau, v)
						}
						nt := normalize(strings.Join(newTableau, "\n"))
						neighborRepresentations = append(neighborRepresentations, nt)
					}
				}

			}
		}
		for _, nextState := range neighborRepresentations {
			// replacement for "still in q"
			// previouslyVisited := false
			if used.has(nextState) {
				continue
			}
			newVertex := Vertex{prev: currentItem, dist: currentSolution + 1}
			q[nextState] = &newVertex
			if nextState == solution {
				return currentSolution + 1
			}

		}

	}
	return currentSolution
}

// func debugprint(score int, outd string) {
// 	fmt.Println(score)
// 	out := strings.Split(outd, "\n")
// 	Reverse(out)
// 	for _, v := range out {
// 		fmt.Println(v)
// 	}
// }

func possibleFloors(currFloor int) []int {

	t := [][]int{{1}, {0, 2}, {1, 3}, {2}}
	return t[currFloor]

}

func currentFloor(s []string) int {
	for k, v := range s {
		if strings.Contains(v, "E") {
			return k
		}
	}
	return 40
}

func sortFloor(s string) string {
	if strings.Contains(s, ",") {
		t := strings.Split(s, ",")
		sort.Strings(t)
		nf := []string{}
		for _, v := range t {
			if len(v) > 0 {
				nf = append(nf, v)
			}
		}
		return strings.Join(nf, ",")
	}
	return s
}

// func makeTableau(data []string) string {
// 	out := []string{}
// 	for _, row := range data {
// 		sp := sortFloor(row)
// 		out = append(out, sp)
// 	}
// 	return strings.Join(out, "\n")
// }

func normalize(tableau string) string {
	t := []string{}
	for _, v := range strings.Split(tableau, "\n") {
		t = append(t, sortFloor(v))
	}
	return strings.Join(t, "\n")
}

func fromTableau(s string) []string {
	return strings.Split(s, "\n")
}

func isSafeFloor(s string) bool {
	t := strings.Split(s, ",")
	// collect the generators
	generators := ""
	for _, item := range t {
		if len(item) == 2 && string(item[1]) == "G" {
			generators += string(item[0])
		}
	}
	if len(generators) == 0 {
		return true
	}
	// collect the microchips.
	chips := ""
	for _, item := range t {
		if len(item) == 2 && string(item[1]) == "M" {
			chips += string(item[0])
		}
	}

	for _, item := range chips {
		if !strings.Contains(generators, string(item)) {
			return false
		}
	}

	return true
}

// https://www.sobyte.net/post/2022-01/go-slice/
func Combinations2(set []string, n int) (subsets [][]string) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

// dijkstra
// for {
// 	// check the nodes around the current node
// 	currentValue := unvisited[currentNode.String()]
// 	neighbors := []Coordinate{
// 		{currentNode.x, currentNode.y + 1},
// 		{currentNode.x, currentNode.y - 1},
// 		{currentNode.x + 1, currentNode.y},
// 		{currentNode.x - 1, currentNode.y}}
// 	for _, node := range neighbors {
// 		x := node.x
// 		y := node.y
// 		if x >= 0 && y >= 0 && x <= xMax && y <= yMax {
// 			testKey := node.String()
// 			if unvisited[testKey] > 0 {
// 				risk := riskLevels[fmt.Sprintf("%d,%d", x, y)]
// 				tRisk := currentValue + risk
// 				if tRisk < unvisited[testKey] {
// 					unvisited[testKey] = tRisk
// 					prevs[testKey] = currentNode.String()
// 				}
// 			}
// 		}
// 	}
// 	delete(unvisited, currentNode.String())
// 	min := math.MaxInt
// 	if len(unvisited) == 0 {
// 		break
// 	}
// 	for key, val := range unvisited {
// 		if val < min {
// 			min = val
// 			currentNode = coordinateFromString(key)
// 		}
// 	}
