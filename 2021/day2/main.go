package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jwashin/aoc/2021/day2/piloting"
)

func main() {
	// https://zetcode.com/golang/readfile/
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(piloting.Navigate(input))
}
