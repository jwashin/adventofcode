package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("1.", part2())
}

func part1() int {
	input, _ := os.ReadFile("input.txt")
	return severity(string(input))
}

func part2() int {
	input, _ := os.ReadFile("input.txt")
	return calcDelay(string(input))
}

type layer struct {
	Range     int
	depth     int
	scanner   int
	direction int
}

func (j *layer) inc() {

	if j.scanner == j.Range-1 {
		j.direction = -1
	}
	if j.scanner == 0 {
		j.direction = 1
	}
	j.scanner += j.direction
}

type firewall map[int]*layer

func (f firewall) inc() {
	for _, v := range f {
		v.inc()
	}
}
func (f firewall) getItem(i int) *layer {
	for v := range f {
		if v == i {
			return f[i]
		}
	}
	layer := layer{scanner: 1}
	return &layer
}

func calcDelay(s string) int {
	f := firewall{}
	max := 0
	for _, k := range strings.Split(s, "\n") {
		k = strings.Replace(k, ":", " :#", 1)
		d := strings.Fields(k)
		depth, _ := strconv.Atoi(d[0])
		Range, _ := strconv.Atoi(d[2])
		item := layer{depth: depth, Range: Range, scanner: 0, direction: 1}
		f[depth] = &item
		if depth > max {
			max = depth
		}
	}
	severity := 1
	delay := 0
	for severity > 0 {
		severity = 0
		delay += 1
		if delay%10000 == 0 {
			fmt.Println(delay)
		}
		f.inc()
		w := firewall{}
		for k, v := range f {
			lyr := layer{depth: v.depth, direction: v.direction, Range: v.Range, scanner: v.scanner}
			w[k] = &lyr
		}
		for i := 0; i <= max; i++ {
			layr := w.getItem(i)
			if layr.scanner == 0 {
				severity = 1
				break
			}
			w.inc()
		}
	}
	return delay
}

func severity(s string) int {
	f := firewall{}
	max := 0
	for _, k := range strings.Split(s, "\n") {
		k = strings.Replace(k, ":", " :#", 1)
		d := strings.Fields(k)
		depth, _ := strconv.Atoi(d[0])
		Range, _ := strconv.Atoi(d[2])
		item := layer{depth: depth, Range: Range, scanner: 0, direction: 1}
		f[depth] = &item
		if depth > max {
			max = depth
		}
	}
	severity := 0
	for i := 0; i <= max; i++ {
		layr := f.getItem(i)
		if layr.scanner == 0 {
			severity += layr.depth * layr.Range
		}
		f.inc()
	}
	return severity
}
