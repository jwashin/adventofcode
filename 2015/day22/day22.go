package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

// pdrpsrpdm
// 2. 1289
//  took a while...

var verbose = false

type spell struct {
	name     string
	cost     int
	effect   bool
	duration int
}

var spells = map[string]spell{
	"m": {name: "Magic Missile", effect: false, cost: 53},
	"d": {name: "Drain", effect: false, cost: 73},
	"s": {name: "Shield", effect: true, duration: 6, cost: 113},
	"p": {name: "Poison", effect: true, duration: 6, cost: 173},
	"r": {name: "Recharge", effect: true, duration: 5, cost: 229},
}

type player struct {
	hitPoints int
	armor     int
	mana      int
	spent     int
}
type boss struct {
	hitPoints int
	damage    int
}

type timerList map[string]int

func (t timerList) inc() {

	for k := range t {
		t[k] -= 1
	}
}

// 1295 too high

func part2() int {

	// spells := spells
	minCost := 0
	q := map[string]int{}
	for _, start := range "mdspr" {
		q[string(start)] = spells[string(start)].cost
	}
	currentPath := ""
	// spellsInUse := timerList{}
	for len(q) > 0 {
		// find lowest q
		minCost = math.MaxInt
		for k, v := range q {
			if v < minCost {
				minCost = v
				currentPath = k
			}
		}
		delete(q, currentPath)
		bigBoss := boss{hitPoints: 55, damage: 8}
		thePlayer := player{hitPoints: 50, mana: 500}
		playerAlive, bossAlive := playGame2(&thePlayer, &bigBoss, currentPath)
		if bossAlive && !playerAlive {
			continue
		}
		cost := thePlayer.spent
		if playerAlive && !bossAlive {
			fmt.Println(currentPath)
			return cost
		}
		// ns := ""
		// for k, v := range activeSpells {
		// 	if v >= 0 {
		// 		ns += string(k)
		// 	}
		// }

		for _, k := range "mdspr" {
			newSpell := string(k)
			q[currentPath+newSpell] = cost + spells[string(k)].cost
		}

	}
	return minCost
}

func doSpell(spell string, p *player, boss *boss) {
	if spell == "p" {
		boss.hitPoints -= 3
		if verbose {
			fmt.Println("Poison deals 3 damage")
		}

	}
	if spell == "r" {
		p.mana += 101
		if verbose {
			fmt.Println("Recharge added 101 mana")
		}

	}
	if spell == "m" {
		if verbose {
			fmt.Println("Magic Missile deals 4 damage")
		}
		boss.hitPoints -= 4

	}
	if spell == "d" {
		boss.hitPoints -= 2
		p.hitPoints += 2
	}

}

func castSpell(play string, p *player, b *boss, timers timerList) timerList {
	spell := spells[play]
	p.mana -= spell.cost
	p.spent += spell.cost

	if play == "m" {

		if verbose {
			fmt.Println("Magic Missile deals 4 damage;")
		}

		b.hitPoints -= 4
		return timers
	}
	if play == "d" {
		if verbose {
			fmt.Println("Drain deals 2 damage and heals 2")
		}
		b.hitPoints -= 2
		p.hitPoints += 2
		return timers
	}
	if play == "s" {
		if verbose {
			fmt.Println("Shield activated")
		}
		p.armor = 7
	}

	timers[play] = spell.duration
	return timers

}

