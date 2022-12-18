package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1:", doTetris(string(input), 2022))
	fmt.Println("Part 2:", doTetris(string(input), 1000000000000))
}

func getRocks() map[int][]string {
	rocks := map[int][]string{}

	rocks[0] = []string{"####"}

	rocks[1] = strings.Split(`.#.
###
.#.`, "\n")

	rocks[2] = strings.Split(`..#
..#
###`, "\n")
	rocks[3] = strings.Split(`#
#
#
#`, "\n")
	rocks[4] = strings.Split(`##
##`, "\n")
	return rocks
}

type tile struct {
	configuration []string
	// coordinate of bottom left corner
	loc coord
}

func (t tile) members() []coord {
	out := []coord{}
	for k, v := range t.configuration {
		cy := len(t.configuration) - k + t.loc.y - 1
		for w, c := range v {
			if c == '#' {
				x := t.loc.x + w
				out = append(out, coord{x, cy})
			}
		}
	}
	return out
}

type coord struct {
	x int
	y int
}

// set implementation
// https://linuxhint.com/golang-set/
type void struct{}

var member void

type chamber struct {
	fills     map[coord]void
	leftWall  int
	rightWall int
	minY      int
}

func (c chamber) numberAt(a int) int {
	y := a
	s := ""
	for x := 1; x < c.rightWall; x++ {
		if c.contains(coord{x, y}) {
			s += "1"
		} else {
			s += "0"
		}
	}
	value, _ := strconv.ParseInt(s, 2, 64)
	return int(value)
}

func (c *chamber) reorg(d coord) {
	got := map[int]void{}
	minY := 0
	for y := d.y; y > c.minY; y-- {
		for x := range []int{1, 2, 3, 4, 5, 6} {
			if c.contains(coord{x, y}) {
				got[x] = member
				if len(got) == 5 {
					minY = y
					break
				}
			}
		}
	}
	if minY > 0 {
		for v := range c.fills {
			if v.y < minY {
				delete(c.fills, v)
			}
		}
	}
}
func (c chamber) String() string {
	yMax := 0
	for k := range c.fills {

		if k.y > yMax {
			yMax = k.y
		}

	}
	s := []string{}
	for y := yMax; y >= 0; y-- {
		t := ""
		for x := c.leftWall; x <= c.rightWall; x++ {
			if x == c.leftWall || x == c.rightWall {
				if y == 0 {
					t += "+"
				} else {
					t += "|"
				}

			} else if c.contains(coord{x, y}) {
				t += "#"
			} else if y != 0 {
				t += "."
			} else {
				t += "-"
			}

		}
		s = append(s, t)
	}
	return strings.Join(s, "\n")
}

func (c chamber) contains(m coord) bool {
	if _, ok := c.fills[m]; ok {
		return true
	} else {
		return false
	}
}

func (c chamber) height() int {
	max := 0
	for v := range c.fills {
		if v.y > max {
			max = v.y
		}
	}
	return max
}

func (c chamber) rightAdjacent(r tile) bool {
	for _, v := range r.members() {
		if v.x+1 == c.rightWall {
			return true
		}
		if c.contains(coord{v.x + 1, v.y}) {
			return true
		}
	}
	return false
}

func (c chamber) leftAdjacent(r tile) bool {
	for _, v := range r.members() {
		if v.x-1 == c.leftWall {
			return true
		}
		if c.contains(coord{v.x - 1, v.y}) {
			return true
		}
	}
	return false
}

func (c chamber) canMoveDown(r tile) bool {
	for _, v := range r.members() {
		if v.y == 1 {
			return false
		}
		if c.contains(coord{v.x, v.y - 1}) {
			return false
		}
	}
	return true
}

func (c *chamber) add(r tile) {
	for _, v := range r.members() {
		c.fills[v] = member
	}
}

func doTetris(s string, count int) int {
	leftWall := 0
	rightWall := 8
	chamber := chamber{leftWall: leftWall, rightWall: rightWall, fills: map[coord]void{}}
	tiles := getRocks()
	jetIdx := 0
	outputs := 0
	for rockCount := 0; rockCount < count; rockCount++ {
		newC := coord{leftWall + 3, chamber.height() + 4}
		rock := tile{configuration: tiles[rockCount%len(tiles)], loc: newC}
		for {
			if jetIdx == len(s) {
				outputs += 1
				if outputs == 25 {
					ch := chamber.height()
					newS := ""
					for y := 0; y < ch; y++ {
						i := chamber.numberAt(y) + 48
						b := string(byte(i))
						newS += b
					}
					rge := len(s)
					first, second := 0, 0
					for ix := range newS {
						z := newS[ix : ix+rge]
						first = strings.Index(newS, z)
						if first > 0 {
							second = strings.Index(newS[first+1:], z)
						}
						if second > 0 {
							break
						}
					}
					repeats := (count - first) / (second - first)
					remainder := (count - first) % (second - first)
					return (count - first) + repeats*(second-first) + remainder
				}
				jetIdx = 0
			}
			jet := string(s[jetIdx])
			// fmt.Println(jet)
			if jet == ">" {
				// move right
				if !chamber.rightAdjacent(rock) {
					testC := coord{rock.loc.x + 1, rock.loc.y}
					rock = tile{configuration: rock.configuration, loc: testC}
				}
			} else if jet == "<" {
				// move left
				if !chamber.leftAdjacent(rock) {
					testC := coord{rock.loc.x - 1, rock.loc.y}
					rock = tile{configuration: rock.configuration, loc: testC}
				}
			}
			jetIdx += 1
			// move down
			if !chamber.canMoveDown(rock) {
				chamber.add(rock)
				// if rockCount%10000 == 0 {
				// 	chamber.reorg(rock.loc)
				// }
				// fmt.Println(chamber)
				// fmt.Println()
				break
			}
			newLoc := coord{rock.loc.x, rock.loc.y - 1}
			rock = tile{configuration: rock.configuration, loc: newLoc}
		}
	}
	return chamber.height()
}
