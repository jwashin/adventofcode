package main

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type valve struct {
	name         string
	open         bool
	flowRate     int
	destinations []string
}

type valveMap map[string]valve

// the order in which to open the valves is the real question.

func (m valveMap) bestPath(a string, b string) ([]string, error) {
	currLoc := m[a]
	q := [][]string{{currLoc.name}}
	used := map[string]int{}
	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			return len(q[i]) < len(q[j])
		})
		currPath := q[0]
		q = q[1:]
		currId := currPath[len(currPath)-1]
		if currId == b {
			return currPath, nil
		}
		for _, v := range m[currId].destinations {
			if used[v] == 0 {
				used[v] = 1
				newPath := currPath
				newPath = append(newPath, v)
				q = append(q, newPath)
			}
		}
	}
	return []string{}, errors.New("no path")
}

func getValves(s []string) []valve {
	valves := []valve{}
	for _, v := range s {
		v = strings.ReplaceAll(v, "=", " ")
		v = strings.ReplaceAll(v, ";", " ")
		v = strings.ReplaceAll(v, ", ", ":")
		f := strings.Fields(v)
		if len(f) > 0 {
			dests := strings.Split(f[10], ":")
			rate, _ := strconv.Atoi(f[5])
			t := valve{name: f[1], open: false, flowRate: rate, destinations: dests}
			valves = append(valves, t)
		}
	}
	return valves
}

func part1(s string) int {
	input := strings.Split(s, "\n")
	valves := getValves(input)
	valveData := valveMap{}
	for _, v := range valves {
		valveData[v.name] = v
	}
	paths := map[string][]string{}
	travelCost := map[string]int{}
	for k, v := range valves {
		if v.flowRate > 0 || k == 0 {
			for _, w := range valves {
				if v.name != w.name && w.flowRate > 0 {
					path, err := valveData.bestPath(v.name, w.name)
					if err == nil {
						paths[v.name+" "+w.name] = path
						// paths[v.name + " " + w.name] = path
						travelCost[v.name+" "+w.name] = len(path)
						// travelCost[stringPair{w.name, v.name}] = len(path)
					}
				}
			}
		}
	}
	toVisit := []string{}
	for _, v := range valves {
		if v.flowRate > 0 {
			toVisit = append(toVisit, v.name)
		}
	}
	timeLimit := 30
	currLoc := valves[0].name
	openValves := []string{}
	t := 0
	visited := map[string]int{}
	visited[currLoc] = 1
	totalPressureReleased := 0
	moves := []string{}

	for t <= timeLimit {
		t += 1
		for _, v := range openValves {
			totalPressureReleased += valveData[v].flowRate
		}
		for len(moves) > 0 {
			move := moves[0]
			moves = moves[1:]
			if move == "open" {
				openValves = append(openValves, currLoc)
			} else {
				currLoc = move
			}
		}
		newLoc := ""
		timeFactor := 1
		for newLoc == "" && len(toVisit) > 0 {
			timeFactor += 1
			max := 0
			for _, v := range toVisit {
				ky := currLoc + " " + v
				tp := paths[ky]
				tc := len(tp)
				// travelCost[stringPair{currLoc, v}]
				if tc <= timeFactor {
					tr := timeLimit - t
					pTot := (tr - tc - 2) * valveData[v].flowRate
					if pTot > max && tc <= timeFactor {
						max = pTot
						newLoc = v
						timeFactor = 1
					}
				}
			}
		}
		toVisitNew := []string{}
		for _, v := range toVisit {
			if v != newLoc {
				toVisitNew = append(toVisitNew, v)
			}
		}
		toVisit = toVisitNew
		moves = paths[currLoc+" "+newLoc]
		moves = append(moves, "open")
	}
	return totalPressureReleased

}
