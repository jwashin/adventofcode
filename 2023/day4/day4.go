package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1: ", part1(string(data)))
	fmt.Println("Part 2: ", part2(string(data)))
}

type Card struct {
	cardNumber int
	winners    []string
	candidates []string
	value      int
}

func part2(s string) int {
	// 1520588 too low
	s = strings.TrimSpace(s)
	cardData := strings.Split(s, "\n")

	cardCount := 0
	cardMap := map[int]Card{}
	cards := []Card{}
	for _, card := range cardData {
		data := strings.Fields(card)
		winners := []string{}
		candidates := []string{}
		colon, bar, cardNumber := 0, 0, 0
		for index, value := range data {
			if strings.Contains(value, ":") {
				colon = index
				cardNumber, _ = strconv.Atoi(value[:len(value)-1])
			}
			if strings.Contains(value, "|") {
				bar = index
				break
			}
		}
		for index, value := range data {
			// v, _ := strconv.Atoi(value)
			if index >= colon+1 && index < bar {
				winners = append(winners, value)
			} else if index > bar {
				candidates = append(candidates, value)
			}
		}
		cardMap[cardNumber] = Card{cardNumber: cardNumber, winners: winners, candidates: candidates}
		cards = append(cards, cardMap[cardNumber])
	}

	for len(cards) > 0 {
		// Card   1: 58  6 71 93 96 38 25 29 17  8 | 79 33 93 58 53 96 71  8 67 90 17  6 46 85 64 25 73 32 18 52 77 16 63  2 38
		card := cards[0]
		gameCount := 0
		if cardMap[card.cardNumber].value > 0 {
			gameCount = cardMap[card.cardNumber].value
		} else {
			gameCount = countMatches(card.winners, card.candidates)
			card.value = gameCount
			cardMap[card.cardNumber] = card
		}
		if gameCount > 0 {
			count := 0
			for count < gameCount {
				count += 1
				idno := card.cardNumber + count
				cards = append(cards, cardMap[idno])
			}
		}
		cards = cards[1:]
		cardCount += 1
	}
	return cardCount
}

func part1(s string) int {
	cards := strings.Split(s, "\n")
	grandTotal := 0
	for _, card := range cards {
		// Card   1: 58  6 71 93 96 38 25 29 17  8 | 79 33 93 58 53 96 71  8 67 90 17  6 46 85 64 25 73 32 18 52 77 16 63  2 38
		data := strings.Fields(card)
		winners := []string{}
		candidates := []string{}
		colon, bar := 0, 0
		for index, value := range data {
			if strings.Contains(value, ":") {
				colon = index
			}
			if strings.Contains(value, "|") {
				bar = index
				break
			}
		}
		for index, value := range data {
			// v, _ := strconv.Atoi(value)
			if index >= colon+1 && index < bar {
				winners = append(winners, value)
			} else if index > bar {
				candidates = append(candidates, value)
			}
		}
		gameCount := countMatches(winners, candidates)
		if gameCount > 0 {
			grandTotal += int(math.Pow(2, float64(gameCount-1)))
		}
	}
	return grandTotal
}

func countMatches(winners []string, candidates []string) int {
	count := 0
	for _, w := range winners {
		for _, c := range candidates {
			if w == c {
				count += 1
			}
		}
	}

	return count

}
