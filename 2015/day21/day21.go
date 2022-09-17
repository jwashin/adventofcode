package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

const Weapons = `Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0`

const Armor = `Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5`

const Rings = `Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`

type item struct {
	name   string
	cost   int
	damage int
	armor  int
}

type player struct {
	hitPoints int
	items     []item
}

func (p player) damage() int {
	d := 0
	for _, v := range p.items {
		d += v.damage
	}
	return d
}

func (p *player) addItem(i item) {
	found := false
	for _, v := range p.items {
		if v.name == i.name {
			found = true
		}
	}
	if !found {
		p.items = append(p.items, i)
	}
}

func (p player) cost() int {
	d := 0
	for _, v := range p.items {
		d += v.cost
	}
	return d
}

func (p player) armor() int {
	d := 0
	for _, v := range p.items {
		d += v.armor
	}
	return d
}

func part1() int {
	bigBoss := player{hitPoints: 108, items: []item{{damage: 8}, {armor: 2}}}
	weapons, armors, rings := getStore()

	armors = append(armors, item{"none", 0, 0, 0})
	rings = append(rings, item{"none", 0, 0, 0})

	minCost := math.MaxInt

	for _, w := range weapons {
		for _, a := range armors {
			for _, r := range rings {
				for _, r2 := range rings {
					p := player{hitPoints: 100}
					p.addItem(w)
					p.addItem(a)
					p.addItem(r)
					p.addItem(r2)
					if playGame(p, bigBoss) {
						if p.cost() < minCost {
							minCost = p.cost()
						}
					}
				}
			}

		}

	}
	return minCost
}

func part2() int {
	bigBoss := player{hitPoints: 108, items: []item{{damage: 8}, {armor: 2}}}
	weapons, armors, rings := getStore()

	armors = append(armors, item{"none", 0, 0, 0})
	rings = append(rings, item{"none", 0, 0, 0})

	maxCost := 0

	for _, w := range weapons {
		for _, a := range armors {
			for _, r := range rings {
				for _, r2 := range rings {
					p := player{hitPoints: 100}
					p.addItem(w)
					p.addItem(a)
					p.addItem(r)
					p.addItem(r2)
					if !playGame(p, bigBoss) {
						if p.cost() > maxCost {
							maxCost = p.cost()
						}
					}
				}
			}

		}

	}
	return maxCost
}

func playGame(player player, boss player) bool {

	for player.hitPoints > 0 && boss.hitPoints > 0 {
		boss.hitPoints -= player.damage() - boss.armor()
		if boss.hitPoints <= 0 {
			break
		}
		player.hitPoints -= boss.damage() - player.armor()
	}
	return player.hitPoints > 0
}

func storeParse(s string) (string, int, int, int) {
	s = strings.Replace(s, " +", "+", 1)
	name := ""
	cost, damage, armor := 0, 0, 0

	fmt.Sscanf(s, "%s %d %d %d", &name, &cost, &damage, &armor)

	return name, cost, damage, armor
}

func getStore() ([]item, []item, []item) {

	weapons := []item{}
	for _, v := range strings.Split(Weapons, "\n") {
		name, cost, damage, armor := storeParse(v)
		weapons = append(weapons, item{name, cost, damage, armor})
	}

	armors := []item{}
	for _, v := range strings.Split(Armor, "\n") {
		name, cost, damage, armor := storeParse(v)
		armors = append(armors, item{name, cost, damage, armor})
	}
	rings := []item{}
	for _, v := range strings.Split(Rings, "\n") {
		name, cost, damage, armor := storeParse(v)
		rings = append(rings, item{name, cost, damage, armor})
	}
	return weapons, armors, rings
}
