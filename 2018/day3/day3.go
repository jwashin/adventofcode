package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

type coordinate struct {
	x, y int
}

type region struct {
	ul coordinate
	br coordinate
}

func (r region) contains(c coordinate) bool {
	if c.x >= r.ul.x && c.x <= r.br.x {
		if c.y >= r.ul.y && c.y <= r.br.y {
			return true
		}
	}
	return false
}
func (r region) intersection(g region) []string {
	out := []string{}
	for x := r.ul.x; x <= r.br.x; x++ {
		for y := r.ul.y; y <= r.br.y; y++ {
			c := coordinate{x, y}
			if g.contains(c) {
				out = append(out, c.String())
			}
		}
	}
	return out
}

func (c coordinate) String() string {
	return fmt.Sprintf("%d %d", c.x, c.y)
}

func (a claim) makeExtents() region {
	ul := coordinate{a.x, a.y}
	br := coordinate{a.x + a.width - 1, a.y + a.height - 1}
	return region{ul, br}
}

func (a claim) intersection(b claim) []string {
	regiona := a.makeExtents()
	regionb := b.makeExtents()
	return regiona.intersection(regionb)
}

func makeClaim(s string) claim {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, "#", "")
	s = strings.ReplaceAll(s, "@", "")
	t := strings.Fields(s)
	id, _ := strconv.Atoi(t[0])
	c := strings.Split(t[1], ",")
	x, _ := strconv.Atoi(c[0])
	y, _ := strconv.Atoi(c[1])
	d := strings.Split(t[2], "x")
	width, _ := strconv.Atoi(d[0])
	height, _ := strconv.Atoi(d[1])
	return claim{id: id, x: x, y: y, width: width, height: height}
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")
	overlaps := map[string]int{}
	claims := []claim{}
	for _, v := range data {
		c := makeClaim(v)
		claims = append(claims, c)
	}

	for len(claims) > 0 {
		currentClaim := claims[0]
		claims = claims[1:]
		for _, v := range claims {
			inter := currentClaim.intersection(v)
			for _, v := range inter {
				overlaps[v] = 1
			}
		}
	}
	return len(overlaps)
}
func part2() int {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")
	overlaps := map[string]int{}
	claims := []claim{}
	hasOverlap := map[int]int{}

	for _, v := range data {
		c := makeClaim(v)
		claims = append(claims, c)
	}
	for _, v := range claims {
		hasOverlap[v.id] = 1
	}
	for len(claims) > 0 {
		currentClaim := claims[0]
		claims = claims[1:]
		for _, v := range claims {
			inter := currentClaim.intersection(v)
			if len(inter) > 0 {
				delete(hasOverlap, currentClaim.id)
				delete(hasOverlap, v.id)
			}
			for _, v := range inter {
				overlaps[v] = 1
			}
		}
	}
	out := 0
	if len(hasOverlap) == 1 {
		for k := range hasOverlap {
			out = k
		}
	}
	return out
}
