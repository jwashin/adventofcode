package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jwashin/aoc/2021/day6/lanternfish"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lanternfish.OldCountFish(string(content), 80))
	fmt.Println(lanternfish.CountFish(string(content), 256))
	// fmt.Println(string(content))
}
