package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input)))
}

func getData(s string) ([]int, []int) {
	raw := strings.TrimSpace(s)
	rawData := strings.Split(raw, "\n")
	times := []int{}
	distances := []int{}
	for _, line := range rawData {
		parsed := strings.Fields(line)
		if parsed[0] == "Time:" {
			for _, val := range parsed[1:] {
				t, _ := strconv.Atoi(val)
				times = append(times, t)
			}
		}
		if parsed[0] == "Distance:" {
			for _, val := range parsed[1:] {
				t, _ := strconv.Atoi(val)
				distances = append(distances, t)
			}
		}
	}
	return times, distances
}

func part1(s string) int {
	times, distances := getData(s)
	accumulator := []int{}
	for race := range times {
		time := times[race]
		recordDistance := distances[race]
		winners := countWinners(time, recordDistance)
		accumulator = append(accumulator, winners)
	}
	out := 1
	for _, count := range accumulator {
		out *= count
	}
	return out
}

func part2(s string) int {
	times, distances := getData(s)
	ntstring := ""
	ndstring := ""
	for i := range times {
		ntstring += fmt.Sprint(times[i])
		ndstring += fmt.Sprint(distances[i])
	}
	// accumulator := []int{}
	recordDistance, _ := strconv.Atoi(ndstring)
	time, _ := strconv.Atoi(ntstring)
	winners := countWinners(time, recordDistance)
	// accumulator = append(accumulator, winners)
	return winners
}

func countWinners(time int, recordDistance int) int {
	hold := 0
	count := 0
	for hold <= time {
		distance := (time - hold) * hold
		if distance > recordDistance {
			count += 1
		}
		hold += 1
	}
	return count
}
