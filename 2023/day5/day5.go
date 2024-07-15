package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input)))
}

func getData(s string) ([]int, map[string][][]int) {
	s = strings.TrimSpace(s)
	seeds := []int{}
	translations := map[string][][]int{}
	mapName := ""
	mapLine := []int{}
	for _, line := range strings.Split(s, "\n") {
		t := strings.TrimSpace(line)
		if len(t) > 0 {
			flds := strings.Fields(t)
			if flds[0] == "seeds:" {
				for _, v := range flds[1:] {
					a, _ := strconv.Atoi(v)
					seeds = append(seeds, a)
				}
			} else if flds[1] == "map:" {
				mapName = flds[0]
				translations[mapName] = [][]int{}
			} else {
				for _, v := range flds {
					d, _ := strconv.Atoi(v)
					mapLine = append(mapLine, d)
				}

				translations[mapName] = append(translations[mapName], mapLine)
				mapLine = []int{}
			}
		}
	}
	return seeds, translations
}

func part2(s string) int {
	// 167883009 too high
	locationNumber := math.MaxInt
	s = strings.TrimSpace(s)
	// memo := map[int]int{}
	seeds, translations := getData(s)
	location := 0
	// seeds = part2seeds(seeds)
	for len(seeds) > 0 {
		fmt.Println(len(seeds))
		seed := seeds[0] - 1
		num := seeds[1]
		seeds = seeds[2:]
		for num > 0 {
			num -= 1
			seed += 1
			soil := translator(seed, translations, "seed-to-soil")
			fertilizer := translator(soil, translations, "soil-to-fertilizer")
			water := translator(fertilizer, translations, "fertilizer-to-water")
			light := translator(water, translations, "water-to-light")
			temperature := translator(light, translations, "light-to-temperature")
			humidity := translator(temperature, translations, "temperature-to-humidity")
			location = translator(humidity, translations, "humidity-to-location")

			if location < locationNumber {
				locationNumber = location
			}
		}
	}
	return locationNumber
}

func part1(s string) int {
	locationNumber := math.MaxInt
	s = strings.TrimSpace(s)
	seeds, translations := getData(s)
	for _, seed := range seeds {
		soil := translator(seed, translations, "seed-to-soil")
		fertilizer := translator(soil, translations, "soil-to-fertilizer")
		water := translator(fertilizer, translations, "fertilizer-to-water")
		light := translator(water, translations, "water-to-light")
		temperature := translator(light, translations, "light-to-temperature")
		humidity := translator(temperature, translations, "temperature-to-humidity")
		location := translator(humidity, translations, "humidity-to-location")

		if location < locationNumber {
			locationNumber = location
		}

	}
	return locationNumber
}

func translator(seed int, translator map[string][][]int, dialect string) int {
	key := translator[dialect]
	for _, rge := range key {
		destinationRangeStart := rge[0]
		sourceRangeStart := rge[1]
		rangeLength := rge[2]
		if seed >= sourceRangeStart && seed < sourceRangeStart+rangeLength {
			diff := seed - sourceRangeStart
			out := destinationRangeStart + diff
			return out
		}
	}
	return seed
}
