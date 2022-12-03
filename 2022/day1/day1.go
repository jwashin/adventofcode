package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1.", mostCalories(string(input)))
	fmt.Println("Part 2.", mostCalories3(string(input)))
}

// 655546 too high
func mostCalories(data string) int {
	split := strings.Split(data, "\n")
	max := 0
	accum := 0
	for _, v := range split {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			if accum > max {
				max = accum
			}
			accum = 0
			continue
		}
		cals, _ := strconv.Atoi(v)
		accum += cals
	}
	if accum > max {
		max = accum
	}
	return max
}

// 129466 too low
func mostCalories3(data string) int {
	split := strings.Split(data, "\n")
	elfLoads := []int{}
	max := 0
	accum := 0
	for _, v := range split {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			if accum > max {
				max = accum
			}
			elfLoads = append(elfLoads, accum)
			accum = 0
			continue
		}
		cals, _ := strconv.Atoi(v)
		accum += cals
	}
	if accum > max {
		max = accum
	}
	elfLoads = append(elfLoads, accum)

	sort.Ints(elfLoads)
	// reverse!
	for i, j := 0, len(elfLoads)-1; i < j; i, j = i+1, j-1 {
		elfLoads[i], elfLoads[j] = elfLoads[j], elfLoads[i]
	}

	return elfLoads[0] + elfLoads[1] + elfLoads[2]
}
