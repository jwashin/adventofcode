package main

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
	"strings"
)

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
	fmt.Println(doGame(true))
}

func getData(test bool) []string {
	if test {
		return []string{"E,HM,LM\nHG\nLG\n"}
	}
	return []string{"E,PG,TG,TM,XG,RG,RM,CG,CM\nPM,XM\n\n"}
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
	s := makeTableau(data)
	unvisited := map[string]int{s: 0}
	prev := map[string]string{}
	currentSolution := 0
	for len(unvisited) > 0 {
		// find min solution so far in unvisited
		min := math.MaxInt
		currentItem := ""
		// currentSolution := math.MaxInt
		for k, v := range unvisited {
			if v < min {
				min = v
				currentItem = k
			}
		}
		// remove currentItem
		currentSolution = min
		delete(unvisited, currentItem)
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
		possibleNewTableaux := []string{}
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

						if nt == "\n\n\nE,HG,HM,LG,LM" {
							debugprint(currentSolution, nt)
							return currentSolution + 1
						}
						possibleNewTableaux = append(possibleNewTableaux, nt)
					}
				}

			}
		}
		for _, newItem := range possibleNewTableaux {
			if unvisited[newItem] == 0 {
				if prev[newItem] == "" {
					unvisited[newItem] = currentSolution + 1
				}
				prev[newItem] = currentItem
				continue
			}
			if unvisited[newItem] > 0 {
				if currentSolution+1 < unvisited[newItem] {
					if prev[newItem] == "" {
						unvisited[newItem] = currentSolution + 1
					}
					prev[newItem] = currentItem
				}
			}

		}

	}

	return currentSolution
}

func debugprint(score int, outd string) {
	fmt.Println(score)
	out := strings.Split(outd, "\n")
	Reverse(out)
	for _, v := range out {
		fmt.Println(v)
	}
}

func possibleFloors(currFloor int) []int {
	if currFloor == 0 {
		return []int{1}
	}
	if currFloor == 2 {
		return []int{1, 3}
	}
	if currFloor == 1 {
		return []int{0, 2}
	}
	// 3
	return []int{2}

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

func makeTableau(data []string) string {
	out := []string{}
	for _, row := range data {
		sp := sortFloor(row)
		out = append(out, sp)
	}
	return strings.Join(out, "\n")
}

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
	for _, m := range t {
		if len(m) < 2 {
			continue
		}
		for _, g := range t {
			if len(g) < 2 {
				continue
			}
			if string(m[1]) == "M" && string(g[1]) == "G" && string(m[0]) != string(g[0]) {
				return false
			}

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
