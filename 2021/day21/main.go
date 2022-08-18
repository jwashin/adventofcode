package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Player struct {
	id       int
	position int
	score    int
}

type DPlayer struct {
	id     int
	states map[DState]int
}

type DState struct {
	position int
	score    int
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
	value int
	count int
}

func (d *Die) roll() int {
	if d.value == 100 {
		d.value = 1
	} else {
		d.value += 1
	}
	d.count += 1
	return d.value
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Println(start(string(data), 1000))
	fmt.Println(dirac(string(data), 21))
}

// state = [defaultdict(int), defaultdict(int)]
// state[0][(0, pos1)] = 1   # this is player 1's initial state
// state[1][(0, pos2)] = 1   # this is player 2's initial state

// r, other_ct, wins = 0, 1, [0, 0]
// while other_ct:
//     new_state = defaultdict(int)
//     for score, pos in state[r%2]:
//         for die, ct in ((3, 1), (4, 3), (5, 6), (6, 7), (7, 6), (8, 3), (9, 1)):
//             new_pos = (pos + die - 1) % 10 + 1
//             new_score = score + new_pos
//             if new_score < WIN_SCORE:
//                 new_state[(new_score, new_pos)] += ct*state[r%2][(score, pos)]
//             else:
//                 wins[r%2]+= ct*other_ct*state[r%2][(score, pos)]
//     state[r%2] = new_state
//     r += 1
//     other_ct = sum(state[(r+1)%2].values())
// print("2:", max(wins))

func dirac(aString string, winningScore int) int {

	data := strings.Split(aString, "\n")
	playerPositions := []int{}
	for _, p := range data {
		var xid, position int
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			fmt.Sscanf(p, "Player %d starting position: %d", &xid, &position)
			playerPositions = append(playerPositions, position)
		}
	}

	players := []DPlayer{{id: 1, states: map[DState]int{{score: 0, position: playerPositions[0]}: 1}},
		{id: 2, states: map[DState]int{{score: 0, position: playerPositions[1]}: 1}}}
	// state := []Player{{score: 0,position: 5}, {score: 0,position: 9}}

	// state := map[Player]int{players[0]: 1, players[1]: 1}

	wins := []int{0, 0}
	other_ct := 1

	outcomes := [][]int{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}}

	for other_ct > 0 {
		for ix, player := range players {
			new_states := map[DState]int{}
			for state, count := range player.states {
				score := state.score
				pos := state.position
				for _, outcome := range outcomes {
					die := outcome[0]
					times := outcome[1]
					new_pos := (pos+die-1)%10 + 1
					new_score := score + new_pos
					if new_score < winningScore {
						new_states[DState{score: new_score, position: new_pos}] += times * count
					} else {
						wins[player.id] += times * other_ct * count
					}
				}

			}
			player.states = new_states
			other_ct = 0
			for _, v := range players[abs(ix-1)].states {
				other_ct += v
			}
		}
	}
	return max(wins)
}

func max(a []int) int {
	max := a[0]
	for _, val := range a {
		if val > max {
			max = val
		}
	}
	return max
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func start(aString string, winningScore int) int {
	players := []*Player{}
	data := strings.Split(aString, "\n")
	for _, p := range data {
		var xid, position int
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			fmt.Sscanf(p, "Player %d starting position: %d", &xid, &position)
			players = append(players, &Player{id: xid, position: position, score: 0})
		}
	}
	return playGame(players, 1000)
}

func playGame(players []*Player, winningScore int) int {

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
		if k.score < minScore {
			minScore = int(k.score)
		}
	}
	return minScore * rollCount

}

// func dirac(aString string, winningScore uint64) uint64 {
// 	players := []*Player{}
// 	data := strings.Split(aString, "\n")
// 	for _, p := range data {
// 		var xid, position uint64
// 		p = strings.TrimSpace(p)
// 		if len(p) > 0 {
// 			fmt.Sscanf(p, "Player %d starting position: %d", &xid, &position)
// 			players = append(players, &Player{id: xid, position: position, score: 0})
// 		}
// 	}
// 	return playDirac(players, winningScore)
// }

// func playDirac(players []*Player, winningScore uint64) uint64 {

// 	// key is "location, score", and value is how many are there at the moment
// 	// player1Data := map[string]uint64{}
// 	// player2Data := map[string]uint64{}

// 	playerData := map[string]uint64{}

// 	// playerID, location, score
// 	ft := "%d,%d,%d"

