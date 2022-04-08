package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jwashin/aoc/2021/day4/bingo"
)

func main() {
	input := []string{}
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(bingo.GetResult(input))
	fmt.Println(bingo.GetLastWinner(input))
}
