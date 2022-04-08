package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	astr := `#############
	#...........#
	###D#A#D#C###
	#D#C#B#A#
  #D#B#A#C#
	  #C#A#B#B#
	  #########`
	fmt.Println(findMinCost(astr))
}

type Coordinate struct {
	row int
	col int
}

type Tableau struct {
	data   []string
	cost   int
	parent *Tableau
}

type Move struct {
	fromColumn, toColumn int
}

func findMinCost(aString string) int {
	t := Tableau{cost: 0, parent: nil}
	t.parse(aString)
	// dijkstra
	q := map[string]*Tableau{}
	q[t.asString()] = &t
	minCost := math.MaxInt
	var winner Tableau
	for len(q) > 0 {
		tst := findMin(q)
		delete(q, tst.asString())
		moves := tst.possibleMoves()
		for _, v := range moves {
			newItem := Tableau{data: strings.Split(tst.asString(), "\n"), cost: tst.cost, parent: tst}
			newItem.doMove(v)
			c := newItem.cost
			key := newItem.asString()
			if newItem.isWinner() {
				// fmt.Println(newItem.asString(), c)
				// return newItem.cost !nope, not good enough
				if c < minCost {
					minCost = c
					winner = newItem
				}
			}
			if q[key] == nil {
				q[key] = &newItem
			} else {
				if q[key].cost > c {
					q[key].cost = c
					q[key].parent = tst
				}
			}
		}
	}
	output := []Tableau{winner}
	for winner.parent != nil {
		output = append(output, *winner.parent)
		winner = *winner.parent
	}
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	for _, k := range output {
		fmt.Println(k.asString())
	}
	return minCost
}

func findMin(q map[string]*Tableau) *Tableau {
	minCost := math.MaxInt
	var ret *Tableau
	for _, v := range q {
		if v.cost < minCost {
			minCost = v.cost
			ret = v
		}
	}
	return ret
}

func (t Tableau) possibleMoves() []Move {
	// fromColumns := []int{}
	// toColumns := []int{}
	homeColumns := map[string]int{"A": 3, "B": 5, "C": 7, "D": 9}
	moves := []Move{}
	critters := t.getCritters()
	for _, k := range critters {
		critterType := t.getCell(k.row, k.col) // A, B, C, or D
		if k.row == 1 {
			canMove := false
			// critter is in hallway. only move is to home column, if can move there
			if t.isHomey(homeColumns[critterType]) {
				// home column is clear to go
				canMove = true
				if homeColumns[critterType] < k.col {
					// move left
					for col := k.col - 1; col >= homeColumns[critterType]; col-- {
						if t.getCell(1, col) != "." {
							canMove = false
							break
						}
					}
				} else {
					// move right
					for col := k.col + 1; col <= homeColumns[critterType]; col++ {
						if t.getCell(1, col) != "." {
							canMove = false
							break
						}
					}
				}
			}
			if canMove {
				moves = append(moves, Move{k.col, homeColumns[critterType]})
				continue
			}

		} else {
			// in a home column. if all below are same, done.
			if t.isHomey(k.col) && homeColumns[critterType] == k.col {
				continue
			}
			// otherwise, may move to any legal hallway cell if can move there
			// first, can we go up to the hallway?
			canGetToHallway := true

			for row := k.row - 1; row >= 1; row-- {
				if t.getCell(row, k.col) != "." {
					canGetToHallway = false
					break
				}
			}
			if !canGetToHallway {
				continue
			}

			// then look for accessible open hallway cells
			for _, val := range []int{1, 2, 4, 6, 8, 10, 11} {
				// go left
				if k.col > val {
					for x := k.col - 1; x >= 1; x-- {
						if t.getCell(1, x) != "." {
							break
						}
						if t.getCell(1, x) == "." && x == val {
							moves = append(moves, Move{k.col, x})
						}
					}

				} else {
					// go right
					for x := k.col + 1; x <= 11; x++ {
						if t.getCell(1, x) != "." {
							break
						}
						if t.getCell(1, x) == "." && x == val {
							moves = append(moves, Move{k.col, x})
						}
					}
				}

			}
		}
	}
	return moves
}

func (t Tableau) isHomey(aColumn int) bool {
	// true if the column is empty except for critters of the right type
	for k := 1; k < len(t.data); k++ {
		d := t.getCell(k, aColumn)
		if d == "#" || d == "." {
			continue
		}
		critterHomes := map[int]string{3: "A", 5: "B", 7: "C", 9: "D"}
		if d != critterHomes[aColumn] {
			return false
		}
	}
	return true
}

func (t Tableau) isWinner() bool {
	homes := map[int]string{3: "A", 5: "B", 7: "C", 9: "D"}
	for _, v := range t.getCritters() {
		if v.row == 1 {
			return false
		}

		critterType := t.getCell(v.row, v.col)
		if homes[v.col] != critterType {
			return false
		}
	}
	return true
}

func (t *Tableau) doMove(m Move) int {
	// return the cost of the move

	fromColumn := m.fromColumn
	fromRow := t.firstAmphipod(fromColumn)
	critter := t.getCell(fromRow, fromColumn)
	toRow := t.drop(m.toColumn)
	toColumn := m.toColumn
	t.moveCritterTo(critter, toRow, toColumn, fromRow, fromColumn)
	costs := map[string]int{"A": 1, "B": 10, "C": 100, "D": 1000}
	t.cost += (intAbs(fromColumn-toColumn) + intAbs(fromRow-toRow)) * costs[critter]
	return t.cost
}

func (t *Tableau) moveCritterTo(critter string, toRow int, toColumn int, fromRow int, fromColumn int) {
	// replace the critter with open cell
	from := t.data[fromRow]
	newFromString := from[:fromColumn] + "." + from[fromColumn+1:]
	t.data[fromRow] = newFromString

	to := t.data[toRow]
	newtoString := to[:toColumn] + critter + to[toColumn+1:]
	t.data[toRow] = newtoString
}

func intAbs(aValue int) int {
	if aValue < 0 {
		return -aValue
	}
	return aValue
}

func (t Tableau) drop(column int) int {
	row := 1
	for t.getCell(row, column) == "." {
		row += 1
	}
	return row - 1
}

func (t Tableau) firstAmphipod(column int) int {
	for row := 1; row < len(t.data); row++ {
		val := t.getCell(row, column)
		if strings.Contains("ABCD", val) {
			return row
		}
	}
	return -1
}

func (t *Tableau) parse(inputString string) {
	t.data = []string{}
	split := strings.Split(inputString, "\n")
	for _, k := range split {
		d := strings.TrimSpace(k)
		if len(d) > 0 {
			if len(d) < 13 {
				d = "  " + d + "  "
			}
			t.data = append(t.data, d)
		}
	}
}

func (t Tableau) getCell(row int, col int) string {
	return string(t.data[row][col])
}

func (t Tableau) getCritters() []Coordinate {
	out := []Coordinate{}
	for row := 1; row < len(t.data)-1; row++ {
		for col := 1; col < len(t.data[0])-1; col++ {

			c := t.getCell(row, col)
			if c == "A" || c == "B" || c == "C" || c == "D" {
				out = append(out, Coordinate{row, col})
			}
		}
	}
	return out
}

func (t Tableau) asString() string {
	return strings.Join(t.data, "\n")
}

func min(aList []int) int {
	min := math.MaxInt
	for _, v := range aList {
		if v < min {
			min = v
		}
	}
	return min
}