// 	for _, k := range players {
// 		playerData[fmt.Sprintf(ft, k.id, k.position, 0)] = 1
// 		// if k.id == 1 {
// 		// 	player1Data[fmt.Sprintf(ft, k.position, 0)] = 1
// 		// } else {
// 		// 	player2Data[fmt.Sprintf(ft, k.position, 0)] = 1
// 		// }
// 	}
// 	wins := map[uint64]uint64{}
// 	// return the number of universes where the most-winning player wins
// 	// player1 start position is 4(0)
// 	// player1 start position is 8(0)

// 	// after roll1, player1 is at
// 	//     clock10Add(startposition +1)   5 (9)
// 	// and clock10Add(startposition +2)   6 (10)
// 	// and clock10Add(startposition +3)   7 (11)
// 	// # of universes is 3
// 	// player2 is at 8(0)

// 	// player2 rolls.
// 	// player2 is now at 9(17) 10(18) 1(9)
// 	// # of universes is 9

// 	// after roll3, player1 loc (and score) is:
// 	// 6 (15), 7 (16), 8 (17),
// 	// 7 (17), 8 (18), 9(19),
// 	// 8 (19), 9(20) ,10(21)

// 	// # of universes is 27
// 	// player1 has won 1, player2 has won 0

// 	// after roll 4, player1 loc(score) is
// 	allWon := false

// 	round := 0
// 	for !allWon {
// 		round += 1
// 		newData := map[string]uint64{}
// 		for _, playerId := range []uint64{1, 2} {
// 			for player, count := range playerData {
// 				var id, oldLoc, oldScore uint64
// 				fmt.Sscanf(player, ft, &id, &oldLoc, &oldScore)
// 				if id == playerId {
// 					// newData[fmt.Sprintf(ft, id, oldLoc, oldScore)] += count
// 					for _, roll := range []uint64{1, 2, 3} {

// 						newLoc := clock10Add(oldLoc, roll)
// 						newScore := oldScore + newLoc
// 						if newScore >= 21 {
// 							wins[id] += count
// 						} else {
// 							// if newData[fmt.Sprintf(ft, id, oldLoc, oldScore)] == 0 {
// 							// 	newData[fmt.Sprintf(ft, id, oldLoc, oldScore)] = count
// 							// }
// 							newData[fmt.Sprintf(ft, id, newLoc, newScore)] += count
// 						}
// 					}
// 				}
// 			}
// 		}
// 		playerData = newData

// 		for k := range playerData {
// 			var id, oldLoc, oldScore uint64
// 			fmt.Sscanf(k, ft, &id, &oldLoc, &oldScore)
// 			if oldScore == 0 {
// 				delete(playerData, k)
// 			}
// 		}

// 		// newData1 := map[string]int{}
// 		// for player, count := range player2Data {
// 		// 	newCount := count * 3
// 		// 	for _, roll := range []int{1, 2, 3} {

// 		// 		var oldLoc, oldScore int
// 		// 		fmt.Sscanf(player, ft, &oldLoc, &oldScore)
// 		// 		newLoc := clock10Add(oldLoc, roll)
// 		// 		newScore := oldScore + newLoc
// 		// 		newData1[fmt.Sprintf(ft, newLoc, newScore)] = newCount + 1
// 		// 	}
// 		// }
// 		// universes *= 3
// 		// player2Data = newData1
// 		allWon = allWinners(playerData)
// 	}

// 	for k, v := range playerData {
// 		var player, loc, score uint64
// 		fmt.Sscanf(k, ft, &player, &loc, &score)
// 		wins[player] += v
// 	}

// 	// for _, k := range player2Data {
// 	// 	wins2 += k
// 	// }

// 	if wins[1] > wins[2] {
// 		return wins[1]
// 	}
// 	return wins[2]
// }

// // func areWeDone(p1 map[string]uint64, p2 map[string]uint64) bool {
// // 	return allWinners(p1) && allWinners(p2)
// // }

// func allWinners(p1 map[string]int) bool {
// 	for k := range p1 {
// 		spl := strings.Split(k, ",")
// 		score, _ := strconv.Atoi(spl[len(spl)-1])
// 		if score < 21 {
// 			return false
// 		}
// 	}
// 	return true

// }

func clock10Add(start int, anInt int) int {
	sum := start + anInt
	for sum >= 11 {
		sum -= 10
	}
	return sum
}
