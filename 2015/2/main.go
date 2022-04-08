package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/aoc/2015/2/wrapping"
)

func main() {
	paperTotal := 0
	ribbonTotal := 0
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		txt := scanner.Text()
		item := wrapping.MakeBox(txt)
		paperTotal += item.PaperTotal()
		ribbonTotal += item.RibbonNeeded()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(paperTotal)
	fmt.Println(ribbonTotal)
}
