package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	j := sumSeverity(string(input))
	return j
}

// 156372 too low
// 636852 too low
func part2() int {
	input, _ := os.ReadFile("input.txt")
	j := calcDelay2(string(input))
	return j
}

type layer struct {
	depth            int
	rge              int
	scannerLoc       int
	scannerIncrement int
}

func (j *layer) inc() {
	loc := j.scannerLoc
	if loc == 0 {
		j.scannerIncrement = 1
	}
	if loc == j.rge-1 {
		j.scannerIncrement = -1
	}
	j.scannerLoc += j.scannerIncrement

}

func calcDelay2(s string) int {
	data := strings.Split(s, "\n")
	layers := map[int]*layer{}
	maxLayer := 0
	for _, v := range data {
		j := strings.Replace(v, ":", " :#", 1)
		d := strings.Fields(j)
		a, _ := strconv.Atoi(d[0])
		b, _ := strconv.Atoi(d[2])
		l := layer{depth: a, rge: b, scannerLoc: 0, scannerIncrement: 1}
		if a > maxLayer {
			maxLayer = a
		}
		layers[a] = &l
	}
	delay := 0

	for d := 0; d < delay; d++ {
		for _, v := range layers {
			v.inc()
		}
	}
	delayedState := layers
	severity := 1
	for severity > 0 {
		severity = 0
		// lastDelay := []
		delay += 1
		if delay%25000 == 0 {
			fmt.Println(delay)
		}
		for _, v := range delayedState {
			v.inc()
		}
		workingState := map[int]*layer{}
		for i, k := range delayedState {
			newLayer := layer{depth: k.depth, rge: k.rge, scannerLoc: k.scannerLoc, scannerIncrement: k.scannerIncrement}
			workingState[i] = &newLayer
		}
		for jump := 0; jump <= maxLayer; jump++ {
			found := false
			for x := range workingState {
				if jump == x {
					found = true
					break
				}
			}
			if found {
				layr := workingState[jump]
				sLoc := layr.scannerLoc
				if sLoc == 0 {
					severity += layr.depth * layr.rge
					if severity > 0 {
						break
					}
				}
			}
			for _, v := range workingState {
				v.inc()
			}
		}
		if severity == 0 {
			for _, v := range workingState {
				fmt.Println(v.depth, v.rge, v.scannerLoc, v.scannerIncrement)
			}
		}

	}

	return delay

}

func printLayers(layers map[int]*layer) {
	max := 0
	for _, v := range layers {
		if v.depth > max {
			max = v.depth
		}
	}
	for i := 0; i < max; i++ {
		fmt.Print(" ", i, " ")

	}
	fmt.Println()

}

func calcDelay(s string) int {
	data := strings.Split(s, "\n")
	severity := math.MaxInt
	layers := map[int]*layer{}
	maxLayer := 0
	for _, v := range data {
		j := strings.Replace(v, ":", " :#", 1)
		d := strings.Fields(j)
		a, _ := strconv.Atoi(d[0])
		b, _ := strconv.Atoi(d[2])
		l := layer{depth: a, rge: b, scannerLoc: 0, scannerIncrement: 1}
		if a > maxLayer {
			maxLayer = a
		}
		layers[a] = &l
	}
	delay := 0
	for severity > 0 {
		severity = 0
		layer := 0
		for layer <= maxLayer {
			found := false
			for x := range layers {
				if layer == x {
					found = true
					break
				}
			}
			if found {
				layr := layers[layer]
				sLoc := layr.scannerLoc
				if sLoc == 0 {
					severity += layr.depth * layr.rge
				}
			}
			layer += 1
		}
		for _, v := range layers {
			v.inc()
		}
		delay += 1
	}
	return delay - 1
}

func scannerLocation(time int, rge int) int {
	n := 0
	loc := 0
	inc := -1
	for n < time {
		if loc == rge-1 || loc == 0 {
			inc = -inc
		}
		loc += inc
		n += 1
	}
	return loc
}

func sumSeverity(s string) int {
	data := strings.Split(s, "\n")
	severity := 0
	layers := map[int]layer{}
	maxLayer := 0
	for _, v := range data {
		j := strings.Replace(v, ":", " :#", 1)
		d := strings.Fields(j)
		a, _ := strconv.Atoi(d[0])
		b, _ := strconv.Atoi(d[2])
		l := layer{depth: a, rge: b}
		if a > maxLayer {
			maxLayer = a
		}
		layers[a] = l
	}

	move := -1
	for move <= maxLayer {
		move += 1
		dep := layers[move]

		if dep.scannerLocation(move) == 0 {
			severity += dep.depth * dep.rge
		}
	}

	return severity
}

func (j layer) scannerLocation(time int) int {
	n := 0
	loc := 0
	inc := -1
	for n < time {
		if loc == j.rge-1 || loc == 0 {
			inc = -inc
		}
		loc += inc
		n += 1
	}
	return loc
}
