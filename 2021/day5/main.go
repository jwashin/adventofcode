package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jwashin/aoc/2021/day5/lines"
)

func main() {
	data := []string{}
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data = append(data, scanner.Text())
		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines.CountIntersections(data))

}
