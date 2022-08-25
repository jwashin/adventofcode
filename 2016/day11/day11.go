package main

import (
	"math"
	"math/bits"
	"sort"
	"strings"
)

func getData(test bool) []string {
	if test {
		return []string{"E,HM,LM\nHG\nLG\n"}
	}
	return []string{"PG,TG,TM,XG,RG,RM,CG,CM\nPM,XM\n\n"}
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
	s := strings.Split(f, ",")
	for _, v := range r {
		if !contains(s, v) {
			s = append(s, v)
		}
	}
	return strings.Join(s, ",")
}

func doGame(test bool) {
	data := getData(true)
	s := makeTableau(data)
	unvisited := map[string]int{s: 0}
	prev := map[string]string{}
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
		currentSolution := min
		delete(unvisited, currentItem)

		// item is a string. enumerate next possible moves
		floors := fromTableau(currentItem)
		floorIdx := currentFloor(floors)
		floorItems := strings.Split(floors[floorIdx], ",")
		toAdd := [][]string{}
		for _, v := range floorItems {
			if v != "E" {
				toAdd = append(toAdd, []string{v})
			}
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
								newTableau = append(newTableau, "E,"+floorx)
								continue
							}
							if kx == floorIdx {
								newTableau = append(newTableau, floorT)
								continue
							}
							newTableau = append(newTableau, v)
						}
						nt := normalize(strings.Join(newTableau, "\n"))
						if nt == "E,HG,HM,LG,LM\n\n\n"
						possibleNewTableaux = append(possibleNewTableaux, nt)
					}
				}

			}
		}
		for _, item := range possibleNewTableaux {
			if unvisited[item] == 0 {
				unvisited[item] = currentSolution + 1
				prev[item] = currentItem
				continue
			}
			if unvisited[item] > 0 {
				if currentSolution+1 < unvisited[item] {
					unvisited[item] = currentSolution + 1
					prev[item] = currentItem
				}
			}

		}

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
	return 0
}

func sortFloor(s string) string {
	t := strings.Split(s, ",")
	sort.Strings(t)
	return strings.Join(t, ",")
}

func makeTableau(data []string) string {
	out := []string{}
	for _, row := range data {
		sp := strings.Split(row, ",")
		sort.Slice(sp, func(i, j int) bool {
			return sp[i] < sp[j]
		})
		out = append(out, strings.Join(sp, ","))
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
