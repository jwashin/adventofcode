package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	x     int
	y     int
	used  int
	avail int
}

func (n Node) id() string {
	return fmt.Sprintf("%d,%d", n.x, n.y)
}

func main() {
	nodes := getNodes()
	fmt.Println("1.", countViablePairs(nodes))
}

// /dev/grid/node-x0-y2     86T   68T    18T   79%
func getNodes() map[string]Node {
	nodes := map[string]Node{}
	r, _ := regexp.Compile(`/dev/grid/node-x(\d*)-y(\d*)\s*(\d*)T\s*(\d*)T\s*(\d*)T\s*(\d*)\%`)

	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")
	for _, v := range data {
		t := r.FindStringSubmatch(v)
		if len(t) > 0 {
			n := Node{}
			x, _ := strconv.Atoi(t[1])
			n.x = x
			y, _ := strconv.Atoi(t[2])
			n.y = y
			ky := fmt.Sprintf("%d,%d", x, y)
			n.used, _ = strconv.Atoi(t[4])
			n.avail, _ = strconv.Atoi(t[5])
			nodes[ky] = n
		}
	}
	return nodes
}

// 1007764 too high

func countViablePairs(nodes map[string]Node) int {
	count := 0
	minx := math.MaxInt
	miny := math.MaxInt
	maxx := 0
	maxy := 0
	for _, val := range nodes {
		if val.x > maxx {
			maxx = val.x
		}
		if val.x < minx {
			minx = val.x
		}
		if val.y > maxy {
			maxy = val.y
		}
		if val.y < miny {
			miny = val.y
		}
		if val.avail == 0 {
			fmt.Printf("0 at %s\n", val.id())
		}
	}
	pout := []string{}
	for y := miny; y <= maxy; y++ {
		newString := ""
		for x := minx; x <= maxx; x++ {
			ky := fmt.Sprintf("%d,%d", x, y)
			if nodes[ky].used == 0 {
				newString += "_"
				continue
			}
			if nodes[ky].used > 400 {
				newString += "#"
				continue
			}
			if y == miny && x == maxx {
				newString += "G"
				continue
			}
			newString += "."
		}
		pout = append(pout, newString)
	}
	for _, v := range pout {
		fmt.Println(v)
	}
	for _, a := range nodes {
		for _, b := range nodes {
			if a.used > 0 && a.id() != b.id() && a.used <= b.avail {
				count += 1
			}
		}
	}
	return count

}

func connected(a Node, b Node, grid []Node) bool {
	return true
}
