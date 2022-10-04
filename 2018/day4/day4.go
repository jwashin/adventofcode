package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	part1()
}

type logEntry struct {
	dt   time.Time
	item string
}

type guard struct {
	id    int
	log   []logEntry
	sleep map[int]int
}

func (g *guard) totalSleep() int {
	sum := 0
	for _, v := range g.sleep {
		sum += v
	}
	return sum
}
func (g *guard) calculateSleep() {
	g.sleep = map[int]int{}
	var startSleep, stopSleep int
	total := 0
	for _, v := range g.log {
		if strings.Contains(v.item, "asleep") {
			startSleep = v.dt.Minute()
		}
		if strings.Contains(v.item, "wake") {
			stopSleep = v.dt.Minute()
			total += stopSleep - startSleep
			for i := startSleep; i < stopSleep; i++ {
				g.sleep[i] += 1
			}
		}
	}
}

// fmt string is how the following date would look
// in the format we want to parse
// Mon Jan 2 15:04:05 -0700 MST 2006
func parseDateTime(s string) time.Time {
	fmt := "2006-01-02 15:04"
	t, _ := time.Parse(fmt, s)
	return t
}

func part1() {
	input, _ := os.ReadFile("input.txt")
	xdata := strings.Split(string(input), "\n")
	data := []logEntry{}
	for _, v := range xdata {
		st := strings.Index(v, "[")
		e := strings.Index(v, "]")
		dt := v[st+1 : e]
		t := parseDateTime(dt)
		data = append(data, logEntry{dt: t, item: v[e+1:]})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].dt.Before(data[j].dt)
	})
	guards := map[int]*guard{}
	guardNumber := 0
	for _, v := range data {
		if strings.Contains(v.item, "Guard") {
			t := strings.ReplaceAll(v.item, "#", "")
			guardNumber, _ = strconv.Atoi(strings.Fields(t)[1])
			found := false
			for k := range guards {
				if k == guardNumber {
					found = true
					break
				}
			}
			if !found {
				g := guard{id: guardNumber, log: []logEntry{}}
				guards[guardNumber] = &g

			}
			continue
		}
		guards[guardNumber].log = append(guards[guardNumber].log, v)
	}
	for _, v := range guards {
		v.calculateSleep()
	}
	max := 0
	guard := -1
	for k, v := range guards {
		if v.totalSleep() > max {
			max = v.totalSleep()
			guard = k
		}
	}
	max = 0
	f := guards[guard]
	minute := 0
	for k, v := range f.sleep {
		if v > max {
			max = v
			minute = k
		}
	}
	fmt.Println("1.", minute*guard)
	max = 0
	theGuard := 0
	for guardId := range guards {
		for k, v := range guards[guardId].sleep {
			if v > max {
				theGuard = guardId
				minute = k
				max = v
			}
		}
	}
	fmt.Println("2.", minute*theGuard)
}
