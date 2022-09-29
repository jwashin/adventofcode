package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

type coordinate struct {
	x, y, z int
}

type particle struct {
	position, velocity, acceleration coordinate
	id                               int
}

func (p *particle) inc() {
	p.velocity.x += p.acceleration.x
	p.velocity.y += p.acceleration.y
	p.velocity.z += p.acceleration.z
	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
	p.position.z += p.velocity.z
}

func (p particle) distance() int {
	return abs(p.position.x) + abs(p.position.y) + abs(p.position.z)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 1000 too high
// 571 too high

// 505 too high

func part2() int {
	particlea := getData()
	// oldlen := 0
	timeSinceCollision := 0
	particles := map[int]*particle{}
	for _, v := range particlea {
		particles[v.id] = v
	}
	for timeSinceCollision < 100 {
		// oldlen = len(particles)
		timeSinceCollision += 1
		deleteSet := map[string]string{}
		for i, v := range particles {
			for j, test := range particles {
				if i != j {
					if v.position.x == test.position.x && v.position.y == test.position.y && v.position.z == test.position.z {
						timeSinceCollision = 0
						deleteSet[fmt.Sprintf("%d,%d,%d", v.position.x, v.position.y, v.position.z)] += fmt.Sprintf(" %d %d", v.id, test.id)
					}
				}
			}
		}
		for _, v := range deleteSet {
			if len(v) > 0 {
				d := strings.Fields(v)
				for _, x := range d {
					idx, _ := strconv.Atoi(x)
					delete(particles, idx)
				}
			}
		}
		for _, v := range particles {
			v.inc()
		}
	}
	return len(particles)
}

func part1() int {
	particles := getData()
	for i := 0; i < 10000; i++ {
		for _, v := range particles {
			v.inc()
		}
	}
	sort.Slice(particles, func(i, j int) bool {
		return particles[i].distance() < particles[j].distance()
	})
	return particles[0].id
}

func getData() []*particle {
	input, _ := os.ReadFile("input.txt")
	p1 := strings.Split(string(input), "\n")
	z := []*particle{}
	for k, s := range p1 {
		par := particle{id: k}
		h := strings.Split(s, ", ")
		for i, v := range h {
			d := v[strings.Index(v, "<")+1 : strings.Index(v, ">")]
			f := strings.Split(d, ",")
			x, _ := strconv.Atoi(f[0])
			y, _ := strconv.Atoi(f[1])
			z, _ := strconv.Atoi(f[2])
			if i == 0 {
				// position
				par.position = coordinate{x, y, z}
			}
			if i == 1 {
				// velocity
				par.velocity = coordinate{x, y, z}
			}
			if i == 2 {
				// acceleration
				par.acceleration = coordinate{x, y, z}
			}
		}
		z = append(z, &par)
	}
	return z
}
