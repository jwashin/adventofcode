package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jwashin/aoc/2021/day3/diagnostic"
)

func main() {
	data := []string{}
	// https://zetcode.com/golang/readfile/
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

	fmt.Println(diagnostic.GetLifeSupportRating(data))
}
