package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func fitsCriteria(s string, criterion []int) bool {
// 	s = strings.ReplaceAll(strings.TrimSpace(s), ".", " ")
// 	t := strings.Fields(s)
// 	if len(t) != len(criterion) {
// 		return false
// 	}
// 	olist := []int{}
// 	for _, v := range t {
// 		olist = append(olist, len(v))
// 	}
// 	return listEqual(olist, criterion)
// }

// func listEqual(a []int, b []int) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i, v := range a {
// 		if v != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part1:", part1(string(input)))
}

func part1(s string) int {
	// 7017 too low
	s = strings.TrimSpace(s)
	data := strings.Split(s, "\n")
	count := 0
	for _, v := range data {
		count += getResult(v)
	}
	return count
}

// func part2(s string) int {
// 	// 7017 too low
// 	s = strings.TrimSpace(s)
// 	data := strings.Split(s, "\n")
// 	count := 0
// 	for _, v := range data {
// 		count += countTheWays2(v)
// 	}
// 	return count
// }

func getResult(s string) int {
	s = strings.TrimSpace(s)
	d2 := strings.Split(s, " ")
	s = d2[0]
	ind := d2[1]
	isplit := strings.Split(ind, ",")
	indicator := []int{}
	for _, v := range isplit {
		length, _ := strconv.Atoi(v)
		indicator = append(indicator, length)
	}

	// we should now be able to
	// countTheWays(s string, indicator []int)
	return countArrangements(s, indicator)
}

type cacher struct {
	s      string
	damage string
}

func asString(lst []int) string {
	s := []string{}
	for _, v := range lst {
		s = append(s, fmt.Sprint(v))
	}
	return strings.Join(s, ",")
}

func countArrangements(springs string, damage []int) int {

	// key := cacher{springs, asString(damage)}
	// value, ok := cache[key]
	// if ok {
	// 	return value
	// }
	// if (springs.isEmpty()) return if (damage.isEmpty()) 1 else 0
	if len(springs) == 0 {
		if len(damage) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if springs[0] == '.' {
		for springs[0] == '.' {
			springs = springs[1:]
		}
		return countArrangements(springs, damage)
	}
	if springs[0] == '?' {
		docked := springs[1:]
		s1 := countArrangements("#"+docked, damage)
		s2 := countArrangements(docked, damage)
		return s1 + s2
	}
	if springs[0] == '#' {
		if len(damage) == 0 {
			return 0
		}
		// val thisDamage = damage.first()
		desiredCount := damage[0]
		if len(springs) >= desiredCount {
			start := springs[:desiredCount]
			if strings.Contains(start, ".") {
				return 0
			}
		} else {
			return 0
		}
		end := springs[desiredCount:]
		if len(end) == 0 {
			// add to cache
			return 1
		}
		if end[0] == '#' {
			return 0
		}
		return countArrangements(end, damage[1:])

		//         val remainingDamage = damage.drop(1)

		//         if (thisDamage <= springs.length && springs.take(thisDamage).none { it == '.' }) {
		//             when {
		//                 thisDamage == springs.length -> if (remainingDamage.isEmpty()) 1 else 0
		//                 springs[thisDamage] == '#' -> 0
		//                 else -> countArrangements(springs.drop(thisDamage + 1), remainingDamage, cache)
		//             }
		//         } else 0

	}
	return 0
}
