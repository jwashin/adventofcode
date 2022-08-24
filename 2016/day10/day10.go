package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	bots := initBots(strings.TrimSpace(string(input)))
	fmt.Println(doInstructions(bots))
}

type Move struct {
	from  string
	to    string
	value int
}

func doInstructions(bots map[string]*Bot) int {

	moreToDo := true

	for moreToDo {
		// get moves
		moves := []Move{}
		for k, v := range bots {
			if strings.Contains(k, "output") {
				continue
			}
			if len(v.purse) == 2 {
				hi := 0
				lo := 0
				p := v.purse
				a, b := 61, 17
				if slices.Contains(v.purse, a) &&
					slices.Contains(v.purse, b) {
					fmt.Printf("%v compares %v and %v.", k, a, b)
				}
				if p[1] < p[0] {
					lo = p[1]
					hi = p[0]
				} else {
					lo = p[0]
					hi = p[1]
				}

				moves = append(moves, Move{from: k, to: v.lowDest, value: lo})
				moves = append(moves, Move{from: k, to: v.highDest, value: hi})

			}
		}
		if len(moves) == 0 {
			moreToDo = false
		}
		for _, v := range moves {
			bots[v.from].remove(v.value)
			bots[v.to].take(v.value)
		}
	}
	a := bots["output 0"].purse[0]
	b := bots["output 1"].purse[0]
	c := bots["output 2"].purse[0]

	return a * b * c
}

type Bot struct {
	purse    []int
	lowDest  string
	highDest string
}

func (b *Bot) take(v int) {
	b.purse = append(b.purse, v)
}

func (b *Bot) remove(t int) {
	newList := []int{}
	for _, v := range b.purse {
		if v != t {
			newList = append(newList, v)
		}
	}
	b.purse = newList
}

func getValue(s string) (int, string) {
	rvalue, _ := regexp.Compile(
		`value (\d{1,3}) goes to (bot \d{1,3})`)
	t := rvalue.FindAllStringSubmatch(s, 1)
	if len(t[0]) > 0 {
		v, _ := strconv.Atoi(t[0][1])
		return v, t[0][2]
	}
	return 0, ""
}

// rvalue, _ := regexp.Compile(
// `(bot\s*\d{1,3}) gives low to (bot|output\s*\d{1,3}) and high to (bot|output\s*\d{1,4})`)
func getInstructions(s string) (string, string, string) {
	rvalue, _ := regexp.Compile(`(bot \d{1,}) gives low to (bot \d{1,}|output \d{1,}) and high to (bot \d{1,}|output \d{1,})`)
	t := rvalue.FindStringSubmatch(s)
	if len(t) > 0 {
		return t[1], t[2], t[3]
	}
	return "", "", ""
}

func initBots(instructions string) map[string]*Bot {
	// containers := map[string][]int{}
	// rbot, _ := regexp.Compile(`(bot \d{1,3})`)
	// routput, _ := regexp.Compile(`(output \d{1,3})`)

	bots := map[string]*Bot{}
	for _, instruction := range strings.Split(instructions, "\n") {
		if strings.Contains(instruction, "gives") {
			// var botNumber, first, second string
			// fmt.Sscanf(instruction, "%s gives low to %s and high to %s", &botNumber, &first, &second)
			botNumber, first, second := getInstructions(instruction)
			for _, v := range []string{first, second} {
				if strings.Contains(v, "output") {
					bots[v] = &Bot{purse: []int{}}
					continue
				}
			}
			bots[botNumber] = &Bot{highDest: second, lowDest: first, purse: []int{}}
		}
	}
	for _, instruction := range strings.Split(instructions, "\n") {
		if strings.Contains(instruction, "value") {
			// var botNumber, value string
			// fmt.Sscanf(instruction, "value %s goes to %s", &value, &botNumber)
			value, botNumber := getValue(instruction)
			bots[botNumber].take(value)
		}
	}
	return bots
}
