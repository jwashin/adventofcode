package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Command struct {
	instruction                        string
	minx, maxx, miny, maxy, minz, maxz int
}

func (c Command) cuboid() Cuboid {
	istr := false
	if c.instruction == "on" {
		istr = true
	}
	return Cuboid{c.minx, c.maxx, c.miny, c.maxy, c.minz, c.maxz, istr}
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	fmt.Println(doCommands2(string(data)))
}

func parseInput(aString string) []Command {
	commands := []Command{}
	lines := strings.Split(aString, "\n")
	for _, line := range lines {
		t := strings.TrimSpace(line)
		if len(t) > 0 {
			var instruction string
			var minx, miny, minz, maxx, maxy, maxz int
			fmt.Sscanf(t, "%s x=%d..%d,y=%d..%d,z=%d..%d", &instruction, &minx, &maxx, &miny, &maxy, &minz, &maxz)
			if maxx < minx {
				maxx, minx = minx, maxx
			}
			if maxy < miny {
				maxy, miny = miny, maxy
			}
			if maxz < minz {
				maxz, minz = minz, maxz
			}
			c := Command{instruction, minx, maxx, miny, maxy, minz, maxz}
			commands = append(commands, c)
		}
	}
	return commands
}

type Coord struct {
	x, y, z int
}

type Cuboid struct {
	minx, maxx, miny, maxy, minz, maxz int
	isOn                               bool
}

func (c Cuboid) count() int {
	tot := (c.maxx - c.minx + 1) * (c.maxy - c.miny + 1) * (c.maxz - c.minz + 1)
	if !c.isOn {
		tot = -tot
	}
	return tot
}

func (c Cuboid) corners() []Coord {
	return []Coord{
		{c.minx, c.miny, c.minz},
		{c.maxx, c.maxy, c.minz},
		{c.minx, c.maxy, c.minz},
		{c.maxx, c.miny, c.minz},

		{c.minx, c.miny, c.maxz},
		{c.maxx, c.maxy, c.maxz},
		{c.minx, c.maxy, c.maxz},
		{c.maxx, c.miny, c.maxz},
	}
}

func (a Cuboid) intersects(b Cuboid) bool {
	for _, g := range b.corners() {
		if a.containsPoint(g) {
			return true
		}
	}
	return false
}

func (a Cuboid) intersection(b Cuboid) Cuboid {
	newOn := b.isOn
	if a.isOn && b.isOn {
		newOn = false
	} else if !a.isOn && !b.isOn {
		newOn = true
	}
	minx, maxx := mids(a.maxx, b.maxx, a.minx, b.minx)
	miny, maxy := mids(a.maxy, b.maxy, a.miny, b.miny)
	minz, maxz := mids(a.maxz, b.maxz, a.minz, b.minz)
	return Cuboid{minx, maxx, miny, maxy, minz, maxz, newOn}
}

func mids(a int, b int, c int, d int) (int, int) {
	t := []int{a, b, c, d}
	sort.Ints(t)
	return t[1], t[2]
}

func (c Cuboid) containsPoint(loc Coord) bool {
	if loc.x >= c.minx && loc.x <= c.maxx &&
		loc.y >= c.miny && loc.y <= c.maxy &&
		loc.z >= c.minz && loc.z <= c.maxz {
		return true
	}
	return false
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://github.com/alexchao26/advent-of-code-go/blob/main/2021/day22/main.go
func (c Cuboid) getIntersection(c2 Cuboid) (intersection Cuboid, hasIntersection bool) {
	// larger of x1s has to be smaller than smaller of x2s for there to be an overlap
	x1 := maxInt(c.minx, c2.minx)
	x2 := minInt(c.maxx, c2.maxx)
	y1 := maxInt(c.miny, c2.miny)
	y2 := minInt(c.maxy, c2.maxy)
	z1 := maxInt(c.minz, c2.minz)
	z2 := minInt(c.maxz, c2.maxz)

	if x1 > x2 || y1 > y2 || z1 > z2 {
		return Cuboid{}, false
	}

	var intersectionState bool
	if c.isOn && c2.isOn {
		intersectionState = false
	} else if !c.isOn && !c2.isOn {
		intersectionState = true
	} else {
		// ! default to second Cuboid's on/off state. This makes the order of which Cuboid's method is
		// called very important. but that's what unit tests are for :)
		// alternatively the caller could deal with it.. that might be more clear...
		intersectionState = c2.isOn
	}

	return Cuboid{
		isOn: intersectionState,
		minx: x1, maxx: x2,
		miny: y1, maxy: y2,
		minz: z1, maxz: z2,
	}, true
}

// https://github.com/alexchao26/advent-of-code-go/blob/main/2021/day22/main.go
func doCommands2(aString string) int {
	// return the number of 'on' cubes
	commands := parseInput(aString)
	core := []Cuboid{}
	for _, command := range commands {
		// newCubes := []Cuboid{}
		addMe := []Cuboid{}
		cmdcube := command.cuboid()
		for _, corecube := range core {
			intersection, didIntersect := corecube.getIntersection(cmdcube)
			if didIntersect {
				// fmt.Println("we have intersection")
				addMe = append(addMe, intersection)
			}
		}
		if cmdcube.isOn {
			addMe = append(addMe, cmdcube)
		}
		core = append(core, addMe...)

	}
	value := 0
	for _, v := range core {
		value += v.count()
	}
	return value
}

func doCommands2Mine(aString string) int {
	// return the number of 'on' cubes
	commands := parseInput(aString)
	core := []Cuboid{}
	for _, command := range commands {
		// newCubes := []Cuboid{}
		addMe := []Cuboid{}
		cmdcube := command.cuboid()
		for _, corecube := range core {
			if corecube.intersects(cmdcube) {
				// fmt.Println("we have intersection")
				intersection := corecube.intersection(command.cuboid())
				addMe = append(addMe, intersection)
			}
		}
		if cmdcube.isOn {
			addMe = append(addMe, cmdcube)
		}
		core = append(core, addMe...)

	}
	value := 0
	for _, v := range core {
		value += v.count()
	}
	return value
}

func doCommands(aString string) int {
	commands := parseInput(aString)
	core := map[string]int{}
	inclusion := Command{"include", -50, 50, -50, 50, -50, 50}
	for _, c := range commands {
		for x := c.minx; x <= c.maxx; x++ {
			if x <= inclusion.maxx && x >= inclusion.minx {
				for y := c.miny; y <= c.maxy; y++ {
					if y <= inclusion.maxy && y >= inclusion.miny {
						for z := c.minz; z <= c.maxz; z++ {
							if z <= inclusion.maxz && z >= inclusion.minz {
								coord := fmt.Sprintf("%d,%d,%d", x, y, z)
								if c.instruction == "on" {
									core[coord] = 1
								} else {
									delete(core, coord)
								}
							}
						}
					}
				}
			}
		}
	}
	return len(core)
}
