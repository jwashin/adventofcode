package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", surfaceArea(string(input)))
	fmt.Println("Part 2:", externalSurfaceArea(string(input)))

}

type cube struct {
	x int
	y int
	z int
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func (a cube) isAdjacent(b cube) bool {
	if a.isEqual(b) {
		return false
	}
	if abs(a.x-b.x) == 1 && !(abs(a.y-b.y) >= 1) && !(abs(a.z-b.z) >= 1) {

		return true
	}
	if abs(a.y-b.y) == 1 && !(abs(a.x-b.x) >= 1) && !(abs(a.z-b.z) >= 1) {
		return true
	}
	if abs(a.z-b.z) == 1 && !(abs(a.x-b.x) >= 1) && !(abs(a.y-b.y) >= 1) {
		return true
	}
	return false
}

func (a cube) isEqual(b cube) bool {
	return a.x == b.x && a.y == b.y && a.z == b.z
}

func (a cube) xyNeighbors() cubeList {
	return cubeList{
		cube{a.x + 1, a.y, a.z},
		cube{a.x - 1, a.y, a.z},
		cube{a.x, a.y - 1, a.z},
		cube{a.x, a.y + 1, a.z},
	}
}
func (a cube) xyzNeighbors() cubeList {
	return cubeList{
		cube{a.x + 1, a.y, a.z},
		cube{a.x - 1, a.y, a.z},
		cube{a.x, a.y - 1, a.z},
		cube{a.x, a.y + 1, a.z},
		cube{a.x + 1, a.y, a.z - 1},
		cube{a.x - 1, a.y, a.z + 1},
		cube{a.x, a.y - 1, a.z - 1},
		cube{a.x, a.y + 1, a.z + 1},
	}
}

func surfaceArea(s string) int {
	data := strings.Split(s, "\n")
	cubes := []cube{}
	for _, v := range data {
		v = strings.TrimSpace(v)
		f := strings.Split(v, ",")
		x, _ := strconv.Atoi(f[0])
		y, _ := strconv.Atoi(f[1])
		z, _ := strconv.Atoi(f[2])
		t := cube{x, y, z}
		cubes = append(cubes, t)
	}
	surfaceCount := 0
	for _, v := range cubes {
		adjacentCubes := 0
		for _, w := range cubes {
			if v.isAdjacent(w) {
				adjacentCubes += 1
			}
		}
		surfaceCount += 6 - adjacentCubes
	}
	return surfaceCount
}

// set implementation
// https://linuxhint.com/golang-set/
type void struct{}

var member void

// func (c chamber) contains(m coord) bool {
// 	if _, ok := c.fills[m]; ok {
// 		return true
// 	} else {
// 		return false
// 	}
// }

type cubeList []cube

type cubeSet map[cube]void

func (a cubeSet) exists(c cube) bool {
	if _, ok := a[c]; ok {
		return true
	}
	return false
}

// func (a cubeList) fillInterior() cubeList {

// 	out := a
// 	t := a[0]
// 	set := map[cube]void{}
// 	xmax, xmin, ymax, ymin, zmax, zmin := t.x, t.x, t.y, t.y, t.z, t.z
// 	for _, v := range a {
// 		set[v] = member
// 		if v.x < xmin {
// 			xmin = v.x
// 		}
// 		if v.y < ymin {
// 			ymin = v.y
// 		}
// 		if v.z < zmin {
// 			zmin = v.z
// 		}
// 		if v.x > xmax {
// 			xmax = v.x
// 		}
// 		if v.y > ymax {
// 			ymax = v.y
// 		}
// 		if v.z > zmax {
// 			zmax = v.z
// 		}
// 	}
// 	for x := xmin; x <= xmax; x++ {
// 		for y := ymin; y <= ymax; y++ {
// 			line := cubeList{}
// 			for z := zmin; z < zmax; z++ {
// 				c := cube{x, y, z}
// 				if _, ok := set[c]; ok {
// 					line = append(line, c)
// 				}
// 			}
// 			out = append(out, a.getInsideCubes(line)...)
// 		}
// 	}
// 	return out
// }

func (a cubeList) exists(c cube) bool {
	for _, v := range a {
		if v.isEqual(c) {
			return true
		}
	}
	return false
}

// func (a cubeList) getInsideCubes(line cubeList) cubeList {
// 	f := cubeList{}
// 	if len(line) < 2 {
// 		return f
// 	}
// 	inside := true
// 	max := line[0].z
// 	min := line[0].z
// 	for _, v := range line {
// 		if v.z > max {
// 			max = v.z
// 		}
// 		if v.z < min {
// 			min = v.z
// 		}
// 	}
// 	for z := min + 1; z < max; z++ {
// 		c := cube{line[0].x, line[0].y, z}
// 		if line.exists(c) {
// 			// && !line.exists(cube{line[0].x,line[0].y,z}){
// 			// check next item
// 			inside = !inside
// 			continue
// 		}
// 		if inside {
// 			f = append(f, c)
// 		}
// 	}
// 	return f
// }

// 2406 too low

func externalSurfaceArea(s string) int {
	data := strings.Split(s, "\n")
	cubes := cubeSet{}
	for _, v := range data {
		v = strings.TrimSpace(v)
		f := strings.Split(v, ",")
		x, _ := strconv.Atoi(f[0])
		y, _ := strconv.Atoi(f[1])
		z, _ := strconv.Atoi(f[2])
		t := cube{x, y, z}
		cubes[t] = member
	}
	return p2(cubes)
}

func p2(cubes cubeSet) int {
	var xmin int
	var xmax int
	var ymin int
	var ymax int
	var zmin int
	var zmax int

	for v := range cubes {
		if v.x < xmin {
			xmin = v.x
		}
		if v.x > xmax {
			xmax = v.x
		}
		if v.y < ymin {
			ymin = v.y
		}
		if v.y > ymax {
			ymax = v.y
		}
		if v.z < zmin {
			zmin = v.z
		}
		if v.z > zmax {
			zmax = v.z
		}
	}
	zmax = zmax + 1
	xmax = xmax + 1
	ymax = ymax + 1
	// zmin = zmin - 1
	// xmin = xmin - 1
	// ymin = ymin - 1

	// floodfill everything outside!
	outside := cubeSet{}
	currCube := cube{xmin, ymin, zmin}
	q := []cube{currCube}
	for len(q) > 0 {
		currCube = q[0]
		q = q[1:]
		if cubes.exists(currCube) {
			// ooh! this cube is on the outside boundary!
			continue
		}
		outside[currCube] = member
		for _, v := range currCube.xyzNeighbors() {
			if v.x >= xmin && v.x <= xmax && v.y >= ymin && v.y <= ymax && v.z >= zmin && v.z <= zmax {
				if !outside.exists(v) {
					q = append(q, v)
				}
			}
		}
	}
	newCubes := cubeSet{}
	for x := xmin; x < xmax; x++ {
		for y := ymin; y < ymax; y++ {
			for z := zmin; z < zmax; z++ {
				c := cube{x, y, z}
				if !outside.exists(c) {
					newCubes[c] = member
				}
			}
		}
	}
	surfaceCount := 0
	for v := range newCubes {
		adjacentCubes := 0
		for w := range newCubes {
			if v.isAdjacent(w) {
				adjacentCubes += 1
			}
		}
		surfaceCount += 6 - adjacentCubes
	}
	return surfaceCount
}

func floodFill3d(cubes map[cube]void) map[cube]void {
	f := map[cube]void{}
	minz := math.MaxInt
	maxz := 0
	for v := range cubes {
		f[v] = member
		if v.z < minz {
			minz = v.z
		}
		if v.z > maxz {
			maxz = v.z
		}
	}
	for z := minz; z <= maxz; z++ {
		// Everything at a particular z
		// should be a point or a loop of some sort.
		// We flood fill that plane.
		g := map[cube]void{}
		var xmin int
		var xmax int
		var ymin int
		var ymax int
		for v := range f {
			if v.z == z {
				g[v] = member
				if v.x < xmin {
					xmin = v.x
				}
				if v.x > xmax {
					xmax = v.x
				}
				if v.y < ymin {
					ymin = v.y
				}
				if v.y > ymax {
					ymax = v.y
				}
			}
		}
		loop := []string{}
		for y := ymin; y <= ymax; y++ {
			s := ""
			for x := xmin; x <= xmax; x++ {
				if _, ok := g[cube{x, y, z}]; ok {
					s += "#"
				}
				s += "."
			}
			loop = append(loop, s)
		}
		for _, v := range loop {
			fmt.Println(v)
		}
		fmt.Println()
		if len(g) < 2 {
			continue
		}
		// pick a spot inside the circle
		// easy pick in center. this wouldn't be harder?
		count := 0
		c := cube{xmin + (xmax-xmin)/2, ymin + (ymax-ymin)/2, z}
		for x := xmin; x < xmax; x++ {
			for y := ymin; y < ymax; y++ {
				if _, ok := f[cube{x, y, z}]; ok {
					count += 1
				}
				if count == 1 {
					c = cube{x, y + 1, z}
				}
			}
		}

		q := []cube{c}
		for len(q) > 0 {
			currCube := q[0]
			q = q[1:]
			if _, ok := g[currCube]; ok {
				continue
			}
			g[currCube] = member
			for _, v := range currCube.xyNeighbors() {
				if _, ok := g[v]; ok {
					continue
				}
				q = append(q, v)
			}
			for v := range g {
				f[v] = member
			}
		}
	}
	return f
}
