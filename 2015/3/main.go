package main

import (
	"fmt"

	"github.com/jwashin/aoc/2015/day25/weathercode"
)

func main() {

	start := weathercode.Coordinate{Rank: 1, Row: 1, Column: 1, Value: 20151125}
	fmt.Println(weathercode.ValueAt(start, 2981, 3075))

}
