package main

import (
	"fmt"
	"os"
	"strings"
)

const flying = "Flying"
const resting = "Resting"

type reindeer struct {
	name      string
	speed     int
	endurance int
	rest      int
	state     string
	duration  int
	distance  int
	points    int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	r, d := winningReindeer(string(input), 2503)
	fmt.Println("1.", r, "wins, traveling", d, "km.")
	r, d = winningReindeer2(string(input), 2503)
	fmt.Println("2.", r, "wins, with", d, "points.")
}

func winningReindeer(data string, time int) (string, int) {
	j := makeReindeers(data)
	t := 0
	for t < time {
		for _, v := range j {
			v.inc()
		}
		t += 1
	}
	max := 0
	name := ""
	for _, v := range j {
		if v.distance > max {
			max = v.distance
			name = v.name
		}
	}
	return name, max

}

func doPoints(r []*reindeer) {
	max := 0
	for _, v := range r {
		if v.distance > max {
			max = v.distance
		}
	}
	for _, v := range r {
		if v.distance == max {
			v.points += 1
		}
	}
}

func winningReindeer2(data string, time int) (string, int) {
	j := makeReindeers(data)
	t := 0
	for t < time {
		for _, v := range j {
			v.inc()
		}
		doPoints(j)
		t += 1
	}
	max := 0
	name := ""
	for _, v := range j {
		if v.points > max {
			max = v.points
			name = v.name
		}
	}
	return name, max

}

func makeReindeer(name string, speed int, endurance int, rest int) reindeer {
	return reindeer{name: name, speed: speed, endurance: endurance, rest: rest, state: flying, duration: 0, distance: 0}
}

func makeReindeers(s string) []*reindeer {
	reindeers := []*reindeer{}
	t := strings.Split(s, "\n")
	for _, v := range t {
		name := ""
		speed, endurance, rest := 0, 0, 0
		fmt.Sscanf(v, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds", &name, &speed, &endurance, &rest)
		rudolph := makeReindeer(name, speed, endurance, rest)
		reindeers = append(reindeers, &rudolph)
	}
	return reindeers
}

func (r *reindeer) inc() {
	r.duration += 1
	if r.state == flying {
		r.distance += r.speed
		if r.duration == r.endurance {
			r.state = resting
			r.duration = 0
		}
	}
	if r.state == resting {
		if r.duration == r.rest {
			r.state = flying
			r.duration = 0
		}
	}
}