func playGame2(player *player, boss *boss, plays string) (bool, bool) {
	// spells := spells
	// plays is a list of plays in "mdspr", the first letters of the spells
	// player plays the plays in turn,
	// verbose := true
	timers := timerList{}
	// t := 0
	for _, p := range plays {

		// player turn

		play := string(p)
		// newSpell := spells[play]

		// one line different
		// player.hitPoints -= 1
		// if player.hitPoints <= 0 {
		// 	return false, true, timerList{}
		// }

		if player.hitPoints > 0 && boss.hitPoints > 0 {
			// t += 1
			if verbose {
				fmt.Println("-- Player turn --")
				fmt.Println("- Player has", player.hitPoints, "hit points,", player.armor, "armor,", player.mana, "mana")
				fmt.Println("- Boss has", boss.hitPoints, "hit points")
			}
			player.hitPoints -= 1
			if player.hitPoints <= 0 {
				return false, true
			}
			// do timed spells
			timers.inc()
			for k, v := range timers {
				if verbose {
					fmt.Println(spells[k].name, " timer is now ", v)
				}
				doSpell(k, player, boss)
			}
			for k, v := range timers {
				if v <= 0 {
					delete(timers, k)
					if k == "s" {
						player.armor = 0
					}
					if verbose {
						fmt.Println(spells[k].name, " timer expires")
					}

				}
			}
			if boss.hitPoints <= 0 {
				if verbose {
					fmt.Println("- Boss dies.")
					fmt.Println("")
				}
				return player.hitPoints > 0, false
			}

			for k := range timers {
				if play == k {
					return false, true
				}
			}

			if verbose {
				fmt.Println("player casts", spells[play].name+".")
			}

			timers = castSpell(play, player, boss, timers)
			if player.mana < 0 {
				return false, true
			}

			// boss turn

			if verbose {
				fmt.Println("")
				fmt.Println("-- Boss turn --")
				fmt.Println("- Player has", player.hitPoints, "hit points,", player.armor, "armor,", player.mana, "mana")
				fmt.Println("- Boss has", boss.hitPoints, "hit points.")
			}
			// do timed spells
			timers.inc()
			for k, v := range timers {
				if verbose {
					fmt.Println(spells[k].name, " timer is now ", v)
				}
				doSpell(k, player, boss)
			}
			for k, v := range timers {
				if v <= 0 {
					delete(timers, k)
					if k == "s" {
						player.armor = 0
					}
					if verbose {
						fmt.Println(spells[k].name, " timer expires")
					}
				}
			}

			if boss.hitPoints <= 0 {
				if verbose {
					fmt.Println("- Boss dies.")
					fmt.Println("")
				}
				break
				// return player.hitPoints > 0, false, timers
			}
			playerDamage := boss.damage - player.armor
			if playerDamage < 1 {
				playerDamage = 1
			}
			if verbose {
				fmt.Println("Boss attacks for ", playerDamage, "damage.")
				fmt.Println("")
			}
			player.hitPoints -= playerDamage
			if player.hitPoints <= 0 {
				return false, boss.hitPoints > 0
			}
		}

	}
	return player.hitPoints > 0, boss.hitPoints > 0
}

// 1481 too high
// answer, fyi
// rpsmmpmmm
// 1. 953

// func part1() int {

// 	// spells := spells
// 	minCost := 0
// 	q := map[string]int{"m": 0, "d": 0, "s": 0, "p": 0, "r": 0}
// 	for _, start := range "mdspr" {
// 		bigBoss := boss{hitPoints: 55, damage: 8}
// 		thePlayer := player{hitPoints: 50, mana: 500}
// 		playGame(&thePlayer, &bigBoss, string(start))
// 		q[string(start)] = thePlayer.spent
// 	}
// 	currentPath := ""
// 	// spellsInUse := timerList{}
// 	for len(q) > 0 {
// 		// find lowest q
// 		minCost = math.MaxInt
// 		for k, v := range q {
// 			if v < minCost {
// 				minCost = v
// 				currentPath = k
// 			}
// 		}
// 		delete(q, currentPath)
// 		bigBoss := boss{hitPoints: 55, damage: 8}
// 		thePlayer := player{hitPoints: 50, mana: 500}
// 		playerAlive, bossAlive, activeSpells := playGame(&thePlayer, &bigBoss, currentPath)
// 		if bossAlive && !playerAlive {
// 			continue
// 		}
// 		cost := thePlayer.spent
// 		if playerAlive && !bossAlive {
// 			fmt.Println(currentPath)
// 			return cost
// 		}
// 		ns := ""
// 		for k := range activeSpells {
// 			ns += string(k)
// 		}
// 		for _, k := range "mdspr" {
// 			newSpell := string(k)
// 			if !strings.Contains(ns, newSpell) {
// 				q[currentPath+newSpell] = cost + spells[newSpell].cost
// 			}
// 		}

// 	}
// 	return minCost
// }

// func playGame(player *player, boss *boss, plays string) (bool, bool, timerList) {
// 	// spells := spells
// 	// plays is a list of plays in "mdspr", the first letters of the spells
// 	// player plays the plays in turn,
// 	timers := timerList{}
// 	t := 0
// 	for _, p := range plays {
// 		play := string(p)
// 		newSpell := spells[play]
// 		player.mana -= newSpell.cost
// 		if player.mana <= 0 {
// 			return false, true, timerList{}
// 		}
// 		player.spent += newSpell.cost
// 		if player.hitPoints > 0 && boss.hitPoints > 0 {
// 			t += 1
// 			// do timed spells
// 			for k := range timers {
// 				doSpell(k, player, boss)
// 			}
// 			z := timers.inc()
// 			if z == "s" {
// 				undoShield(player)
// 			}
// 			// do the new spell
// 			if newSpell.effect {
// 				timers[play] = newSpell.duration
// 				if play == "s" {
// 					doShield(player)
// 				}
// 			} else {
// 				doSpell(play, player, boss)
// 			}
// 			// boss turn
// 			t += 1
// 			// do timed spells
// 			for k := range timers {
// 				doSpell(k, player, boss)
// 			}
// 			z = timers.inc()
// 			if z == "s" {
// 				undoShield(player)
// 			}

// 			if boss.hitPoints <= 0 {
// 				break
// 				// return player.hitPoints > 0, false, timers
// 			}
// 			player.hitPoints -= boss.damage - player.armor
// 		}

// 	}
// 	return player.hitPoints > 0, boss.hitPoints > 0, timers
// }
