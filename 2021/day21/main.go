package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Player struct {
	id       uint64
	position uint64
	score    uint64
}

func (p *Player) play(die *Die) {
	roll1 := die.roll()
	roll2 := die.roll()
	roll3 := die.roll()
	rollSum := roll1 + roll2 + roll3
	p.position = clock10Add(p.position, rollSum)
	p.score += p.position
	// fmt.Printf("Player %d rolls %d+%d+%d and moves to space %d for a total score of %d\n", p.id, roll1, roll2, roll3, p.position, p.score)
}

type Die struct {
	value uint64
	count uint64
}

func (d *Die) roll() uint64 {
	if d.value == 100 {
		d.value = 1
	} else {
		d.value += 1
	}
	d.count += 1
	return d.value
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	fmt.Println(start(string(data), 1000))
	fmt.Println(dirac(string(data), 21))
}

func start(aString string, winningScore uint64) uint64 {
	players := []*Player{}
	data := strings.Split(aString, "\n")
	for _, p := range data {
		var xid, position uint64
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			fmt.Sscanf(p, "Player %d starting position: %d", &xid, &position)
			players = append(players, &Player{id: xid, position: position, score: 0})
		}
	}
	return playGame(players, 1000)
}

func playGame(players []*Player, winningScore uint64) uint64 {

	gameover := false

	die := Die{value: 0, count: 0}

	for !gameover {
		for _, player := range players {
			player.play(&die)
			if player.score >= winningScore {
				gameover = true
				break
			}
		}
	}
	rollCount := die.count
	minScore := int(math.MaxInt64)
	for _, k := range players {
		if k.score < uint64(minScore) {
			minScore = int(k.score)
		}
	}
	return uint64(minScore) * uint64(rollCount)

}

func dirac(aString string, winningScore uint64) uint64 {
	players := []*Player{}
	data := strings.Split(aString, "\n")
	for _, p := range data {
		var xid, position uint64
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			fmt.Sscanf(p, "Player %d starting position: %d", &xid, &position)
			players = append(players, &Player{id: xid, position: position, score: 0})
		}
	}
	return playDirac(players, winningScore)
}

func playDirac(players []*Player, winningScore uint64) uint64 {

	// key is "location, score", and value is how many are there at the moment
	// player1Data := map[string]uint64{}
	// player2Data := map[string]uint64{}

	playerData := map[string]uint64{}

	// playerID, location, score
	ft := "%d,%d,%d"

	for _, k := range players {
		playerData[fmt.Sprintf(ft, k.id, k.position, 0)] = 1
		// if k.id == 1 {
		// 	player1Data[fmt.Sprintf(ft, k.position, 0)] = 1
		// } else {
		// 	player2Data[fmt.Sprintf(ft, k.position, 0)] = 1
		// }
	}
	wins := map[uint64]uint64{}
	// return the number of universes where the most-winning player wins
	// player1 start position is 4(0)
	// player1 start position is 8(0)

	// after roll1, player1 is at
	//     clock10Add(startposition +1)   5 (9)
	// and clock10Add(startposition +2)   6 (10)
	// and clock10Add(startposition +3)   7 (11)
	// # of universes is 3
	// player2 is at 8(0)

	// player2 rolls.
	// player2 is now at 9(17) 10(18) 1(9)
	// # of universes is 9

	// after roll3, player1 loc (and score) is:
	// 6 (15), 7 (16), 8 (17),
	// 7 (17), 8 (18), 9(19),
	// 8 (19), 9(20) ,10(21)

	// # of universes is 27
	// player1 has won 1, player2 has won 0

	// after roll 4, player1 loc(score) is
	allWon := false

	round := 0
	for !allWon {
		round += 1
		newData := map[string]uint64{}
		for _, playerId := range []uint64{1, 2} {
			for player, count := range playerData {
				var id, oldLoc, oldScore uint64
				fmt.Sscanf(player, ft, &id, &oldLoc, &oldScore)
				if id == playerId {
					// newData[fmt.Sprintf(ft, id, oldLoc, oldScore)] += count
					for _, roll := range []uint64{1, 2, 3} {

						newLoc := clock10Add(oldLoc, roll)
						newScore := oldScore + newLoc
						if newScore >= 21 {
							wins[id] += count
						} else {
							// if newData[fmt.Sprintf(ft, id, oldLoc, oldScore)] == 0 {
							// 	newData[fmt.Sprintf(ft, id, oldLoc, oldScore)] = count
							// }
							newData[fmt.Sprintf(ft, id, newLoc, newScore)] += count
						}
					}
				}
			}
		}
		playerData = newData

		for k := range playerData {
			var id, oldLoc, oldScore uint64
			fmt.Sscanf(k, ft, &id, &oldLoc, &oldScore)
			if oldScore == 0 {
				delete(playerData, k)
			}
		}

		// newData1 := map[string]int{}
		// for player, count := range player2Data {
		// 	newCount := count * 3
		// 	for _, roll := range []int{1, 2, 3} {

		// 		var oldLoc, oldScore int
		// 		fmt.Sscanf(player, ft, &oldLoc, &oldScore)
		// 		newLoc := clock10Add(oldLoc, roll)
		// 		newScore := oldScore + newLoc
		// 		newData1[fmt.Sprintf(ft, newLoc, newScore)] = newCount + 1
		// 	}
		// }
		// universes *= 3
		// player2Data = newData1
		allWon = allWinners(playerData)
	}

	for k, v := range playerData {
		var player, loc, score uint64
		fmt.Sscanf(k, ft, &player, &loc, &score)
		wins[player] += v
	}

	// for _, k := range player2Data {
	// 	wins2 += k
	// }

	if wins[1] > wins[2] {
		return wins[1]
	}
	return wins[2]
}

func areWeDone(p1 map[string]uint64, p2 map[string]uint64) bool {
	return allWinners(p1) && allWinners(p2)
}

func allWinners(p1 map[string]uint64) bool {
	for k := range p1 {
		spl := strings.Split(k, ",")
		score, _ := strconv.Atoi(spl[len(spl)-1])
		if score < 21 {
			return false
		}
	}
	return true

}

func clock10Add(start uint64, anInt uint64) uint64 {
	sum := start + anInt
	for sum >= 11 {
		sum -= 10
	}
	return sum
}
