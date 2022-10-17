package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	fmt.Println("1.", part1(data))
}

type coordinate struct {
	x, y int
}

func (c coordinate) neighbors() coordinateList {
	return coordinateList{{c.x + 1, c.y}, {c.x - 1, c.y}, {c.x, c.y - 1}, {c.x, c.y + 1}}
}

func (c coordinate) dNeighbors() map[string]coordinate {
	return map[string]coordinate{
		"e": {c.x + 1, c.y},
		"w": {c.x - 1, c.y},
		"n": {c.x, c.y - 1},
		"s": {c.x, c.y + 1}}
}

func (c coordinate) follow(s string) coordinate {
	decode := map[string]coordinate{
		"e": {c.x + 1, c.y},
		"w": {c.x - 1, c.y},
		"n": {c.x, c.y - 1},
		"s": {c.x, c.y + 1}}
	return decode[s]
}

func (c coordinate) pathResult(s string) coordinate {
	new := c
	for _, v := range s {
		p := string(v)
		new = new.follow(p)
	}
	return new
}

type coordinateList []coordinate

func (c coordinateList) Sort() {
	sort.Slice(c, func(i, j int) bool {
		if c[i].y == c[j].y {
			return c[i].x < c[j].x
		}
		return c[i].y < c[j].y
	})
}

func (c coordinateList) contains(loc coordinate) bool {
	for _, v := range c {
		if v.x == loc.x && v.y == loc.y {
			return true
		}
	}
	return false
}

type Unit struct {
	disp      string
	hitPoints int
}

type tableau map[coordinate]*Unit

func (t tableau) clear(c coordinate) *Unit {
	s := t[c]
	tmp := Unit{disp: ".", hitPoints: 0}
	t[c] = &tmp
	return s
}

func (t tableau) move(unit *Unit, from coordinate, to coordinate) {
	t[to] = unit
	t.clear(from)
}

func (t tableau) unitLocs() coordinateList {
	n := coordinateList{}
	for k, v := range t {
		if v.hitPoints > 0 {
			n = append(n, k)
		}
	}
	n.Sort()
	return n
}

func (t tableau) listEnemies(loc coordinate) coordinateList {
	enemies := coordinateList{}
	for _, e := range t.unitLocs() {
		if t[e].disp != t[loc].disp && t[loc].hitPoints > 0 {
			enemies = append(enemies, e)
		}
	}
	return enemies
}

func (t tableau) attack(loc coordinate) bool {
	inRange := loc.neighbors()
	attackList := coordinateList{}
	for _, v := range t.listEnemies(loc) {
		if inRange.contains(v) {
			attackList = append(attackList, v)
		}
	}
	if len(attackList) == 0 {
		return false
	}
	if len(attackList) == 1 {
		v := attackList[0]
		t[v].hitPoints -= 3
		if t[v].hitPoints <= 0 {
			t[v].hitPoints = 0
			t[v].disp = "."
		}
		return true
	}
	min := t[attackList[0]].hitPoints
	for _, v := range attackList {
		if t[v].hitPoints < min {
			min = t[v].hitPoints
		}
	}
	hitList := coordinateList{}
	for _, v := range attackList {
		if t[v].hitPoints == min {
			hitList = append(hitList, v)
		}
	}
	// sort, then attack the first one
	hitList.Sort()
	v := hitList[0]
	t[v].hitPoints -= 3
	if t[v].hitPoints <= 0 {
		t[v].disp = "."
		t[v].hitPoints = 0
	}
	return true
}

func (t tableau) contains(c coordinate) bool {
	for k := range t {
		if k == c {
			return true
		}
	}
	return false
}

