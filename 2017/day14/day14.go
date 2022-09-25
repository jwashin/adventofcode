package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1(false))
	fmt.Println("2.", part2(false))
}

func part1(test bool) int {
	data := "amgozmfv"
	if test {
		data = "flqrgnkx"
	}
	rows := []string{}
	for i := 0; i < 128; i++ {
		rows = append(rows, knotHash(fmt.Sprintf("%s-%d", data, i)))
	}
	count := 0
	for _, v := range rows {
		count += strings.Count(hash2binaryString(v), "1")
	}
	return count
}

type coordinate struct {
	y int
	x int
}

type coordinateList []coordinate

func (c coordinateList) contains(n coordinate) bool {
	for _, v := range c {
		if v.x == n.x && v.y == n.y {
			return true
		}
	}
	return false
}

type stringGrid []string

func (g stringGrid) toHashes() stringGrid {
	newGrid := stringGrid{}
	for _, row := range g {
		newRow := ""
		for _, v := range row {
			if v == '0' {
				newRow += "."
				continue
			}
			newRow += "#"
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func (g stringGrid) makeRegion(c coordinate) coordinateList {
	region := coordinateList{}
	if g.getItem(c) == "." {
		return region
	}
	q := coordinateList{c}
	for len(q) > 0 {
		currentItem := q[0]
		region = append(region, currentItem)
		q = q[1:]
		for _, v := range g.neighbors(currentItem) {
			if g.getItem(v) == "#" {
				if !region.contains(v) && !q.contains(v) {
					q = append(q, v)
				}
			}
		}
	}
	return region
}

func (g stringGrid) countRegions() int {
	used := coordinateList{}
	count := 0
	for y := range g {
		for x := range g[0] {
			c := coordinate{y: y, x: x}
			if !used.contains(c) {
				region := g.makeRegion(c)
				if len(region) > 0 {
					count += 1
				}
				used = append(used, region...)
			}
		}
	}
	return count
}

func (g stringGrid) getItem(c coordinate) string {
	return string(g[c.y][c.x])
}

// func (g stringGrid) setItem(c coordinate, s byte) {
// 	row := []byte(g[c.y])
// 	idx := c.x
// 	row[idx] = s
// 	g[c.y] = string(row)
// }

// func (g stringGrid) print() {
// 	for _, row := range g {
// 		fmt.Println(row)
// 	}
// }

func (g stringGrid) neighbors(c coordinate) coordinateList {
	theList := coordinateList{}
	y := c.y
	x := c.x
	if y > 0 {
		theList = append(theList, coordinate{y: y - 1, x: x})
	}
	if x > 0 {
		theList = append(theList, coordinate{y: y, x: x - 1})
	}
	if y < len(g)-1 {
		theList = append(theList, coordinate{y: y + 1, x: x})
	}
	if x < len(g[0])-1 {
		theList = append(theList, coordinate{y: y, x: x + 1})
	}
	return theList

}

func part2(test bool) int {
	data := "amgozmfv"
	if test {
		data = "flqrgnkx"
	}
	datas := []string{}
	for i := 0; i < 128; i++ {
		datas = append(datas, knotHash(fmt.Sprintf("%s-%d", data, i)))
	}
	rows := stringGrid{}
	for _, v := range datas {
		rows = append(rows, hash2binaryString(v))
	}
	rows = rows.toHashes()
	return rows.countRegions()

}

func hash2binaryString(s string) string {
	out := ""
	for _, v := range s {
		t, _ := strconv.ParseInt(string(v), 16, 0)
		n := strconv.FormatInt(t, 2)
		for len(n) < 4 {
			n = "0" + n
		}
		out += n
	}
	return out
}

// knot hash was done on day 10.

func knotHash(data string) string {
	lengths := []int{}
	for _, v := range data {
		lengths = append(lengths, int(v))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)
	z := knotHash64(lengths, false)
	newS := []int{}
	hxs := []int{}
	for _, v := range z {
		newS = append(newS, v)
		if len(newS) == 16 {
			hxs = append(hxs, xorList(newS))
			newS = []int{}
		}
	}
	d := ""
	for _, v := range hxs {
		r := fmt.Sprintf("%x", v)
		if len(r) == 1 {
			r = "0" + r
		}
		d += r
	}
	return d
}

func makeList(test bool) []int {
	out := []int{}
	n := 256
	if test {
		n = 5
	}
	x := 0
	for x < n {
		out = append(out, x)
		x += 1
	}
	return out
}

func xorList(s []int) int {
	x := s[0]
	for _, v := range s[1:] {
		x = x ^ v
	}
	return x
}

func knotHash64(lengths []int, test bool) []int {
	s := makeList(test)
	currentPosition := 0
	skipSize := 0
	for n := 0; n < 64; n++ {
		for _, length := range lengths {
			var selection []int
			if currentPosition+length > len(s)-1 {
				selection = s[currentPosition:]
				selection = append(selection, s[0:length-len(selection)]...)
			} else {
				selection = s[currentPosition : currentPosition+length]
			}
			for i, j := 0, len(selection)-1; i < j; i, j = i+1, j-1 {
				selection[i], selection[j] = selection[j], selection[i]
			}
			for i, v := range selection {
				idx := i + currentPosition
				if idx > len(s)-1 {
					idx -= len(s)
				}
				s[idx] = v
			}
			currentPosition += length + skipSize
			for currentPosition > len(s)-1 {
				currentPosition -= len(s)
			}
			skipSize += 1
		}
	}
	return s
}
