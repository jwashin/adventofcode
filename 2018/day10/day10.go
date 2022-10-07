package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1(false)
}

type point struct {
	x, y, xvel, yvel int
}

func (p *point) iter() {
	p.x += p.xvel
	p.y += p.yvel
}

type points []*point

func (p points) iter() {
	for _, v := range p {
		v.iter()
	}
}

func (p points) get(x, y int) *point {
	for _, v := range p {
		if x == v.x && y == v.y {
			return v
		}
	}
	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p points) zrange() int {
	xmax, ymax := -math.MaxInt, -math.MaxInt
	xmin, ymin := math.MaxInt, math.MaxInt
	for _, v := range p {
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
	return abs(xmax-xmin) + abs(ymax-ymin)
}
func (p points) view() []string {
	xmax, ymax := -math.MaxInt, -math.MaxInt
	xmin, ymin := math.MaxInt, math.MaxInt
	for _, v := range p {
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
	out := []string{}
	for y := 0; y <= ymax; y++ {
		newLine := ""
		for x := 0; x <= xmax; x++ {
			c := p.get(x, y)
			if c == nil {
				newLine += " "
			} else {
				newLine += "#"
			}
		}
		out = append(out, newLine)
	}
	return out
}

type stringGrid []string

func part1(test bool) {
	theFile := "input.txt"
	if test {
		theFile = "test.txt"
	}
	input, _ := os.ReadFile(theFile)
	data := strings.Split(string(input), "\n")
	f := points{}
	for _, v := range data {
		v = strings.ReplaceAll(v, "<", " ")
		v = strings.ReplaceAll(v, ">", " ")
		v = strings.ReplaceAll(v, ",", " ")
		t := strings.Fields(v)
		x, _ := strconv.Atoi(t[1])
		y, _ := strconv.Atoi(t[2])
		xvel, _ := strconv.Atoi(t[4])
		yvel, _ := strconv.Atoi(t[5])

		p := point{x: x, y: y, xvel: xvel, yvel: yvel}
		f = append(f, &p)
	}
	count := 0
	for f.zrange() > 2000 {

		f.iter()
		count += 1
	}
	//  TODO stop iterating when zrange starts to go up again
	for {
		d := f.view()
		for _, v := range d {
			fmt.Println(v)
		}
		fmt.Println(count)
		fmt.Println()
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		f.iter()
		count += 1
		// clear screen
		fmt.Println("\033[2J")
	}

}