func (t tableau) path(c coordinate, b coordinate) coordinateList {
	start := c
	prevs := map[coordinate]coordinate{}
	frontier := coordinateList{start}
	expanded := coordinateList{}
	for len(frontier) > 0 {
		a := frontier[0]
		frontier = frontier[1:]
		if a.x == b.x && a.y == b.y {
			path := coordinateList{b}
			for path[len(path)-1] != start {
				path = append(path, prevs[path[len(path)-1]])
			}
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path
		}
		expanded = append(expanded, a)
		for _, v := range a.neighbors() {
			if t.contains(v) {
				cv := t[v].disp
				if cv == "." && !frontier.contains(v) && !expanded.contains(v) {
					frontier = append(frontier, v)
					prevs[v] = a
				}
			}
		}
	}
	return coordinateList{}
}

// 	for len(q) > 0 {
// 		sort.Slice(q, func(i, j int) bool {
// 			return len(q[i]) < len(q[j])
// 		})
// 		currentPath := q[0]
// 		q = q[1:]
// 		loc := a.pathResult(currentPath)
// 		used = append(used, loc)
// 		if loc == b {
// 			return currentPath
// 		}
// 		for k, v := range loc.dNeighbors() {
// 			if t[v] != nil {
// 				if !used.contains(v) {
// 					cv := t[v].disp
// 					if cv != "E" && cv != "G" && cv != "#" {
// 						found := false
// 						testPath := currentPath + k
// 						for _, v := range q {
// 							if v == testPath {
// 								found = true
// 								break
// 							}
// 						}
// 						if !found {
// 							q = append(q, currentPath+k)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }
// return ""
// }

func (t tableau) totalHitPoints() int {
	count := 0
	for _, v := range t {
		if v.hitPoints > 0 {
			count += v.hitPoints
		}
	}
	return count
}

func (t tableau) String() string {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for k := range t {
		if k.x < minX {
			minX = k.x
		}
		if k.y < minY {
			minY = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	out := []string{}
	for y := minY; y <= maxY; y++ {
		s := ""
		for x := minX; x <= maxX; x++ {
			s += t[coordinate{x, y}].disp
		}
		out = append(out, s)
	}
	return strings.Join(out, "\n")
}

// 250410 too high

func part1(s string) int {
	t := tableau{}
	for y, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		for x, c := range line {
			d := string(c)
			j := Unit{disp: d}
			coord := coordinate{x: x, y: y}
			if d == "G" || d == "E" {
				j.hitPoints = 200
			}
			t[coord] = &j
		}
	}
	round := 0

	for {
		fmt.Println(t)
		fmt.Println(round, t.totalHitPoints())
		unitlocs := t.unitLocs()
		for _, currentLoc := range unitlocs {
			unit := t[currentLoc]
			if unit.hitPoints <= 0 {
				continue
			}
			// unitType := unit.disp
			enemies := t.listEnemies(currentLoc)
			if len(enemies) == 0 {
				return round * t.totalHitPoints()
			}
			done := t.attack(currentLoc)
			if !done {
				enemies = t.listEnemies(currentLoc)
				// move, maybe. Collect in-range targets
				inRangeTargets := coordinateList{}
				for _, v := range enemies {
					for _, z := range v.neighbors() {
						if t.contains(z) {
							item := t[z].disp
							if item == "." && !inRangeTargets.contains(z) {
								inRangeTargets = append(inRangeTargets, z)
							}
						}
					}
				}
				// min path to in-range targets if possible
				paths := map[coordinate]coordinateList{}
				for _, target := range inRangeTargets {
					c := t.path(currentLoc, target)
					if len(c) > 0 {
						paths[target] = c
					}
				}

				min := math.MaxInt
				for _, v := range paths {
					if len(v) < min {
						min = len(v)
					}
				}

				moves := []coordinateList{}
				for _, v := range paths {
					if len(v) == min {
						moves = append(moves, v)
					}
				}
				if len(moves) == 0 {
					// moves.Sort()
					newlocation := moves[0][1]
					t.move(t[currentLoc], currentLoc, newlocation)
					t.attack(newlocation)
				}
				// moves
			}
		}
		round += 1
	}

}
