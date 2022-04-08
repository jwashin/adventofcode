package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jwashin/aoc/2021/day1/sonar"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	theList := []string{}
	for scanner.Scan() {
		theList = append(theList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sonar.CountIncreasesbyThrees(theList))
}
