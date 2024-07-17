package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards    string
	bet      int
	handType int
}

var cardOrder = "23456789TJQKA"
var cardOrder2 = "J23456789TQKA"

var handTypes = map[string]int{"Five of a kind": 7,
	"Four of a kind":  6,
	"Full house":      5,
	"Three of a kind": 4,
	"Two pair":        3,
	"One pair":        2,
	"High card":       1}

func getHandType(s string) string {
	bin := map[string]int{}
	for _, v := range s {
		bin[string(v)] += 1
	}
	bin32 := map[int]int{}
	for _, count := range bin {
		if count == 5 {
			return "Five of a kind"
		}
		if count == 4 {
			return "Four of a kind"
		}
		if count == 3 {
			bin32[3] += 1
		}
		if count == 2 {
			bin32[2] += 1
		}
	}
	if bin32[3] == 1 && bin32[2] == 1 {
		return "Full house"
	}
	if bin32[3] == 1 {
		return "Three of a kind"
	}
	if bin32[2] == 2 {
		return "Two pair"
	}
	if bin32[2] == 1 {
		return "One pair"
	}
	return "High card"
}

func getHandType2(s string) string {
	bin := map[string]int{}
	for _, v := range s {
		bin[string(v)] += 1
	}
	jokerCount := bin["J"]
	delete(bin, "J")
	bin32 := map[int]int{}
	if jokerCount == 0 {
		for _, count := range bin {
			if count == 5 {
				return "Five of a kind"
			}
			if count == 4 {
				return "Four of a kind"
			}
			if count == 3 {
				bin32[3] += 1
			}
			if count == 2 {
				bin32[2] += 1
			}
		}
		if bin32[3] == 1 && bin32[2] == 1 {
			return "Full house"
		}
		if bin32[3] == 1 {
			return "Three of a kind"
		}
		if bin32[2] == 2 {
			return "Two pair"
		}
		if bin32[2] == 1 {
			return "One pair"
		}

	} else if jokerCount == 1 {
		for _, count := range bin {
			if count == 4 {
				return "Five of a kind"
			}
			if count == 3 {
				return "Four of a kind"
			}

			if count == 2 {
				bin32[2] += 1
			}
		}
		if bin32[2] == 2 {
			return "Full house"
		}
		if bin32[2] == 1 {
			return "Three of a kind"
		}
		return "One pair"
	} else if jokerCount == 2 {
		for _, count := range bin {
			if count == 3 {
				bin32[3] += 1
			}
			if count == 2 {
				bin32[2] += 1
			}
		}
		if bin32[3] == 1 {
			return "Five of a kind"
		} else if bin32[2] == 1 {
			return "Four of a kind"
		} else {
			return "Three of a kind"
		}

	} else if jokerCount == 3 {
		for _, count := range bin {
			if count == 2 {
				return "Five of a kind"
			}
			return "Four of a kind"
		}

	} else if jokerCount >= 4 {
		return "Five of a kind"
	}
	return "High card"
}

func getData(s string) []Hand {
	output := []Hand{}
	s = strings.TrimSpace(s)
	data := strings.Split(s, "\n")
	for _, v := range data {
		parsed := strings.Fields(v)
		hand := parsed[0]
		handType := handTypes[getHandType(hand)]
		bet, _ := strconv.Atoi(parsed[1])
		output = append(output, Hand{cards: hand, handType: handType, bet: bet})
	}
	return output
}

func getData2(s string) []Hand {
	output := []Hand{}
	s = strings.TrimSpace(s)
	data := strings.Split(s, "\n")
	for _, v := range data {
		parsed := strings.Fields(v)
		hand := parsed[0]
		handType := handTypes[getHandType2(hand)]
		bet, _ := strconv.Atoi(parsed[1])
		output = append(output, Hand{cards: hand, handType: handType, bet: bet})
	}
	return output
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Println("Part1", part1(string(data)))
	fmt.Println("Part2", part2(string(data)))
}

func part1(s string) int {
	hands := getData(s)
	slices.SortFunc(hands, cmpHand)
	out := 0
	for index, hand := range hands {
		out += (index + 1) * hand.bet
	}
	return out
}

func part2(s string) int {
	// 248396258 too high
	// 246760320 too high
	// 246621542 too high
	hands := getData2(s)
	slices.SortFunc(hands, cmpHand2)
	out := 0
	for index, hand := range hands {
		out += (index + 1) * hand.bet
	}
	return out
}

func cmpHand(a, b Hand) int {
	handA := a.cards
	handB := b.cards
	if a.handType == b.handType {
		for item := range handA {
			if handA[item] != handB[item] {
				cardRankA := strings.Index(cardOrder, string(handA[item]))
				cardRankB := strings.Index(cardOrder, string(handB[item]))
				return cmp.Compare(cardRankA, cardRankB)
			}
		}
	}
	return cmp.Compare(a.handType, b.handType)
}
func cmpHand2(a, b Hand) int {
	handA := a.cards
	handB := b.cards
	if a.handType == b.handType {
		for item := range handA {
			if handA[item] != handB[item] {
				cardRankA := strings.Index(cardOrder2, string(handA[item]))
				cardRankB := strings.Index(cardOrder2, string(handB[item]))
				return cmp.Compare(cardRankA, cardRankB)
			}
		}
	}
	return cmp.Compare(a.handType, b.handType)
}
